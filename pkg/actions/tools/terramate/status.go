package terramate

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionCloudStatus completes cloud status
//
//	ok (The stack is not drifted and the last deployment succeeded)
//	failed (The last deployment of the stack failed so the status is unknown)
func ActionCloudStatus() carapace.Action {
	return carapace.ActionStyledValuesDescribed(
		"ok", "The stack is not drifted and the last deployment succeeded", style.Carapace.KeywordPositive,
		"failed", "The last deployment of the stack failed so the status is unknown", style.Carapace.KeywordNegative,
		"drifted", "The actual state is different from that defined in the code of the stack", style.Carapace.KeywordAmbiguous,
		"unhealthy", "This meta state matches any undesirable state (failed, drifted etc)", style.Carapace.KeywordNegative,
		"healthy", "This meta state matches stacks that have no undesireable state", style.Carapace.KeywordPositive,
	)
}
