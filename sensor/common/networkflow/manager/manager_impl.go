package manager

import (
	"errors"
	"fmt"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/internalapi/sensor"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/centralsensor"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/net"
	"github.com/stackrox/rox/pkg/netutil"
	"github.com/stackrox/rox/pkg/networkgraph"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/timestamp"
	"github.com/stackrox/rox/sensor/common/clusterentities"
	"github.com/stackrox/rox/sensor/common/detector"
	"github.com/stackrox/rox/sensor/common/externalsrcs"
	"github.com/stackrox/rox/sensor/common/metrics"
	flowMetrics "github.com/stackrox/rox/sensor/common/networkflow/metrics"
)

const (
	// Wait at least this long before determining that an unresolvable IP is "outside of the cluster".
	clusterEntityResolutionWaitPeriod = 10 * time.Second
	// Wait at least this long before giving up on resolving the container for a connection
	maxContainerResolutionWaitPeriod = 2 * time.Minute

	connectionDeletionGracePeriod = 5 * time.Minute
)

var (
	// these are "canonical" external addresses sent by collector when we don't care about the precise IP address.
	externalIPv4Addr = net.ParseIP("255.255.255.255")
	externalIPv6Addr = net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
)

type hostConnections struct {
	hostname           string
	connections        map[connection]*connStatus
	endpoints          map[containerEndpoint]*connStatus
	lastKnownTimestamp timestamp.MicroTS

	// connectionsSequenceID is the sequence ID of the current active connections state
	connectionsSequenceID int64
	// currentSequenceID is the sequence ID of the most recent `Register` call
	currentSequenceID int64

	pendingDeletion *time.Timer

	mutex sync.Mutex
}

type connStatus struct {
	firstSeen timestamp.MicroTS
	lastSeen  timestamp.MicroTS
	used      bool
	// rotten implies we expected to correlate the flow with a container, but were unable to
	rotten bool
}

type networkConnIndicator struct {
	srcEntity networkgraph.Entity
	dstEntity networkgraph.Entity
	dstPort   uint16
	protocol  storage.L4Protocol
}

func (i *networkConnIndicator) toProto(ts timestamp.MicroTS) *storage.NetworkFlow {
	proto := &storage.NetworkFlow{
		Props: &storage.NetworkFlowProperties{
			SrcEntity:  i.srcEntity.ToProto(),
			DstEntity:  i.dstEntity.ToProto(),
			DstPort:    uint32(i.dstPort),
			L4Protocol: i.protocol,
		},
	}

	if ts != timestamp.InfiniteFuture {
		proto.LastSeenTimestamp = ts.GogoProtobuf()
	}
	return proto
}

type containerEndpointIndicator struct {
	entity   networkgraph.Entity
	port     uint16
	protocol storage.L4Protocol
}

func (i *containerEndpointIndicator) toProto(ts timestamp.MicroTS) *storage.NetworkEndpoint {
	proto := &storage.NetworkEndpoint{
		Props: &storage.NetworkEndpointProperties{
			Entity:     i.entity.ToProto(),
			Port:       uint32(i.port),
			L4Protocol: i.protocol,
		},
	}

	if ts != timestamp.InfiniteFuture {
		proto.LastActiveTimestamp = ts.GogoProtobuf()
	}
	return proto
}

type processUniqueKey struct {
	podID         string
	containerName string
	process       processInfo
}

type processListeningIndicator struct {
	key      processUniqueKey
	port     uint16
	protocol storage.L4Protocol
}

func (i *processListeningIndicator) toProto(ts timestamp.MicroTS) *storage.ProcessListeningOnPortFromSensor {
	proto := &storage.ProcessListeningOnPortFromSensor{
		Port:     uint32(i.port),
		Protocol: i.protocol,
		Process: &storage.ProcessIndicatorUniqueKey{
			PodId:               i.key.podID,
			ContainerName:       i.key.containerName,
			ProcessName:         i.key.process.processName,
			ProcessExecFilePath: i.key.process.processExec,
			ProcessArgs:         i.key.process.processArgs,
		},
	}

	return proto
}

