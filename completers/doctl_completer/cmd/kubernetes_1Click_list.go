package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_1Click_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of Kubernetes 1-Click applications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_1Click_listCmd).Standalone()
	kubernetes_1Click_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `SLUG`, `TYPE`")
	kubernetes_1Click_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	kubernetes_1ClickCmd.AddCommand(kubernetes_1Click_listCmd)
}
