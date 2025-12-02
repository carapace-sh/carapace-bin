package fs

import "github.com/carapace-sh/carapace"

func actionBlockdevices(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("diskutil", "list", "-plist")(func(output []byte) carapace.Action {
		return carapace.ActionValues("TODO")
	})
}
