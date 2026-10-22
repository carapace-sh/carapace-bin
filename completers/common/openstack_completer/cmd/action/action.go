package action

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
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

func ActionCloudAuthURLs(cmd *cobra.Command) carapace.Action {
	return cloudValues(cmd, "auth_url")
}

func ActionCloudDomainIDs(cmd *cobra.Command) carapace.Action {
	return cloudValues(cmd, "domain_id", "user_domain_id", "project_domain_id")
}

func ActionCloudDomainNames(cmd *cobra.Command) carapace.Action {
	return cloudValues(cmd, "domain_name", "user_domain_name", "project_domain_name")
}

func ActionCloudProjectIDs(cmd *cobra.Command) carapace.Action {
	return cloudValues(cmd, "project_id")
}

func ActionCloudProjectNames(cmd *cobra.Command) carapace.Action {
	return cloudValues(cmd, "project_name")
}

func ActionCloudRegions(cmd *cobra.Command) carapace.Action {
	return cloudValues(cmd, "region_name")
}

func ActionCloudUserIDs(cmd *cobra.Command) carapace.Action {
	return cloudValues(cmd, "user_id")
}

func ActionCloudUsernames(cmd *cobra.Command) carapace.Action {
	return cloudValues(cmd, "username")
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

func cloudValues(cmd *cobra.Command, keys ...string) carapace.Action {
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
			if flag := cmd.Root().Flags().Lookup("os-cloud"); flag != nil && flag.Changed {
				selected = flag.Value.String()
			}
			if selected == "" {
				selected = os.Getenv("OS_CLOUD")
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

func actionNamed(cmd *cobra.Command, command ...string) carapace.Action {
	return actionList(cmd, command, "ID", "Name")
}

func actionID(cmd *cobra.Command, command ...string) carapace.Action {
	return actionList(cmd, command, "ID")
}

func actionList(cmd *cobra.Command, command []string, columns ...string) carapace.Action {
	return actionListArgs(cmd, append(append([]string{}, command...), "list"), columns...)
}

func actionListArgs(cmd *cobra.Command, command []string, columns ...string) carapace.Action {
	if len(columns) == 0 {
		columns = []string{"ID", "Name"}
	}

	args := append([]string{}, command...)
	args = append(args, "-f", "csv")

	return actionExec(cmd, args, func(output []byte, err error) carapace.Action {
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

// ActionColumns completes column names from the corresponding list command.
func ActionColumns(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		list := columnList(cmd)
		if list == nil {
			return carapace.ActionValues()
		}

		path := strings.Fields(list.CommandPath())
		if len(path) < 2 {
			return carapace.ActionValues()
		}
		return actionColumns(cmd, path[1:])
	})
}

func actionColumns(cmd *cobra.Command, command []string) carapace.Action {
	args := append([]string{}, command...)
	args = append(args, "-f", "csv")

	return actionExec(cmd, args, func(output []byte, err error) carapace.Action {
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

func columnList(cmd *cobra.Command) *cobra.Command {
	if isLister(cmd) {
		return cmd
	}
	for parent := cmd.Parent(); parent != nil; parent = parent.Parent() {
		for _, child := range parent.Commands() {
			if child.Name() == "list" && isLister(child) {
				return child
			}
		}
	}
	return nil
}

func isLister(cmd *cobra.Command) bool {
	return cmd.Flags().Lookup("sort-column") != nil
}

func actionExec(cmd *cobra.Command, args []string, f func(output []byte, err error) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if overrides := actionEnv(cmd.Root()); len(overrides) > 0 {
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

		return carapace.ActionExecCommandE("openstack", args...)(f).Invoke(c).ToA()
	})
}

func actionEnv(cmd *cobra.Command) []string {
	env := []string{}
	add := func(name, flagName string) {
		if flag := cmd.Flags().Lookup(flagName); flag != nil && flag.Changed {
			env = append(env, name+"="+flag.Value.String())
		}
	}

	add("OS_CLOUD", "os-cloud")
	add("OS_REGION_NAME", "os-region-name")
	add("OS_INTERFACE", "os-interface")
	add("OS_AUTH_TYPE", "os-auth-type")
	add("OS_AUTH_URL", "os-auth-url")
	add("OS_IDENTITY_API_VERSION", "os-identity-api-version")
	add("OS_USERNAME", "os-username")
	add("OS_USER_ID", "os-user-id")
	add("OS_PASSWORD", "os-password")
	add("OS_TOKEN", "os-token")
	add("OS_PROJECT_NAME", "os-project-name")
	add("OS_PROJECT_ID", "os-project-id")
	add("OS_PROJECT_DOMAIN_NAME", "os-project-domain-name")
	add("OS_PROJECT_DOMAIN_ID", "os-project-domain-id")
	add("OS_USER_DOMAIN_NAME", "os-user-domain-name")
	add("OS_USER_DOMAIN_ID", "os-user-domain-id")
	add("OS_DOMAIN_NAME", "os-domain-name")
	add("OS_DOMAIN_ID", "os-domain-id")
	add("OS_DEFAULT_DOMAIN", "os-default-domain")
	add("OS_DEFAULT_DOMAIN_NAME", "os-default-domain-name")
	add("OS_DEFAULT_DOMAIN_ID", "os-default-domain-id")
	add("OS_TENANT_NAME", "os-tenant-name")
	add("OS_TENANT_ID", "os-tenant-id")
	add("OS_TRUST_ID", "os-trust-id")
	add("OS_APPLICATION_CREDENTIAL_ID", "os-application-credential-id")
	add("OS_APPLICATION_CREDENTIAL_NAME", "os-application-credential-name")
	add("OS_APPLICATION_CREDENTIAL_SECRET", "os-application-credential-secret")
	add("OS_ACCESS_TOKEN", "os-access-token")
	add("OS_SYSTEM_SCOPE", "os-system-scope")
	add("OS_CACERT", "os-cacert")
	add("OS_CERT", "os-cert")
	add("OS_KEY", "os-key")
	add("OS_ENDPOINT_OVERRIDE", "os-endpoint-override")
	add("OS_SERVICE_PROVIDER", "os-service-provider")
	add("OS_REMOTE_PROJECT_ID", "os-remote-project-id")
	add("OS_REMOTE_PROJECT_NAME", "os-remote-project-name")
	add("OS_REMOTE_PROJECT_DOMAIN_ID", "os-remote-project-domain-id")
	add("OS_REMOTE_PROJECT_DOMAIN_NAME", "os-remote-project-domain-name")
	add("OS_COMPUTE_API_VERSION", "os-compute-api-version")
	add("OS_NETWORK_API_VERSION", "os-network-api-version")
	add("OS_VOLUME_API_VERSION", "os-volume-api-version")
	add("OS_IMAGE_API_VERSION", "os-image-api-version")
	add("OS_OBJECT_API_VERSION", "os-object-api-version")
	add("OS_SHARE_API_VERSION", "os-share-api-version")
	add("OS_ACCELERATOR_API_VERSION", "os-accelerator-api-version")
	add("OS_INSECURE", "insecure")
	add("OS_VERIFY", "verify")

	return env
}

func actionScoped(cmd *cobra.Command, argIndex int, command []string, columns ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if argIndex < 0 || argIndex >= len(c.Args) {
			return carapace.ActionValues()
		}

		args := append([]string{}, command...)
		args = append(args, c.Args[argIndex])
		return actionListArgs(cmd, args, columns...)
	})
}

func ActionAggregates(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "aggregate")
}

func ActionApplicationCredentials(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "application", "credential")
}

func ActionAvailabilityZones(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"availability", "zone"}, "Zone Name")
}

