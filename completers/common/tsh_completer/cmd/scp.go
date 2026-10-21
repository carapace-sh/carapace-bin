package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var scpCmd = &cobra.Command{
	Use:   "scp",
	Short: "Transfer files to a remote SSH node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scpCmd).Standalone()

	scpCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	scpCmd.Flags().Bool("no-no-resume", false, "Disable SSH connection resumption.")
	scpCmd.Flags().Bool("no-preserve", false, "Preserves access and modification times from the original file.")
	scpCmd.Flags().Bool("no-quiet", false, "Quiet mode.")
	scpCmd.Flags().Bool("no-recursive", false, "Recursive copy of subdirectories.")
	scpCmd.Flags().Bool("no-relogin", false, "Permit performing an authentication attempt on a failed command.")
	scpCmd.Flags().Bool("no-resume", false, "Disable SSH connection resumption.")
	scpCmd.Flags().StringP("port", "P", "", "Port to connect to on the remote host.")
	scpCmd.Flags().BoolP("preserve", "p", false, "Preserves access and modification times from the original file.")
	scpCmd.Flags().BoolP("quiet", "q", false, "Quiet mode.")
	scpCmd.Flags().BoolP("recursive", "r", false, "Recursive copy of subdirectories.")
	scpCmd.Flags().Bool("relogin", true, "Permit performing an authentication attempt on a failed command.")
	scpCmd.Flag("no-no-resume").Hidden = true
	scpCmd.Flag("no-preserve").Hidden = true
	scpCmd.Flag("no-quiet").Hidden = true
	scpCmd.Flag("no-recursive").Hidden = true
	scpCmd.Flag("no-relogin").Hidden = true
	rootCmd.AddCommand(scpCmd)

	carapace.Gen(scpCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"port":    net.ActionPorts(),
	})

	carapace.Gen(scpCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.Batch(
					tsh.ActionNodes(),
					carapace.ActionFiles(),
				).ToA()
			default:
				return carapace.ActionFiles()
			}
		}),
	)
}
