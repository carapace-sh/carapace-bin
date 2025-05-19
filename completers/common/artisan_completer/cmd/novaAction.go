package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaActionCmd = &cobra.Command{
	Use:   "nova:action [--destructive] [--queued] [--] <name>",
	Short: "Create a new action class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaActionCmd).Standalone()

	novaActionCmd.Flags().Bool("destructive", false, "Indicate that the action deletes / destroys resources")
	novaActionCmd.Flags().Bool("queued", false, "Indicates the action should be queued")
	rootCmd.AddCommand(novaActionCmd)
}
