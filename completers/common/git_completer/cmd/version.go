package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Display version information about Git",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().Bool("build-options", false, "include additional information about how git was built")
	rootCmd.AddCommand(versionCmd)
}
