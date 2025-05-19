package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeInterfaceCmd = &cobra.Command{
	Use:   "make:interface [-f|--force] [--] <name>",
	Short: "Create a new interface",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeInterfaceCmd).Standalone()

	makeInterfaceCmd.Flags().Bool("force", false, "Create the interface even if the interface already exists")
	rootCmd.AddCommand(makeInterfaceCmd)
}
