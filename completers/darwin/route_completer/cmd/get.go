package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Look up and display the route for a destination",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()
	getCmd.Flags().Bool("host", false, "Destination is a host")
	getCmd.Flags().Bool("net", false, "Destination is a network")
	rootCmd.AddCommand(getCmd)
}
