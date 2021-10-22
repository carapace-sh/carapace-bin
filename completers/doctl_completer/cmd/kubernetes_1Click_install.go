package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_1Click_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install 1-click apps on a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_1Click_installCmd).Standalone()
	kubernetes_1Click_installCmd.Flags().StringSlice("1-clicks", []string{}, "1-clicks to be installed on a Kubernetes cluster. Multiple 1-clicks can be added at once. Example: --1-clicks moon,loki,netdata")
	kubernetes_1Click_installCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `SLUG`, `TYPE`")
	kubernetes_1Click_installCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	kubernetes_1ClickCmd.AddCommand(kubernetes_1Click_installCmd)
}
