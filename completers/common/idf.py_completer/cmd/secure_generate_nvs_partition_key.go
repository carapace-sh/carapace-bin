package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secureGenerateNvsPartitionKeyCmd = &cobra.Command{
	Use:   "secure-generate-nvs-partition-key",
	Short: "Generate NVS partition key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureGenerateNvsPartitionKeyCmd).Standalone()
	secureGenerateNvsPartitionKeyCmd.Flags().StringP("encryption-scheme", "s", "", "Encryption scheme to use")
	secureGenerateNvsPartitionKeyCmd.Flags().StringP("hmac-keyfile", "l", "", "File to store the generated HMAC key")
	secureGenerateNvsPartitionKeyCmd.Flags().StringP("keyfile", "k", "", "File to store the generated key")
	carapace.Gen(secureGenerateNvsPartitionKeyCmd).FlagCompletion(carapace.ActionMap{
		"encryption-scheme": carapace.ActionValues("Flash", "HMAC"),
		"hmac-keyfile":      carapace.ActionFiles(),
		"keyfile":           carapace.ActionFiles(),
	})
	rootCmd.AddCommand(secureGenerateNvsPartitionKeyCmd)
}
