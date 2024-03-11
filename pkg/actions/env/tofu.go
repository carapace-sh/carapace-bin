package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tofu"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
)

func init() {
	knownVariables["tofu"] = func() variables {
		_bool := carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
		return variables{
			Condition: conditions.ConditionPath("tofu"),
			Variables: map[string]string{
				"TF_LOG":                      "Enables detailed logs to appear on stderr which is useful for debugging",
				"TF_LOG_PATH":                 "This specifies where the log should persist its output to",
				"TF_INPUT":                    "If set to \"false\" or \"0\", causes tofu commands to behave as if the -input=false flag was specified",
				"TF_CLI_ARGS":                 "Additional cli arguments",
				"TF_CLI_ARGS_init":            "Additional cli arguments to init",
				"TF_CLI_ARGS_validate":        "Additional cli arguments to validate",
				"TF_CLI_ARGS_plan":            "Additional cli arguments to plan",
				"TF_CLI_ARGS_apply":           "Additional cli arguments to apply",
				"TF_CLI_ARGS_destroy":         "Additional cli arguments to destroy",
				"TF_CLI_ARGS_console":         "Additional cli arguments to console",
				"TF_CLI_ARGS_fmt":             "Additional cli arguments to fmt",
				"TF_CLI_ARGS_force-unlock":    "Additional cli arguments to force-unlock",
				"TF_CLI_ARGS_get":             "Additional cli arguments to get",
				"TF_CLI_ARGS_graph":           "Additional cli arguments to graph",
				"TF_CLI_ARGS_import":          "Additional cli arguments to import",
				"TF_CLI_ARGS_login":           "Additional cli arguments to login",
				"TF_CLI_ARGS_logout":          "Additional cli arguments to logout",
				"TF_CLI_ARGS_metadata":        "Additional cli arguments to metadata",
				"TF_CLI_ARGS_output":          "Additional cli arguments to output",
				"TF_CLI_ARGS_providers":       "Additional cli arguments to providers",
				"TF_CLI_ARGS_refresh":         "Additional cli arguments to refresh",
				"TF_CLI_ARGS_show":            "Additional cli arguments to show",
				"TF_CLI_ARGS_state":           "Additional cli arguments to state",
				"TF_CLI_ARGS_taint":           "Additional cli arguments to taint",
				"TF_CLI_ARGS_test":            "Additional cli arguments to test",
				"TF_CLI_ARGS_untaint":         "Additional cli arguments to untaint",
				"TF_CLI_ARGS_version":         "Additional cli arguments to version",
				"TF_CLI_ARGS_workspace":       "Additional cli arguments to workspace",
				"TF_DATA_DIR":                 "Changes the location where tofu keeps its per-working-directory data",
				"TF_WORKSPACE":                "Selects the workspace",
				"TF_IN_AUTOMATION":            "Adjusts its output to avoid suggesting specific commands to run next",
				"TF_REGISTRY_DISCOVERY_RETRY": "Max number of request retries the remote registry client will attempt",
				"TF_REGISTRY_CLIENT_TIMEOUT":  "Client timeout for requests ",
				"TF_CLI_CONFIG_FILE":          "Location of the tofu CLI configuration file",
				"TF_PLUGIN_CACHE_DIR":         "Plugin cache directory",
				"TF_IGNORE":                   "When set to \"trace\", tofu will output debug messages to display ignored files and folder",
			},
			VariableCompletion: map[string]carapace.Action{
				"TF_LOG":                   _bool,
				"TF_LOG_PATH":              carapace.ActionValues(),
				"TF_INPUT":                 carapace.ActionFiles(),
				"TF_CLI_ARGS":              bridge.ActionCarapaceBin("tofu").Split(),
				"TF_CLI_ARGS_init":         bridge.ActionCarapaceBin("tofu", "init").Split(),
				"TF_CLI_ARGS_validate":     bridge.ActionCarapaceBin("tofu", "validate").Split(),
				"TF_CLI_ARGS_plan":         bridge.ActionCarapaceBin("tofu", "plan").Split(),
				"TF_CLI_ARGS_apply":        bridge.ActionCarapaceBin("tofu", "apply").Split(),
				"TF_CLI_ARGS_destroy":      bridge.ActionCarapaceBin("tofu", "destroy").Split(),
				"TF_CLI_ARGS_console":      bridge.ActionCarapaceBin("tofu", "console").Split(),
				"TF_CLI_ARGS_fmt":          bridge.ActionCarapaceBin("tofu", "fmt").Split(),
				"TF_CLI_ARGS_force-unlock": bridge.ActionCarapaceBin("tofu", "force-unlock").Split(),
				"TF_CLI_ARGS_get":          bridge.ActionCarapaceBin("tofu", "get").Split(),
				"TF_CLI_ARGS_graph":        bridge.ActionCarapaceBin("tofu", "graph").Split(),
				"TF_CLI_ARGS_import":       bridge.ActionCarapaceBin("tofu", "import").Split(),
				"TF_CLI_ARGS_login":        bridge.ActionCarapaceBin("tofu", "login").Split(),
				"TF_CLI_ARGS_logout":       bridge.ActionCarapaceBin("tofu", "logout").Split(),
				"TF_CLI_ARGS_metadata":     bridge.ActionCarapaceBin("tofu", "metadata").Split(),
				"TF_CLI_ARGS_output":       bridge.ActionCarapaceBin("tofu", "output").Split(),
				"TF_CLI_ARGS_providers":    bridge.ActionCarapaceBin("tofu", "providers").Split(),
				"TF_CLI_ARGS_refresh":      bridge.ActionCarapaceBin("tofu", "refresh").Split(),
				"TF_CLI_ARGS_show":         bridge.ActionCarapaceBin("tofu", "show").Split(),
				"TF_CLI_ARGS_state":        bridge.ActionCarapaceBin("tofu", "state").Split(),
				"TF_CLI_ARGS_taint":        bridge.ActionCarapaceBin("tofu", "taint").Split(),
				"TF_CLI_ARGS_test":         bridge.ActionCarapaceBin("tofu", "test").Split(),
				"TF_CLI_ARGS_untaint":      bridge.ActionCarapaceBin("tofu", "untaint").Split(),
				"TF_CLI_ARGS_version":      bridge.ActionCarapaceBin("tofu", "version").Split(),
				"TF_CLI_ARGS_workspace":    bridge.ActionCarapaceBin("tofu", "workspace").Split(),
				"TF_DATA_DIR":              carapace.ActionDirectories(),
				"TF_WORKSPACE":             tofu.ActionWorkspaces(),
				"TF_IN_AUTOMATION":         _bool,
				"TF_CLI_CONFIG_FILE":       carapace.ActionFiles(),
				"TF_PLUGIN_CACHE_DIR":      carapace.ActionDirectories(),
				"TF_IGNORE":                carapace.ActionValues("trace").StyleF(style.ForLogLevel),
			},
		}
	}
}
