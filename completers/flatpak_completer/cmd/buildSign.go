package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var buildSignCmd = &cobra.Command{
	Use:     "build-sign [OPTION…] LOCATION [ID [BRANCH]]",
	Short:   "Sign an application or runtime",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildSignCmd).Standalone()

	buildSignCmd.Flags().String("arch", "", "Arch to install for")
	buildSignCmd.Flags().String("gpg-homedir", "", "GPG Homedir to use when looking for keyrings")
	buildSignCmd.Flags().String("gpg-sign", "", "GPG Key ID to sign the commit with")
	buildSignCmd.Flags().BoolP("help", "h", false, "Show help options")
	buildSignCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	buildSignCmd.Flags().Bool("runtime", false, "Look for runtime with the specified name")
	buildSignCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(buildSignCmd)

	carapace.Gen(buildSignCmd).FlagCompletion(carapace.ActionMap{
		"arch":        flatpak.ActionArches(),
		"gpg-homedir": carapace.ActionDirectories(), // TODO handle gpg-sign
		"gpg-sign":    os.ActionGpgKeyIds(),
	})

	// TODO positional

}
