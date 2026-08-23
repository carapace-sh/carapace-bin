package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var calculatedriftCmd = &cobra.Command{
	Use:   "calculatedrift",
	Short: "analyze backups and determine change between them",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(calculatedriftCmd).Standalone()
	rootCmd.AddCommand(calculatedriftCmd)
}