// connection is an instance of a connection as reported by collector
type connection struct {
	local       net.NetworkPeerID
	remote      net.NumericEndpoint
	containerID string
	incoming    bool
}

func (c *connection) String() string {
	var arrow string
	if c.incoming {
		arrow = "<-"
	} else {
		arrow = "->"
	}
	return fmt.Sprintf("%s: %s %s %s", c.containerID, c.local, arrow, c.remote)
}

type processInfo struct {
	processName string
	processArgs string
	processExec string
}

func (p *processInfo) String() string {
	return fmt.Sprintf("%s: %s %s", p.processExec, p.processName, p.processArgs)
}

type containerEndpoint struct {
	endpoint    net.NumericEndpoint
	containerID string
	processKey  *processInfo
}

func (e *containerEndpoint) String() string {
	return fmt.Sprintf("%s %s: %s", e.containerID, e.processKey, e.endpoint)
}

// NewManager creates a new instance of network flow manager
func NewManager(
	clusterEntities *clusterentities.Store,
	externalSrcs externalsrcs.Store,
	policyDetector detector.Detector,
) Manager {
	mgr := &networkFlowManager{
		done:              concurrency.NewSignal(),
		connectionsByHost: make(map[string]*hostConnections),
		clusterEntities:   clusterEntities,
		sensorUpdates:     make(chan *central.MsgFromSensor),
		publicIPs:         newPublicIPsManager(),
		externalSrcs:      externalSrcs,
		policyDetector:    policyDetector,
	}

	return mgr
}

type networkFlowManager struct {
	connectionsByHost      map[string]*hostConnections
	connectionsByHostMutex sync.Mutex

	clusterEntities *clusterentities.Store
	externalSrcs    externalsrcs.Store

	enrichedConnsLastSentState     map[networkConnIndicator]timestamp.MicroTS
	enrichedEndpointsLastSentState map[containerEndpointIndicator]timestamp.MicroTS
	enrichedProcessesLastSentState map[processListeningIndicator]timestamp.MicroTS

	done          concurrency.Signal
	sensorUpdates chan *central.MsgFromSensor

	publicIPs *publicIPsManager

	policyDetector detector.Detector
}

func (m *networkFlowManager) ProcessMessage(msg *central.MsgToSensor) error {
	return nil
}

func (m *networkFlowManager) Start() error {
	go m.enrichConnections()
	go m.publicIPs.Run(&m.done, m.clusterEntities)
	return nil
}

func (m *networkFlowManager) Stop(_ error) {
	m.done.Signal()
}

func (m *networkFlowManager) Capabilities() []centralsensor.SensorCapability {
	return nil
}

func (m *networkFlowManager) ResponsesC() <-chan *central.MsgFromSensor {
	return m.sensorUpdates
}

func (m *networkFlowManager) enrichConnections() {
	ticker := time.NewTicker(time.Second * 30)

	for {
		select {
		case <-m.done.WaitC():
			return
		case <-ticker.C:
			m.enrichAndSend()

			if features.ProcessesListeningOnPort.Enabled() {
				m.enrichAndSendProcesses()
			}
		}
	}
}

