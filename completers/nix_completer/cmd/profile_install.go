package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var profile_installCmd = &cobra.Command{
	Use:   "install",
	Short: "install a package into a profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_installCmd).Standalone()

	profile_installCmd.Flags().String("priority", "", "The priority of the package to install")
	profile_installCmd.Flags().String("profile", "", "The profile to operate on")
	profileCmd.AddCommand(profile_installCmd)

	addEvaluationFlags(profile_installCmd)
	addFlakeFlags(profile_installCmd)
	addLoggingFlags(profile_installCmd)

	carapace.Gen(profile_installCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionDirectories(),
	})

	carapace.Gen(profile_installCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("#", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return nix.ActionLocalChannels().Suffix("#")
					case 1:
						return carapace.ActionMessage("TODO: support branch completion") // TODO support branch completion
					default:
						return carapace.ActionValues()
					}
				})

			case 1:
				return nix.ActionPackages(strings.SplitN(c.Parts[0], "/", 2)[0])

			default:
				return carapace.ActionValues()
			}
		}),
	)
}