func ActionComputeAgents(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"compute", "agent"}, "Agent ID")
}

func ActionCredentials(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "credential")
}

func ActionDomains(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "domain")
}

func ActionEndpoints(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "endpoint")
}

func ActionFlavors(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "flavor")
}

func ActionGroups(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "group")
}

func ActionHypervisors(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"hypervisor"}, "ID", "Hypervisor Hostname")
}

func ActionIdentityProviders(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "identity", "provider")
}

func ActionImages(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "image")
}

func ActionKeypairs(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"keypair"}, "Name")
}

func ActionLimits(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "limit")
}

func ActionMappings(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "mapping")
}

func ActionProjectIDs(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "project")
}

func ActionProjects(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "project")
}

func ActionRegions(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"region"}, "Region")
}

func ActionRegisteredLimits(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "registered", "limit")
}

func ActionRoles(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "role")
}

func ActionServerGroups(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "server", "group")
}

func ActionServerIDs(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "server")
}

func ActionServers(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "server")
}

func ActionServices(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "service")
}

func ActionTrusts(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "trust")
}

func ActionUsers(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "user")
}

func ActionVolumeAttachments(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "volume", "attachment")
}

func ActionVolumeBackups(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "volume", "backup")
}

func ActionVolumeGroupSnapshots(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "volume", "group", "snapshot")
}

func ActionVolumeGroups(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "volume", "group")
}

func ActionVolumeQosSpecs(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "volume", "qos")
}

func ActionVolumeSnapshots(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "volume", "snapshot")
}

func ActionVolumeTransfers(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "volume", "transfer", "request")
}

func ActionVolumeTypes(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "volume", "type")
}

func ActionVolumeIDs(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "volume")
}

func ActionVolumes(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "volume")
}

func ActionShareAccessRules(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) == 0 {
			return carapace.ActionValues()
		}
		return actionListArgs(cmd, []string{"share", "access", "list", c.Args[0]}, "ID")
	})
}

func ActionShareBackups(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "backup")
}

func ActionShareGroupSnapshots(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "group", "snapshot")
}

func ActionShareGroups(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "group")
}

func ActionShareInstances(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "share", "instance")
}

func ActionShareMessages(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "share", "message")
}

func ActionShareNetworks(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "network")
}

func ActionShareQosTypes(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "qos", "type")
}

func ActionShareReplicas(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "replica")
}

func ActionShareSecurityServices(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "security", "service")
}

func ActionShareServers(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "share", "server")
}

