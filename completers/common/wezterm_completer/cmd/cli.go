package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli [OPTIONS] <COMMAND>",
	Short: "Interact with experimental mux server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cliCmd).Standalone()

	cliCmd.Flags().String("class", "", "Class of the gui instance")
	cliCmd.Flags().BoolP("help", "h", false, "Print help")
	cliCmd.Flags().Bool("no-auto-start", false, "Don't automatically start the server")
	cliCmd.Flags().Bool("prefer-mux", false, "Prefer connecting to a background mux server")
	rootCmd.AddCommand(cliCmd)
}
