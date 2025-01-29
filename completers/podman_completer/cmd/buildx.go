package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildxCmd = &cobra.Command{
	Use:     "buildx",
	Short:   "Build images",
	Aliases: []string{"builder"},
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildxCmd).Standalone()

	rootCmd.AddCommand(buildxCmd)
}
