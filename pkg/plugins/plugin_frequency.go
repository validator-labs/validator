package plugins

import (
	"strconv"
	"time"

	"github.com/go-logr/logr"
	"github.com/validator-labs/validator/pkg/constants"
	ctrl "sigs.k8s.io/controller-runtime"
)

// FrequencyFromAnnotations calculates reconciliation frequency from annotations of the plugin custom resource.
// Defaults to 120 seconds if annotation is not found.
func FrequencyFromAnnotations(l logr.Logger, annotations map[string]string) ctrl.Result {
	var frequency time.Duration
	if secondsString, ok := annotations[constants.ReconciliationFrequencyAnnotation]; ok {
		seconds, err := strconv.Atoi(secondsString)
		if err != nil {
			l.Error(err, "Failed to convert frequency annotation", "rescheduleSeconds", 120)
			frequency = time.Second * 120
		} else {
			l.Info("Reconciliation frequency annotation found", "rescheduleSeconds", seconds)
			frequency = time.Second * time.Duration(seconds)
		}
	} else {
		l.Info("Frequency annotation not found: defaulting to 120 seconds")
		frequency = time.Second * 120
	}

	return ctrl.Result{RequeueAfter: frequency}
}