func ActionShareSnapshotIDs(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "share", "snapshot")
}

func ActionShareSnapshots(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "snapshot")
}

func ActionShareTypes(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "type")
}

func ActionShares(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share")
}

func ActionAddressGroups(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "address", "group")
}

func ActionAddressScopes(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "address", "scope")
}

func ActionFloatingIPs(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"floating", "ip"}, "ID", "Floating IP Address")
}

func ActionLocalIPs(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "local", "ip")
}

func ActionNetworkAgents(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "network", "agent")
}

func ActionNetworkFlavors(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "network", "flavor")
}

func ActionNetworkQosPolicies(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "network", "qos", "policy")
}

func ActionNetworkRbacPolicies(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "network", "rbac")
}

func ActionNetworkSegments(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "network", "segment")
}

func ActionNetworkTrunks(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "network", "trunk")
}

func ActionNetworks(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "network")
}

func ActionPorts(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "port")
}

func ActionRouters(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "router")
}

func ActionSecurityGroupRules(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "security", "group", "rule")
}

func ActionSecurityGroups(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "security", "group")
}

func ActionSubnetPools(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "subnet", "pool")
}

func ActionSubnets(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "subnet")
}

func ActionFirewallGroups(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "firewall", "group")
}

func ActionFirewallPolicies(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "firewall", "group", "policy")
}

func ActionFirewallRules(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "firewall", "group", "rule")
}

func ActionBgpSpeakers(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "bgp", "speaker")
}

func ActionBgpPeers(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "bgp", "peer")
}

func ActionBgpDragents(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "bgp", "dragent")
}

func ActionComputeHosts(cmd *cobra.Command) carapace.Action {
	return actionListArgs(cmd, []string{"compute", "service", "list"}, "Host")
}

func ActionFloatingIPPortForwardings(cmd *cobra.Command) carapace.Action {
	return actionScoped(cmd, 0, []string{"floating", "ip", "port", "forwarding", "list"}, "ID")
}

func ActionImageMetadefObjects(cmd *cobra.Command) carapace.Action {
	return actionScoped(cmd, 0, []string{"image", "metadef", "object", "list"}, "Name")
}

func ActionImageMetadefProperties(cmd *cobra.Command) carapace.Action {
	return actionScoped(cmd, 0, []string{"image", "metadef", "property", "list"}, "Name")
}

func ActionImageMetadefResourceTypes(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"image", "metadef", "resource", "type"}, "Name")
}

func ActionNetworkFlavorProfiles(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "network", "flavor", "profile")
}

func ActionNetworkQosRules(cmd *cobra.Command) carapace.Action {
	return actionScoped(cmd, 0, []string{"network", "qos", "rule", "list"}, "ID")
}

func ActionObjects(cmd *cobra.Command) carapace.Action {
	return ActionObjectsFromArg(cmd, 0)
}

func ActionObjectsFromArg(cmd *cobra.Command, index int) carapace.Action {
	return actionScoped(cmd, index, []string{"object", "list"}, "Name")
}

func ActionServerMigrations(cmd *cobra.Command) carapace.Action {
	return actionScoped(cmd, 0, []string{"server", "migration", "list", "--server"}, "ID")
}

func ActionVolumeHosts(cmd *cobra.Command) carapace.Action {
	return actionListArgs(cmd, []string{"volume", "service", "list"}, "Host")
}

func ActionVolumeServiceBinaries(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"volume", "service", "list"}
		if len(c.Args) > 0 {
			args = append(args, "--host", c.Args[0])
		}
		return actionListArgs(cmd, args, "Binary")
	})
}

func ActionBgpvpns(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "bgpvpn")
}

func ActionAcceleratorArqs(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"accelerator", "arq"}, "uuid")
}

func ActionAcceleratorDeployables(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"accelerator", "deployable"}, "uuid", "name")
}

func ActionAcceleratorDevices(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"accelerator", "device"}, "uuid")
}

func ActionAcceleratorDeviceProfiles(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"accelerator", "device", "profile"}, "uuid", "name")
}

func ActionContainers(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"container"}, "Name")
}

func ActionCachedImages(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"cached", "image"}, "image_id")
}

func ActionConsistencyGroups(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "consistency", "group")
}

func ActionConsistencyGroupSnapshots(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "consistency", "group", "snapshot")
}

func ActionExtensions(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"extension"}, "Name")
}

func ActionImageMetadefNamespaces(cmd *cobra.Command) carapace.Action {
	return actionList(cmd, []string{"image", "metadef", "namespace"}, "Namespace")
}

func ActionImageTasks(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "image", "task")
}

func ActionNetworkMeters(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "network", "meter")
}

func ActionNetworkSegmentRanges(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "network", "segment", "range")
}

func ActionShareGroupTypes(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "group", "type")
}

func ActionShareLocks(cmd *cobra.Command) carapace.Action {
	return actionID(cmd, "share", "lock")
}

func ActionShareTransfers(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "share", "transfer")
}

func ActionVolumeGroupTypes(cmd *cobra.Command) carapace.Action {
	return actionNamed(cmd, "volume", "group", "type")
}
