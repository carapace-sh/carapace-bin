package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an individual value in a kubeconfig file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()
	config_setCmd.Flags().Bool("set-raw-bytes", false, "When writing a []byte PROPERTY_VALUE, write the given string directly without base64 decoding.")
	config_setCmd.Flag("set-raw-bytes").NoOptDefVal = "true"
	configCmd.AddCommand(config_setCmd)
}
