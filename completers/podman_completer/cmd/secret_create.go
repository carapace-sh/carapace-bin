package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secret_createCmd = &cobra.Command{
	Use:   "create [options] NAME FILE|-",
	Short: "Create a new secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_createCmd).Standalone()

	secret_createCmd.Flags().StringP("driver", "d", "", "Specify secret driver")
	secret_createCmd.Flags().String("driver-opts", "", "Specify driver specific options")
	secret_createCmd.Flags().Bool("env", false, "Read secret data from environment variable")
	secret_createCmd.Flags().StringSliceP("label", "l", []string{}, "Specify labels on the secret")
	secret_createCmd.Flags().Bool("replace", false, "If a secret with the same name exists, replace it")
	secretCmd.AddCommand(secret_createCmd)
}
