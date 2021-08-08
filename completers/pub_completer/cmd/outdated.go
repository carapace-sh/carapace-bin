package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Analyze your dependencies to find which ones can be upgraded",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()

	outdatedCmd.Flags().Bool("color", false, "Color the output.")
	outdatedCmd.Flags().Bool("dependency-overrides", false, "Show resolutions with `dependency_overrides`.")
	outdatedCmd.Flags().Bool("dev-dependencies", false, "Take dev dependencies into account.")
	outdatedCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	outdatedCmd.Flags().Bool("json", false, "Output the results using a json format.")
	outdatedCmd.Flags().String("mode", "", "Highlight versions with PROPERTY.")
	outdatedCmd.Flags().Bool("no-color", false, "Do not color the output.")
	outdatedCmd.Flags().Bool("no-dependency-overrides", false, "Do not show resolutions with `dependency_overrides`.")
	outdatedCmd.Flags().Bool("no-dev-dependencies", false, "Do not take dev dependencies into account.")
	outdatedCmd.Flags().Bool("no-prereleases", false, "Do not include prereleases in latest version.")
	outdatedCmd.Flags().Bool("no-show-all", false, "Do not include dependencies that are already fullfilling --mode.")
	outdatedCmd.Flags().Bool("no-transitive", false, "Do not show transitive dependencies.")
	outdatedCmd.Flags().Bool("prereleases", false, "Include prereleases in latest version.")
	outdatedCmd.Flags().Bool("show-all", false, "Include dependencies that are already fullfilling --mode.")
	outdatedCmd.Flags().Bool("transitive", false, "Show transitive dependencies.")
	rootCmd.AddCommand(outdatedCmd)

	carapace.Gen(outdatedCmd).FlagCompletion(carapace.ActionMap{
		"mode": carapace.ActionValues("outdated", "null-safety"),
	})
}
