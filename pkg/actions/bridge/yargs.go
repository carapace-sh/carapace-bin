package bridge

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionYargs bridges https://github.com/yargs/yargs
//
//	var rootCmd = &cobra.Command{
//		Use:                "example",
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
//			bridge.ActionYargs("example"),
//		)
//	}
func ActionYargs(command string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if _, err := exec.LookPath(command); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		args := c.Args
		current := c.CallbackValue

		compLine := strings.Join(append(args, current), " ") // TODO escape/quote special characters
		return carapace.ActionExecCommand(command, "--get-yargs-completions", compLine)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines {
				if line != "" {
					line = strings.Replace(line, `\:`, "\001", 1)
					splitted := strings.SplitN(line, ":", 2)
					name := strings.Replace(splitted[0], "\001", `\:`, 1)
					description := ""
					if len(splitted) > 1 {
						description = strings.Replace(splitted[1], "\001", `\:`, 1)
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
