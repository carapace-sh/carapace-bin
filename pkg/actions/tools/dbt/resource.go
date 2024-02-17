package dbt

import (
	"encoding/json"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache/key"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/rsteube/carapace/pkg/util"
)

type ResourceOpts struct {
	Analysis bool
	Exposure bool
	Metric   bool
	Model    bool
	Seed     bool
	Snapshot bool
	Source   bool
	Test     bool
}

func (o ResourceOpts) Default() ResourceOpts {
	o.Analysis = true
	o.Exposure = true
	o.Metric = true
	o.Model = true
	o.Seed = true
	o.Snapshot = true
	o.Source = true
	o.Test = true
	return o
}

func (o ResourceOpts) key() key.Key {
	return key.String(
		strconv.FormatBool(o.Analysis),
		strconv.FormatBool(o.Exposure),
		strconv.FormatBool(o.Metric),
		strconv.FormatBool(o.Model),
		strconv.FormatBool(o.Seed),
		strconv.FormatBool(o.Snapshot),
		strconv.FormatBool(o.Source),
		strconv.FormatBool(o.Test),
	)
}

func (o ResourceOpts) includes(t string) bool {
	switch t {
	case "analysis":
		return o.Analysis
	case "exposure":
		return o.Exposure
	case "metric":
		return o.Metric
	case "model":
		return o.Model
	case "seed":
		return o.Seed
	case "snapshot":
		return o.Snapshot
	case "source":
		return o.Source
	case "test":
		return o.Test
	default:
		return false
	}
}

type resource struct {
	Name         string
	ResourceType string `json:"resource_type"`
}

func (r resource) style() string {
	i := map[string]int{
		"analysis": 0,
		"exposure": 1,
		"metric":   2,
		"model":    3,
		"seed":     4,
		"snapshot": 5,
		"source":   6,
		"test":     7,
	}
	return style.Carapace.Highlight(i[r.ResourceType])
}

// ActionResources completes resources
//
//	redshift_admin_dependencies
//	redshift_columns
func ActionResources(opts ResourceOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path, err := util.FindReverse(c.Dir, "dbt_project.yml")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		path = filepath.Dir(path) + "/target/manifest.json"

		return carapace.ActionExecCommand("dbt", "list", "--output", "json")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				var r resource
				if err := json.Unmarshal([]byte(line), &r); err != nil {
					return carapace.ActionMessage(err.Error())
				}

				if opts.includes(r.ResourceType) {
					vals = append(vals, r.Name, r.ResourceType, r.style())
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		}).Cache(24*time.Hour, opts.key(), key.FileStats(path)) // TODO opts , manifest.json constantly gets new modification time??
	}).Tag("resources")
}
