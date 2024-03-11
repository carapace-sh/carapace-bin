package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dagger"
	"github.com/spf13/cobra"
)

var developCmd = &cobra.Command{
	Use:     "develop",
	Short:   "Setup or update all the resources needed to develop on a module locally",
	GroupID: "module",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(developCmd).Standalone()

	developCmd.PersistentFlags().Bool("focus", false, "Only show output for focused commands")
	developCmd.PersistentFlags().StringP("mod", "m", "", "Path to dagger.json config file for the module or a directory containing that file. Either local path (e.g. \"/path/to/some/dir\") or a github repo (e.g. \"github.com/dagger/dagger/path/to/some/subdir\")")
	developCmd.Flags().String("sdk", "", "New SDK for the module")
	developCmd.Flags().String("source", "", "Directory to store the module implementation source code in")
	rootCmd.AddCommand(developCmd)

	carapace.Gen(developCmd).FlagCompletion(carapace.ActionMap{
		"mod":    dagger.ActionMods(),
		"sdk":    dagger.ActionSdks(),
		"source": carapace.ActionDirectories(),
	})
}
