package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy all resources in the deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destroyCmd).Standalone()

	destroyCmd.Flags().Bool("all", false, "Destroy all deployments")
	destroyCmd.Flags().StringSlice("exclude", nil, "Destroy all except the specified machines")
	destroyCmd.Flags().StringSlice("include", nil, "Destroy only the specified machines")
	destroyCmd.Flags().Bool("wipe", false, "Securely wipe data on the machines")

	rootCmd.AddCommand(destroyCmd)
}
