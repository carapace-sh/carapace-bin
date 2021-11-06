// package os contains number related actions
package number

import (
	"fmt"

	"github.com/rsteube/carapace"
)

// ActionRangeF completes a number range formatted with givien format specifier
//   ActionRangeF("%02d", 0, 59)
func ActionRangeF(format string, start int, end int) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)

		for i := start; i <= end; i = i + 1 {
			vals = append(vals, fmt.Sprintf(format, i))
		}
		return carapace.ActionValues(vals...)
	})
}
