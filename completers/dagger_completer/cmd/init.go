package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dagger"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init [flags] [PATH]",
	Short:   "Initialize a new Dagger module",
	GroupID: "module",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().String("license", "", "License identifier to generate - see https://spdx.org/licenses/")
	initCmd.Flags().String("name", "", "Name of the new module (defaults to parent directory name)")
	initCmd.Flags().String("sdk", "", "Optionally initialize module for development in the given SDK")
	initCmd.Flags().String("source", "", "Directory to store the module implementation source code in (defaults to \"dagger/ if \"--sdk\" is provided)")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"license": dagger.ActionLicenses(),
		"sdk":     dagger.ActionSdks(),
		"source":  carapace.ActionDirectories(),
	})

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
