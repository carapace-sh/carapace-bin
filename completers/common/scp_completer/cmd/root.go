package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scp",
	Short: "OpenSSH secure file copy",
	Long:  "https://linux.die.net/man/1/scp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("3", "3", false, "Copies between two remote hosts are transferred through the local host")
	rootCmd.Flags().BoolS("4", "4", false, "Forces scp to use IPv4 addresses only")
	rootCmd.Flags().BoolS("6", "6", false, "Forces scp to use IPv6 addresses only")
	rootCmd.Flags().BoolS("A", "A", false, "Allows forwarding of ssh-agent to the remote system")
	rootCmd.Flags().BoolS("B", "B", false, "Selects batch mode")
	rootCmd.Flags().BoolS("C", "C", false, "Compression enable")
	rootCmd.Flags().StringS("D", "D", "", "Connect directly to a local SFTP server program")
	rootCmd.Flags().StringS("F", "F", "", "Specifies an alternative per-user configuration file for ssh")
	rootCmd.Flags().StringS("J", "J", "", "Connect to the target host by first making an scp connection to the jump host")
	rootCmd.Flags().BoolS("O", "O", false, "Use the legacy SCP protocol for file transfers instead of the SFTP protocol")
	rootCmd.Flags().StringS("P", "P", "", "Specifies the port to connect to on the remote host")
	rootCmd.Flags().BoolS("R", "R", false, "Copies between two remote hosts are performed by connecting to the origin host and executing scp there")
	rootCmd.Flags().StringS("S", "S", "", "Name of program to use for the encrypted connection")
	rootCmd.Flags().BoolS("T", "T", false, "Disable strict filename checking")
	rootCmd.Flags().StringSliceS("X", "X", nil, "Specify an option that controls aspects of SFTP protocol behaviour")
	rootCmd.Flags().StringS("c", "c", "", "Selects the cipher to use for encrypting the data transfer")
	rootCmd.Flags().StringS("i", "i", "", "Selects the file from which the identity for public key authentication is read")
	rootCmd.Flags().StringS("l", "l", "", "Limits the used bandwidth, specified in Kbit/s")
	rootCmd.Flags().StringSliceS("o", "o", nil, "Can be used to pass options to ssh in the format used in ssh_config")
	rootCmd.Flags().BoolS("p", "p", false, "Preserves modification times, access times, and file mode bits from the source file")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().BoolS("r", "r", false, "Recursively copy entire directories")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"D": carapace.ActionFiles(),
		"F": carapace.ActionFiles(),
		"J": net.ActionHosts(),
		"S": carapace.ActionFiles(),
		"X": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"nrequests", "Controls how many concurrent SFTP read or write requests may be in progress",
					"buffer", "Controls the maximum buffer size for a single SFTP read/write operation",
				).Suffix("=")
			default:
				return carapace.ActionValues()
			}
		}),
		"c": ssh.ActionCiphers(),
		"i": carapace.Batch(
			ssh.ActionPrivateKeys(),
			carapace.ActionFiles(),
		).ToA(),
		"o": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ssh.ActionOptions().NoSpace()
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.Batch(
					net.ActionHosts(),
					ssh.ActionHosts(rootCmd.Flag("F").Value.String()).Style("yellow"),
					carapace.ActionFiles(),
				).ToA().NoSpace()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
