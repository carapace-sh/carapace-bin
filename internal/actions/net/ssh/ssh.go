package ssh

import "github.com/carapace-sh/carapace"

// ActionHosts is a circular dependency workaround
var ActionHosts func() carapace.Action
