package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var contentsCmd = &cobra.Command{
	Use:   "contents",
	Short: "List files installed by a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(contentsCmd).Standalone()
	contentsCmd.Flags().Bool("size", false, "Print file sizes")
	contentsCmd.Flags().String("units", "", "Size unit (B, K, Ki, Mi, Gi, kB, MB, GB)")
	rootCmd.AddCommand(contentsCmd)
}
