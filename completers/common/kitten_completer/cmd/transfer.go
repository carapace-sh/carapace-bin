package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer files easily over the TTY device",
}

func init() {
	rootCmd.AddCommand(transferCmd)
	carapace.Gen(transferCmd).Standalone()

	transferCmd.Flags().String("compress", "", "Whether to compress data being sent. By default compression is enabled based on the type of file being sent. For files recognized as being already compressed, compression is turned off as it just wastes CPU cycles.")
	transferCmd.Flags().StringP("confirm-paths", "c", "", "Before actually transferring files, show a mapping of local file names to remote file names and ask for confirmation.")
	transferCmd.Flags().StringP("direction", "d", "", "Whether to send or receive files. send or download copy files from the computer on which the kitten is running (usually the remote computer) to the local computer. receive or upload copy files from the local computer to the remote computer.")
	transferCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	transferCmd.Flags().StringP("mode", "m", "", "How to interpret command line arguments. In mirror mode all arguments are assumed to be files/dirs on the sending computer and they are mirrored onto the receiving computer. Files under the HOME directory are copied to the HOME directory on the receiving computer even if the HOME directory is different. In normal mode the last argument is assumed to be a destination path on the receiving computer. The last argument must be an existing directory unless copying a single file. When it is a directory it should end with a trailing slash.")
	transferCmd.Flags().StringP("permissions-bypass", "p", "", "The password to use to skip the transfer confirmation popup in kitty. Must match the password set for the file_transfer_confirmation_bypass option in kitty.conf. Note that leading and trailing whitespace is removed from the password. A password starting with ., / or ~ characters is assumed to be a file name to read the password from. A value of - means read the password from STDIN. A password that is purely a number less than 256 is assumed to be the number of a file descriptor from which to read the actual password.")
	transferCmd.Flags().BoolP("transmit-deltas", "x", false, "If a file on the receiving side already exists, use the rsync algorithm to update it to match the file on the sending side, potentially saving lots of bandwidth and also automatically resuming partial transfers. Note that this will actually degrade performance on fast links or with small files, so use with care.")

	carapace.Gen(transferCmd).FlagCompletion(carapace.ActionMap{
		"direction":       carapace.ActionValues("download", "receive", "send", "upload"),
		"mode":            carapace.ActionValues("normal", "mirror"),
		"compress":        carapace.ActionValues("auto", "always", "never"),
		"confirm-paths":   carapace.ActionValues("yes", "no"),
		"transmit-deltas": carapace.ActionValues("yes", "no"),
	})

	carapace.Gen(transferCmd).PositionalAnyCompletion(carapace.Batch(
		carapace.ActionFiles(),
		carapace.ActionDirectories(),
	).ToA())
}
