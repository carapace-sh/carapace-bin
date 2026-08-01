package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zpool",
	Short: "configure ZFS storage pools",
	Long:  "https://openzfs.github.io/openzfs-docs/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "creation", Title: "Pool Creation Commands"},
		&cobra.Group{ID: "vdev", Title: "Virtual Device Commands"},
		&cobra.Group{ID: "property", Title: "Property Commands"},
		&cobra.Group{ID: "monitoring", Title: "Monitoring Commands"},
		&cobra.Group{ID: "maintenance", Title: "Maintenance Commands"},
		&cobra.Group{ID: "fault", Title: "Fault Resolution Commands"},
		&cobra.Group{ID: "io", Title: "Import/Export Commands"},
	)
}
