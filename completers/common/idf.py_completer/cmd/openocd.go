package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var openocdCmd = &cobra.Command{
	Use:   "openocd",
	Short: "Run openocd from current path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openocdCmd).Standalone()

	openocdCmd.Flags().String("openocd-commands", "", "Command line arguments for openocd")
	openocdCmd.Flags().String("openocd-scripts", "", "Script directory for openocd cfg files")
	rootCmd.AddCommand(openocdCmd)

	carapace.Gen(openocdCmd).FlagCompletion(carapace.ActionMap{
		"openocd-scripts": carapace.ActionDirectories(),
	})
}
