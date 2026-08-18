package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Compress unhandled.log files in .git/svn and remove index files in .git/svn",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gcCmd).Standalone()
	rootCmd.AddCommand(gcCmd)
}
