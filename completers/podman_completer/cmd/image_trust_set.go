package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_trust_setCmd = &cobra.Command{
	Use:   "set [options] REGISTRY",
	Short: "Set default trust policy or a new trust policy for a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_trust_setCmd).Standalone()

	image_trust_setCmd.Flags().String("policypath", "", "")
	image_trust_setCmd.Flags().StringSliceP("pubkeysfile", "f", []string{}, "Path of installed public key(s) to trust for TARGET.")
	image_trust_setCmd.Flags().StringP("type", "t", "", "Trust type, accept values: signedBy(default), accept, reject")
	image_trust_setCmd.Flag("policypath").Hidden = true
	image_trustCmd.AddCommand(image_trust_setCmd)
}
