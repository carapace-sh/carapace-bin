package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mvn_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mvn"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mvn",
	Short: "Apache Maven is a software project management and comprehension tool",
	Long:  "https://maven.apache.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("activate-profiles", "P", "", "Comma-delimited list of profiles to activate")
	rootCmd.Flags().BoolP("also-make", "am", false, "also build projects required by the list")
	rootCmd.Flags().BoolP("also-make-dependents", "amd", false, "also build projects that depend on projects on the list")
	rootCmd.Flags().BoolP("batch-mode", "B", false, "Run in non-interactive (batch) mode")
	rootCmd.Flags().StringP("builder", "b", "", "The id of the build strategy to use")
	rootCmd.Flags().String("color", "", "Defines the color mode of the output")
	rootCmd.Flags().BoolP("debug", "X", false, "Produce execution debug output")
	rootCmd.Flags().StringP("define", "D", "", "Define a system property")
	rootCmd.Flags().StringP("encrypt-master-password", "emp", "", "Encrypt master security password")
	rootCmd.Flags().StringP("encrypt-password", "ep", "", "Encrypt server password")
	rootCmd.Flags().BoolP("errors", "e", false, "Produce execution error messages")
	rootCmd.Flags().BoolP("fail-at-end", "fae", false, "Only fail the build afterwards")
	rootCmd.Flags().BoolP("fail-fast", "ff", false, "Stop at first failure in reactorized builds")
	rootCmd.Flags().BoolP("fail-never", "fn", false, "NEVER fail the build, regardless of project result")
	rootCmd.Flags().StringP("file", "f", "", "Force the use of an alternate POM file")
	rootCmd.Flags().StringP("global-settings", "gs", "", "Alternate path for the global settings file")
	rootCmd.Flags().StringP("global-toolchains", "gt", "", "Alternate path for the global toolchains file")
	rootCmd.Flags().BoolP("help", "h", false, "Display help information")
	rootCmd.Flags().BoolP("lax-checksums", "c", false, "Warn if checksums don't match")
	rootCmd.Flags().BoolP("legacy-local-repository", "llr", false, "Use Maven 2 Legacy Local Repository behaviour")
	rootCmd.Flags().StringP("log-file", "l", "", "Log file where all build output will go")
	rootCmd.Flags().BoolP("no-snapshot-updates", "nsu", false, "Suppress SNAPSHOT updates")
	rootCmd.Flags().BoolP("no-transfer-progress", "ntp", false, "Do not display transfer progress when downloading or uploading")
	rootCmd.Flags().BoolP("non-recursive", "N", false, "Do not recurse into sub-projects")
	rootCmd.Flags().BoolP("offline", "o", false, "Work offline")
	rootCmd.Flags().StringP("projects", "pl", "", "Comma-delimited list of specified reactor projects to build instead of all projects")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet output - only show errors")
	rootCmd.Flags().StringP("resume-from", "rf", "", "Resume reactor from specified project")
	rootCmd.Flags().StringP("settings", "s", "", "Alternate path for the user settings file")
	rootCmd.Flags().BoolP("show-version", "V", false, "Display version information WITHOUT stopping build")
	rootCmd.Flags().BoolP("strict-checksums", "C", false, "Fail the build if checksums don't match")
	rootCmd.Flags().StringP("threads", "T", "", "Thread count")
	rootCmd.Flags().StringP("toolchains", "t", "", "Alternate path for the user toolchains file")
	rootCmd.Flags().BoolP("update-snapshots", "U", false, "Forces a check for missing releases and updated snapshots on remote repositories")
	rootCmd.Flags().BoolP("version", "v", false, "Display version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"activate-profiles": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return mvn.ActionProfiles(rootCmd.Flag("file").Value.String()).UniqueList(",")
		}),
		"color":             carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"file":              carapace.ActionFiles(".xml"),
		"global-settings":   carapace.ActionFiles(".xml"),
		"global-toolchains": carapace.ActionFiles(".xml"),
		"log-file":          carapace.ActionFiles(),
		"projects":          action.ActionProjects(rootCmd).UniqueList(","),
		"resume-from":       action.ActionProjects(rootCmd),
		"settings":          carapace.ActionFiles(".xml"),
		"toolchains":        carapace.ActionFiles(".xml"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return mvn.ActionGoalsAndPhases(rootCmd.Flag("file").Value.String()).Invoke(c).Filter(c.Args...).ToA().MultiParts(":") // TODO use FilterArgs
		}),
	)
}