func (m *networkFlowManager) enrichAndSend() {
	currentConns, currentEndpoints := m.currentEnrichedConnsAndEndpoints()

	updatedConns := computeUpdatedConns(currentConns, m.enrichedConnsLastSentState)
	updatedEndpoints := computeUpdatedEndpoints(currentEndpoints, m.enrichedEndpointsLastSentState)

	m.enrichedConnsLastSentState = currentConns
	m.enrichedEndpointsLastSentState = currentEndpoints

	if len(updatedConns)+len(updatedEndpoints) == 0 {
		return
	}

	protoToSend := &central.NetworkFlowUpdate{
		Updated:          updatedConns,
		UpdatedEndpoints: updatedEndpoints,
		Time:             types.TimestampNow(),
	}

	for _, conn := range protoToSend.Updated {
		log.Debugw("debug_updated_conn", conn.Props)
	}

	for _, endpoint := range protoToSend.UpdatedEndpoints {
		log.Debugw("debug_updated_endpoint", endpoint.Props)
	}

	// Before sending, run the flows through policies asynchronously
	for _, flow := range updatedConns {
		m.policyDetector.ProcessNetworkFlow(flow)
	}

	metrics.IncrementTotalNetworkFlowsSentCounter(len(protoToSend.Updated))
	metrics.IncrementTotalNetworkEndpointsSentCounter(len(protoToSend.UpdatedEndpoints))
	log.Debugf("Flow update : %v", protoToSend)
	select {
	case <-m.done.Done():
		return
	case m.sensorUpdates <- &central.MsgFromSensor{
		Msg: &central.MsgFromSensor_NetworkFlowUpdate{
			NetworkFlowUpdate: protoToSend,
		},
	}:
		return
	}
}

func (m *networkFlowManager) enrichAndSendProcesses() {
	currentProcesses := m.currentEnrichedProcesses()

	updatedProcesses := computeUpdatedProcesses(currentProcesses, m.enrichedProcessesLastSentState)

	m.enrichedProcessesLastSentState = currentProcesses

	if len(updatedProcesses) == 0 {
		return
	}

	processesToSend := &central.ProcessListeningOnPortsUpdate{
		ProcessesListeningOnPorts: updatedProcesses,
		Time:                      types.TimestampNow(),
	}

	metrics.IncrementTotalProcessesSentCounter(len(processesToSend.ProcessesListeningOnPorts))
	select {
	case <-m.done.Done():
		return
	case m.sensorUpdates <- &central.MsgFromSensor{
		Msg: &central.MsgFromSensor_ProcessListeningOnPortUpdate{
			ProcessListeningOnPortUpdate: processesToSend,
		},
	}:
		return
	}
}

func (m *networkFlowManager) enrichConnection(conn *connection, status *connStatus, enrichedConnections map[networkConnIndicator]timestamp.MicroTS) {
	timeElapsedSinceFirstSeen := timestamp.Now().ElapsedSince(status.firstSeen)
	isFresh := timeElapsedSinceFirstSeen < clusterEntityResolutionWaitPeriod

	container, ok := m.clusterEntities.LookupByContainerID(conn.containerID)
	if !ok {
		// Expire the connection if the container cannot be found within the clusterEntityResolutionWaitPeriod
		if timeElapsedSinceFirstSeen > maxContainerResolutionWaitPeriod {
			status.rotten = true
			// Only increment metric once the connection is marked rotten
			flowMetrics.ContainerIDMisses.Inc()
			log.Debugf("Unable to fetch deployment information for container %s: no deployment found", conn.containerID)
		}
		return
	}

	var lookupResults []clusterentities.LookupResult

	// Check if the remote address represents the de-facto INTERNET entity.
	if conn.remote.IPAndPort.Address == externalIPv4Addr || conn.remote.IPAndPort.Address == externalIPv6Addr {
		isFresh = false
	} else {
		// Otherwise, check if the remote entity is actually a cluster entity.
		lookupResults = m.clusterEntities.LookupByEndpoint(conn.remote)
	}

	if len(lookupResults) == 0 {
		// If the address is set and is not resolvable, we want to we wait for `clusterEntityResolutionWaitPeriod` time
		// before associating it to a known network or INTERNET.
		if isFresh && conn.remote.IPAndPort.Address.IsValid() {
			return
		}

		extSrc := m.externalSrcs.LookupByNetwork(conn.remote.IPAndPort.IPNetwork)
		if extSrc != nil {
			isFresh = false
		}

		if isFresh {
			return
		}
		status.used = true

		flowMetrics.ExternalFlowCounter.Inc()

		var port uint16
		if conn.incoming {
			port = conn.local.Port
		} else {
			port = conn.remote.IPAndPort.Port
		}

		if extSrc == nil {
			// Fake a lookup result.
			lookupResults = []clusterentities.LookupResult{
				{
					Entity:         networkgraph.InternetEntity(),
					ContainerPorts: []uint16{port},
				},
			}
		} else {
			lookupResults = []clusterentities.LookupResult{
				{
					Entity:         networkgraph.EntityFromProto(extSrc),
					ContainerPorts: []uint16{port},
				},
			}
		}
	} else {
		status.used = true
		if conn.incoming {
			// Only report incoming connections from outside of the cluster. These are already taken care of by the
			// corresponding outgoing connection from the other end.
			return
		}
	}

	for _, lookupResult := range lookupResults {
		for _, port := range lookupResult.ContainerPorts {
			indicator := networkConnIndicator{
				dstPort:  port,
				protocol: conn.remote.L4Proto.ToProtobuf(),
			}

			if conn.incoming {
				indicator.srcEntity = lookupResult.Entity
				indicator.dstEntity = networkgraph.EntityForDeployment(container.DeploymentID)
			} else {
				indicator.srcEntity = networkgraph.EntityForDeployment(container.DeploymentID)
				indicator.dstEntity = lookupResult.Entity
			}

			// Multiple connections from a collector can result in a single enriched connection
			// hence update the timestamp only if we have a more recent connection than the one we have already enriched.
			if oldTS, found := enrichedConnections[indicator]; !found || oldTS < status.lastSeen {
				enrichedConnections[indicator] = status.lastSeen
			}
		}
	}
}

