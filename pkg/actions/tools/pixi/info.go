package pixi

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type pixiInfo struct {
	EnvironmentsInfo []environmentInfo `json:"environments_info"`
}

type environmentInfo struct {
	Name         string   `json:"name"`
	Features     []string `json:"features"`
	Dependencies []string `json:"dependencies"`
	Platforms    []string `json:"platforms"`
	Tasks        []string `json:"tasks"`
	Channels     []string `json:"channels"`
}

func actionExecPixi(arg ...string) func(f func(output []byte) carapace.Action) carapace.Action {
	return func(f func(output []byte) carapace.Action) carapace.Action {
		return carapace.ActionExecCommand("pixi", arg...)(func(output []byte) carapace.Action {
			return f(output)
		})
	}
}

func parseInfo(output []byte) (pixiInfo, error) {
	var info pixiInfo
	err := json.Unmarshal(output, &info)
	return info, err
}
