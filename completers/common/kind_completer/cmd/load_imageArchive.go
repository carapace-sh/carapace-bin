package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kind_completer/cmd/action"
	"github.com/spf13/cobra"
)

var load_imageArchiveCmd = &cobra.Command{
	Use:   "image-archive <IMAGE.tar>",
	Short: "Loads docker image from archive into nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(load_imageArchiveCmd).Standalone()

	load_imageArchiveCmd.Flags().StringP("name", "n", "kind", "the cluster context name")
	load_imageArchiveCmd.Flags().StringSlice("nodes", nil, "comma separated list of nodes to load images into")

	carapace.Gen(load_imageArchiveCmd).FlagCompletion(carapace.ActionMap{
		"name":  action.ActionClusterNames(),
		"nodes": action.ActionNodes(load_imageArchiveCmd),
	})

	carapace.Gen(load_imageArchiveCmd).PositionalCompletion(
		carapace.ActionFiles(".tar", ".tar.gz"),
	)

	loadCmd.AddCommand(load_imageArchiveCmd)
}
