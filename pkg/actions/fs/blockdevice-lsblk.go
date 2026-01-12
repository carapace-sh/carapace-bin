package fs

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type blockdevice struct {
	Kname        string `json:"Kname"`
	Label        string `json:"Label"`
	Partlabel    string `json:"Partlabel"`
	Partuuid     string `json:"Partuuid"`
	Parttypename string `json:"Parttypename"`
	Path         string `json:"Path"`
	Size         string `json:"Size"`
	Type         string `json:"Type"`
	Uuid         string `json:"Uuid"`
}

func actionBlockdevicesLsblk(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("lsblk", "--json", "-o", "KNAME,LABEL,PARTLABEL,PARTUUID,PATH,SIZE,PARTTYPENAME,TYPE,UUID")(func(output []byte) carapace.Action {
		var b struct {
			Blockdevices []blockdevice
		}
		if err := json.Unmarshal(output, &b); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(b.Blockdevices)
	})
}
