package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the language server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverCmd).Standalone()

	serverCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	serverCmd.Flags().Bool("no-preview", false, "")
	serverCmd.Flags().Bool("preview", false, "Enable preview mode. Use `--no-preview` to disable")
	serverCmd.Flags().Bool("untrusted-workspace", false, "Treat all workspaces as untrusted, disabling the uv formatter backend")
	serverCmd.Flag("no-preview").Hidden = true
	rootCmd.AddCommand(serverCmd)
}
