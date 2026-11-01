package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "Evaluate conditional expression",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-test",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValuesDescribed(
			"!=", "strings are not equal",
			"-G", "file is effectively owned by your group",
			"-L", "file is a symbolic link",
			"-N", "file has been modified since it was last read",
			"-O", "file is effectively owned by you",
			"-R", "shell variable is set and is a name reference",
			"-S", "file is a socket",
			"-a", "file exists",
			"-b", "file is block special",
			"-c", "file is character special",
			"-d", "file is a directory",
			"-e", "file exists",
			"-ef", "files are hard links to the same file",
			"-eq", "integers are equal",
			"-f", "file exists and is a regular file",
			"-g", "file is set-group-id",
			"-ge", "first integer is greater than or equal to second",
			"-gt", "first integer is greater than second",
			"-h", "file is a symbolic link",
			"-k", "file has its sticky bit set",
			"-le", "first integer is less than or equal to second",
			"-lt", "first integer is less than second",
			"-n", "string is not empty",
			"-ne", "integers are not equal",
			"-nt", "first file is newer than second",
			"-o", "shell option is enabled",
			"-ot", "first file is older than second",
			"-p", "file is a named pipe",
			"-r", "file is readable",
			"-s", "file exists and is not empty",
			"-t", "file descriptor is opened on a terminal",
			"-u", "file is set-user-id",
			"-v", "shell variable is set",
			"-w", "file is writable",
			"-x", "file is executable",
			"-z", "string is empty",
			"=", "strings are equal",
		),
	)
}
