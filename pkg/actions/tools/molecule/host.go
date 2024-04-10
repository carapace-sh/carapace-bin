package molecule

import (
	"os"
	"path"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/carapace-sh/carapace/pkg/xdg"
	"gopkg.in/yaml.v3"
)

type InstanceOpts struct {
	Scenario string
	EnvFile  string
}

// ActionInstances completes known host names and known molecule instance hosts
func ActionInstances(opts InstanceOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		ephemeralPath := getMoleculeEphemeralDir(opts.EnvFile)
		roleName := path.Base(c.Dir)

		if opts.Scenario == "" {
			opts.Scenario = "default"
		}

		inventoryPath := path.Join(ephemeralPath, roleName, opts.Scenario, "inventory", "ansible_inventory.yml")
		inventoryData, err := os.ReadFile(inventoryPath)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var inventoryMap struct {
			all struct {
				hosts map[string]struct {
					ansibleHost string `yaml:"ansible_host"`
				}
			}
		}
		if err := yaml.Unmarshal(inventoryData, &inventoryMap); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		inventory := make([]string, 0)
		for host, hostData := range inventoryMap.all.hosts {
			inventory = append(inventory, hostData.ansibleHost, host)
		}

		return carapace.ActionValues(inventory...)
	})
}

func traverseEphemeralDir(envFilePath string) func(tc traverse.Context) (string, error) {
	return func(tc traverse.Context) (string, error) {
		return "", nil // TODO
	}
}

// Find the location molecule is storing ephemeral data
func getMoleculeEphemeralDir(envFilePath string) string {
	cacheDir, err := xdg.UserCacheDir()
	if err != nil {
		return "" // TODO error
	}
	ephemeralPath := path.Join(cacheDir, "molecule")

	// Check for specific molecule environment variable
	if moleculeEphemeral, ok := os.LookupEnv("MOLECULE_EPHEMERAL_DIRECTORY"); ok {
		ephemeralPath = moleculeEphemeral
	}

	// Check for path specified in molecule .env.yml file
	// GetString defaults to the correct value (.env.yml)
	if envFilePath == "" {
		envFilePath = ".env.yml"
	}

	envData, err := os.ReadFile(envFilePath)
	if err != nil {
		return ephemeralPath
	}

	envMap := make(map[string]any)
	if err := yaml.Unmarshal(envData, envMap); err != nil {
		return ephemeralPath
	}

	if moleculeEphemeral, ok := envMap["MOLECULE_EPHEMERAL_DIRECTORY"]; ok {
		if s, ok := moleculeEphemeral.(string); ok {
			ephemeralPath = s
		}
	}

	return ephemeralPath
}
