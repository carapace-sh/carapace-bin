package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dagger"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:     "publish",
	Short:   "Publish a Dagger module to the Daggerverse",
	GroupID: "module",
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	publishCmd.Flags().Bool("focus", false, "Only show output for focused commands")
	publishCmd.Flags().BoolP("force", "f", false, "Force publish even if the git repository is not clean")
	publishCmd.Flags().StringP("mod", "m", "", "Path to dagger.json config file for the module or a directory containing that file. Either local path (e.g. \"/path/to/some/dir\") or a github repo (e.g. \"github.com/dagger/dagger/path/to/some/subdir\")")
	rootCmd.AddCommand(publishCmd)

	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"mod": dagger.ActionMods(),
	})
}
