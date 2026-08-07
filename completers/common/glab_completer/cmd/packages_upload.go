package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packages_uploadCmd = &cobra.Command{
	Use:     "upload <file> --name <package> --version <version> [flags]",
	Short:   "Upload a file to a project's package registry.",
	Aliases: []string{"ul"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packages_uploadCmd).Standalone()

	packages_uploadCmd.Flags().String("filename", "", "Name to store the file under. Defaults to the local file name.")
	packages_uploadCmd.Flags().StringP("name", "n", "", "Name of the package.")
	packages_uploadCmd.Flags().StringP("version", "v", "", "Version of the package.")
	packages_uploadCmd.MarkFlagRequired("name")
	packages_uploadCmd.MarkFlagRequired("version")
	packagesCmd.AddCommand(packages_uploadCmd)
}