func (m *networkFlowManager) enrichContainerEndpoint(ep *containerEndpoint, status *connStatus, enrichedEndpoints map[containerEndpointIndicator]timestamp.MicroTS) {
	timeElapsedSinceFirstSeen := timestamp.Now().ElapsedSince(status.firstSeen)
	isFresh := timeElapsedSinceFirstSeen < clusterEntityResolutionWaitPeriod
	if !isFresh {
		status.used = true
	}

	container, ok := m.clusterEntities.LookupByContainerID(ep.containerID)
	if !ok {
		// Expire the connection if the container cannot be found within the clusterEntityResolutionWaitPeriod
		if timeElapsedSinceFirstSeen > maxContainerResolutionWaitPeriod {
			status.rotten = true
			// Only increment metric once the connection is marked rotten
			flowMetrics.ContainerIDMisses.Inc()
			log.Debugf("Unable to fetch deployment information for container %s: no deployment found", ep.containerID)
		}
		return
	}

	status.used = true

	indicator := containerEndpointIndicator{
		entity:   networkgraph.EntityForDeployment(container.DeploymentID),
		port:     ep.endpoint.IPAndPort.Port,
		protocol: ep.endpoint.L4Proto.ToProtobuf(),
	}

	// Multiple endpoints from a collector can result in a single enriched endpoint,
	// hence update the timestamp only if we have a more recent endpoint than the one we have already enriched.
	if oldTS, found := enrichedEndpoints[indicator]; !found || oldTS < status.lastSeen {
		enrichedEndpoints[indicator] = status.lastSeen
	}
}

func (m *networkFlowManager) enrichProcessListening(ep *containerEndpoint, status *connStatus, processesListening map[processListeningIndicator]timestamp.MicroTS) {
	timeElapsedSinceFirstSeen := timestamp.Now().ElapsedSince(status.firstSeen)
	isFresh := timeElapsedSinceFirstSeen < clusterEntityResolutionWaitPeriod
	if !isFresh {
		status.used = true
	}

	container, ok := m.clusterEntities.LookupByContainerID(ep.containerID)
	if !ok {
		// Expire the process if the container cannot be found within the clusterEntityResolutionWaitPeriod
		if timeElapsedSinceFirstSeen > maxContainerResolutionWaitPeriod {
			status.rotten = true
			// Only increment metric once the connection is marked rotten
			flowMetrics.ContainerIDMisses.Inc()
			log.Debugf("Unable to fetch deployment information for container %s: no deployment found", ep.containerID)
		}
		return
	}

	status.used = true

	indicator := processListeningIndicator{
		key: processUniqueKey{
			podID:         container.PodID,
			containerName: container.ContainerName,
			process:       *ep.processKey,
		},
		port:     ep.endpoint.IPAndPort.Port,
		protocol: ep.endpoint.L4Proto.ToProtobuf(),
	}

	// Multiple endpoints from a collector can result in a single enriched endpoint,
	// hence update the timestamp only if we have a more recent endpoint than the one we have already enriched.
	if oldTS, found := processesListening[indicator]; !found || oldTS < status.lastSeen {
		processesListening[indicator] = status.lastSeen
	}
}

