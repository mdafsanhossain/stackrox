package processor

import (
	"context"
	"fmt"

	timestamp "github.com/gogo/protobuf/types"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/integrationhealth"
	"github.com/stackrox/rox/pkg/notifiers"
	"github.com/stackrox/rox/pkg/set"
)

var (
	// Replacing with a background context such that outside context cancellation
	// does not affect long running go routines.
	ctxBackground = context.Background()
)

// Processor takes in alerts and sends the notifications tied to that alert
type processorImpl struct {
	ns       NotifierSet
	reporter integrationhealth.Reporter
}

func (p *processorImpl) HasNotifiers() bool {
	return p.ns.HasNotifiers()
}

func (p *processorImpl) HasEnabledAuditNotifiers() bool {
	return p.ns.HasEnabledAuditNotifiers()
}

// RemoveNotifier removes the in memory copy of the specified notifier
func (p *processorImpl) RemoveNotifier(ctx context.Context, id string) {
	p.ns.RemoveNotifier(ctx, id)
}

// GetNotifier gets the in memory copy of the specified notifier id
func (p *processorImpl) GetNotifier(ctx context.Context, id string) (notifier notifiers.Notifier) {
	return p.ns.GetNotifier(ctx, id)
}

// GetNotifiers gets the in memory copies of all notifiers
func (p *processorImpl) GetNotifiers(ctx context.Context) (notifiers []notifiers.Notifier) {
	return p.ns.GetNotifiers(ctx)
}

// UpdateNotifier updates or adds the passed notifier into memory
func (p *processorImpl) UpdateNotifier(ctx context.Context, notifier notifiers.Notifier) {
	p.ns.UpsertNotifier(ctx, notifier)
}

// IsSecuredClusterNotifier returns true if this is a notifier that can be accessed by the secured cluster
func (p *processorImpl) IsSecuredClusterNotifier(notifier notifiers.Notifier) bool {
	if !env.SecuredClusterNotifiers.BooleanSetting() {
		return false
	}
	if _, ok := notifier.ProtoNotifier().Config.(*storage.Notifier_Jira); ok {
		return true
	}
	if _, ok := notifier.ProtoNotifier().Config.(*storage.Notifier_Generic); ok {
		return true
	}
	if _, ok := notifier.ProtoNotifier().Config.(*storage.Notifier_Syslog); ok {
		return true
	}
	return false
}

// ProcessAlert pushes the alert into a channel to be processed
func (p *processorImpl) ProcessAlert(ctx context.Context, alert *storage.Alert) {
	if len(alert.GetPolicy().GetNotifiers()) == 0 {
		return
	}
	alertNotifiers := set.NewStringSet(alert.GetPolicy().GetNotifiers()...)

	p.ns.ForEach(ctx, func(ctx context.Context, notifier notifiers.Notifier, failures AlertSet) {
		if alertNotifiers.Contains(notifier.ProtoNotifier().GetId()) {
			// If this is a secured cluster notifier the notification for this alert has already been processed in the secured
			// cluster before the alert reached here for processing. Hence, skip the notifier and continue with the rest
			// of the notifiers configured for the policy that generated the alert
			if p.IsSecuredClusterNotifier(notifier) {
				return
			}
			go func() {
				err := tryToAlert(ctx, notifier, alert)
				if err != nil {
					p.UpdateNotifierHealthStatus(notifier, storage.IntegrationHealth_UNHEALTHY, err.Error())
					failures.Add(alert)
				} else {
					p.UpdateNotifierHealthStatus(notifier, storage.IntegrationHealth_HEALTHY, "")
				}
			}()
		}
	})
}

// ProcessAuditMessage sends the audit message with all applicable notifiers.
func (p *processorImpl) ProcessAuditMessage(ctx context.Context, msg *v1.Audit_Message) {
	// TODO: Turn processorImpl into a work queue and introduce func (p *processorImpl) run(context.Context) error.
	// With that, we wouldn't have to fan out n go routines (n = # notifiers in p.ns) and ensure ordering
	// of audit messages.
	p.ns.ForEach(ctx, func(_ context.Context, notifier notifiers.Notifier, _ AlertSet) {
		go p.tryToSendAudit(ctxBackground, notifier, msg)
	})
}

func (p *processorImpl) UpdateNotifierHealthStatus(notifier notifiers.Notifier, healthStatus storage.IntegrationHealth_Status, errMessage string) {
	p.reporter.UpdateIntegrationHealthAsync(&storage.IntegrationHealth{
		Id:            notifier.ProtoNotifier().Id,
		Name:          notifier.ProtoNotifier().Id,
		Type:          storage.IntegrationHealth_NOTIFIER,
		Status:        healthStatus,
		LastTimestamp: timestamp.TimestampNow(),
		ErrorMessage:  errMessage,
	})
}

func (p *processorImpl) tryToSendAudit(ctx context.Context, notifier notifiers.Notifier, msg *v1.Audit_Message) {
	auditNotifier, ok := notifier.(notifiers.AuditNotifier)
	if ok {
		if err := auditNotifier.SendAuditMessage(ctx, msg); err != nil {
			protoNotifier := notifier.ProtoNotifier()
			log.Errorf("Unable to send audit msg to %s (%s): %v", protoNotifier.GetName(), protoNotifier.GetType(), err)
			p.UpdateNotifierHealthStatus(notifier, storage.IntegrationHealth_UNHEALTHY, fmt.Sprintf("Unable to send audit msg: %v", err))
		}
		p.UpdateNotifierHealthStatus(notifier, storage.IntegrationHealth_HEALTHY, "")
	}
}

// Used for testing.
func (p *processorImpl) processAlertSync(ctx context.Context, alert *storage.Alert) {
	alertNotifiers := set.NewStringSet(alert.GetPolicy().GetNotifiers()...)
	p.ns.ForEach(ctx, func(ctx context.Context, notifier notifiers.Notifier, failures AlertSet) {
		if alertNotifiers.Contains(notifier.ProtoNotifier().GetId()) {
			if p.IsSecuredClusterNotifier(notifier) {
				return
			}
			err := tryToAlert(ctx, notifier, alert)
			if err != nil {
				failures.Add(alert)
			}
		}
	})
}
