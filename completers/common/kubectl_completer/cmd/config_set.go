package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set PROPERTY_NAME PROPERTY_VALUE",
	Short: "Set an individual value in a kubeconfig file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().String("set-raw-bytes", "", "When writing a []byte PROPERTY_VALUE, write the given string directly without base64 decoding.")
	config_setCmd.Flag("set-raw-bytes").NoOptDefVal = " "
	configCmd.AddCommand(config_setCmd)
}
