package splunk

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/pkg/errors"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/internalapi/wrapper"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/httputil/proxy"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/notifiers"
	"github.com/stackrox/rox/pkg/protoutils"
	"github.com/stackrox/rox/pkg/retry"
	"github.com/stackrox/rox/pkg/urlfmt"
	"github.com/stackrox/rox/pkg/utils"
)

const (
	integrationType = "splunk"

	source                    = "stackrox"
	splunkHECDefaultDataLimit = 10000
	splunkHECHealthEndpoint   = "/services/collector/health/1.0"
	splunkHECEventEndpoint    = "/services/collector/event/1.0"

	alertSourceTypeKey = "alert"
	auditSourceTypeKey = "audit"
	jsonSourceType     = "_json"
)

var (
	log = logging.LoggerForModule()

	timeout = 5 * time.Second

	baseURLPattern = regexp.MustCompile(`^(https?://)?[^/]+/*$`)

	defaultSourceTypeMap = map[string]string{
		alertSourceTypeKey: "stackrox-alert",
		auditSourceTypeKey: "stackrox-audit-message",
	}
)

type splunk struct {
	client *http.Client

	eventEndpoint  string
	healthEndpoint string
	conf           *storage.Splunk

	*storage.Notifier
}

func (s *splunk) AlertNotify(ctx context.Context, alert *storage.Alert) error {
	return s.postAlert(ctx, alert)
}

func (s *splunk) ProtoNotifier() *storage.Notifier {
	return s.Notifier
}

func (s *splunk) Test(ctx context.Context) error {
	if s.healthEndpoint != "" {
		return s.sendHTTPPayload(ctx, http.MethodGet, s.healthEndpoint, nil)
	}
	alert := &storage.Alert{
		Policy: &storage.Policy{Name: "Test Policy"},
		Entity: &storage.Alert_Deployment_{Deployment: &storage.Alert_Deployment{Name: "Test Deployment"}},
		Violations: []*storage.Alert_Violation{
			{Message: "This is a sample Splunk alert message created to test integration with StackRox."},
		},
	}
	return s.postAlert(ctx, alert)
}

func (s *splunk) postAlert(ctx context.Context, alert *storage.Alert) error {
	clonedAlert := alert.Clone()
	// Splunk's HEC by default has a limitation of data size == 10KB
	// Removing some of the fields here to make it smaller
	// More details on HEC limitation: https://developers.perfectomobile.com/display/TT/Splunk+-+Configure+HTTP+Event+Collector
	// Check section on "Increasing the Event Data Truncate Limit"
	clonedAlert.Policy.Description = ""
	clonedAlert.Policy.Rationale = ""
	clonedAlert.Policy.Remediation = ""
	clonedAlert.Policy.Exclusions = nil
	notifiers.PruneAlert(clonedAlert, int(s.conf.GetTruncate()))

	return retry.WithRetry(
		func() error {
			return s.sendEvent(ctx, clonedAlert, alertSourceTypeKey)
		},
		retry.OnlyRetryableErrors(),
		retry.Tries(3),
		retry.BetweenAttempts(func(previousAttempt int) {
			wait := time.Duration(previousAttempt * previousAttempt * 100)
			time.Sleep(wait * time.Millisecond)
		}),
	)
}

func (s *splunk) getSplunkEvent(msg proto.Message, sourceTypeKey string) (*wrapper.SplunkEvent, error) {
	e, err := protoutils.MarshalAny(msg)
	if err != nil {
		return nil, err
	}

	return &wrapper.SplunkEvent{
		Event:      e,
		Source:     source,
		Sourcetype: s.conf.SourceTypes[sourceTypeKey],
	}, nil
}

func (*splunk) Close(_ context.Context) error {
	return nil
}

func (s *splunk) SendAuditMessage(ctx context.Context, msg *v1.Audit_Message) error {
	if !s.AuditLoggingEnabled() {
		return nil
	}

	return retry.WithRetry(
		func() error {
			return s.sendEvent(ctx, msg, auditSourceTypeKey)
		},
		retry.OnlyRetryableErrors(),
		retry.Tries(3),
		retry.BetweenAttempts(func(previousAttempt int) {
			wait := time.Duration(previousAttempt * previousAttempt * 100)
			time.Sleep(wait * time.Millisecond)
		}),
	)
}