func (m *networkFlowManager) enrichHostConnections(hostConns *hostConnections, enrichedConnections map[networkConnIndicator]timestamp.MicroTS) {
	hostConns.mutex.Lock()
	defer hostConns.mutex.Unlock()

	prevSize := len(hostConns.connections)
	for conn, status := range hostConns.connections {
		m.enrichConnection(&conn, status, enrichedConnections)
		if status.rotten || (status.used && status.lastSeen != timestamp.InfiniteFuture) {
			// connections that are no longer active and have already been used can be deleted.
			delete(hostConns.connections, conn)
		}
	}
	flowMetrics.HostConnectionsRemoved.Add(float64(prevSize - len(hostConns.connections)))
}

func (m *networkFlowManager) enrichHostContainerEndpoints(hostConns *hostConnections, enrichedEndpoints map[containerEndpointIndicator]timestamp.MicroTS) {
	hostConns.mutex.Lock()
	defer hostConns.mutex.Unlock()

	prevSize := len(hostConns.endpoints)
	for ep, status := range hostConns.endpoints {
		m.enrichContainerEndpoint(&ep, status, enrichedEndpoints)
		if status.used && status.lastSeen != timestamp.InfiniteFuture {
			// endpoints that are no longer active and have already been used can be deleted.
			delete(hostConns.endpoints, ep)
		}
	}
	flowMetrics.HostEndpointsRemoved.Add(float64(prevSize - len(hostConns.endpoints)))
}

func (m *networkFlowManager) enrichProcessesListening(hostConns *hostConnections, processesListening map[processListeningIndicator]timestamp.MicroTS) {
	hostConns.mutex.Lock()
	defer hostConns.mutex.Unlock()

	prevSize := len(hostConns.endpoints)
	for ep, status := range hostConns.endpoints {
		if ep.processKey == nil {
			// No way to update a process if the data isn't there
			continue
		}

		m.enrichProcessListening(&ep, status, processesListening)
		if status.used && status.lastSeen != timestamp.InfiniteFuture {
			// endpoints that are no longer active and have already been used can be deleted.
			delete(hostConns.endpoints, ep)
		}
	}
	flowMetrics.HostProcessesRemoved.Add(float64(prevSize - len(hostConns.endpoints)))
}

func (m *networkFlowManager) currentEnrichedConnsAndEndpoints() (map[networkConnIndicator]timestamp.MicroTS, map[containerEndpointIndicator]timestamp.MicroTS) {
	allHostConns := m.getAllHostConnections()

	enrichedConnections := make(map[networkConnIndicator]timestamp.MicroTS)
	enrichedEndpoints := make(map[containerEndpointIndicator]timestamp.MicroTS)
	for _, hostConns := range allHostConns {
		m.enrichHostConnections(hostConns, enrichedConnections)
		m.enrichHostContainerEndpoints(hostConns, enrichedEndpoints)
	}

	return enrichedConnections, enrichedEndpoints
}

func (m *networkFlowManager) currentEnrichedProcesses() map[processListeningIndicator]timestamp.MicroTS {
	allHostConns := m.getAllHostConnections()

	enrichedProcesses := make(map[processListeningIndicator]timestamp.MicroTS)
	for _, hostConns := range allHostConns {
		m.enrichProcessesListening(hostConns, enrichedProcesses)
	}

	return enrichedProcesses
}

