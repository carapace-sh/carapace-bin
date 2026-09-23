package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getInstallLocationCmd = &cobra.Command{
	Use:   "get-install-location",
	Short: "Return the current install location",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getInstallLocationCmd).Standalone()

	rootCmd.AddCommand(getInstallLocationCmd)

}
