package action

import (
	"bytes"
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"go.yaml.in/yaml/v3"
)

// ActionClouds completes cloud names from the os-client-config files.
func ActionClouds() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		paths := []string{}
		if path := os.Getenv("OS_CLIENT_CONFIG_FILE"); path != "" {
			paths = append(paths, path)
		}
		if configDir, err := os.UserConfigDir(); err == nil {
			paths = append(paths,
				filepath.Join(configDir, "openstack", "clouds.yaml"),
				filepath.Join(configDir, "openstack", "clouds-public.yaml"),
			)
		}
		if homeDir, err := os.UserHomeDir(); err == nil {
			paths = append(paths, filepath.Join(homeDir, ".openstack", "clouds.yaml"))
		}
		paths = append(paths,
			"/etc/openstack/clouds.yaml",
			"/etc/openstack/clouds-public.yaml",
		)

		values := make([]string, 0)
		for _, path := range paths {
			content, err := os.ReadFile(path)
			if err != nil {
				continue
			}
			var config struct {
				Clouds map[string]any `yaml:"clouds"`
			}
			if err := yaml.Unmarshal(content, &config); err != nil {
				continue
			}
			for name := range config.Clouds {
				values = append(values, name)
			}
		}

		return carapace.ActionValues(values...).Unique()
	})
}

func actionNamed(command ...string) carapace.Action {
	return actionList(command, "ID", "Name")
}

func actionID(command ...string) carapace.Action {
	return actionList(command, "ID")
}

var globalEnv func() []string

// SetGlobalEnv registers a provider for environment variables derived from global flags.
func SetGlobalEnv(f func() []string) {
	globalEnv = f
}

func actionList(command []string, columns ...string) carapace.Action {
	return actionListArgs(append(append([]string{}, command...), "list"), columns...)
}

func actionListArgs(command []string, columns ...string) carapace.Action {
	args := append([]string{}, command...)
	args = append(args, "-f", "csv")
	for _, column := range columns {
		args = append(args, "-c", column)
	}

	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if globalEnv != nil {
			if overrides := globalEnv(); len(overrides) > 0 {
				env := c.Env
				if len(env) == 0 {
					env = os.Environ()
				}

				keys := make(map[string]struct{}, len(overrides))
				for _, entry := range overrides {
					key, _, _ := strings.Cut(entry, "=")
					keys[key] = struct{}{}
				}

				filtered := make([]string, 0, len(env)+len(overrides))
				for _, entry := range env {
					key, _, _ := strings.Cut(entry, "=")
					if _, ok := keys[key]; !ok {
						filtered = append(filtered, entry)
					}
				}
				c.Env = append(filtered, overrides...)
			}
		}

		return carapace.ActionExecCommandE("openstack", args...)(func(output []byte, err error) carapace.Action {
			if err != nil {
				return carapace.ActionValues()
			}

			reader := csv.NewReader(bytes.NewReader(output))
			reader.FieldsPerRecord = -1
			records, err := reader.ReadAll()
			if err != nil || len(records) <= 1 {
				return carapace.ActionValues()
			}

			values := make([]string, 0, (len(records)-1)*2)
			for _, record := range records[1:] {
				if len(record) == 0 || record[0] == "" {
					continue
				}
				switch {
				case len(record) > 1 && record[1] != "":
					values = append(values, record[1], record[0])
				default:
					values = append(values, record[0])
				}
			}
			return carapace.ActionValuesDescribed(values...).Unique()
		}).Invoke(c).ToA()
	})
}

func ActionAggregates() carapace.Action {
	return actionNamed("aggregate")
}

func ActionApplicationCredentials() carapace.Action {
	return actionNamed("application", "credential")
}

func ActionAvailabilityZones() carapace.Action {
	return actionList([]string{"availability", "zone"}, "Zone Name")
}

func ActionComputeAgents() carapace.Action {
	return actionList([]string{"compute", "agent"}, "Agent ID")
}

