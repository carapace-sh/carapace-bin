package action

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"go.yaml.in/yaml/v3"
)

// ActionClouds completes cloud names from the os-client-config files.
func ActionClouds() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		values := make([]string, 0)
		for _, path := range cloudConfigPaths() {
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

func ActionCloudAuthURLs() carapace.Action {
	return cloudValues("auth_url")
}

func ActionCloudDomainIDs() carapace.Action {
	return cloudValues("domain_id", "user_domain_id", "project_domain_id")
}

func ActionCloudDomainNames() carapace.Action {
	return cloudValues("domain_name", "user_domain_name", "project_domain_name")
}

func ActionCloudProjectIDs() carapace.Action {
	return cloudValues("project_id")
}

func ActionCloudProjectNames() carapace.Action {
	return cloudValues("project_name")
}

func ActionCloudRegions() carapace.Action {
	return cloudValues("region_name")
}

func ActionCloudUserIDs() carapace.Action {
	return cloudValues("user_id")
}

func ActionCloudUsernames() carapace.Action {
	return cloudValues("username")
}

func cloudConfigPaths() []string {
	paths := []string{}
	if path := os.Getenv("OS_CLIENT_CONFIG_FILE"); path != "" {
		paths = append(paths, path)
	}

	dirs := []string{}
	if configDir, err := os.UserConfigDir(); err == nil {
		dirs = append(dirs, filepath.Join(configDir, "openstack"))
	}
	if homeDir, err := os.UserHomeDir(); err == nil {
		dirs = append(dirs, filepath.Join(homeDir, ".openstack"))
	}
	if wd, err := os.Getwd(); err == nil {
		dirs = append(dirs, wd)
	}
	dirs = append(dirs, "/etc/openstack", "/etc/xdg/openstack")

	for _, dir := range dirs {
		for _, name := range []string{
			"clouds.yaml", "clouds.yml", "clouds.json",
			"clouds-public.yaml", "clouds-public.yml", "clouds-public.json",
		} {
			paths = append(paths, filepath.Join(dir, name))
		}
	}
	return paths
}

func cloudValues(keys ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		values := make([]string, 0)
		for _, path := range cloudConfigPaths() {
			content, err := os.ReadFile(path)
			if err != nil {
				continue
			}
			var config struct {
				Clouds map[string]map[string]any `yaml:"clouds"`
			}
			if err := yaml.Unmarshal(content, &config); err != nil {
				continue
			}

			selected := ""
			if globalCloud != nil {
				selected = globalCloud()
			}

			for name, cloud := range config.Clouds {
				if selected != "" && name != selected {
					continue
				}
				for _, key := range keys {
					if value, ok := cloudString(cloud[key]); ok {
						values = append(values, value)
					}
				}
				if auth, ok := cloud["auth"].(map[string]any); ok {
					for _, key := range keys {
						if value, ok := cloudString(auth[key]); ok {
							values = append(values, value)
						}
					}
				}
			}
		}

		return carapace.ActionValues(values...).Unique()
	})
}

func cloudString(value any) (string, bool) {
	switch v := value.(type) {
	case string:
		return v, v != ""
	case int, int64, float64, bool:
		return fmt.Sprint(v), true
	}
	return "", false
}

func actionNamed(command ...string) carapace.Action {
	return actionList(command, "ID", "Name")
}

func actionID(command ...string) carapace.Action {
	return actionList(command, "ID")
}

var globalCloud func() string

var globalEnv func() []string

// SetGlobalCloud registers a provider for the cloud selected via global flags.
func SetGlobalCloud(f func() string) {
	globalCloud = f
}

// SetGlobalEnv registers a provider for environment variables derived from global flags.
func SetGlobalEnv(f func() []string) {
	globalEnv = f
}

func actionList(command []string, columns ...string) carapace.Action {
	return actionListArgs(append(append([]string{}, command...), "list"), columns...)
}

