// package os contains time related actions
package time

import (
	"fmt"
	"strconv"
	"time"

	"github.com/rsteube/carapace"
)

// ActionDate completes `yyyy-MM-dd` dates separately
//   2020-12-20
//   2014-05-02
func ActionDate() carapace.Action {
	return carapace.ActionMultiParts("-", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			from := time.Now().Year() - 100
			to := time.Now().Year()
			vals := make([]string, 0)
			for i := from; i <= to; i = i + 1 {
				vals = append(vals, strconv.Itoa(i))
			}
			return carapace.ActionValues(vals...).Invoke(c).Suffix("-").ToA()
		case 1:
			return ActionMonths().Invoke(c).Suffix("-").ToA()
		case 2:
			year, err := strconv.Atoi(c.Parts[0])
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			month, err := strconv.Atoi(c.Parts[1])
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return ActionDays(year, month)
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionMonths completes `MM` months
//   01 (Januar)
//   11 (November)
func ActionMonths() carapace.Action {
	return carapace.ActionValuesDescribed(
		"01", "Januar",
		"02", "Februar",
		"03", "March",
		"04", "April",
		"05", "May",
		"06", "June",
		"07", "July",
		"08", "August",
		"09", "September",
		"10", "Oktober",
		"11", "November",
		"12", "December",
	)
}

// ActionDays completes `dd` days for a month
//   01
//   28
func ActionDays(year int, month int) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		month = month + 1
		if month > 12 {
			year = year + 1
			month = month - 12
		}
		date := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)

		vals := make([]string, 0)
		for i := 1; i <= date.Day(); i = i + 1 {
			vals = append(vals, fmt.Sprintf("%02d", i))
		}
		return carapace.ActionValues(vals...)
	})
}
