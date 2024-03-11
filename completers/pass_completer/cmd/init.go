package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize new password storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().StringP("path", "p", "", "subfolder")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).PositionalAnyCompletion(
		os.ActionGpgKeyIds().FilterArgs(),
	)
}
