// package os contains number related actions
package number

import (
	"fmt"

	"github.com/carapace-sh/carapace"
)

type RangeOpts struct {
	Format string
	Start  int
	End    int
}

// ActionRange completes a number range formatted with given (optional) format specifier
//
//	ActionRange(RangeOpts{Format: "%02d", Start: 0, End: 59})
func ActionRange(opts RangeOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if opts.Format == "" {
			opts.Format = "%d"
		}

		vals := make([]string, 0)
		for i := opts.Start; i <= opts.End; i = i + 1 {
			vals = append(vals, fmt.Sprintf(opts.Format, i))
		}
		return carapace.ActionValues(vals...)
	})
}
