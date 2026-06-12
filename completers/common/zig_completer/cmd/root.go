package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zig",
	Short: "General-purpose programming language and toolchain",
	Long:  "https://ziglang.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.AddGroup(&cobra.Group{ID: "project", Title: "Project Commands"})
	rootCmd.AddGroup(&cobra.Group{ID: "build", Title: "Build Commands"})
	rootCmd.AddGroup(&cobra.Group{ID: "source", Title: "Source Commands"})
	rootCmd.AddGroup(&cobra.Group{ID: "wrapper", Title: "Wrapper Commands"})

	rootCmd.Flags().String("color", "", "Enable or disable colored error messages")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("version", "v", false, "Print version number and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
	})
}
