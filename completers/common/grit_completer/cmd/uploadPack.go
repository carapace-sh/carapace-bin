package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uploadPackCmd = &cobra.Command{
	Use:    "upload-pack",
	Short:  "Serve a fetch or clone over stdin/stdout (run by ssh and grit-http-server)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uploadPackCmd).Standalone()

	uploadPackCmd.Flags().Bool("advertise-refs", false, "Print the ref advertisement and exit (smart HTTP discovery)")
	uploadPackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	uploadPackCmd.Flags().Bool("http-backend-info-refs", false, "Print the ref advertisement and exit (smart HTTP discovery)")
	uploadPackCmd.Flags().Bool("stateless-rpc", false, "Answer a single request without advertising refs first (smart HTTP)")
	uploadPackCmd.Flag("http-backend-info-refs").Hidden = true
	rootCmd.AddCommand(uploadPackCmd)
}
