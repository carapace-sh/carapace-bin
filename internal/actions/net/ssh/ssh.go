package ssh

import "github.com/rsteube/carapace"

// ActionHosts is a circular dependency workaround
var ActionHosts func() carapace.Action
