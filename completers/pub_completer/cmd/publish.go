package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish the current package to pub.dartlang.org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	publishCmd.Flags().BoolP("dry-run", "n", false, "Validate but do not publish the package.")
	publishCmd.Flags().BoolP("force", "f", false, "Publish without confirmation if there are no errors.")
	publishCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(publishCmd)
}
