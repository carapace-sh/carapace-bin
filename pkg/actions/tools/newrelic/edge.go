package newrelic

import (
	"encoding/json"
	"strconv"

	"github.com/rsteube/carapace"
)

type traceObserver struct {
	Id   int
	Name string
}

// ActionTraceObservers completes trace observers
//
//	1 (observer1)
//	2 (observer2)
func ActionTraceObservers(profile string) carapace.Action {
	return carapace.ActionCallback(func(carapace.Context) carapace.Action {
		args := []string{"edge", "trace-observer", "list"}
		if profile != "" {
			args = append(args, profile)
		}

		return carapace.ActionExecCommand("newrelic", args...)(func(output []byte) carapace.Action {
			var observers []traceObserver
			if err := json.Unmarshal(output, &observers); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, observer := range observers {
				vals = append(vals, strconv.Itoa(observer.Id), observer.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
