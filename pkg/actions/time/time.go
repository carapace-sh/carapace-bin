// package os contains time related actions
package time

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/number"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionDate completes `yyyy-MM-dd` dates separately
//
//	2020-12-20
//	2014-05-02
func ActionDate() carapace.Action {
	return carapace.ActionMultiParts("-", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			from := time.Now().Year() - 5
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
			return ActionDays(DaysOpts{Year: year, Month: month})
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionMonths completes `MM` months
//
//	01 (Januar)
//	11 (November)
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

type DaysOpts struct {
	Year  int
	Month int
}

// ActionDays completes `dd` days for a month
//
//	01 (Friday)
//	28 (Thursday)
func ActionDays(opts DaysOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		date := time.Date(opts.Year, time.Month(opts.Month)+1, 0, 0, 0, 0, 0, time.UTC)

		vals := make([]string, 0)
		for i := 1; i <= date.Day(); i = i + 1 {
			weekday := time.Date(opts.Year, time.Month(opts.Month), i, 0, 0, 0, 0, time.UTC).Weekday()
			if weekday == time.Saturday || weekday == time.Sunday {
				vals = append(vals, fmt.Sprintf("%02d", i), weekday.String(), style.Blue)
			} else {
				vals = append(vals, fmt.Sprintf("%02d", i), weekday.String(), style.Default)
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}

// ActionTime completes `hh:mm` time
//
//	00:30
//	14:47
func ActionTime() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return number.ActionRange(number.RangeOpts{Format: "%02d", Start: 0, End: 23}).Invoke(c).Suffix(":").ToA()
		case 1:
			return number.ActionRange(number.RangeOpts{Format: "%02d", Start: 0, End: 59})
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionTimeS completes `hh:mm:ss` time
//
//	00:30:19
//	14:47:14
func ActionTimeS() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return number.ActionRange(number.RangeOpts{Format: "%02d", Start: 0, End: 23}).Invoke(c).Suffix(":").ToA()
		case 1:
			return number.ActionRange(number.RangeOpts{Format: "%02d", Start: 0, End: 59}).Invoke(c).Suffix(":").ToA()
		case 2:
			return number.ActionRange(number.RangeOpts{Format: "%02d", Start: 0, End: 59})
		default:
			return carapace.ActionValues()
		}
	})
}

type DateTimeOpts struct {
	Strict bool
}

// ActionDateTime completes `yyyy-MM-dd hh:mm:ss` datetime
//
//	2021-11-11 04:02:12
//	2021-04-02 16:11:33
func ActionDateTime(opts DateTimeOpts) carapace.Action { // TODO might be best to just accept a pattern as in go lib here
	delimiter := " "
	if opts.Strict {
		delimiter = "T"
	}
	return carapace.ActionMultiParts(delimiter, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			if strings.Count(c.Value, "-") == 2 {
				return ActionDate().Invoke(c).Suffix(delimiter).ToA()
			}
			return ActionDate()
		case 1:
			return ActionTimeS()
		default:
			return carapace.ActionValues()
		}
	})
}