func computeUpdatedConns(current map[networkConnIndicator]timestamp.MicroTS, previous map[networkConnIndicator]timestamp.MicroTS) []*storage.NetworkFlow {
	var updates []*storage.NetworkFlow

	for conn, currTS := range current {
		prevTS, ok := previous[conn]
		if !ok || currTS > prevTS {
			updates = append(updates, conn.toProto(currTS))
		}
	}

	for conn, prevTS := range previous {
		if _, ok := current[conn]; !ok {
			updates = append(updates, conn.toProto(prevTS))
		}
	}

	return updates
}

func computeUpdatedEndpoints(current map[containerEndpointIndicator]timestamp.MicroTS, previous map[containerEndpointIndicator]timestamp.MicroTS) []*storage.NetworkEndpoint {
	var updates []*storage.NetworkEndpoint

	for ep, currTS := range current {
		prevTS, ok := previous[ep]
		if !ok || currTS > prevTS {
			updates = append(updates, ep.toProto(currTS))
		}
	}

	for ep, prevTS := range previous {
		if _, ok := current[ep]; !ok {
			updates = append(updates, ep.toProto(prevTS))
		}
	}

	return updates
}

func computeUpdatedProcesses(current map[processListeningIndicator]timestamp.MicroTS, previous map[processListeningIndicator]timestamp.MicroTS) []*storage.ProcessListeningOnPortFromSensor {
	var updates []*storage.ProcessListeningOnPortFromSensor

	for pl, currTS := range current {
		prevTS, ok := previous[pl]
		if !ok || currTS > prevTS {
			updates = append(updates, pl.toProto(currTS))
		}
	}

	for ep, prevTS := range previous {
		if _, ok := current[ep]; !ok {
			updates = append(updates, ep.toProto(prevTS))
		}
	}

	return updates
}

func (m *networkFlowManager) getAllHostConnections() []*hostConnections {
	// Get a snapshot of all *hostConnections. This allows us to lock the individual mutexes without having to hold
	// two locks simultaneously.
	m.connectionsByHostMutex.Lock()
	defer m.connectionsByHostMutex.Unlock()

	allHostConns := make([]*hostConnections, 0, len(m.connectionsByHost))
	for _, hostConns := range m.connectionsByHost {
		allHostConns = append(allHostConns, hostConns)
	}

	return allHostConns
}

func (m *networkFlowManager) RegisterCollector(hostname string) (HostNetworkInfo, int64) {
	m.connectionsByHostMutex.Lock()
	defer m.connectionsByHostMutex.Unlock()

	conns := m.connectionsByHost[hostname]

	if conns == nil {
		conns = &hostConnections{
			hostname:    hostname,
			connections: make(map[connection]*connStatus),
			endpoints:   make(map[containerEndpoint]*connStatus),
		}
		m.connectionsByHost[hostname] = conns
	}

	conns.mutex.Lock()
	defer conns.mutex.Unlock()

	if conns.pendingDeletion != nil {
		// Note that we don't need to check the return value, since `deleteHostConnections` needs to acquire
		// m.connectionsByHostMutex. It can therefore only proceed once this function returns, in which case it will be
		// a no-op due to `pendingDeletion` being `nil`.
		conns.pendingDeletion.Stop()
		conns.pendingDeletion = nil
	}

	conns.currentSequenceID++

	return conns, conns.currentSequenceID
}

func (m *networkFlowManager) deleteHostConnections(hostname string) {
	m.connectionsByHostMutex.Lock()
	defer m.connectionsByHostMutex.Unlock()

	conns := m.connectionsByHost[hostname]
	if conns == nil {
		return
	}

	conns.mutex.Lock()
	defer conns.mutex.Unlock()

	if conns.pendingDeletion == nil {
		return
	}
	flowMetrics.HostConnectionsRemoved.Add(float64(len(conns.connections)))
	delete(m.connectionsByHost, hostname)
}

