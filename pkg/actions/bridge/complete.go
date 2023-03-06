package bridge

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionComplete bridges https://github.com/posener/complete
//
//	var rootCmd = &cobra.Command{
//		Use:                "vault",
//		Short:              "A tool for secrets management",
//		Long:               "https://www.vaultproject.io/",
//		Run:                func(cmd *cobra.Command, args []string) {},
//		DisableFlagParsing: true,
//	}
//
//	func Execute() error {
//		return rootCmd.Execute()
//	}
//	func init() {
//		carapace.Gen(rootCmd).Standalone()
//
//		carapace.Gen(rootCmd).PositionalAnyCompletion(
//			bridge.ActionComplete("vault"),
//		)
//	}
func ActionComplete(cmd string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		c.Setenv("COMP_LINE", fmt.Sprintf("%v %v %v", cmd, strings.Join(c.Args, " "), c.CallbackValue))
		return carapace.ActionExecCommand(cmd)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			a := carapace.ActionValues(lines[:len(lines)-1]...)
			for _, line := range lines[:len(lines)-1] {
				if len(line) > 0 && strings.ContainsAny(line[:len(line)-1], "/=@:.,") {
					a = a.NoSpace()
					break
				}
			}
			return a
		}).Invoke(c).ToA()
	})
}
