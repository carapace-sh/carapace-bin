package configs

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/config"
)

type gh struct {
  LabelCache time.Duration
}

var GH = gh {
  LabelCache: 24*time.Hour,
}

func (_gh gh) Completion() carapace.ActionMap {
  return carapace.ActionMap{
      "LabelCache": carapace.ActionValues("24m", "7d"),
  }
}

func init() {
  config.Register("gh", &GH)
}
