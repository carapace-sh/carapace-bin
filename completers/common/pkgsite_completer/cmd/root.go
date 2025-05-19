package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pkgsite",
	Short: "Pkgsite extracts and generates documentation for Go programs",
	Long:  "https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("cache", "cache", false, "fetch from the module cache")
	rootCmd.Flags().StringS("cachedir", "cachedir", "", "module cache directory")
	rootCmd.Flags().BoolS("dev", "dev", false, "enable developer mode")
	rootCmd.Flags().BoolS("gopath_mode", "gopath_mode", false, "assume that local modules' paths are relative to GOPATH/src")
	rootCmd.Flags().StringS("gorepo", "gorepo", "", "path to Go repo on local filesystem")
	rootCmd.Flags().StringS("http", "http", "", "HTTP service address to listen for incoming requests on")
	rootCmd.Flags().BoolS("list", "list", false, "for each path, serve all modules in build list")
	rootCmd.Flags().BoolS("open", "open", false, "open a browser window to the server's address")
	rootCmd.Flags().BoolS("proxy", "proxy", false, "fetch from GOPROXY if not found locally")
	rootCmd.Flags().StringS("static", "static", "", "path to folder containing static files served")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cachedir": carapace.ActionDirectories(),
		"gorepo":   carapace.ActionDirectories(),
		"http": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues()
			default:
				return net.ActionPorts()
			}
		}),
		"static": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories().List(","),
	)
}
