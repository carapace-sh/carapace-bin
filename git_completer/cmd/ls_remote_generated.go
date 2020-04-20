package cmd

import (
	"github.com/spf13/cobra"
)

var ls_remoteCmd = &cobra.Command{
	Use: "ls-remote",
	Short: "List references in a remote repository",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
	ls_remoteCmd.Flags().Bool("exit-code", false, "exit with exit code 2 if no matching refs are found")
	ls_remoteCmd.Flags().Bool("get-url", false, "take url.<base>.insteadOf into account")
	ls_remoteCmd.Flags().BoolP("heads", "h", false, "limit to heads")
	ls_remoteCmd.Flags().BoolP("server-option", "o", false, "<server-specific>    option to transmit")
	ls_remoteCmd.Flags().BoolP("quiet", "q", false, "do not print remote URL")
	ls_remoteCmd.Flags().Bool("refs", false, "do not show peeled tags")
	ls_remoteCmd.Flags().String("sort", "", "field name to sort on")
	ls_remoteCmd.Flags().Bool("symref", false, "show underlying ref in addition to the object pointed by it")
	ls_remoteCmd.Flags().BoolP("tags", "t", false, "limit to tags")
	ls_remoteCmd.Flags().String("upload-pack", "", "path of git-upload-pack on the remote host")
    rootCmd.AddCommand(ls_remoteCmd)
}
