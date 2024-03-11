package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func init() {
	knownVariables["git"] = func() variables {
		_bool := carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
		return variables{
			Condition: conditions.ConditionPath("git"),
			Variables: map[string]string{
				"GIT_ALTERNATE_OBJECT_DIRECTORIES": "is a colon-separated list which tells Git where to check for objects",
				"GIT_ASKPASS":                      "is an override for the core.askpass configuration value",
				"GIT_AUTHOR_DATE":                  "is the timestamp used for the “author” field",
				"GIT_AUTHOR_EMAIL":                 "is the email for the “author” field",
				"GIT_AUTHOR_NAME":                  "is the human-readable name in the “author” field",
				"GIT_CEILING_DIRECTORIES":          "controls the behavior of searching for a .git directory",
				"GIT_COMMITTER_DATE":               "is used for the timestamp in the “committer” field",
				"GIT_COMMITTER_EMAIL":              "is the email address for the “committer” field",
				"GIT_COMMITTER_NAME":               "sets the human name for the “committer” field",
				"GIT_CONFIG_NOSYSTEM":              "if set, disables the use of the system-wide configuration file",
				"GIT_DIFF_OPTS":                    "controls the number of context lines shown in a git diff command",
				"GIT_DIFF_PATH_COUNTER":            "represents which file in a series is being diffed (starting with 1)",
				"GIT_DIFF_PATH_TOTAL":              "total number of files in the batch",
				"GIT_DIR":                          "is the location of the .git folder",
				"GIT_EDITOR":                       "is the editor Git will launch when the user needs to edit some text",
				"GIT_EXEC_PATH":                    "determines where Git looks for its sub-programs",
				"GIT_EXTERNAL_DIFF":                "is used as an override for the diff.external configuration value",
				"GIT_FLUSH":                        "can be used to force Git to use non-buffered I/O when writing incrementally to stdout",
				"GIT_GLOB_PATHSPECS":               "controls the default behavior of wildcards in pathspecs",
				"GIT_HTTP_USER_AGENT":              "sets the user-agent string used by Git when communicating over HTTP",
				"GIT_ICASE_PATHSPECS":              "sets all pathspecs to work in a case-insensitive manner",
				"GIT_INDEX_FILE":                   "is the path to the index file (non-bare repositories only)",
				"GIT_LITERAL_PATHSPECS":            "disables both of the above behaviors",
				"GIT_MERGE_VERBOSITY":              "controls the output for the recursive merge strategy",
				"GIT_NAMESPACE":                    "controls access to namespaced refs",
				"GIT_NOGLOB_PATHSPECS":             "controls the default behavior of wildcards in pathspecs",
				"GIT_OBJECT_DIRECTORY":             "can be used to specify the location of the directory that usually resides at .git/objects",
				"GIT_PAGER":                        "controls the program used to display multi-page output on the command line",
				"GIT_REFLOG_ACTION":                "lets you specify the descriptive text written to the reflog",
				"GIT_SSH":                          "a program that is invoked instead of ssh when Git tries to connect to an SSH host",
				"GIT_SSH_COMMAND":                  "sets the SSH command used when Git tries to connect to an SSH host",
				"GIT_SSL_NO_VERIFY":                "tells Git not to verify SSL certificates",
				"GIT_TRACE":                        "controls general traces, which don’t fit into any specific category",
				"GIT_TRACE_PACK_ACCESS":            "controls tracing of packfile access",
				"GIT_TRACE_PACKET":                 "enables packet-level tracing for network operations",
				"GIT_TRACE_PERFORMANCE":            "controls logging of performance data",
				"GIT_TRACE_SETUP":                  "shows information about what Git is discovering about the repository and environment it’s interacting with",
				"GIT_WORK_TREE":                    "is the location of the root of the working directory for a non-bare repository",
			},
			VariableCompletion: map[string]carapace.Action{
				// TODO complete more variables
				"GIT_ALTERNATE_OBJECT_DIRECTORIES": carapace.ActionDirectories().MultiParts(":"),
				"GIT_ASKPASS":                      git.ActionConfigValues("core.askpass"),
				"GIT_CEILING_DIRECTORIES":          carapace.ActionDirectories().MultiParts("").NoSpace(),
				"GIT_COMMITTER_DATE":               time.ActionDate(),
				"GIT_CONFIG_NOSYSTEM":              _bool,
				"GIT_DIR":                          carapace.ActionDirectories(),
				"GIT_EXTERNAL_DIFF":                git.ActionConfigValues("diff.external"),
				"GIT_DIFF_OPTS": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					cmd := &cobra.Command{}
					carapace.Gen(cmd).Standalone()
					cmd.Flags().StringP("unified", "U", "", "Generate diffs with <n> lines of context instead of the usual three")
					return carapace.ActionExecute(cmd)
				}).Split(),
				"GIT_EDITOR":    git.ActionConfigValues("core.editor"),
				"GIT_EXEC_PATH": carapace.ActionDirectories(),
				"GIT_FLUSH": carapace.ActionValuesDescribed(
					"0", "buffer all output",
					"1", "flush more often",
				),
				"GIT_HTTP_USER_AGENT":  http.ActionUserAgents(),
				"GIT_INDEX_FILE":       carapace.ActionFiles(),
				"GIT_OBJECT_DIRECTORY": carapace.ActionDirectories(),
				"GIT_PAGER":            git.ActionConfigValues("core.pager"),
				"GIT_WORK_TREE":        carapace.ActionDirectories(),
			},
		}
	}
}
