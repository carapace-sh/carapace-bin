package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopls",
	Short: "gopls is a Go language server",
	Long:  "https://pkg.go.dev/golang.org/x/tools/gopls",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("ocagent", "ocagent", "", "the address of the ocagent (e.g. http://localhost:55678), or off (default \"off\")")
	rootCmd.Flags().StringS("profile.alloc", "profile.alloc", "", "write alloc profile to this file")
	rootCmd.Flags().StringS("profile.cpu", "profile.cpu", "", "write CPU profile to this file")
	rootCmd.Flags().StringS("profile.mem", "profile.mem", "", "write memory profile to this file")
	rootCmd.Flags().StringS("profile.trace", "profile.trace", "", "write trace log to this file")
	rootCmd.Flags().StringS("remote", "remote", "", "forward all commands to a remote lsp specified by this flag. With no special prefix, this is assumed to be a TCP address. If prefixed by 'unix;', the subsequent address is assumed to be a unix domain socket. If 'auto', or prefixed by 'auto;', the remote address is automatically resolved based on the executing environment.")
	rootCmd.Flags().BoolN("verbose", "v", false, "verbose output")
	rootCmd.Flags().BoolN("veryverbose", "vv", false, "very verbose output")
	addServerFlags(rootCmd)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"profile.alloc": carapace.ActionFiles(),
		"profile.cpu":   carapace.ActionFiles(),
		"profile.mem":   carapace.ActionFiles(),
		"profile.trace": carapace.ActionFiles(),
	})
}
