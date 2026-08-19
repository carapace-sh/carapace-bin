package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ModifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ModifyCmd).Standalone()
	rootCmd.AddCommand(ModifyCmd)
}
