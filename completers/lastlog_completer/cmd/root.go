package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lastlog",
	Short: "reports the most recent login of all users or of a given user",
	Long:  "https://linux.die.net/man/8/lastlog",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("before", "b", "", "print only lastlog records older than DAYS")
	rootCmd.Flags().BoolP("clear", "C", false, "clear lastlog record of an user (usable only with -u)")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().StringP("root", "R", "", "directory to chroot into")
	rootCmd.Flags().BoolP("set", "S", false, "set lastlog record to current time (usable only with -u)")
	rootCmd.Flags().StringP("time", "t", "", "print only lastlog records more recent than DAYS")
	rootCmd.Flags().StringP("user", "u", "", "print lastlog record of the specified LOGIN")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"root": carapace.ActionDirectories(),
		"user": os.ActionUsers(),
	})
}
