package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flatpak",
	Short: "Linux application sandboxing and distribution framework",
	Long:  "https://flatpak.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "install", Title: "Manage installed applications and runtimes"},
		&cobra.Group{ID: "search", Title: "Find applications and runtimes"},
		&cobra.Group{ID: "run", Title: "Manage running applications"},
		&cobra.Group{ID: "access", Title: "Manage file access"},
		&cobra.Group{ID: "permission", Title: "Manage dynamic permissions"},
		&cobra.Group{ID: "repository", Title: "Manage remote repositories"},
		&cobra.Group{ID: "build", Title: "Build applications"},
	)

	rootCmd.Flags().Bool("default-arch", false, "Print default arch and exit")
	rootCmd.Flags().Bool("gl-drivers", false, "Print active gl drivers and exit")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().Bool("installations", false, "Print paths for system installations and exit")
	rootCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	rootCmd.Flags().Bool("print-system-only", false, "Only include the system installation with --print-updated-env")
	rootCmd.Flags().Bool("print-updated-env", false, "Print the updated environment needed to run flatpaks")
	rootCmd.Flags().Bool("supported-arches", false, "Print supported arches and exit")
	rootCmd.Flags().CountP("verbose", "v", "Show debug information, -vv for more detail")
	rootCmd.Flags().Bool("version", false, "Print version information and exit")
}

func columns(a carapace.Action) carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return a
		default:
			return flatpak.ActionEllipsizations()
		}
	}).UniqueListF(",", func(s string) string {
		return strings.SplitN(s, ":", 2)[0]
	})
}
