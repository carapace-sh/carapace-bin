package action

import (
	"os"
	"path"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// ActionInstances completes known host names and known molecule instance hosts
func ActionInstances(cmd *cobra.Command) carapace.Action {
	return carapace.Batch(
		net.ActionHosts(),
		carapace.ActionValuesDescribed(getInstances(cmd)...).Cache(time.Minute),
	).ToA()
}

// Get active instances
func getInstances(cmd *cobra.Command) []string {
	ephemeralPath := getMoleculeEphemeralDir(cmd)

	cwd, err := os.Getwd()
	if err != nil {
		return []string{}
	}
	roleName := path.Base(cwd)

	scenarioName, err := cmd.Flags().GetString("scenario-name")
	if err != nil {
		scenarioName = "default"
	}

	inventoryPath := path.Join(ephemeralPath, roleName, scenarioName, "inventory", "ansible_inventory.yml")
	inventoryData, err := os.ReadFile(inventoryPath)
	if err != nil {
		return []string{}
	}

	var inventoryMap struct {
		all struct {
			hosts map[string]struct {
				ansibleHost string `yaml:"ansible_host"`
			}
		}
	}
	if err := yaml.Unmarshal(inventoryData, &inventoryMap); err != nil {
		return []string{}
	}

	inventory := make([]string, 0)
	for host, hostData := range inventoryMap.all.hosts {
		inventory = append(inventory, hostData.ansibleHost, host)
	}

	return inventory
}

// Find the location molecule is storing ephemeral data
func getMoleculeEphemeralDir(cmd *cobra.Command) string {
	// Default to the platform equivalent to "$HOME" joined with .cache/molecule"
	// (molecule does not support windows)
	ephemeralPath, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	ephemeralPath = path.Join(ephemeralPath, ".cache", "molecule")

	// Check for more specific cache directory environment variable
	if cacheHome, ok := os.LookupEnv("XDG_CACHE_HOME"); ok {
		ephemeralPath = path.Join(cacheHome, "molecule")
	}

	// Check for specific molecule environment variable
	if moleculeEphemeral, ok := os.LookupEnv("MOLECULE_EPHEMERAL_DIRECTORY"); ok {
		ephemeralPath = moleculeEphemeral
	}

	// Check for path specified in molecule .env.yml file
	// GetString defaults to the correct value (.env.yml)
	envFilePath, err := cmd.Flags().GetString("env-file")
	if err != nil {
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
