package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update client tools (tsh, tctl) to the latest version defined by the cluster configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("clear", false, "Removes locally installed client tools updates from the Teleport home directory.")
	updateCmd.Flags().Bool("no-clear", false, "Removes locally installed client tools updates from the Teleport home directory.")
	updateCmd.Flag("no-clear").Hidden = true
	rootCmd.AddCommand(updateCmd)
}
