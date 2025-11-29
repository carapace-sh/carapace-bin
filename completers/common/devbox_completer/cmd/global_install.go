package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install all packages mentioned in devbox.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_installCmd).Standalone()

	global_installCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	global_installCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	global_installCmd.Flags().Bool("tidy-lockfile", false, "Fix missing store paths in the devbox.lock file.")
	global_installCmd.Flag("config").Hidden = true
	globalCmd.AddCommand(global_installCmd)

	// TODO environment
	carapace.Gen(global_installCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})
}
