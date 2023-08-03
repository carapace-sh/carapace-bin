package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var importEnvironmentCmd = &cobra.Command{
	Use:     "import-environment",
	Short:   "Import all or some environment variables",
	GroupID: "environment",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importEnvironmentCmd).Standalone()

	rootCmd.AddCommand(importEnvironmentCmd)

	carapace.Gen(importEnvironmentCmd).PositionalAnyCompletion(
		os.ActionEnvironmentVariables().FilterArgs(),
	)
}
