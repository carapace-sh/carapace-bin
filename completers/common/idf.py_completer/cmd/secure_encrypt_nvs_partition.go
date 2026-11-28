package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secureEncryptNvsPartitionCmd = &cobra.Command{
	Use:   "secure-encrypt-nvs-partition",
	Short: "Encrypt NVS partition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureEncryptNvsPartitionCmd).Standalone()

	secureEncryptNvsPartitionCmd.Flags().StringP("keyfile", "k", "", "File with NVS partition key")

	rootCmd.AddCommand(secureEncryptNvsPartitionCmd)

	carapace.Gen(secureEncryptNvsPartitionCmd).FlagCompletion(carapace.ActionMap{
		"keyfile": carapace.ActionFiles(),
	})

	carapace.Gen(secureEncryptNvsPartitionCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
