package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ip",
	Short: "show / manipulate routing, network devices, interfaces and tunnels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("6", "6", false, "shortcut for -family inet6")
	rootCmd.Flags().BoolS("0", "0", false, "shortcut for -family link")
	rootCmd.Flags().BoolS("M", "M", false, "shortcut for -family mpls")
	rootCmd.Flags().BoolS("B", "B", false, "shortcut for -family bridge")
	rootCmd.Flags().BoolS("4", "4", false, "shortcut for -family inet")
	rootCmd.Flags().BoolP("Numeric", "N", false, "Print the number of protocol, scope, dsfield, etc directly")
	rootCmd.Flags().BoolP("Version", "V", false, "Print the version of the ip utility and exit")
	rootCmd.Flags().BoolP("all", "a", false, "executes specified command over all objects")
	rootCmd.Flags().StringP("batch", "b", "", "Read commands from provided file or standard input and invoke them")
	rootCmd.Flags().Bool("br", false, "Print only basic information in a tabular format for better readability")
	rootCmd.Flags().Bool("brief", false, "Print only basic information in a tabular format for better readability")
	rootCmd.Flags().StringP("color", "c", "", "Configure color output")
	rootCmd.Flags().BoolP("details", "d", false, "Output more detailed information")
	rootCmd.Flags().StringP("family", "f", "", "Specifies the protocol family to use")
	rootCmd.Flags().Bool("force", false, "Don't terminate ip on errors in batch mode")
	rootCmd.Flags().BoolP("human", "h", false, "output statistics with human readable values followed by suffix")
	rootCmd.Flags().Bool("human-readable", false, "output statistics with human readable values followed by suffix")
	rootCmd.Flags().Bool("iec", false, "print human readable rates in IEC units (e.g. 1Ki = 1024)")
	rootCmd.Flags().BoolP("json", "j", false, "Output results in JavaScript Object Notation (JSON)")
	rootCmd.Flags().StringP("loops", "l", "", "Specify maximum number of loops")
	rootCmd.Flags().StringP("netns", "n", "", "switches ip to the specified network namespace NETNS")
	rootCmd.Flags().BoolP("oneline", "o", false, "output each record on a single line")
	rootCmd.Flags().BoolP("pretty", "p", false, "This flag adds indentation for readability")
	rootCmd.Flags().String("rc", "", "Set the netlink socket receive buffer size, defaults to 1MB")
	rootCmd.Flags().String("rcvbuf", "", "Set the netlink socket receive buffer size, defaults to 1MB")
	rootCmd.Flags().BoolP("resolve", "r", false, "use the system's name resolver to print DNS names instead of host addresses")
	rootCmd.Flags().Bool("statistics", false, "Output more information")
	rootCmd.Flags().BoolP("stats", "s", false, "Output more information")
	rootCmd.Flags().Bool("ts", false, "Like -timestamp, but use shorter format")
	rootCmd.Flags().Bool("tshort", false, "Like -timestamp, but use shorter format")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"batch":  carapace.ActionFiles(),
		"color":  carapace.ActionValues("always", "auto", "never"),
		"family": carapace.ActionValues("inet", "inet6", "bridge", "mpls", "link"),
	})
}
