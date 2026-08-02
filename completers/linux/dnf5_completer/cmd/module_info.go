package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleInfoCmd = &cobra.Command{
	Use:   "info [options] [<module-spec>...]",
	Short: "print details about module streams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleInfoCmd).Standalone()

	moduleInfoCmd.Flags().Bool("disabled", false, "Show disabled modules")
	moduleInfoCmd.Flags().Bool("enabled", false, "Show enabled modules")

	moduleCmd.AddCommand(moduleInfoCmd)
}