func ActionCredentials() carapace.Action {
	return actionID("credential")
}

func ActionDomains() carapace.Action {
	return actionNamed("domain")
}

func ActionEndpoints() carapace.Action {
	return actionID("endpoint")
}

func ActionFlavors() carapace.Action {
	return actionNamed("flavor")
}

func ActionGroups() carapace.Action {
	return actionNamed("group")
}

func ActionHypervisors() carapace.Action {
	return actionList([]string{"hypervisor"}, "ID", "Hypervisor Hostname")
}

func ActionIdentityProviders() carapace.Action {
	return actionID("identity", "provider")
}

func ActionImages() carapace.Action {
	return actionNamed("image")
}

func ActionKeypairs() carapace.Action {
	return actionList([]string{"keypair"}, "Name")
}

func ActionLimits() carapace.Action {
	return actionID("limit")
}

func ActionMappings() carapace.Action {
	return actionID("mapping")
}

func ActionProjects() carapace.Action {
	return actionNamed("project")
}

func ActionRegions() carapace.Action {
	return actionList([]string{"region"}, "Region")
}

func ActionRegisteredLimits() carapace.Action {
	return actionID("registered", "limit")
}

func ActionRoles() carapace.Action {
	return actionNamed("role")
}

func ActionServerGroups() carapace.Action {
	return actionNamed("server", "group")
}

func ActionServers() carapace.Action {
	return actionNamed("server")
}

func ActionServices() carapace.Action {
	return actionNamed("service")
}

func ActionTrusts() carapace.Action {
	return actionID("trust")
}

func ActionUsers() carapace.Action {
	return actionNamed("user")
}

func ActionVolumeAttachments() carapace.Action {
	return actionID("volume", "attachment")
}

func ActionVolumeBackups() carapace.Action {
	return actionNamed("volume", "backup")
}

func ActionVolumeGroupSnapshots() carapace.Action {
	return actionNamed("volume", "group", "snapshot")
}

func ActionVolumeGroups() carapace.Action {
	return actionNamed("volume", "group")
}

func ActionVolumeQosSpecs() carapace.Action {
	return actionNamed("volume", "qos")
}

func ActionVolumeSnapshots() carapace.Action {
	return actionNamed("volume", "snapshot")
}

func ActionVolumeTransfers() carapace.Action {
	return actionNamed("volume", "transfer", "request")
}

func ActionVolumeTypes() carapace.Action {
	return actionNamed("volume", "type")
}

func ActionVolumes() carapace.Action {
	return actionNamed("volume")
}

func ActionShareAccessRules() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) == 0 {
			return carapace.ActionValues()
		}
		return actionListArgs([]string{"share", "access", "list", c.Args[0]}, "ID")
	})
}

func ActionShareBackups() carapace.Action {
	return actionNamed("share", "backup")
}

func ActionShareGroupSnapshots() carapace.Action {
	return actionNamed("share", "group", "snapshot")
}

func ActionShareGroups() carapace.Action {
	return actionNamed("share", "group")
}

func ActionShareInstances() carapace.Action {
	return actionID("share", "instance")
}

func ActionShareMessages() carapace.Action {
	return actionID("share", "message")
}

func ActionShareNetworks() carapace.Action {
	return actionNamed("share", "network")
}

func ActionShareQosTypes() carapace.Action {
	return actionNamed("share", "qos", "type")
}

func ActionShareReplicas() carapace.Action {
	return actionNamed("share", "replica")
}

func ActionShareSecurityServices() carapace.Action {
	return actionNamed("share", "security", "service")
}

func ActionShareServers() carapace.Action {
	return actionID("share", "server")
}

func ActionShareSnapshots() carapace.Action {
	return actionNamed("share", "snapshot")
}

func ActionShareTypes() carapace.Action {
	return actionNamed("share", "type")
}

func ActionShares() carapace.Action {
	return actionNamed("share")
}

func ActionAddressGroups() carapace.Action {
	return actionNamed("address", "group")
}

