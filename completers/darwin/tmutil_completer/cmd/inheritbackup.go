package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(inheritbackupCmd)
}

var inheritbackupCmd = &cobra.Command{
	Use:   "inheritbackup",
	Short: "claim a machine directory or sparsebundle",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inheritbackupCmd).Standalone()
}