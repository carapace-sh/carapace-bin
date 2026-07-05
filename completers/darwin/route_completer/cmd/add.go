package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a route",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()
	addCmd.Flags().Bool("blackhole", false, "Install a blackhole route")
	addCmd.Flags().Bool("cloning", false, "Enable cloning of new routes")
	addCmd.Flags().Bool("host", false, "Destination is a host")
	addCmd.Flags().String("ifscope", "", "Bind to the specified interface")
	addCmd.Flags().Bool("interface", false, "Destination is directly reachable via an interface")
	addCmd.Flags().Bool("net", false, "Destination is a network")
	addCmd.Flags().String("netmask", "", "Specify the network mask")
	addCmd.Flags().String("prefixlen", "", "Specify the prefix length (IPv6)")
	addCmd.Flags().Bool("reject", false, "Install a rejecting route")
	rootCmd.AddCommand(addCmd)
}
