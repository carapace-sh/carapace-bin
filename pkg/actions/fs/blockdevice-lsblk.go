package fs

import (
	"encoding/json"
	"runtime"

	"github.com/carapace-sh/carapace"
)

func actionBlockdevicesLsblk(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if runtime.GOOS == "darwin" {
			return carapace.ActionValues()
		}
		return carapace.ActionExecCommand("lsblk", "--json", "-o", "KNAME,LABEL,PARTLABEL,PARTUUID,PATH,SIZE,PARTTYPENAME,TYPE,UUID")(func(output []byte) carapace.Action {
			var b struct {
				Blockdevices []blockdevice
			}
			if err := json.Unmarshal(output, &b); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return f(b.Blockdevices)
		})
	})
}