func actionListArgs(command []string, columns ...string) carapace.Action {
	if len(columns) == 0 {
		columns = []string{"ID", "Name"}
	}

	args := append([]string{}, command...)
	args = append(args, "-f", "csv")

	return actionExec(args, func(output []byte, err error) carapace.Action {
		if err != nil {
			return carapace.ActionValues()
		}

		reader := csv.NewReader(bytes.NewReader(output))
		reader.FieldsPerRecord = -1
		records, err := reader.ReadAll()
		if err != nil || len(records) <= 1 {
			return carapace.ActionValues()
		}

		indices := make([]int, len(columns))
		for i, column := range columns {
			indices[i] = -1
			for j, header := range records[0] {
				if strings.EqualFold(strings.TrimSpace(header), column) {
					indices[i] = j
					break
				}
			}
			if indices[i] == -1 && i < len(records[0]) {
				indices[i] = i
			}
		}

		field := func(record []string, index int) string {
			if index >= 0 && index < len(record) {
				return record[index]
			}
			return ""
		}

		described := len(indices) > 1
		values := make([]string, 0, (len(records)-1)*2)
		for _, record := range records[1:] {
			id := field(record, indices[0])
			if id == "" {
				continue
			}
			if described {
				if name := field(record, indices[1]); name != "" {
					values = append(values, name, id)
				} else {
					values = append(values, id, "")
				}
				continue
			}
			values = append(values, id)
		}

		if described {
			return carapace.ActionValuesDescribed(values...).Unique()
		}
		return carapace.ActionValues(values...).Unique()
	})
}

// ActionColumns completes column names from the output of the given list command.
func ActionColumns(command ...string) carapace.Action {
	args := append([]string{}, command...)
	args = append(args, "-f", "csv")

	return actionExec(args, func(output []byte, err error) carapace.Action {
		if err != nil {
			return carapace.ActionValues()
		}

		reader := csv.NewReader(bytes.NewReader(output))
		reader.FieldsPerRecord = -1
		header, err := reader.Read()
		if err != nil {
			return carapace.ActionValues()
		}

		values := make([]string, 0, len(header))
		for _, column := range header {
			if column = strings.TrimSpace(column); column != "" {
				values = append(values, column)
			}
		}
		return carapace.ActionValues(values...).Unique()
	})
}

func actionExec(args []string, f func(output []byte, err error) carapace.Action) carapace.Action {
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

		return carapace.ActionExecCommandE("openstack", args...)(f).Invoke(c).ToA()
	})
}

func actionScoped(argIndex int, command []string, columns ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if argIndex < 0 || argIndex >= len(c.Args) {
			return carapace.ActionValues()
		}

		args := append([]string{}, command...)
		args = append(args, c.Args[argIndex])
		return actionListArgs(args, columns...)
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

func ActionProjectIDs() carapace.Action {
	return actionID("project")
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

func ActionServerIDs() carapace.Action {
	return actionID("server")
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

func ActionVolumeIDs() carapace.Action {
	return actionID("volume")
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

func ActionShareSnapshotIDs() carapace.Action {
	return actionID("share", "snapshot")
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

func ActionBgpDragents() carapace.Action {
	return actionID("bgp", "dragent")
}

func ActionComputeHosts() carapace.Action {
	return actionListArgs([]string{"compute", "service", "list"}, "Host")
}

func ActionFloatingIPPortForwardings() carapace.Action {
	return actionScoped(0, []string{"floating", "ip", "port", "forwarding", "list"}, "ID")
}

func ActionImageMetadefObjects() carapace.Action {
	return actionScoped(0, []string{"image", "metadef", "object", "list"}, "Name")
}

func ActionImageMetadefProperties() carapace.Action {
	return actionScoped(0, []string{"image", "metadef", "property", "list"}, "Name")
}

func ActionImageMetadefResourceTypes() carapace.Action {
	return actionList([]string{"image", "metadef", "resource", "type"}, "Name")
}

func ActionNetworkFlavorProfiles() carapace.Action {
	return actionNamed("network", "flavor", "profile")
}

func ActionNetworkQosRules() carapace.Action {
	return actionScoped(0, []string{"network", "qos", "rule", "list"}, "ID")
}

func ActionObjects() carapace.Action {
	return ActionObjectsFromArg(0)
}

func ActionObjectsFromArg(index int) carapace.Action {
	return actionScoped(index, []string{"object", "list"}, "Name")
}

func ActionServerMigrations() carapace.Action {
	return actionScoped(0, []string{"server", "migration", "list", "--server"}, "ID")
}

func ActionVolumeHosts() carapace.Action {
	return actionListArgs([]string{"volume", "service", "list"}, "Host")
}

func ActionVolumeServiceBinaries() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"volume", "service", "list"}
		if len(c.Args) > 0 {
			args = append(args, "--host", c.Args[0])
		}
		return actionListArgs(args, "Binary")
	})
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
