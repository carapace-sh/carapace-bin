package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inertiaStartSsrCmd = &cobra.Command{
	Use:   "inertia:start-ssr [--runtime [RUNTIME]]",
	Short: "Start the Inertia SSR server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inertiaStartSsrCmd).Standalone()

	inertiaStartSsrCmd.Flags().String("runtime", "", "The runtime to use (`node` or `bun`)")
	rootCmd.AddCommand(inertiaStartSsrCmd)
}
