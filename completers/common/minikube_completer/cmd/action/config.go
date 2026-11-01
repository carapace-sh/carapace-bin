package action

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

func ActionConfigNames() carapace.Action {
	return carapace.ActionValues(
		"ReminderWaitPeriodInHours",
		"ShowBootstrapperDeprecationNotification",
		"ShowDriverDeprecationNotification",
		"WantBetaUpdateNotification",
		"WantKubectlDownloadMsg",
		"WantNoneDriverWarning",
		"WantReportError",
		"WantReportErrorPrompt",
		"WantUpdateNotification",
		"bootstrapper",
		"cache",
		"container-runtime",
		"cpus",
		"disable-driver-mounts",
		"disk-size",
		"driver",
		"embed-certs",
		"feature-gates",
		"host-only-cidr",
		"hyperv-virtual-switch",
		"insecure-registry",
		"iso-url",
		"kubernetes-version",
		"log_dir",
		"memory",
		"profile",
		"v",
		"vm-driver",
	)
}

func ActionConfigDefaults(config string) carapace.Action {
	return carapace.ActionExecCommand("minikube", "config", "defaults", config, "--output", "json")(func(output []byte) carapace.Action {
		var vals []string
		if err := json.Unmarshal(output, &vals); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionValues(vals...)
	})
}
