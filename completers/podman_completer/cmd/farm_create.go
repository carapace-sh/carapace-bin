package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var farm_createCmd = &cobra.Command{
	Use:   "create NAME [CONNECTIONS...]",
	Short: "Create a new farm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(farm_createCmd).Standalone()

	farmCmd.AddCommand(farm_createCmd)
}
