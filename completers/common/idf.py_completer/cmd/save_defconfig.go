package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var saveDefconfigCmd = &cobra.Command{
	Use:   "save-defconfig",
	Short: "Generate a sdkconfig.defaults with minimal options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(saveDefconfigCmd).Standalone()

	saveDefconfigCmd.Flags().Bool("add-menu-labels", false, "Add menu labels to minimal config")
	rootCmd.AddCommand(saveDefconfigCmd)
}