func (m *networkFlowManager) UnregisterCollector(hostname string, sequenceID int64) {
	m.connectionsByHostMutex.Lock()
	defer m.connectionsByHostMutex.Unlock()

	conns := m.connectionsByHost[hostname]
	if conns == nil {
		return
	}

	conns.mutex.Lock()
	defer conns.mutex.Unlock()
	if conns.currentSequenceID != sequenceID {
		// Skip deletion if there has been a more recent Register call than the corresponding Unregister call
		return
	}

	if conns.pendingDeletion != nil {
		// Cancel any pending deletions there might be. See `RegisterCollector` on why we do not need to check for the
		// return value of Stop.
		conns.pendingDeletion.Stop()
	}
	conns.pendingDeletion = time.AfterFunc(connectionDeletionGracePeriod, func() {
		m.deleteHostConnections(hostname)
	})
}

func (h *hostConnections) Process(networkInfo *sensor.NetworkConnectionInfo, nowTimestamp timestamp.MicroTS, sequenceID int64) error {
	updatedConnections := getUpdatedConnections(h.hostname, networkInfo)
	updatedEndpoints := getUpdatedContainerEndpoints(h.hostname, networkInfo)

	collectorTS := timestamp.FromProtobuf(networkInfo.GetTime())
	tsOffset := nowTimestamp - collectorTS

	h.mutex.Lock()
	defer h.mutex.Unlock()

	if sequenceID != h.currentSequenceID {
		return errors.New("replaced by newer connection")
	} else if sequenceID != h.connectionsSequenceID {
		// This is the first message of the new connection.
		for _, status := range h.connections {
			// Mark all connections as closed this is the first update
			// after a connection went down and came back up again.
			status.lastSeen = h.lastKnownTimestamp
		}
		for _, status := range h.endpoints {
			status.lastSeen = h.lastKnownTimestamp
		}
		h.connectionsSequenceID = sequenceID
	}

	{
		prevSize := len(h.connections)
		for c, t := range updatedConnections {
			// timestamp = zero implies the connection is newly added. Add new connections, update existing ones to mark them closed
			if t != timestamp.InfiniteFuture { // adjust timestamp if not zero.
				t += tsOffset
			}
			status := h.connections[c]
			if status == nil {
				status = &connStatus{
					firstSeen: timestamp.Now(),
				}
				if t < status.firstSeen {
					status.firstSeen = t
				}
				h.connections[c] = status
			}
			status.lastSeen = t
		}

		h.lastKnownTimestamp = nowTimestamp
		flowMetrics.HostConnectionsAdded.Add(float64(len(h.connections) - prevSize))
	}

	{
		prevSize := len(h.endpoints)
		for ep, t := range updatedEndpoints {
			// timestamp = zero implies the endpoint is newly added. Add new endpoints, update existing ones to mark them closed
			if t != timestamp.InfiniteFuture { // adjust timestamp if not zero.
				t += tsOffset
			}
			status := h.endpoints[ep]
			if status == nil {
				status = &connStatus{
					firstSeen: timestamp.Now(),
				}
				if t < status.firstSeen {
					status.firstSeen = t
				}
				h.endpoints[ep] = status
			}
			status.lastSeen = t
		}

		h.lastKnownTimestamp = nowTimestamp
		flowMetrics.HostEndpointsAdded.Add(float64(len(h.endpoints) - prevSize))
	}

	return nil
}

func getProcessKey(originator *storage.NetworkProcessUniqueKey) *processInfo {
	if originator == nil {
		return nil
	}

	return &processInfo{
		processName: originator.ProcessName,
		processArgs: originator.ProcessArgs,
		processExec: originator.ProcessExecFilePath,
	}
}

func getIPAndPort(address *sensor.NetworkAddress) net.NetworkPeerID {
	tuple := net.NetworkPeerID{
		// For private address, both address and IPNetwork are expected to be set by Collector.
		// If not set, this will be invalid i.e. `IPNetwork{}`.
		IPNetwork: net.IPNetworkFromCIDRBytes(address.GetIpNetwork()),
		// If not set, this will be invalid i.e. `IPAddress{}`.
		Address: net.IPFromBytes(address.GetAddressData()),
		Port:    uint16(address.GetPort()),
	}
	return tuple
}

