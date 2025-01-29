package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_createCmd = &cobra.Command{
	Use:   "create [options] [NAME]",
	Short: "Create networks for containers and pods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_createCmd).Standalone()

	network_createCmd.Flags().Bool("disable-dns", false, "disable dns plugin")
	network_createCmd.Flags().StringSlice("dns", []string{}, "DNS servers this network will use")
	network_createCmd.Flags().StringP("driver", "d", "", "driver to manage the network")
	network_createCmd.Flags().StringSlice("gateway", []string{}, "IPv4 or IPv6 gateway for the subnet")
	network_createCmd.Flags().Bool("ignore", false, "Don't fail if network already exists")
	network_createCmd.Flags().String("interface-name", "", "interface name which is used by the driver")
	network_createCmd.Flags().Bool("internal", false, "restrict external access from this network")
	network_createCmd.Flags().StringSlice("ip-range", []string{}, "allocate container IP from range")
	network_createCmd.Flags().String("ipam-driver", "", "IP Address Management Driver")
	network_createCmd.Flags().Bool("ipv6", false, "enable IPv6 networking")
	network_createCmd.Flags().StringSlice("label", []string{}, "set metadata on a network")
	network_createCmd.Flags().String("macvlan", "", "create a Macvlan connection based on this device")
	network_createCmd.Flags().StringSliceP("opt", "o", []string{}, "Set driver specific options (default [])")
	network_createCmd.Flags().StringSlice("route", []string{}, "static routes")
	network_createCmd.Flags().StringSlice("subnet", []string{}, "subnets in CIDR format")
	network_createCmd.Flag("macvlan").Hidden = true
	networkCmd.AddCommand(network_createCmd)
}
