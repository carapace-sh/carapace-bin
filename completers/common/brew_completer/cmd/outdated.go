package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:     "outdated",
	Short:   "List installed casks and formulae that have an updated version available",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()

	outdatedCmd.Flags().Bool("cask", false, "List only outdated casks.")
	outdatedCmd.Flags().Bool("debug", false, "Display any debugging information.")
	outdatedCmd.Flags().Bool("fetch-HEAD", false, "Fetch the upstream repository to detect if the HEAD installation of the formula is outdated. Otherwise, the repository's HEAD will only be checked for updates when a new stable or development version has been released.")
	outdatedCmd.Flags().Bool("formula", false, "List only outdated formulae.")
	outdatedCmd.Flags().Bool("greedy", false, "Also include outdated casks with `auto_updates true` or `version :latest`.")
	outdatedCmd.Flags().Bool("greedy-auto-updates", false, "Also include outdated casks including those with `auto_updates true`.")
	outdatedCmd.Flags().Bool("greedy-latest", false, "Also include outdated casks including those with `version :latest`.")
	outdatedCmd.Flags().Bool("help", false, "Show this message.")
	outdatedCmd.Flags().Bool("json", false, "Print output in JSON format. There are two versions: `v1` and `v2`. `v1` is deprecated and is currently the default if no version is specified. `v2` prints outdated formulae and casks.")
	outdatedCmd.Flags().Bool("quiet", false, "List only the names of outdated kegs (takes precedence over `--verbose`).")
	outdatedCmd.Flags().Bool("verbose", false, "Include detailed version information.")
	rootCmd.AddCommand(outdatedCmd)
}
