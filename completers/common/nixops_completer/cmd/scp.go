package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ScpCmd = &cobra.Command{
	Use:   "scp",
	Short: "Scp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ScpCmd).Standalone()
	rootCmd.AddCommand(ScpCmd)
}
