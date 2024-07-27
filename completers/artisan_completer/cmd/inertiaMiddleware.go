package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inertiaMiddlewareCmd = &cobra.Command{
	Use:   "inertia:middleware [--force] [--] [<name>]",
	Short: "Create a new Inertia middleware",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inertiaMiddlewareCmd).Standalone()

	inertiaMiddlewareCmd.Flags().Bool("force", false, "Create the class even if the Middleware already exists")
	rootCmd.AddCommand(inertiaMiddlewareCmd)
}
