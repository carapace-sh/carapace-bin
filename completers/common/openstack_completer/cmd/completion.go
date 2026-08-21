package cmd

import "github.com/carapace-sh/carapace"

func init() {
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"os-auth-type":         carapace.ActionValues("v3applicationcredential", "v3oidcaccesstoken", "v3oidcdeviceauthz", "password", "v3password", "token", "admin_token", "v2password", "v3token", "v3multifactor", "v3oauth2clientcredential", "v3tokenlessauth", "v3oidcauthcode", "http_basic", "v3totp", "none", "v3oidcclientcredentials", "v2token", "v3oauth2mtlsclientcredential", "v3oidcpassword"),
		"os-interface":         carapace.ActionValues("admin", "public", "internal"),
		"os-share-api-version": carapace.ActionValues("2.0", "2.1", "2.2", "2.3", "2.4", "2.5", "2.6", "2.7", "2.8", "2.9", "2.10", "2.11", "2.12", "2.13", "2.14", "2.15", "2.16", "2.17", "2.18", "2.19", "2.20", "2.21", "2.22", "2.23", "2.24", "2.25", "2.26", "2.27", "2.28", "2.29", "2.30", "2.31", "2.32", "2.33", "2.34", "2.35", "2.36", "2.37", "2.38", "2.39", "2.40", "2.41", "2.42", "2.43", "2.44", "2.45", "2.46", "2.47", "2.48", "2.49", "2.50", "2.51", "2.52", "2.53", "2.54", "2.55", "2.56", "2.57", "2.58", "2.59", "2.60", "2.61", "2.62", "2.63", "2.64", "2.65", "2.66", "2.67", "2.68", "2.69", "2.70", "2.71", "2.72", "2.73", "2.74", "2.75", "2.76", "2.77", "2.78", "2.79", "2.80", "2.81", "2.82", "2.83", "2.84", "2.85", "2.86", "2.87", "2.88", "2.89", "2.90", "2.91", "2.92", "2.93", "2.94", "2.95", "2.96"),
	})
	carapace.Gen(access_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(access_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(access_token_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(address_group_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(address_group_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(address_group_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(address_scope_createCmd).FlagCompletion(carapace.ActionMap{
		"format":     carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"ip-version": carapace.ActionValues("4", "6"),
	})
	carapace.Gen(address_scope_listCmd).FlagCompletion(carapace.ActionMap{
		"format":     carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"ip-version": carapace.ActionValues("4", "6"),
		"quote":      carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(address_scope_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(aggregate_add_hostCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(aggregate_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(aggregate_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(aggregate_remove_hostCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(aggregate_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(application_credential_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(application_credential_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(application_credential_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(availability_zone_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(bgp_dragent_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(bgp_peer_createCmd).FlagCompletion(carapace.ActionMap{
		"auth-type": carapace.ActionValues("none", "md5"),
		"format":    carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(bgp_peer_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(bgp_peer_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(bgp_speaker_createCmd).FlagCompletion(carapace.ActionMap{
		"format":     carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"ip-version": carapace.ActionValues("4", "6"),
	})
	carapace.Gen(bgp_speaker_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(bgp_speaker_list_advertised_routesCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(bgp_speaker_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(bgpvpn_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"type":   carapace.ActionValues("l2", "l3"),
	})
	carapace.Gen(bgpvpn_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(bgpvpn_network_association_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(bgpvpn_network_association_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(bgpvpn_network_association_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(bgpvpn_port_association_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(bgpvpn_port_association_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(bgpvpn_port_association_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(bgpvpn_router_association_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(bgpvpn_router_association_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(bgpvpn_router_association_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(bgpvpn_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(block_storage_cleanupCmd).FlagCompletion(carapace.ActionMap{
		"format":        carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":         carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"resource-type": carapace.ActionValues("Volume", "Snapshot"),
	})
	carapace.Gen(block_storage_cluster_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(block_storage_cluster_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(block_storage_log_level_listCmd).FlagCompletion(carapace.ActionMap{
		"format":  carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":   carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"service": carapace.ActionValues("None", "*", "cinder-api", "cinder-volume", "cinder-scheduler", "cinder-backup"),
	})
	carapace.Gen(block_storage_log_level_setCmd).FlagCompletion(carapace.ActionMap{
		"service": carapace.ActionValues("None", "*", "cinder-api", "cinder-volume", "cinder-scheduler", "cinder-backup"),
	})
	carapace.Gen(block_storage_resource_filter_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(block_storage_resource_filter_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(block_storage_snapshot_manageable_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(block_storage_volume_manageable_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(cached_image_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(catalog_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(catalog_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(command_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(completeCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "none"),
	})
	carapace.Gen(compute_agent_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(compute_agent_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(compute_service_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(configuration_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(consistency_group_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(consistency_group_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(consistency_group_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(consistency_group_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(consistency_group_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"status": carapace.ActionValues("available", "error", "creating", "deleting", "error_deleting"),
	})
	carapace.Gen(consistency_group_snapshot_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(console_connection_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(console_url_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(consumer_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(consumer_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(consumer_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(container_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(container_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(container_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(credential_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(credential_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(credential_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(default_security_group_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"ethertype": carapace.ActionValues("IPv4", "IPv6"),
		"format":    carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(default_security_group_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(default_security_group_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(domain_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(domain_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(domain_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(ec2_credentials_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(ec2_credentials_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(ec2_credentials_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(endpoint_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(endpoint_group_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(endpoint_group_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(endpoint_group_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(endpoint_listCmd).FlagCompletion(carapace.ActionMap{
		"format":    carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"interface": carapace.ActionValues("admin", "public", "internal"),
		"quote":     carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(endpoint_setCmd).FlagCompletion(carapace.ActionMap{
		"interface": carapace.ActionValues("admin", "public", "internal"),
	})
	carapace.Gen(endpoint_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(extension_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(extension_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(federation_domain_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(federation_project_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(federation_protocol_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(federation_protocol_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(federation_protocol_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(firewall_group_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(firewall_group_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(firewall_group_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(firewall_group_policy_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(firewall_group_policy_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(firewall_group_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"action":     carapace.ActionValues("allow", "deny", "reject"),
		"format":     carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"ip-version": carapace.ActionValues("4", "6"),
	})
	carapace.Gen(firewall_group_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(firewall_group_rule_setCmd).FlagCompletion(carapace.ActionMap{
		"action":     carapace.ActionValues("allow", "deny", "reject"),
		"ip-version": carapace.ActionValues("4", "6"),
	})
	carapace.Gen(firewall_group_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(firewall_group_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(flavor_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(flavor_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(flavor_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(floating_ip_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(floating_ip_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"status": carapace.ActionValues("ACTIVE", "DOWN"),
	})
	carapace.Gen(floating_ip_port_forwarding_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(floating_ip_port_forwarding_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(floating_ip_port_forwarding_setCmd).FlagCompletion(carapace.ActionMap{
		"protocol": carapace.ActionValues("tcp", "udp"),
	})
	carapace.Gen(floating_ip_port_forwarding_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(floating_ip_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(group_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(group_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(group_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(host_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(host_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(hypervisor_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(hypervisor_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(hypervisor_stats_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(identity_provider_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(identity_provider_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(identity_provider_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_add_projectCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_createCmd).FlagCompletion(carapace.ActionMap{
		"container-format": carapace.ActionValues("ami", "ari", "aki", "bare", "docker", "ova", "ovf"),
		"disk-format":      carapace.ActionValues("ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vhdx", "vdi", "iso", "ploop"),
		"format":           carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_importCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"method": carapace.ActionValues("glance-direct", "web-download", "glance-download", "copy-image"),
	})
	carapace.Gen(image_import_infoCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_listCmd).FlagCompletion(carapace.ActionMap{
		"format":        carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"member-status": carapace.ActionValues("accepted", "pending", "rejected", "all"),
		"quote":         carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(image_member_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_member_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(image_metadef_namespace_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_metadef_namespace_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(image_metadef_namespace_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_metadef_object_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_metadef_object_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(image_metadef_object_property_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_metadef_object_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_metadef_property_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_metadef_property_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(image_metadef_property_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_metadef_resource_type_association_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_metadef_resource_type_association_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(image_metadef_resource_type_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(image_setCmd).FlagCompletion(carapace.ActionMap{
		"container-format": carapace.ActionValues("ami", "ari", "aki", "bare", "docker", "ova", "ovf"),
		"disk-format":      carapace.ActionValues("ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vhdx", "vdi", "iso", "ploop"),
	})
	carapace.Gen(image_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(image_stores_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(image_task_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"status": carapace.ActionValues("pending", "processing", "success", "failure"),
		"type":   carapace.ActionValues("import"),
	})
	carapace.Gen(image_task_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(implied_role_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(implied_role_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(ip_availability_listCmd).FlagCompletion(carapace.ActionMap{
		"format":     carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"ip-version": carapace.ActionValues("4", "6"),
		"quote":      carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(ip_availability_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(keypair_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"type":   carapace.ActionValues("ssh", "x509"),
	})
	carapace.Gen(keypair_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(keypair_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(limit_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(limit_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(limit_setCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(limit_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(limits_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(local_ip_association_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(local_ip_association_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(local_ip_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(local_ip_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(local_ip_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(mapping_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(mapping_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(mapping_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(module_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_agent_listCmd).FlagCompletion(carapace.ActionMap{
		"agent-type": carapace.ActionValues("bgp", "dhcp", "open-vswitch", "linux-bridge", "ofa", "l3", "loadbalancer", "metering", "metadata", "macvtap", "nic", "baremetal", "ovn-controller", "ovn-controller-gateway", "ovn-metadata", "ovn-agent"),
		"format":     carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":      carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_agent_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_auto_allocated_topology_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_flavor_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_flavor_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_flavor_profile_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_flavor_profile_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_flavor_profile_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_flavor_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_l3_conntrack_helper_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_l3_conntrack_helper_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_l3_conntrack_helper_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_listCmd).FlagCompletion(carapace.ActionMap{
		"format":                carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"provider-network-type": carapace.ActionValues("flat", "geneve", "gre", "local", "vlan", "vxlan"),
		"quote":                 carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"status":                carapace.ActionValues("ACTIVE", "BUILD", "DOWN", "ERROR"),
	})
	carapace.Gen(network_meter_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_meter_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_meter_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_meter_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_meter_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_meter_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_qos_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_qos_policy_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_qos_policy_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_qos_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"type":   carapace.ActionValues("minimum-bandwidth", "minimum-packet-rate", "dscp-marking", "bandwidth-limit"),
	})
	carapace.Gen(network_qos_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_qos_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_qos_rule_type_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_qos_rule_type_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_rbac_createCmd).FlagCompletion(carapace.ActionMap{
		"action": carapace.ActionValues("access_as_external", "access_as_shared"),
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"type":   carapace.ActionValues("address_group", "address_scope", "security_group", "subnetpool", "qos_policy", "network"),
	})
	carapace.Gen(network_rbac_listCmd).FlagCompletion(carapace.ActionMap{
		"action": carapace.ActionValues("access_as_external", "access_as_shared"),
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"type":   carapace.ActionValues("address_group", "address_scope", "security_group", "subnetpool", "qos_policy", "network"),
	})
	carapace.Gen(network_rbac_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_segment_createCmd).FlagCompletion(carapace.ActionMap{
		"format":       carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"network-type": carapace.ActionValues("flat", "geneve", "gre", "local", "vlan", "vxlan"),
	})
	carapace.Gen(network_segment_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_segment_range_createCmd).FlagCompletion(carapace.ActionMap{
		"format":       carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"network-type": carapace.ActionValues("geneve", "gre", "vlan", "vxlan"),
	})
	carapace.Gen(network_segment_range_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_segment_range_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_segment_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_service_provider_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_subport_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_trunk_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(network_trunk_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(network_trunk_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(object_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(object_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(object_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(object_store_account_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(policy_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(policy_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(policy_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(port_createCmd).FlagCompletion(carapace.ActionMap{
		"format":     carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"pvlan-type": carapace.ActionValues("promiscuous", "isolated", "community"),
		"vnic-type":  carapace.ActionValues("direct", "direct-physical", "macvtap", "normal", "baremetal", "virtio-forwarder", "vdpa", "remote-managed"),
	})
	carapace.Gen(port_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"status": carapace.ActionValues("ACTIVE", "BUILD", "DOWN", "ERROR"),
	})
	carapace.Gen(port_setCmd).FlagCompletion(carapace.ActionMap{
		"data-plane-status": carapace.ActionValues("ACTIVE", "DOWN"),
		"pvlan-type":        carapace.ActionValues("promiscuous", "isolated", "community"),
		"vnic-type":         carapace.ActionValues("direct", "direct-physical", "macvtap", "normal", "baremetal", "virtio-forwarder", "vdpa", "remote-managed"),
	})
	carapace.Gen(port_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(project_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(project_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(project_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(quota_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(quota_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(region_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(region_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(region_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(registered_limit_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(registered_limit_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(registered_limit_setCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(registered_limit_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(request_token_authorizeCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(request_token_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(role_assignment_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(role_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(role_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(role_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(router_add_gatewayCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(router_add_routeCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(router_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(router_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(router_ndp_proxy_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(router_ndp_proxy_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(router_ndp_proxy_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(router_remove_gatewayCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(router_remove_routeCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(router_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(security_group_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(security_group_default_statefulness_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(security_group_default_statefulness_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(security_group_default_statefulness_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(security_group_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(security_group_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"ethertype": carapace.ActionValues("IPv4", "IPv6"),
		"format":    carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(security_group_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(security_group_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(security_group_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_add_fixed_ipCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_add_shareCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_add_volumeCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_backup_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_evacuateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_event_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(server_event_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_group_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"policy": carapace.ActionValues("affinity", "anti-affinity", "soft-affinity", "soft-anti-affinity"),
	})
	carapace.Gen(server_group_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(server_group_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_image_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_listCmd).FlagCompletion(carapace.ActionMap{
		"format":      carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"power-state": carapace.ActionValues("pending", "running", "paused", "shutdown", "crashed", "suspended"),
		"quote":       carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"status":      carapace.ActionValues("ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING", "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE", "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED", "VERIFY_RESIZE"),
		"task-state":  carapace.ActionValues("block_device_mapping", "deleting", "image_backup", "image_pending_upload", "image_snapshot", "image_snapshot_pending", "image_uploading", "migrating", "networking", "pausing", "powering-off", "powering-on", "rebooting", "reboot_pending", "reboot_started", "reboot_pending_hard", "reboot_started_hard", "rebooting_hard", "rebuilding", "rebuild_block_device_mapping", "rebuild_spawning", "rescuing", "resize_confirming", "resize_finish", "resize_migrated", "resize_migrating", "resize_prep", "resize_reverting", "restoring", "resuming", "scheduling", "shelving", "shelving_image_pending_upload", "shelving_image_uploading", "shelving_offloading", "soft-deleting", "spawning", "suspending", "updating_password", "unpausing", "unrescuing", "unshelving"),
		"vm-state":    carapace.ActionValues("active", "building", "deleted", "error", "paused", "stopped", "suspended", "rescued", "resized", "shelved", "shelved_offloaded", "soft-delete"),
	})
	carapace.Gen(server_migration_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"type":   carapace.ActionValues("evacuation", "live-migration", "cold-migration", "resize"),
	})
	carapace.Gen(server_migration_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_rebuildCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_setCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionValues("active", "error"),
	})
	carapace.Gen(server_share_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(server_share_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(server_volume_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(service_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(service_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(service_provider_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(service_provider_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(service_provider_setCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(service_provider_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(service_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_access_createCmd).FlagCompletion(carapace.ActionMap{
		"access-level": carapace.ActionValues("rw", "ro"),
		"format":       carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_access_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_access_setCmd).FlagCompletion(carapace.ActionMap{
		"access-level": carapace.ActionValues("rw", "ro"),
	})
	carapace.Gen(share_access_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_adoptCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_availability_zone_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_backup_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_backup_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_backup_setCmd).FlagCompletion(carapace.ActionMap{
		"status": carapace.ActionValues("available", "error", "creating", "deleting", "restoring"),
	})
	carapace.Gen(share_backup_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_export_location_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_export_location_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_group_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_group_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_group_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_group_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_group_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_group_snapshot_members_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_group_snapshot_setCmd).FlagCompletion(carapace.ActionMap{
		"status": carapace.ActionValues("available", "error", "creating", "deleting", "error_deleting"),
	})
	carapace.Gen(share_group_snapshot_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_group_type_access_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_group_type_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_group_type_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_group_type_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_instance_export_location_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_instance_export_location_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_instance_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_instance_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_limits_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_lock_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_lock_listCmd).FlagCompletion(carapace.ActionMap{
		"context":  carapace.ActionValues("user", "admin", "service"),
		"format":   carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":    carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"sort_dir": carapace.ActionValues("asc", "desc"),
		"sort_key": carapace.ActionValues("id", "created_at", "updated_at", "resource_id", "resource_type", "resource_action", "lock_reason"),
	})
	carapace.Gen(share_lock_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_message_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_message_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_migration_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_migration_startCmd).FlagCompletion(carapace.ActionMap{
		"force-host-assisted-migration": carapace.ActionValues("True", "False"),
		"nondisruptive":                 carapace.ActionValues("True", "False"),
		"preserve-metadata":             carapace.ActionValues("True", "False"),
		"preserve-snapshots":            carapace.ActionValues("True", "False"),
		"writable":                      carapace.ActionValues("True", "False"),
	})
	carapace.Gen(share_network_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_network_listCmd).FlagCompletion(carapace.ActionMap{
		"format":     carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"ip-version": carapace.ActionValues("4", "6"),
		"quote":      carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_network_setCmd).FlagCompletion(carapace.ActionMap{
		"status": carapace.ActionValues("active", "error", "network_change"),
	})
	carapace.Gen(share_network_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_network_subnet_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_network_subnet_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_pool_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_properties_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_qos_type_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_qos_type_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_qos_type_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_quota_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_replica_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_replica_export_location_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_replica_export_location_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_replica_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_replica_setCmd).FlagCompletion(carapace.ActionMap{
		"replica-state": carapace.ActionValues("in_sync", "out_of_sync", "active", "error"),
		"status":        carapace.ActionValues("available", "error", "creating", "deleting", "error_deleting"),
	})
	carapace.Gen(share_replica_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_security_service_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_security_service_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_security_service_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_server_adoptCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_server_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_server_migration_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_server_migration_startCmd).FlagCompletion(carapace.ActionMap{
		"format":             carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"nondisruptive":      carapace.ActionValues("True", "False"),
		"preserve-snapshots": carapace.ActionValues("True", "False"),
		"writable":           carapace.ActionValues("True", "False"),
	})
	carapace.Gen(share_server_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_service_listCmd).FlagCompletion(carapace.ActionMap{
		"ensuring": carapace.ActionValues("True", "False"),
		"format":   carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":    carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"state":    carapace.ActionValues("up", "down"),
	})
	carapace.Gen(share_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_snapshot_access_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_snapshot_access_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_snapshot_adoptCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_snapshot_export_location_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_snapshot_export_location_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_snapshot_instance_export_location_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_snapshot_instance_export_location_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_snapshot_instance_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_snapshot_instance_setCmd).FlagCompletion(carapace.ActionMap{
		"status": carapace.ActionValues("available", "error", "creating", "deleting", "error_deleting"),
	})
	carapace.Gen(share_snapshot_instance_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"usage":  carapace.ActionValues("used", "unused"),
	})
	carapace.Gen(share_snapshot_setCmd).FlagCompletion(carapace.ActionMap{
		"status": carapace.ActionValues("available", "error", "creating", "deleting", "manage_starting", "manage_error", "unmanage_starting", "unmanage_error", "error_deleting"),
	})
	carapace.Gen(share_snapshot_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_transfer_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_transfer_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_transfer_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_type_access_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_type_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(share_type_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(share_type_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(subnet_createCmd).FlagCompletion(carapace.ActionMap{
		"format":            carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"ip-version":        carapace.ActionValues("4", "6"),
		"ipv6-address-mode": carapace.ActionValues("dhcpv6-stateful", "dhcpv6-stateless", "slaac"),
		"ipv6-ra-mode":      carapace.ActionValues("dhcpv6-stateful", "dhcpv6-stateless", "slaac"),
	})
	carapace.Gen(subnet_listCmd).FlagCompletion(carapace.ActionMap{
		"format":     carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"ip-version": carapace.ActionValues("4", "6"),
		"quote":      carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(subnet_pool_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(subnet_pool_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(subnet_pool_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(subnet_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(tap_flow_createCmd).FlagCompletion(carapace.ActionMap{
		"direction": carapace.ActionValues("IN", "OUT", "BOTH"),
		"format":    carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(tap_flow_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(tap_flow_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(tap_flow_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(tap_mirror_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(tap_mirror_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(tap_mirror_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(tap_mirror_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(tap_service_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(tap_service_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(tap_service_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(tap_service_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(token_issueCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(trust_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(trust_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(trust_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(usage_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(usage_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(user_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(user_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(user_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(versions_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_attachment_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_attachment_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_attachment_setCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_attachment_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_backend_capability_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_backend_pool_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_backup_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_backup_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"status": carapace.ActionValues("creating", "available", "deleting", "error", "restoring", "error_restoring"),
	})
	carapace.Gen(volume_backup_record_exportCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_backup_record_importCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_backup_restoreCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_backup_setCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionValues("available", "error"),
	})
	carapace.Gen(volume_backup_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_group_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_group_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_group_setCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_group_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_group_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_group_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_group_snapshot_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_group_type_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_group_type_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_group_type_setCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_group_type_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_message_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_message_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_qos_createCmd).FlagCompletion(carapace.ActionMap{
		"consumer": carapace.ActionValues("front-end", "back-end", "both"),
		"format":   carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_qos_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_qos_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_service_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_setCmd).FlagCompletion(carapace.ActionMap{
		"migration-policy": carapace.ActionValues("never", "on-demand"),
		"retype-policy":    carapace.ActionValues("never", "on-demand"),
		"state":            carapace.ActionValues("available", "error", "creating", "deleting", "in-use", "attaching", "detaching", "error_deleting", "maintenance"),
	})
	carapace.Gen(volume_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
		"status": carapace.ActionValues("available", "error", "creating", "deleting", "error_deleting"),
	})
	carapace.Gen(volume_snapshot_setCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionValues("available", "error", "creating", "deleting", "error_deleting"),
	})
	carapace.Gen(volume_snapshot_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_summaryCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_transfer_request_acceptCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_transfer_request_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_transfer_request_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_transfer_request_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_type_createCmd).FlagCompletion(carapace.ActionMap{
		"encryption-control-location": carapace.ActionValues("front-end", "back-end"),
		"format":                      carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(volume_type_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(volume_type_setCmd).FlagCompletion(carapace.ActionMap{
		"encryption-control-location": carapace.ActionValues("front-end", "back-end"),
	})
	carapace.Gen(volume_type_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(vpn_endpoint_group_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"type":   carapace.ActionValues("subnet", "cidr"),
	})
	carapace.Gen(vpn_endpoint_group_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(vpn_endpoint_group_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(vpn_ike_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"auth-algorithm":          carapace.ActionValues("sha1", "sha256", "sha384", "sha512", "aes-xcbc", "aes-cmac"),
		"encryption-algorithm":    carapace.ActionValues("3des", "aes-128", "aes-192", "aes-256", "aes-128-ccm-8", "aes-192-ccm-8", "aes-256-ccm-8", "aes-128-ccm-12", "aes-192-ccm-12", "aes-256-ccm-12", "aes-128-ccm-16", "aes-192-ccm-16", "aes-256-ccm-16", "aes-128-gcm-8", "aes-192-gcm-8", "aes-256-gcm-8", "aes-128-gcm-12", "aes-192-gcm-12", "aes-256-gcm-12", "aes-128-gcm-16", "aes-192-gcm-16", "aes-256-gcm-16", "aes-128-ctr", "aes-192-ctr", "aes-256-ctr"),
		"format":                  carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"ike-version":             carapace.ActionValues("v1", "v2"),
		"pfs":                     carapace.ActionValues("group2", "group5", "group14", "group15", "group16", "group17", "group18", "group19", "group20", "group21", "group22", "group23", "group24", "group25", "group26", "group27", "group28", "group29", "group30", "group31"),
		"phase1-negotiation-mode": carapace.ActionValues("main", "aggressive"),
	})
	carapace.Gen(vpn_ike_policy_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(vpn_ike_policy_setCmd).FlagCompletion(carapace.ActionMap{
		"auth-algorithm":          carapace.ActionValues("sha1", "sha256", "sha384", "sha512", "aes-xcbc", "aes-cmac"),
		"encryption-algorithm":    carapace.ActionValues("3des", "aes-128", "aes-192", "aes-256", "aes-128-ccm-8", "aes-192-ccm-8", "aes-256-ccm-8", "aes-128-ccm-12", "aes-192-ccm-12", "aes-256-ccm-12", "aes-128-ccm-16", "aes-192-ccm-16", "aes-256-ccm-16", "aes-128-gcm-8", "aes-192-gcm-8", "aes-256-gcm-8", "aes-128-gcm-12", "aes-192-gcm-12", "aes-256-gcm-12", "aes-128-gcm-16", "aes-192-gcm-16", "aes-256-gcm-16", "aes-128-ctr", "aes-192-ctr", "aes-256-ctr"),
		"ike-version":             carapace.ActionValues("v1", "v2"),
		"pfs":                     carapace.ActionValues("group2", "group5", "group14", "group15", "group16", "group17", "group18", "group19", "group20", "group21", "group22", "group23", "group24", "group25", "group26", "group27", "group28", "group29", "group30", "group31"),
		"phase1-negotiation-mode": carapace.ActionValues("main", "aggressive"),
	})
	carapace.Gen(vpn_ike_policy_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(vpn_ipsec_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"auth-algorithm":       carapace.ActionValues("sha1", "sha256", "sha384", "sha512", "aes-xcbc", "aes-cmac"),
		"encapsulation-mode":   carapace.ActionValues("tunnel", "transport"),
		"encryption-algorithm": carapace.ActionValues("3des", "aes-128", "aes-192", "aes-256", "aes-128-ccm-8", "aes-192-ccm-8", "aes-256-ccm-8", "aes-128-ccm-12", "aes-192-ccm-12", "aes-256-ccm-12", "aes-128-ccm-16", "aes-192-ccm-16", "aes-256-ccm-16", "aes-128-gcm-8", "aes-192-gcm-8", "aes-256-gcm-8", "aes-128-gcm-12", "aes-192-gcm-12", "aes-256-gcm-12", "aes-128-gcm-16", "aes-192-gcm-16", "aes-256-gcm-16", "aes-128-ctr", "aes-192-ctr", "aes-256-ctr"),
		"format":               carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"pfs":                  carapace.ActionValues("group2", "group5", "group14", "group15", "group16", "group17", "group18", "group19", "group20", "group21", "group22", "group23", "group24", "group25", "group26", "group27", "group28", "group29", "group30", "group31"),
		"transform-protocol":   carapace.ActionValues("esp", "ah", "ah-esp"),
	})
	carapace.Gen(vpn_ipsec_policy_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(vpn_ipsec_policy_setCmd).FlagCompletion(carapace.ActionMap{
		"auth-algorithm":       carapace.ActionValues("sha1", "sha256", "sha384", "sha512", "aes-xcbc", "aes-cmac"),
		"encapsulation-mode":   carapace.ActionValues("tunnel", "transport"),
		"encryption-algorithm": carapace.ActionValues("3des", "aes-128", "aes-192", "aes-256", "aes-128-ccm-8", "aes-192-ccm-8", "aes-256-ccm-8", "aes-128-ccm-12", "aes-192-ccm-12", "aes-256-ccm-12", "aes-128-ccm-16", "aes-192-ccm-16", "aes-256-ccm-16", "aes-128-gcm-8", "aes-192-gcm-8", "aes-256-gcm-8", "aes-128-gcm-12", "aes-192-gcm-12", "aes-256-gcm-12", "aes-128-gcm-16", "aes-192-gcm-16", "aes-256-gcm-16", "aes-128-ctr", "aes-192-ctr", "aes-256-ctr"),
		"pfs":                  carapace.ActionValues("group2", "group5", "group14", "group15", "group16", "group17", "group18", "group19", "group20", "group21", "group22", "group23", "group24", "group25", "group26", "group27", "group28", "group29", "group30", "group31"),
		"transform-protocol":   carapace.ActionValues("esp", "ah", "ah-esp"),
	})
	carapace.Gen(vpn_ipsec_policy_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(vpn_ipsec_site_connection_createCmd).FlagCompletion(carapace.ActionMap{
		"format":    carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"initiator": carapace.ActionValues("bi-directional", "response-only"),
	})
	carapace.Gen(vpn_ipsec_site_connection_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(vpn_ipsec_site_connection_setCmd).FlagCompletion(carapace.ActionMap{
		"initiator": carapace.ActionValues("bi-directional", "response-only"),
	})
	carapace.Gen(vpn_ipsec_site_connection_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(vpn_service_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
	carapace.Gen(vpn_service_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("csv", "json", "table", "value", "yaml"),
		"quote":  carapace.ActionValues("all", "minimal", "none", "nonnumeric"),
	})
	carapace.Gen(vpn_service_showCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
}