func getUpdatedConnections(hostname string, networkInfo *sensor.NetworkConnectionInfo) map[connection]timestamp.MicroTS {
	updatedConnections := make(map[connection]timestamp.MicroTS)

	flowMetrics.NetworkFlowMessagesPerNode.With(prometheus.Labels{"Hostname": hostname}).Inc()

	for _, conn := range networkInfo.GetUpdatedConnections() {
		var incoming bool
		switch conn.Role {
		case sensor.ClientServerRole_ROLE_SERVER:
			flowMetrics.NetworkFlowsPerNodeByType.With(prometheus.Labels{"Hostname": hostname, "Type": "incoming", "Protocol": conn.Protocol.String()}).Inc()
			incoming = true
		case sensor.ClientServerRole_ROLE_CLIENT:
			flowMetrics.NetworkFlowsPerNodeByType.With(prometheus.Labels{"Hostname": hostname, "Type": "outgoing", "Protocol": conn.Protocol.String()}).Inc()
			incoming = false
		default:
			continue
		}

		remote := net.NumericEndpoint{
			IPAndPort: getIPAndPort(conn.GetRemoteAddress()),
			L4Proto:   net.L4ProtoFromProtobuf(conn.GetProtocol()),
		}
		local := getIPAndPort(conn.GetLocalAddress())

		// Special handling for UDP ports - role reported by collector may be unreliable, so look at which port is more
		// likely to be ephemeral.
		if remote.L4Proto == net.UDP {
			incoming = netutil.IsEphemeralPort(remote.IPAndPort.Port) > netutil.IsEphemeralPort(local.Port)
		}

		c := connection{
			local:       local,
			remote:      remote,
			containerID: conn.GetContainerId(),
			incoming:    incoming,
		}

		// timestamp will be set to close timestamp for closed connections, and zero for newly added connection.
		ts := timestamp.FromProtobuf(conn.CloseTimestamp)
		if ts == 0 {
			ts = timestamp.InfiniteFuture
		}
		updatedConnections[c] = ts
	}

	return updatedConnections
}

func getUpdatedContainerEndpoints(hostname string, networkInfo *sensor.NetworkConnectionInfo) map[containerEndpoint]timestamp.MicroTS {
	updatedEndpoints := make(map[containerEndpoint]timestamp.MicroTS)

	flowMetrics.NetworkFlowMessagesPerNode.With(prometheus.Labels{"Hostname": hostname}).Inc()

	for _, endpoint := range networkInfo.GetUpdatedEndpoints() {
		flowMetrics.ContainerEndpointsPerNode.With(prometheus.Labels{"Hostname": hostname, "Protocol": endpoint.Protocol.String()}).Inc()

		ep := containerEndpoint{
			containerID: endpoint.GetContainerId(),
			endpoint: net.NumericEndpoint{
				IPAndPort: getIPAndPort(endpoint.GetListenAddress()),
				L4Proto:   net.L4ProtoFromProtobuf(endpoint.GetProtocol()),
			},
			processKey: getProcessKey(endpoint.GetOriginator()),
		}

		// timestamp will be set to close timestamp for closed connections, and zero for newly added connection.
		ts := timestamp.FromProtobuf(endpoint.GetCloseTimestamp())
		if ts == 0 {
			ts = timestamp.InfiniteFuture
		}
		updatedEndpoints[ep] = ts
	}

	return updatedEndpoints
}

func (m *networkFlowManager) PublicIPsValueStream() concurrency.ReadOnlyValueStream[*sensor.IPAddressList] {
	return m.publicIPs.PublicIPsProtoStream()
}

func (m *networkFlowManager) ExternalSrcsValueStream() concurrency.ReadOnlyValueStream[*sensor.IPNetworkList] {
	return m.externalSrcs.ExternalSrcsValueStream()
}
