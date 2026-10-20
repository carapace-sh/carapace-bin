// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Show the GitHub token mise uses for a host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_githubCmd).Standalone()

	token_githubCmd.Flags().String("git-credential", "", "Speak Git's credential helper protocol")
	token_githubCmd.Flags().BoolP("help", "h", false, "Print help")
	token_githubCmd.Flags().Bool("oauth", false, "Use only native GitHub OAuth (a cached token, a refresh, or the device-code flow), ignoring other token sources")
	token_githubCmd.Flags().Bool("raw", false, "Print only the token, unmasked, for use in scripts")
	token_githubCmd.Flags().Bool("refresh", false, "With --oauth, get a new OAuth token even if the cached one has not expired, through the refresh-token grant or a new device-code flow. Use it after changing the GitHub App's installations or permissions: cached tokens keep their original access until they expire.")
	token_githubCmd.Flags().Bool("unmask", false, "Show the full unmasked token")
	token_githubCmd.Flag("git-credential").Hidden = true
	tokenCmd.AddCommand(token_githubCmd)
}
