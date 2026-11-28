package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hazmat_cdCmd = &cobra.Command{
	Use:   "cd [subdir] [cmd...]",
	Short: "cd to a subdir and run the command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hazmat_cdCmd).Standalone()

	carapace.Gen(hazmat_cdCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(hazmat_cdCmd).PositionalAnyCompletion(
		carapace.ActionExecutables(),
	)

	hazmatCmd.AddCommand(hazmat_cdCmd)
}
