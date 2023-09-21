package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/rsteube/carapace-bin/pkg/conditions"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
)

func init() {
	knownVariables["gh"] = variables{
		Condition: conditions.ConditionPath("gh"),
		Variables: map[string]string{
			"GH_TOKEN":              "an authentication token for github.com API requests",
			"GH_ENTERPRISE_TOKEN":   "an authentication token for API requests to GitHub Enterprise",
			"GH_HOST":               "specify the GitHub hostname",
			"GH_REPO":               "specify the GitHub repository",
			"GH_EDITOR":             "the editor tool to use for authoring text",
			"GH_BROWSER":            "the web browser to use for opening links",
			"GH_DEBUG":              "set to a truthy value to enable verbose output on standard error",
			"GH_PAGER":              "a terminal paging program to send standard output",
			"GH_FORCE_TTY":          "set to any value to force terminal-style output",
			"GH_NO_UPDATE_NOTIFIER": "set to any value to disable update notifications",
			"GH_CONFIG_DIR":         "the directory where gh will store configuration files",
			"GH_PROMPT_DISABLED":    "set to any value to disable interactive prompting in the terminal",
			"GH_PATH":               "set the path to the gh executable",
		},
		VariableCompletion: map[string]carapace.Action{
			"GH_REPO":               gh.ActionOwnerRepositories(gh.HostOpts{}),
			"GH_EDITOR":             bridge.ActionCarapaceBin().Split(),
			"GH_BROWSER":            bridge.ActionCarapaceBin().Split(),
			"GH_DEBUG":              carapace.ActionValues("1"),
			"GH_PAGER":              bridge.ActionCarapaceBin().Split(),
			"GH_FORCE_TTY":          carapace.ActionValues("1"),
			"GH_NO_UPDATE_NOTIFIER": carapace.ActionValues("1"),
			"GH_CONFIG_DIR":         carapace.ActionDirectories(),
			"GH_PATH":               carapace.ActionFiles(),
		},
	}
}
