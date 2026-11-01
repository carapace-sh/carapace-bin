package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "resume BitLocker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resumeCmd).Standalone()
	rootCmd.AddCommand(resumeCmd)
}
