package anchor

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/util"
	"github.com/pelletier/go-toml"
)

type anchorToml struct {
	Programs map[string]map[string]any `toml:"programs"`
	Scripts  map[string]string         `toml:"scripts"`
}

func readConfig(f func(cfg anchorToml) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path, err := util.FindReverse(c.Dir, "Anchor.toml")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		var cfg anchorToml
		if err := toml.Unmarshal(content, &cfg); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(cfg)
	})
}

// ActionPrograms completes workspace programs from Anchor.toml
//
//	escrow (7fkckTm12FmUhVtEeMHF7rijAYjmwEe5WhkUW8DgC9WG)
//	vault (Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS)
func ActionPrograms() carapace.Action {
	return readConfig(func(cfg anchorToml) carapace.Action {
		vals := make([]string, 0)
		for _, programs := range cfg.Programs {
			for name, deployment := range programs {
				// value is either the address string or a table with an address key
				switch d := deployment.(type) {
				case string:
					vals = append(vals, name, d)
				case map[string]any:
					address, _ := d["address"].(string)
					vals = append(vals, name, address)
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...).Unique()
	})
}

// ActionScripts completes scripts from Anchor.toml
//
//	test (yarn run ts-mocha -p ./tsconfig.json -t 1000000 tests/**/*.ts)
//	lint (cargo clippy)
func ActionScripts() carapace.Action {
	return readConfig(func(cfg anchorToml) carapace.Action {
		vals := make([]string, 0)
		for name, command := range cfg.Scripts {
			vals = append(vals, name, command)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
