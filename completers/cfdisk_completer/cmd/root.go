package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cfdisk",
	Short: "display or manipulate a disk partition table",
	Long:  "https://man7.org/linux/man-pages/man8/cfdisk.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("color", "L", "", "colorize output")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().String("lock", "", "use exclusive device lock")
	rootCmd.Flags().BoolP("read-only", "r", false, "forced open cfdisk in read-only mode")
	rootCmd.Flags().StringP("sector-size", "b", "", "physical and logical sector size")
	rootCmd.Flags().BoolP("version", "V", false, "display version")
	rootCmd.Flags().BoolP("zero", "z", false, "start with zeroed partition table")

	rootCmd.Flag("color").NoOptDefVal = " "
	rootCmd.Flag("lock").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":       carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"lock":        carapace.ActionValues("yes", "no", "nonblock").StyleF(style.ForKeyword),
		"sector-size": carapace.ActionValues("512", "1024", "2048", "4096"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			fs.ActionBlockDevices(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
