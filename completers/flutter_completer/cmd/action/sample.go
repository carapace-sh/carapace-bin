package action

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/rsteube/carapace"
)

type sample struct {
	Id          string
	Description string
}

func ActionSamples() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		tmpfile, err := ioutil.TempFile(os.TempDir(), "carapace-flutter-samples.json")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		defer os.Remove(tmpfile.Name())
		return carapace.ActionExecCommand("flutter", "create", "--suppress-analytics", "--list-samples", tmpfile.Name())(func(output []byte) carapace.Action {
			content, err := ioutil.ReadFile(tmpfile.Name())
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			var samples []sample
			if err := json.Unmarshal(content, &samples); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, sample := range samples {
				vals = append(vals, sample.Id, sample.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Invoke(c).ToA() // ensure this is executed before `defer os.Remove`
	}).Cache(1 * time.Hour)
}
