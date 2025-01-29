package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_lsCmd = &cobra.Command{
	Use:     "ls [options]",
	Short:   "List OCI artifacts",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_lsCmd).Standalone()

	artifact_lsCmd.Flags().String("format", "", "Format volume output using JSON or a Go template")
	artifactCmd.AddCommand(artifact_lsCmd)
}
