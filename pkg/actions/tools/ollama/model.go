package ollama

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionModels completes models
func ActionModels() carapace.Action {
	return carapace.ActionExecCommand("ollama", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[1:] {
			if splitted := strings.Split(line, "\t"); len(splitted) >= 4 {
				vals = append(vals, splitted[0], fmt.Sprintf("%v - %v", splitted[2], splitted[3]))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
