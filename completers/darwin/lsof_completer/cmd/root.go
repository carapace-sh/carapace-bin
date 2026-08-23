package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsof",
	Short: "list open files",
	Long:  "https://man.freebsd.org/cgi/man.cgi?lsof",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "help")
	rootCmd.Flags().BoolS("C", "C", false, "disables reporting of path name components from kernel's name cache")
	rootCmd.Flags().BoolS("F", "F", false, "select output fields")
	rootCmd.Flags().BoolS("L", "L", false, "listing of submitted lock leases")
	rootCmd.Flags().BoolS("M", "M", false, "disable portmap registration")
	rootCmd.Flags().BoolS("N", "N", false, "select NFS files")
	rootCmd.Flags().BoolS("P", "P", false, "inhibit conversion of port numbers to port names")
	rootCmd.Flags().BoolS("R", "R", false, "list parent process identification")
	rootCmd.Flags().BoolS("T", "T", false, "disable TCP/TPI information")
	rootCmd.Flags().BoolS("U", "U", false, "select UNIX domain socket files")
	rootCmd.Flags().BoolS("V", "V", false, "indicate options selected")
	rootCmd.Flags().BoolS("a", "a", false, "AND the selections")
	rootCmd.Flags().BoolS("b", "b", false, "avoid kernel functions that might block")
	rootCmd.Flags().BoolS("c", "c", false, "select the listing of files for processes executing the command")
	rootCmd.Flags().BoolS("d", "d", false, "specify file descriptors")
	rootCmd.Flags().BoolS("g", "g", false, "select process group IDs")
	rootCmd.Flags().BoolS("h", "h", false, "help")
	rootCmd.Flags().StringS("i", "i", "", "select the listing of files for processes executing the command")
	rootCmd.Flags().BoolS("l", "l", false, "inhibit conversion of user ID numbers to login names")
	rootCmd.Flags().BoolS("n", "n", false, "inhibit conversion of network numbers to host names")
	rootCmd.Flags().StringS("o", "o", "", "specify the number of decimal digits")
	rootCmd.Flags().StringS("p", "p", "", "select the listing of files for processes whose process ID numbers are in s")
	rootCmd.Flags().StringS("s", "s", "", "select the listing of files for processes whose process ID numbers are in s")
	rootCmd.Flags().BoolS("t", "t", false, "produces a result in terse format")
	rootCmd.Flags().StringS("u", "u", "", "select the listing of files for processes whose user ID numbers or login names are in s")
	rootCmd.Flags().BoolS("v", "v", false, "version")
	rootCmd.Flags().BoolS("w", "w", false, "suppress warning messages")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"i": carapace.ActionValues("4", "6", "TCP", "UDP"),
		"u": carapace.ActionValues(),
	})
}
