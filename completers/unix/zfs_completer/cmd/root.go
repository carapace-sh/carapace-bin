package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zfs",
	Short: "configure ZFS file systems",
	Long:  "https://openzfs.github.io/openzfs-docs/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "dataset", Title: "Dataset Commands"},
		&cobra.Group{ID: "property", Title: "Property Commands"},
		&cobra.Group{ID: "mount", Title: "Mount/Share Commands"},
		&cobra.Group{ID: "send", Title: "Send/Receive Commands"},
		&cobra.Group{ID: "delegation", Title: "Delegation Commands"},
		&cobra.Group{ID: "encryption", Title: "Encryption Commands"},
		&cobra.Group{ID: "misc", Title: "Other Commands"},
	)
}
