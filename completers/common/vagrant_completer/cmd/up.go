package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "starts and provisions the vagrant environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	upCmd.Flags().Bool("destroy-on-error", false, "Destroy machine if any fatal error happens")
	upCmd.Flags().Bool("install-provider", false, "If possible, install the provider if it isn't installed")
	upCmd.Flags().Bool("no-destroy-on-error", false, "Do not destroy machine if any fatal error happens")
	upCmd.Flags().Bool("no-install-provider", false, "Do not install the provider if it isn't installed")
	upCmd.Flags().Bool("no-parallel", false, "Disable parallelism if provider supports it")
	upCmd.Flags().Bool("no-provision", false, "Disable provisioning")
	upCmd.Flags().Bool("parallel", false, "Enable parallelism if provider supports it")
	upCmd.Flags().String("provider", "", "Back the machine with a specific provider")
	upCmd.Flags().Bool("provision", false, "Enable provisioning")
	upCmd.Flags().String("provision-with", "", "Enable only certain provisioners, by type or by name.")
	rootCmd.AddCommand(upCmd)

	carapace.Gen(upCmd).FlagCompletion(carapace.ActionMap{
		"provider":       vagrant.ActionProviders(),
		"provision-with": vagrant.ActionProvisioners().UniqueList(","),
	})

	carapace.Gen(upCmd).PositionalCompletion(
		vagrant.ActionMachines(),
	)
}
