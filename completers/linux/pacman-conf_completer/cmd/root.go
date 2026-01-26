package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pacman"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pacman-conf",
	Short: "query pacman's configuration file",
	Long:  "https://man.archlinux.org/man/pacman-conf.8.en",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config", "c", "", "set an alternate configuration file")
	rootCmd.Flags().BoolP("help", "h", false, "display this help information")
	rootCmd.Flags().StringP("repo", "r", "", "query options for a specific repo")
	rootCmd.Flags().BoolP("repo-list", "l", false, "list configured repositories")
	rootCmd.Flags().StringP("rootdir", "R", "", "set an alternate installation root")
	rootCmd.Flags().StringP("sysroot", "S", "", "set an alternate system root")
	rootCmd.Flags().BoolP("verbose", "v", false, "always show directive names")
	rootCmd.Flags().BoolP("version", "V", false, "display version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":  carapace.ActionFiles(),
		"repo":    pacman.ActionRepositories(),
		"rootdir": carapace.ActionDirectories(),
		"sysroot": carapace.ActionDirectories(),
	})
}
