package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newFetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch [options] <url>",
		Short: "Copy a package into the global cache and print its hash",
	}

	cmd.Flags().Bool("global-cache-dir", false, "Override the global cache directory")
	cmd.Flags().Bool("debug-hash", false, "Print verbose hash information to stdout")
	cmd.Flags().Bool("save", false, "Add the fetched package to build.zig.zon")
	cmd.Flags().String("save", "", "Add the fetched package to build.zig.zon with given name")
	cmd.Flags().Bool("save-exact", false, "Add exact package to build.zig.zon")

	carapace.Gen(cmd).PositionalCompletion(
		carapace.ActionFiles("tar.gz", "tar.xz", "tar.zst", "zip", "gz", "xz"),
	)

	return cmd
}
