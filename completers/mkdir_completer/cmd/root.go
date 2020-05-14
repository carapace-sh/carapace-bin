package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	rootCmd.Flags().BoolP("Z", "Z", false, "set SELinux security context of each created directory to the default type")
	rootCmd.Flags().String("context", "", "like -Z, or if CTX is specified then set the SELinux or SMACK security context to CTX")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("mode", "m", "", "set file mode (as in chmod), not a=rwx - umask")
	rootCmd.Flags().BoolP("parents", "p", false, "no error if existing, make parent directories as needed")
	rootCmd.Flags().BoolP("verbose", "v", false, "print a message for each created directory")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
