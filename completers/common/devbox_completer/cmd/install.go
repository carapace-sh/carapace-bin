package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install all packages mentioned in devbox.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	installCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	installCmd.Flags().Bool("tidy-lockfile", false, "Fix missing store paths in the devbox.lock file.")
	rootCmd.AddCommand(installCmd)

	// TODO environment
	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})
}