func ActionAddressScopes() carapace.Action {
	return actionNamed("address", "scope")
}

func ActionFloatingIPs() carapace.Action {
	return actionList([]string{"floating", "ip"}, "ID", "Floating IP Address")
}

func ActionLocalIPs() carapace.Action {
	return actionNamed("local", "ip")
}

func ActionNetworkAgents() carapace.Action {
	return actionID("network", "agent")
}

func ActionNetworkFlavors() carapace.Action {
	return actionNamed("network", "flavor")
}

func ActionNetworkQosPolicies() carapace.Action {
	return actionNamed("network", "qos", "policy")
}

func ActionNetworkRbacPolicies() carapace.Action {
	return actionID("network", "rbac")
}

func ActionNetworkSegments() carapace.Action {
	return actionNamed("network", "segment")
}

func ActionNetworkTrunks() carapace.Action {
	return actionNamed("network", "trunk")
}

func ActionNetworks() carapace.Action {
	return actionNamed("network")
}

func ActionPorts() carapace.Action {
	return actionNamed("port")
}

func ActionRouters() carapace.Action {
	return actionNamed("router")
}

func ActionSecurityGroupRules() carapace.Action {
	return actionID("security", "group", "rule")
}

func ActionSecurityGroups() carapace.Action {
	return actionNamed("security", "group")
}

func ActionSubnetPools() carapace.Action {
	return actionNamed("subnet", "pool")
}

func ActionSubnets() carapace.Action {
	return actionNamed("subnet")
}

func ActionFirewallGroups() carapace.Action {
	return actionNamed("firewall", "group")
}

func ActionFirewallPolicies() carapace.Action {
	return actionNamed("firewall", "group", "policy")
}

func ActionFirewallRules() carapace.Action {
	return actionNamed("firewall", "group", "rule")
}

func ActionBgpSpeakers() carapace.Action {
	return actionNamed("bgp", "speaker")
}

func ActionBgpPeers() carapace.Action {
	return actionNamed("bgp", "peer")
}

func ActionBgpvpns() carapace.Action {
	return actionNamed("bgpvpn")
}

func ActionAcceleratorArqs() carapace.Action {
	return actionList([]string{"accelerator", "arq"}, "uuid")
}

func ActionAcceleratorDeployables() carapace.Action {
	return actionList([]string{"accelerator", "deployable"}, "uuid", "name")
}

func ActionAcceleratorDevices() carapace.Action {
	return actionList([]string{"accelerator", "device"}, "uuid")
}

func ActionAcceleratorDeviceProfiles() carapace.Action {
	return actionList([]string{"accelerator", "device", "profile"}, "uuid", "name")
}

func ActionContainers() carapace.Action {
	return actionList([]string{"container"}, "Name")
}

func ActionCachedImages() carapace.Action {
	return actionList([]string{"cached", "image"}, "image_id")
}

func ActionConsistencyGroups() carapace.Action {
	return actionNamed("consistency", "group")
}

func ActionConsistencyGroupSnapshots() carapace.Action {
	return actionNamed("consistency", "group", "snapshot")
}

func ActionExtensions() carapace.Action {
	return actionList([]string{"extension"}, "Name")
}

func ActionImageMetadefNamespaces() carapace.Action {
	return actionList([]string{"image", "metadef", "namespace"}, "Namespace")
}

func ActionImageTasks() carapace.Action {
	return actionID("image", "task")
}

func ActionNetworkMeters() carapace.Action {
	return actionNamed("network", "meter")
}

func ActionNetworkSegmentRanges() carapace.Action {
	return actionNamed("network", "segment", "range")
}

func ActionShareGroupTypes() carapace.Action {
	return actionNamed("share", "group", "type")
}

func ActionShareLocks() carapace.Action {
	return actionID("share", "lock")
}

func ActionShareTransfers() carapace.Action {
	return actionNamed("share", "transfer")
}

func ActionVolumeGroupTypes() carapace.Action {
	return actionNamed("volume", "group", "type")
}
