package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "resume a suspended vagrant machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resumeCmd).Standalone()

	resumeCmd.Flags().Bool("no-provision", false, "Disable provisioning")
	resumeCmd.Flags().Bool("provision", false, "Enable provisioning")
	resumeCmd.Flags().String("provision-with", "", "Enable only certain provisioners, by type or by name.")
	rootCmd.AddCommand(resumeCmd)

	carapace.Gen(resumeCmd).FlagCompletion(carapace.ActionMap{
		"provision-with": vagrant.ActionProvisioners().UniqueList(","),
	})

	carapace.Gen(resumeCmd).PositionalCompletion(
		vagrant.ActionMachines(),
	)
}
