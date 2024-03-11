package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golangci-lint",
	Short: "golangci-lint is a smart linters runner.",
	Long:  "https://golangci-lint.run/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("color", "", "Use color when printing; can be 'always', 'auto', or 'never'")
	rootCmd.PersistentFlags().StringP("concurrency", "j", "", "Concurrency (default NumCPU)")
	rootCmd.PersistentFlags().String("cpu-profile-path", "", "Path to CPU profile output file")
	rootCmd.PersistentFlags().String("mem-profile-path", "", "Path to memory profile output file")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "disables congrats outputs")
	rootCmd.PersistentFlags().String("trace-path", "", "Path to trace output file")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().Bool("version", false, "Print version")
	rootCmd.Flag("silent").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":            carapace.ActionValues("always", "auto", "never").StyleF(style.ForKeyword),
		"cpu-profile-path": carapace.ActionFiles(),
		"mem-profile-path": carapace.ActionFiles(),
		"trace-path":       carapace.ActionFiles(),
	})
}
