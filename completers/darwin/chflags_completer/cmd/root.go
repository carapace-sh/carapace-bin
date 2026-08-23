package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chflags",
	Short: "change file flags",
	Long:  "https://keith.github.io/xcode-manpages/chflags.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("crossmount", "x", false, "Do not cross mount points")
	rootCmd.Flags().BoolP("dereference", "L", false, "Follow all symbolic links")
	rootCmd.Flags().BoolP("follow", "H", false, "Follow symbolic links on command line")
	rootCmd.Flags().BoolP("force", "f", false, "Do not display a diagnostic message if chflags could not modify the flags")
	rootCmd.Flags().BoolP("nofollow", "h", false, "Change flags of symbolic link itself")
	rootCmd.Flags().BoolP("physical", "P", false, "Do not follow symbolic links (default)")
	rootCmd.Flags().BoolP("recursive", "R", false, "Change file flags recursively")
	rootCmd.Flags().BoolP("verbose", "v", false, "Be verbose showing filenames as flags are modified")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("arch", "archived", "nodump", "opaque", "sappnd", "sappend", "schg", "schange", "simmutable", "uappnd", "uappend", "uchg", "uchange", "uimmutable", "hidden", "noarch", "noarchived", "nonodump", "noopaque", "nosappnd", "nosappend", "noschg", "noschange", "nosimmutable", "nouappnd", "nouappend", "nouchg", "nouchange", "nouimmutable", "nohidden"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
