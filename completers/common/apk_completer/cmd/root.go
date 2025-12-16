package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apk",
	Short: "Alpine package manager",
	Long:  "https://gitlab.alpinelinux.org/alpine/apk-tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.AddGroup(
		&cobra.Group{ID: "package installation and removal"},
		&cobra.Group{ID: "system maintenance"},
		&cobra.Group{ID: "querying package information"},
		&cobra.Group{ID: "repository maintenance"},
		&cobra.Group{ID: "miscellaneous"},
	)

	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().Bool("version", false, "show version")
}
