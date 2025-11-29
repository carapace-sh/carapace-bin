package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_addCmd = &cobra.Command{
	Use:   "add <pkg>...",
	Short: "Add a new package to your devbox",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_addCmd).Standalone()

	global_addCmd.Flags().StringSlice("allow-insecure", nil, "allow adding packages marked as insecure.")
	global_addCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	global_addCmd.Flags().Bool("disable-plugin", false, "disable plugin (if any) for this package.")
	global_addCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	global_addCmd.Flags().StringSliceP("exclude-platform", "e", nil, "exclude packages from a specific platform.")
	global_addCmd.Flags().StringSliceP("outputs", "o", nil, "specify the outputs to select for the nix package")
	global_addCmd.Flags().String("patch", "", "allow Devbox to patch the package to fix known issues (auto, always, never)")
	global_addCmd.Flags().Bool("patch-glibc", false, "patch any ELF binaries to use the latest glibc version in nixpkgs")
	global_addCmd.Flags().StringSliceP("platform", "p", nil, "add packages to run on only this platform.")
	global_addCmd.Flag("config").Hidden = true
	global_addCmd.Flag("patch-glibc").Hidden = true
	globalCmd.AddCommand(global_addCmd)

	// TODO complete more flags
	carapace.Gen(global_addCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
		"patch":  carapace.ActionValues("auto", "always", "never"),
	})

	// TODO positional completion
}
