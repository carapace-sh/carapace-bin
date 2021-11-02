package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve details about a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_getCmd).Standalone()
	registry_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`, `Endpoint`")
	registry_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	registryCmd.AddCommand(registry_getCmd)
}
