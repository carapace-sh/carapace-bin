package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "present",
	Short: "present implements parsing and rendering of present file",
	Long:  "https://pkg.go.dev/golang.org/x/tools/present",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("base", "base", "", "base path for slide template and static resources")
	rootCmd.Flags().StringS("content", "content", "", "base path for presentation content")
	rootCmd.Flags().BoolN("help", "h", false, "show help")
	rootCmd.Flags().StringS("http", "http", "", "HTTP service address")
	rootCmd.Flags().BoolS("notes", "notes", false, "enable presenter notes")
	rootCmd.Flags().StringS("orighost", "orighost", "", "host component of web origin URL")
	rootCmd.Flags().BoolS("play", "play", false, "enable playground")
	rootCmd.Flags().BoolS("use_playground", "use_playground", false, "run code snippets using play.golang.org")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"base":    carapace.ActionDirectories(),
		"content": carapace.ActionDirectories(),
		"http": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
	})
}
