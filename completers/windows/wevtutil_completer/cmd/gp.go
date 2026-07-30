package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gpCmd = &cobra.Command{
	Use:   "gp",
	Short: "get event publisher metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpCmd).Standalone()
	rootCmd.AddCommand(gpCmd)
}
