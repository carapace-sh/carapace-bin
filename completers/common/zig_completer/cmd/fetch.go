package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch [options] <url>",
	Short:   "Copy a package into global cache and print its hash",
	GroupID: "project",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	rootCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().String("color", "", "Enable or disable colored error messages")
	fetchCmd.Flags().Bool("debug-hash", false, "Print verbose hash information to stdout")
	fetchCmd.Flags().String("global-cache-dir", "", "Override path to global Zig cache directory")
	fetchCmd.Flags().BoolP("help", "h", false, "Print help")
	fetchCmd.Flags().Bool("save", false, "Add the fetched package to build.zig.zon")
	fetchCmd.Flags().Bool("save-exact", false, "Add the fetched package to build.zig.zon, storing the URL verbatim")
	fetchCmd.Flags().String("save-exact-name", "", "Add the fetched package to build.zig.zon as name, storing the URL verbatim")
	fetchCmd.Flags().String("save-name", "", "Add the fetched package to build.zig.zon as name")

	carapace.Gen(fetchCmd).FlagCompletion(carapace.ActionMap{
		"color":            carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
		"global-cache-dir": carapace.ActionDirectories(),
	})
}
