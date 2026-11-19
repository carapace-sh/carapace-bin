package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lspCmd = &cobra.Command{
	Use:    "lsp",
	Short:  "Starts skaffold LSP language server that provides lint suggestions to IDEs",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lspCmd).Standalone()

	lspCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	lspCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	lspCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	lspCmd.Flags().String("port", "4389", "port number that the skaffold lsp will expose for connections, '4389' is the default port value")
	lspCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	lspCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	lspCmd.Flags().StringP("user", "u", "", "")
	lspCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(lspCmd)
	rootCmd.AddCommand(lspCmd)
}
