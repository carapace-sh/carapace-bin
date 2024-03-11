package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Create a new package",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().StringP("install", "i", "latest", "Initialize a package with a specific bundle that will be locked in the project")
	initCmd.Flags().BoolP("private", "p", false, "Initialize a private package")
	initCmd.Flags().BoolP("workspace", "w", false, "Initialize a workspace root with a packages/ directory")
	rootCmd.AddCommand(initCmd)

	initCmd.Flag("install").NoOptDefVal = "latest"
}
