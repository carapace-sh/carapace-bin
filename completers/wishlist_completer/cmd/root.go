package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wishlist",
	Short: "The SSH Directory",
	Long:  "https://github.com/charmbracelet/soft-serve",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().StringP("config", "c", "", "Path to the config file to use. Defaults to, in order of preference: .wishlist/config.yaml, .wishlist/config.yml, .wishlist/config, /home/rsteube/.config/wishlist.yaml, /home/rsteube/.config/wishlist.yml, /home/rsteube/.config/wishlist, /home/rsteube/.ssh/config, /etc/ssh/ssh_config")
	rootCmd.Flags().BoolP("help", "h", false, "help for wishlist")
	rootCmd.Flags().BoolP("version", "v", false, "version for wishlist")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
