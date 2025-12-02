//go:build !darwin && !windows

package fs

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

func actionBlockdevices(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
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
