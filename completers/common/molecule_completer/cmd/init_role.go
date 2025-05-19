package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var initRoleCmd = &cobra.Command{
	Use:   "role [flags] ROLE_NAME",
	Short: "Initialize a new role for use with Molecule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initRoleCmd)

	initRoleCmd.Flags().String("dependency-name", "", "Name of galaxy dependency to initialize")
	initRoleCmd.Flags().StringP("driver-name", "d", "delegated", "Name of the driver to use")
	initRoleCmd.Flags().Bool("help", false, "Show help message and exit")
	initRoleCmd.Flags().String("provisioner-name", "ansible", "Name of provisioner to initialize")
	initRoleCmd.Flags().String("verifier-name", "ansible", "Name of verifier to initialize")

	carapace.Gen(initRoleCmd).FlagCompletion(carapace.ActionMap{
		"driver-name":      molecule.ActionDrivers(),
		"provisioner-name": carapace.ActionExecutables(),
		"verifier-name":    carapace.ActionValues("ansible", "testinfra"),
	})

	initCmd.AddCommand(initRoleCmd)
}
