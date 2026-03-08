package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleInfoCmd = &cobra.Command{
	Use:   "info [module-spec]...",
	Short: "print details about module streams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleInfoCmd).Standalone()

	moduleInfoCmd.Flags().Bool("disabled", false, "show only disabled modules")
	moduleInfoCmd.Flags().Bool("enabled", false, "show only enabled modules")

	moduleCmd.AddCommand(moduleInfoCmd)
}
