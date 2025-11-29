package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <pkg>...",
	Short: "Add a new package to your devbox",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().StringSlice("allow-insecure", nil, "allow adding packages marked as insecure.")
	addCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	addCmd.Flags().Bool("disable-plugin", false, "disable plugin (if any) for this package.")
	addCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	addCmd.Flags().StringSliceP("exclude-platform", "e", nil, "exclude packages from a specific platform.")
	addCmd.Flags().StringSliceP("outputs", "o", nil, "specify the outputs to select for the nix package")
	addCmd.Flags().String("patch", "", "allow Devbox to patch the package to fix known issues (auto, always, never)")
	addCmd.Flags().Bool("patch-glibc", false, "patch any ELF binaries to use the latest glibc version in nixpkgs")
	addCmd.Flags().StringSliceP("platform", "p", nil, "add packages to run on only this platform.")
	addCmd.Flag("patch-glibc").Hidden = true
	rootCmd.AddCommand(addCmd)

	// TODO flag completion
	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
		"patch":  carapace.ActionValues("auto", "always", "never"),
	})

	// TODO devbox is currently hardcoded against `nixpkgs` alias which is implicitly added as prefix.
	// This channel also has no sqlitedb so package names cannot be completed
	// carapace.Gen(addCmd).PositionalAnyCompletion(
	// 	carapace.ActionCallback(func(c carapace.Context) carapace.Action {
	// 		return nix.ActionChannelPackages().Invoke(c).Filter(c.Parts).ToA()
	// 	}),
	// )
}
