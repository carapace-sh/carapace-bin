//go:build !darwin && !windows

package fs

import (
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/sandbox"
)

func TestBlockDevicesLsblk(t *testing.T) {
	var output = `{
   "blockdevices": [
      {
         "kname": "nvme0n1",
         "label": null,
         "partlabel": null,
         "partuuid": null,
         "path": "/dev/nvme0n1",
         "size": "350,4G",
         "parttypename": null,
         "type": "disk",
         "uuid": null
      },{
         "kname": "nvme0n1p1",
         "label": null,
         "partlabel": null,
         "partuuid": "111-1-1111-1-11111111",
         "path": "/dev/nvme0n1p1",
         "size": "150M",
         "parttypename": "EFI System",
         "type": "part",
         "uuid": "AAAA-AAAA"
      },{
         "kname": "nvme0n1p2",
         "label": null,
         "partlabel": "root",
         "partuuid": "22222222-2222-2222-2222-222222222222",
         "path": "/dev/nvme0n1p2",
         "size": "476,6G",
         "parttypename": "Linux filesystem",
         "type": "part",
         "uuid": "33333333-3333-3333-3333-333333333333"
      }
   ]
}`

	sandbox.Action(t, func() carapace.Action {
		return actionBlockdevices(func(blockdevices []blockdevice) carapace.Action {
			return carapace.ActionMessage("TODO")
		})
	})(func(s *sandbox.Sandbox) {
		s.Reply("lsblk", "--json", "-o", "KNAME,LABEL,PARTLABEL,PARTUUID,PATH,SIZE,PARTTYPENAME,TYPE,UUID").With(output)

	})
}
