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

	scpCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	scpCmd.Flags().Bool("no-preserve", false, "Preserves access and modification times from the original file")
	scpCmd.Flags().Bool("no-quiet", false, "Quiet mode")
	scpCmd.Flags().Bool("no-recursive", false, "Recursive copy of subdirectories")
	scpCmd.Flags().StringP("port", "P", "", "Port to connect to on the remote host")
	scpCmd.Flags().BoolP("preserve", "p", false, "Preserves access and modification times from the original file")
	scpCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	scpCmd.Flags().BoolP("recursive", "r", false, "Recursive copy of subdirectories")
	scpCmd.Flag("no-preserve").Hidden = true
	scpCmd.Flag("no-quiet").Hidden = true
	scpCmd.Flag("no-recursive").Hidden = true
	rootCmd.AddCommand(scpCmd)

	carapace.Gen(scpCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"port":    net.ActionPorts(),
	})
}
