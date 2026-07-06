package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
)

// ActionAioModes completes QEMU AIO modes
func ActionAioModes() carapace.Action { return qemu.ActionAioModes() }

// ActionCacheModes completes QEMU cache modes
func ActionCacheModes() carapace.Action { return qemu.ActionCacheModes() }

// ActionImageFormats completes supported QEMU image formats
func ActionImageFormats() carapace.Action { return qemu.ActionImageFormats() }
