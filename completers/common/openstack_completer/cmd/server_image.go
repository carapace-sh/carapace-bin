package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_imageCmd).Standalone()

	serverCmd.AddCommand(server_imageCmd)
}
