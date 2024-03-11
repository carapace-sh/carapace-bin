package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "last",
	Short: "Show a listing of last logged in users",
	Long:  "https://linux.die.net/man/1/last",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("dns", "d", false, "translate the IP number back into a hostname")
	rootCmd.Flags().StringP("file", "f", "", "use a specific file instead of /var/log/wtmp")
	rootCmd.Flags().BoolP("fullnames", "w", false, "display full user and domain names")
	rootCmd.Flags().BoolP("fulltimes", "F", false, "print full login and logout times and dates")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("hostlast", "a", false, "display hostnames in the last column")
	rootCmd.Flags().BoolP("ip", "i", false, "display IP numbers in numbers-and-dots notation")
	rootCmd.Flags().StringP("limit", "n", "", "how many lines to show")
	rootCmd.Flags().BoolP("nohostname", "R", false, "don't display the hostname field")
	rootCmd.Flags().StringP("present", "p", "", "display who were present at the specified time")
	rootCmd.Flags().StringP("since", "s", "", "display the lines since the specified time")
	rootCmd.Flags().BoolP("system", "x", false, "display system shutdown entries and run level changes")
	rootCmd.Flags().String("time-format", "", "show timestamps in the specified <format>")
	rootCmd.Flags().StringP("until", "t", "", "display the lines until the specified time")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file":        carapace.ActionFiles(),
		"present":     actionTime(),
		"since":       actionTime(),
		"time-format": carapace.ActionValues("notime", "short", "full", "iso"),
		"until":       actionTime(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.Batch(
			os.ActionUsers(),
			os.ActionTerminals(),
		).ToA(),
	)
}

func actionTime() carapace.Action {
	return carapace.Batch(
		time.ActionDate(),
		time.ActionDateTime(time.DateTimeOpts{}),
		time.ActionTime(),
		carapace.ActionValues("now", "yesterday", "today", "tomorrow").Style(style.Blue),
	).ToA()
}
