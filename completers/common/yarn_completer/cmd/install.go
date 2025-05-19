package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install the project dependencies",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("check-cache", false, "Always refetch the packages and ensure that their checksums are consistent")
	installCmd.Flags().Bool("immutable", false, "Abort with an error exit code if the lockfile was to be modified")
	installCmd.Flags().Bool("immutable-cache", false, "Abort with an error exit code if the cache folder was to be modified")
	installCmd.Flags().Bool("inline-builds", false, "Verbosely print the output of the build steps of dependencies")
	installCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	installCmd.Flags().String("mode", "", "Change what artifacts installs generate")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"mode": carapace.ActionValuesDescribed(
			"skip-build", "do not run the build scripts at all",
			"update-lockfile", "skip the link step altogether",
		),
	})
}
