package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump repository state information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_dumpCmd).Standalone()

	repository_dumpCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_dumpCmd.Flags().String("max-depth", "", "Optional max depth of tree dump")
	repository_dumpCmd.Flags().String("path", "", "Optional path in the repository to start dumping from")
	repository_dumpCmd.Flags().String("revision", "", "Optional revision to dump")
	repositoryCmd.AddCommand(repository_dumpCmd)
}
