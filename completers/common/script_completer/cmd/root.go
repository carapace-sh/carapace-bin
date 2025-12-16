package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "script",
	Short: "make typescript of terminal session",
	Long:  "https://man7.org/linux/man-pages/man1/script.1.html",
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

	rootCmd.Flags().BoolP("append", "a", false, "append to the log file")
	rootCmd.Flags().StringP("command", "c", "", "run command rather than interactive shell")
	rootCmd.Flags().StringP("echo", "E", "", "echo input in session (auto, always or never)")
	rootCmd.Flags().BoolP("flush", "f", false, "run flush after each write")
	rootCmd.Flags().Bool("force", false, "use output file even when it is a link")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().StringP("log-in", "I", "", "log stdin to file")
	rootCmd.Flags().StringP("log-io", "B", "", "log stdin and stdout to file")
	rootCmd.Flags().StringP("log-out", "O", "", "log stdout to file (default)")
	rootCmd.Flags().StringP("log-timing", "T", "", "log timing information to file")
	rootCmd.Flags().StringP("logging-format", "m", "", "force to 'classic' or 'advanced' format")
	rootCmd.Flags().StringP("output-limit", "o", "", "terminate if output files exceed size")
	rootCmd.Flags().BoolP("quiet", "q", false, "be quiet")
	rootCmd.Flags().BoolP("return", "e", false, "return exit code of the child process")
	rootCmd.Flags().StringP("timing", "t", "", "deprecated alias to -T (default file is stderr)")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	rootCmd.Flag("timing").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"echo":           carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"log-in":         carapace.ActionFiles(),
		"log-io":         carapace.ActionFiles(),
		"log-out":        carapace.ActionFiles(),
		"log-timing":     carapace.ActionFiles(),
		"logging-format": carapace.ActionValues("advanced", "classic"),
		"timing":         carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
