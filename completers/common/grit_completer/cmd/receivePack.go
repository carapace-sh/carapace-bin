package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var receivePackCmd = &cobra.Command{
	Use:     "receive-pack",
	Short:   "Accept a push over stdin/stdout (run by ssh and grit-http-server)",
	GroupID: "plumbing",
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(receivePackCmd).Standalone()

	receivePackCmd.Flags().Bool("advertise-refs", false, "Print the ref advertisement and exit (smart HTTP discovery)")
	receivePackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	receivePackCmd.Flags().Bool("http-backend-info-refs", false, "Print the ref advertisement and exit (smart HTTP discovery)")
	receivePackCmd.Flags().Bool("stateless-rpc", false, "Answer a single request without advertising refs first (smart HTTP)")
	receivePackCmd.Flag("http-backend-info-refs").Hidden = true
	rootCmd.AddCommand(receivePackCmd)
}
