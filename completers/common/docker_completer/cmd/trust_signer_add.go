package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var trust_signer_addCmd = &cobra.Command{
	Use:   "add OPTIONS NAME REPOSITORY [REPOSITORY...] ",
	Short: "Add a signer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_signer_addCmd).Standalone()

	trust_signer_addCmd.Flags().String("key", "", "Path to the signer's public key file")
	trust_signerCmd.AddCommand(trust_signer_addCmd)

	carapace.Gen(trust_signer_addCmd).FlagCompletion(carapace.ActionMap{
		"key": carapace.ActionFiles(),
	})

	carapace.Gen(trust_signer_addCmd).PositionalAnyCompletion(docker.ActionRepositories())
}
