package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var libcCmd = &cobra.Command{
	Use:   "libc",
	Short: "Display native libc paths file or validate one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(libcCmd).Standalone()

	rootCmd.AddCommand(libcCmd)

	libcCmd.Flags().String("color", "", "Enable or disable colored error messages")
	libcCmd.Flags().BoolP("help", "h", false, "Print help")

	carapace.Gen(libcCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
	})

	carapace.Gen(libcCmd).PositionalAnyCompletion(
		carapace.ActionFiles().FilterArgs(),
	)
}
