// package os contains process related actions
package supervisor

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionSignals completes signals
//
//	SIGHUP (supervisord and all its subprocesses will shut down)
//	SIGINT (supervisord and all its subprocesses will shut down)
func ActionSignals() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionStyledValuesDescribed(
			"SIGHUP", "supervisord and all its subprocesses will shut down", styles.CarapaceBin.KillSignalTerm,
			"SIGINT", "supervisord and all its subprocesses will shut down", styles.CarapaceBin.KillSignalTerm,
			"SIGQUIT", "supervisord and all its subprocesses will shut down", styles.CarapaceBin.KillSignalCore,
			"SIGTERM", "supervisord will reload the configuration and restart all processes", styles.CarapaceBin.KillSignalTerm,
			"SIGUSR2", "supervisord will close and reopen the main activity log and all child log files", styles.CarapaceBin.KillSignalTerm,
		)
	}).Tag("signals")
}