func (s *splunk) AuditLoggingEnabled() bool {
	return s.GetSplunk().GetAuditLoggingEnabled()
}

func (s *splunk) sendEvent(ctx context.Context, msg proto.Message, sourceTypeKey string) error {
	splunkEvent, err := s.getSplunkEvent(msg, sourceTypeKey)
	if err != nil {
		return err
	}

	var data bytes.Buffer
	err = new(jsonpb.Marshaler).Marshal(&data, splunkEvent)
	if err != nil {
		return err
	}

	if data.Len() > int(s.conf.GetTruncate()) {
		return fmt.Errorf("Splunk HEC truncate data limit (%d bytes) exceeded: %d", s.conf.GetTruncate(), data.Len())
	}

	return s.sendHTTPPayload(ctx, http.MethodPost, s.eventEndpoint, &data)
}

func (s *splunk) sendHTTPPayload(ctx context.Context, method, path string, data io.Reader) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, path, data)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Splunk %s", s.conf.HttpToken))

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer utils.IgnoreError(resp.Body.Close)

	return notifiers.CreateError("Splunk", resp)
}

func init() {
	notifiers.Add(integrationType, func(notifier *storage.Notifier) (notifiers.Notifier, error) {
		s, err := newSplunk(notifier)
		return s, err
	})
}

func newSplunk(notifier *storage.Notifier) (*splunk, error) {
	conf := notifier.GetSplunk()
	if conf == nil {
		return nil, errors.New("Splunk configuration required")
	}
	if err := validate(conf); err != nil {
		return nil, err
	}
	url := urlfmt.FormatURL(conf.GetHttpEndpoint(), urlfmt.HTTPS, urlfmt.NoTrailingSlash)

	eventEndpoint := url
	var healthEndpoint string
	if baseURLPattern.MatchString(url) {
		eventEndpoint = url + splunkHECEventEndpoint
		healthEndpoint = url + splunkHECHealthEndpoint
	}

	tr := proxy.RoundTripperWithTLSConfig(&tls.Config{InsecureSkipVerify: conf.GetInsecure()})
	client := &http.Client{Transport: tr}

	return &splunk{
		client:         client,
		conf:           conf,
		eventEndpoint:  eventEndpoint,
		healthEndpoint: healthEndpoint,
		Notifier:       notifier,
	}, nil
}

func validate(conf *storage.Splunk) error {
	errorList := errorhelpers.NewErrorList("Splunk config validation")
	if len(conf.HttpToken) == 0 {
		errorList.AddString("Splunk HTTP Event Collector(HEC) token must be specified")
	}
	if len(conf.HttpEndpoint) == 0 {
		errorList.AddString("Splunk HTTP endpoint must be specified")
	}
	if conf.GetTruncate() == 0 {
		conf.Truncate = splunkHECDefaultDataLimit
	}
	for sourceTypeKey := range defaultSourceTypeMap {
		if _, ok := conf.SourceTypes[sourceTypeKey]; !ok {
			errorList.AddStringf("Source type key %s must be specified", sourceTypeKey)
		}
	}
	return errorList.ToError()
}

// UpgradeNotifierConfig applies changes to the current notifier to make it backwards compatible
func UpgradeNotifierConfig(notifier *storage.Notifier) {
	if notifier.GetType() == integrationType {
		if notifier.GetSplunk().GetDerivedSourceTypeDeprecated() != nil {
			splunk := notifier.GetSplunk()
			// Handle backwards compatibility for derived source type field
			if splunk.GetDerivedSourceType() {
				splunk.SourceTypes = defaultSourceTypeMap
			} else {
				splunk.SourceTypes = make(map[string]string)
				for k := range defaultSourceTypeMap {
					splunk.SourceTypes[k] = jsonSourceType
				}
			}
			splunk.DerivedSourceTypeDeprecated = nil
		}
	}
}
