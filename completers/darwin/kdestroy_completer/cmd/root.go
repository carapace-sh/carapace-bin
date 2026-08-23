package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kdestroy",
	Short: "destroy Kerberos tickets",
	Long:  "https://man.freebsd.org/cgi/man.cgi?kdestroy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "All caches")
	rootCmd.Flags().BoolS("a", "a", false, "All")
	rootCmd.Flags().Bool("all", false, "Remove all credential caches")
	rootCmd.Flags().StringS("c", "c", "", "Cache file")
	rootCmd.Flags().String("cache", "", "Cache type:name")
	rootCmd.Flags().String("credential", "", "Principal")
	rootCmd.Flags().Bool("no-delete-v4", false, "No delete v4")
	rootCmd.Flags().Bool("no-unlog", false, "No unlog")
	rootCmd.Flags().StringS("p", "p", "", "Principal")
	rootCmd.Flags().String("principal", "", "Principal")
}
