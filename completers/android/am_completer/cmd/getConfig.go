package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "Retrieve the configuration of the device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getConfigCmd).Standalone()

	getConfigCmd.Flags().Int("days", 0, "also return last `N` days of configurations that have been seen")
	getConfigCmd.Flags().Bool("device", false, "also output global device configuration info")
	getConfigCmd.Flags().String("display", "", "specify for which `DISPLAY_ID` to run the command")
	getConfigCmd.Flags().Bool("proto", false, "return result as proto")

	rootCmd.AddCommand(getConfigCmd)

}
