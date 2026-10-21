package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

// ActionBeams completes beams
//
//	gentle-river (alice)
//	1234-5678 (gentle-river)
func ActionBeams() carapace.Action {
	return carapace.ActionExecCommand("tsh", "beams", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var beams []struct {
			ID    string `json:"id"`
			UUID  string `json:"uuid"`
			Owner string `json:"owner"`
		}

		if err := json.Unmarshal(output, &beams); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, beam := range beams {
			vals = append(vals, beam.ID, beam.Owner)
			vals = append(vals, beam.UUID, beam.ID)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
