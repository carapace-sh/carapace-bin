package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:     "index",
	Short:   "Create repository index file from packages",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "repository maintenance",
}

func init() {
	carapace.Gen(indexCmd).Standalone()

	indexCmd.Flags().StringP("description", "d", "", "Add a description to the index")
	indexCmd.Flags().StringP("index", "x", "", "Read an existing index from INDEX to speed up the creation")
	indexCmd.Flags().Bool("merge", false, "Merge PACKAGES into the existing INDEX")
	indexCmd.Flags().Bool("no-warnings", false, "Disable the warning about missing dependencies")
	indexCmd.Flags().StringP("output", "o", "", "Output generated index to FILE")
	indexCmd.Flags().Bool("prune-origin", false, "Prune packages from the existing INDEX with same origin")
	indexCmd.Flags().String("rewrite-arch", "", "Set all package's architecture to ARCH")
	common.AddGlobalFlags(indexCmd)
	rootCmd.AddCommand(indexCmd)

	carapace.Gen(indexCmd).FlagCompletion(carapace.ActionMap{
		"index":        carapace.ActionValues(), // TODO
		"output":       carapace.ActionFiles(),
		"rewrite-arch": carapace.ActionValues(), // TODO
	})

	carapace.Gen(indexCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
