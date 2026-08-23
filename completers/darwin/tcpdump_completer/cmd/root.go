package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tcpdump",
	Short: "dump traffic on a network",
	Long:  "https://keith.github.io/xcode-manpages/tcpdump.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("#", false, "Print packet number")
	rootCmd.Flags().Bool("A", false, "Print each packet in ASCII")
	rootCmd.Flags().String("B", "", "Set operating system capture buffer size")
	rootCmd.Flags().String("C", "", "Write packets to a file after file_size MB")
	rootCmd.Flags().String("F", "", "Use file as input for filter expression")
	rootCmd.Flags().String("G", "", "Rotate dump file every rotate_seconds")
	rootCmd.Flags().Bool("H", false, "Attempt to detect 802.11s draft mesh headers")
	rootCmd.Flags().Bool("I", false, "Monitor mode")
	rootCmd.Flags().Bool("K", false, "Don't attempt to verify TCP checksums")
	rootCmd.Flags().Bool("L", false, "List known data link types")
	rootCmd.Flags().String("M", "", "Use shared secret for TCP-MD5")
	rootCmd.Flags().Bool("N", false, "Don't print domain name of host names")
	rootCmd.Flags().Bool("O", false, "Do not run packet-matching code optimizer")
	rootCmd.Flags().String("Q", "", "Choose send/receive direction")
	rootCmd.Flags().Bool("S", false, "Print absolute TCP sequence numbers")
	rootCmd.Flags().String("T", false, "Force packet type")
	rootCmd.Flags().Bool("U", false, "Make output packet-buffered")
	rootCmd.Flags().String("W", "", "Used with -C to limit number of files created")
	rootCmd.Flags().Bool("X", false, "Print packets in hex and ASCII")
	rootCmd.Flags().String("Z", "", "Drop privileges to user")
	rootCmd.Flags().Bool("b", false, "Print AS number in BGP packets in ASDOT notation")
	rootCmd.Flags().String("c", "", "Exit after receiving count packets")
	rootCmd.Flags().Bool("count", false, "Print packet count on exit")
	rootCmd.Flags().Bool("d", false, "Dump compiled packet-matching code in human-readable form")
	rootCmd.Flags().Bool("e", false, "Print link-level header")
	rootCmd.Flags().Bool("f", false, "Print foreign IPv4 addresses numerically")
	rootCmd.Flags().Bool("h", false, "Print help")
	rootCmd.Flags().String("i", "", "Listen on interface")
	rootCmd.Flags().Bool("immediate-mode", false, "Capture in immediate mode")
	rootCmd.Flags().Bool("j", false, "Set time stamp type")
	rootCmd.Flags().String("k", "", "Metadata argument")
	rootCmd.Flags().Bool("l", false, "Make stdout line buffered")
	rootCmd.Flags().String("m", "", "Load SMI MIB module")
	rootCmd.Flags().Bool("micro", false, "Use microsecond precision")
	rootCmd.Flags().Bool("n", false, "Don't convert addresses to names")
	rootCmd.Flags().Bool("nano", false, "Use nanosecond precision")
	rootCmd.Flags().Bool("number", false, "Print packet number")
	rootCmd.Flags().Bool("p", false, "Don't put interface into promiscuous mode")
	rootCmd.Flags().Bool("print", false, "Print packet count on exit")
	rootCmd.Flags().Bool("q", false, "Quick output")
	rootCmd.Flags().String("r", "", "Read packets from file")
	rootCmd.Flags().String("s", "", "Set snaplen")
	rootCmd.Flags().Bool("t", false, "Don't print timestamps")
	rootCmd.Flags().String("time-stamp-precision", "", "Set time stamp precision")
	rootCmd.Flags().Bool("u", false, "Print undecoded NFS handles")
	rootCmd.Flags().Bool("v", false, "Verbose output")
	rootCmd.Flags().Bool("w", false, "Write raw packets to file")
	rootCmd.Flags().Bool("x", false, "Print packets in hex")
	rootCmd.Flags().String("y", "", "Set data link type")
	rootCmd.Flags().String("z", "", "Post-rotate command")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"F": carapace.ActionFiles(),
		"Q": carapace.ActionValues("in", "out", "inout"),
		"i": net.ActionInterfaces(),
		"r": carapace.ActionFiles(),
		"time-stamp-precision": carapace.ActionValues("micro", "nano"),
		"y": carapace.ActionValues("en10mb", "en3mb", "linux_sll", "raw", "null"),
	})
}