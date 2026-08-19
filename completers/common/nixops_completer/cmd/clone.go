package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var CloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(CloneCmd).Standalone()
	rootCmd.AddCommand(CloneCmd)
}
