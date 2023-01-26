package bridge

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionYargs bridges https://github.com/yargs/yargs
//
//	var rootCmd = &cobra.Command{
//		Use:                "ng",
//		Short:              "CLI tool for Angular",
//		Run:                func(cmd *cobra.Command, args []string) {},
//		DisableFlagParsing: true,
//	}
//
//	func Execute() error {
//		return rootCmd.Execute()
//	}
//
//	func init() {
//		carapace.Gen(rootCmd).Standalone()
//
//		carapace.Gen(rootCmd).PositionalAnyCompletion(
//			bridge.ActionYargs("ng"),
//		)
//	}
func ActionYargs(command string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if _, err := exec.LookPath(command); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		c.Setenv("SHELL", "zsh")

		args := []string{"--get-yargs-completions"}
		args = append(args, c.Args...)
		args = append(args, c.CallbackValue)
		return carapace.ActionExecCommand(command, args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines {
				if line != "" {
					splitted := strings.SplitN(line, ":", 2)
					name := splitted[0]
					description := ""
					if len(splitted) > 1 {
						description = splitted[1]
					}
					vals = append(vals, name, description)
				}
			}
			if len(vals) == 0 {
				return carapace.ActionFiles()
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Invoke(c).ToA()
	})
}
