// package posenercomplete contains actions related to posener/complete
package posenercomplete

import (
	"fmt"
	"os"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionPosenerComplete bridges posener/complete
//   var rootCmd = &cobra.Command{
//   	Use:                "vault",
//   	Short:              "A tool for secrets management",
//   	Long:               "https://www.vaultproject.io/",
//   	Run:                func(cmd *cobra.Command, args []string) {},
//   	DisableFlagParsing: true,
//   }
//
//   func Execute() error {
//   	return rootCmd.Execute()
//   }
//   func init() {
//   	carapace.Gen(rootCmd).Standalone()
//
//   	carapace.Gen(rootCmd).PositionalAnyCompletion(
//   		posenercomplete.ActionPosenerComplete("vault"),
//   	)
//   }
func ActionPosenerComplete(cmd string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		os.Setenv("COMP_LINE", fmt.Sprintf("%v %v %v", cmd, strings.Join(c.Args, " "), c.CallbackValue))
		return carapace.ActionExecCommand(cmd)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
