// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var github_tokenCmd = &cobra.Command{
	Use:    "token",
	Short:  "Show the GitHub token for a host (deprecated)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(github_tokenCmd).Standalone()

	github_tokenCmd.Flags().BoolP("help", "h", false, "Print help")
	github_tokenCmd.Flags().Bool("oauth", false, "Force native GitHub OAuth device flow instead of normal token resolution")
	github_tokenCmd.Flags().Bool("raw", false, "Print only the token value")
	github_tokenCmd.Flags().Bool("refresh", false, "Mint a fresh OAuth token even if the cached one has not expired, via the refresh-token grant or a new device-code flow")
	github_tokenCmd.Flags().Bool("unmask", false, "Show the full unmasked token")
	githubCmd.AddCommand(github_tokenCmd)
}
