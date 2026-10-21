package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
)

func init() {
	carapace.Gen(accelerator_arq_bindCmd).PositionalCompletion(action.ActionAcceleratorArqs())
	carapace.Gen(accelerator_arq_deleteCmd).PositionalAnyCompletion(action.ActionAcceleratorArqs())
	carapace.Gen(accelerator_arq_showCmd).PositionalCompletion(action.ActionAcceleratorArqs())
	carapace.Gen(accelerator_arq_unbindCmd).PositionalCompletion(action.ActionAcceleratorArqs())
	carapace.Gen(accelerator_deployable_programCmd).PositionalCompletion(action.ActionAcceleratorDeployables())
	carapace.Gen(accelerator_deployable_showCmd).PositionalCompletion(action.ActionAcceleratorDeployables())
	carapace.Gen(accelerator_device_profile_deleteCmd).PositionalAnyCompletion(action.ActionAcceleratorDeviceProfiles())
	carapace.Gen(accelerator_device_profile_showCmd).PositionalCompletion(action.ActionAcceleratorDeviceProfiles())
	carapace.Gen(accelerator_device_showCmd).PositionalCompletion(action.ActionAcceleratorDevices())
	carapace.Gen(address_group_deleteCmd).PositionalAnyCompletion(action.ActionAddressGroups())
	carapace.Gen(address_group_setCmd).PositionalCompletion(action.ActionAddressGroups())
	carapace.Gen(address_group_showCmd).PositionalCompletion(action.ActionAddressGroups())
	carapace.Gen(address_group_unsetCmd).PositionalCompletion(action.ActionAddressGroups())
	carapace.Gen(address_scope_deleteCmd).PositionalAnyCompletion(action.ActionAddressScopes())
	carapace.Gen(address_scope_setCmd).PositionalCompletion(action.ActionAddressScopes())
	carapace.Gen(address_scope_showCmd).PositionalCompletion(action.ActionAddressScopes())
	carapace.Gen(aggregate_add_hostCmd).PositionalCompletion(action.ActionAggregates())
	carapace.Gen(aggregate_cache_imageCmd).PositionalCompletion(action.ActionAggregates())
	carapace.Gen(aggregate_deleteCmd).PositionalAnyCompletion(action.ActionAggregates())
	carapace.Gen(aggregate_remove_hostCmd).PositionalCompletion(action.ActionAggregates())
	carapace.Gen(aggregate_setCmd).PositionalCompletion(action.ActionAggregates())
	carapace.Gen(aggregate_showCmd).PositionalCompletion(action.ActionAggregates())
	carapace.Gen(aggregate_unsetCmd).PositionalCompletion(action.ActionAggregates())
	carapace.Gen(application_credential_deleteCmd).PositionalAnyCompletion(action.ActionApplicationCredentials())
	carapace.Gen(application_credential_showCmd).PositionalCompletion(action.ActionApplicationCredentials())
	carapace.Gen(bgp_peer_deleteCmd).PositionalAnyCompletion(action.ActionBgpPeers())
	carapace.Gen(bgp_peer_setCmd).PositionalCompletion(action.ActionBgpPeers())
	carapace.Gen(bgp_peer_showCmd).PositionalCompletion(action.ActionBgpPeers())
	carapace.Gen(bgp_speaker_deleteCmd).PositionalAnyCompletion(action.ActionBgpSpeakers())
	carapace.Gen(bgp_speaker_setCmd).PositionalCompletion(action.ActionBgpSpeakers())
	carapace.Gen(bgp_speaker_showCmd).PositionalCompletion(action.ActionBgpSpeakers())
	carapace.Gen(bgpvpn_deleteCmd).PositionalAnyCompletion(action.ActionBgpvpns())
	carapace.Gen(bgpvpn_setCmd).PositionalCompletion(action.ActionBgpvpns())
	carapace.Gen(bgpvpn_showCmd).PositionalCompletion(action.ActionBgpvpns())
	carapace.Gen(bgpvpn_unsetCmd).PositionalCompletion(action.ActionBgpvpns())
	carapace.Gen(cached_image_deleteCmd).PositionalAnyCompletion(action.ActionCachedImages())
	carapace.Gen(compute_agent_deleteCmd).PositionalAnyCompletion(action.ActionComputeAgents())
	carapace.Gen(compute_agent_setCmd).PositionalCompletion(action.ActionComputeAgents())
	carapace.Gen(consistency_group_deleteCmd).PositionalAnyCompletion(action.ActionConsistencyGroups())
	carapace.Gen(consistency_group_setCmd).PositionalCompletion(action.ActionConsistencyGroups())
	carapace.Gen(consistency_group_showCmd).PositionalCompletion(action.ActionConsistencyGroups())
	carapace.Gen(consistency_group_snapshot_deleteCmd).PositionalAnyCompletion(action.ActionConsistencyGroupSnapshots())
	carapace.Gen(consistency_group_snapshot_showCmd).PositionalCompletion(action.ActionConsistencyGroupSnapshots())
	carapace.Gen(container_deleteCmd).PositionalAnyCompletion(action.ActionContainers())
	carapace.Gen(container_setCmd).PositionalCompletion(action.ActionContainers())
	carapace.Gen(container_showCmd).PositionalCompletion(action.ActionContainers())
	carapace.Gen(container_unsetCmd).PositionalCompletion(action.ActionContainers())
	carapace.Gen(credential_deleteCmd).PositionalAnyCompletion(action.ActionCredentials())
	carapace.Gen(credential_setCmd).PositionalCompletion(action.ActionCredentials())
	carapace.Gen(credential_showCmd).PositionalCompletion(action.ActionCredentials())
	carapace.Gen(domain_deleteCmd).PositionalAnyCompletion(action.ActionDomains())
	carapace.Gen(domain_setCmd).PositionalCompletion(action.ActionDomains())
	carapace.Gen(domain_showCmd).PositionalCompletion(action.ActionDomains())
	carapace.Gen(endpoint_deleteCmd).PositionalAnyCompletion(action.ActionEndpoints())
	carapace.Gen(endpoint_setCmd).PositionalCompletion(action.ActionEndpoints())
	carapace.Gen(endpoint_showCmd).PositionalCompletion(action.ActionEndpoints())
	carapace.Gen(extension_showCmd).PositionalCompletion(action.ActionExtensions())
	carapace.Gen(firewall_group_deleteCmd).PositionalAnyCompletion(action.ActionFirewallGroups())
	carapace.Gen(firewall_group_policy_deleteCmd).PositionalAnyCompletion(action.ActionFirewallPolicies())
	carapace.Gen(firewall_group_policy_setCmd).PositionalCompletion(action.ActionFirewallPolicies())
	carapace.Gen(firewall_group_policy_showCmd).PositionalCompletion(action.ActionFirewallPolicies())
	carapace.Gen(firewall_group_policy_unsetCmd).PositionalCompletion(action.ActionFirewallPolicies())
	carapace.Gen(firewall_group_rule_deleteCmd).PositionalAnyCompletion(action.ActionFirewallRules())
	carapace.Gen(firewall_group_rule_setCmd).PositionalCompletion(action.ActionFirewallRules())
	carapace.Gen(firewall_group_rule_showCmd).PositionalCompletion(action.ActionFirewallRules())
	carapace.Gen(firewall_group_rule_unsetCmd).PositionalCompletion(action.ActionFirewallRules())
	carapace.Gen(firewall_group_setCmd).PositionalCompletion(action.ActionFirewallGroups())
	carapace.Gen(firewall_group_showCmd).PositionalCompletion(action.ActionFirewallGroups())
	carapace.Gen(firewall_group_unsetCmd).PositionalCompletion(action.ActionFirewallGroups())
	carapace.Gen(flavor_deleteCmd).PositionalAnyCompletion(action.ActionFlavors())
	carapace.Gen(flavor_setCmd).PositionalCompletion(action.ActionFlavors())
	carapace.Gen(flavor_showCmd).PositionalCompletion(action.ActionFlavors())
	carapace.Gen(flavor_unsetCmd).PositionalCompletion(action.ActionFlavors())
	carapace.Gen(floating_ip_deleteCmd).PositionalAnyCompletion(action.ActionFloatingIPs())
	carapace.Gen(floating_ip_setCmd).PositionalCompletion(action.ActionFloatingIPs())
	carapace.Gen(floating_ip_showCmd).PositionalCompletion(action.ActionFloatingIPs())
	carapace.Gen(floating_ip_unsetCmd).PositionalCompletion(action.ActionFloatingIPs())
	carapace.Gen(group_add_userCmd).PositionalCompletion(action.ActionGroups(), action.ActionUsers())
	carapace.Gen(group_contains_userCmd).PositionalCompletion(action.ActionGroups(), action.ActionUsers())
	carapace.Gen(group_deleteCmd).PositionalAnyCompletion(action.ActionGroups())
	carapace.Gen(group_remove_userCmd).PositionalCompletion(action.ActionGroups(), action.ActionUsers())
	carapace.Gen(group_setCmd).PositionalCompletion(action.ActionGroups())
	carapace.Gen(group_showCmd).PositionalCompletion(action.ActionGroups())
	carapace.Gen(hypervisor_showCmd).PositionalCompletion(action.ActionHypervisors())
	carapace.Gen(identity_provider_deleteCmd).PositionalAnyCompletion(action.ActionIdentityProviders())
	carapace.Gen(identity_provider_setCmd).PositionalCompletion(action.ActionIdentityProviders())
	carapace.Gen(identity_provider_showCmd).PositionalCompletion(action.ActionIdentityProviders())
	carapace.Gen(image_add_projectCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_deleteCmd).PositionalAnyCompletion(action.ActionImages())
	carapace.Gen(image_importCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_member_getCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_member_listCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_metadef_namespace_deleteCmd).PositionalAnyCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_namespace_setCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_namespace_showCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_namespace_unsetCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_remove_projectCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_saveCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_setCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_showCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_stageCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_task_showCmd).PositionalCompletion(action.ActionImageTasks())
	carapace.Gen(image_unsetCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(keypair_deleteCmd).PositionalAnyCompletion(action.ActionKeypairs())
	carapace.Gen(keypair_showCmd).PositionalCompletion(action.ActionKeypairs())
	carapace.Gen(limit_deleteCmd).PositionalAnyCompletion(action.ActionLimits())
	carapace.Gen(limit_setCmd).PositionalCompletion(action.ActionLimits())
	carapace.Gen(limit_showCmd).PositionalCompletion(action.ActionLimits())
	carapace.Gen(local_ip_deleteCmd).PositionalAnyCompletion(action.ActionLocalIPs())
	carapace.Gen(local_ip_setCmd).PositionalCompletion(action.ActionLocalIPs())
	carapace.Gen(local_ip_showCmd).PositionalCompletion(action.ActionLocalIPs())
	carapace.Gen(mapping_deleteCmd).PositionalAnyCompletion(action.ActionMappings())
	carapace.Gen(mapping_setCmd).PositionalCompletion(action.ActionMappings())
	carapace.Gen(mapping_showCmd).PositionalCompletion(action.ActionMappings())
	carapace.Gen(network_agent_deleteCmd).PositionalAnyCompletion(action.ActionNetworkAgents())
	carapace.Gen(network_agent_setCmd).PositionalCompletion(action.ActionNetworkAgents())
	carapace.Gen(network_agent_showCmd).PositionalCompletion(action.ActionNetworkAgents())
	carapace.Gen(network_deleteCmd).PositionalAnyCompletion(action.ActionNetworks())
	carapace.Gen(network_flavor_deleteCmd).PositionalAnyCompletion(action.ActionNetworkFlavors())
	carapace.Gen(network_flavor_setCmd).PositionalCompletion(action.ActionNetworkFlavors())
	carapace.Gen(network_flavor_showCmd).PositionalCompletion(action.ActionNetworkFlavors())
	carapace.Gen(network_meter_deleteCmd).PositionalAnyCompletion(action.ActionNetworkMeters())
	carapace.Gen(network_meter_showCmd).PositionalCompletion(action.ActionNetworkMeters())
	carapace.Gen(network_onboard_subnetCmd).PositionalCompletion(action.ActionNetworks(), action.ActionSubnetPools())
	carapace.Gen(network_qos_policy_deleteCmd).PositionalAnyCompletion(action.ActionNetworkQosPolicies())
	carapace.Gen(network_qos_policy_setCmd).PositionalCompletion(action.ActionNetworkQosPolicies())
	carapace.Gen(network_qos_policy_showCmd).PositionalCompletion(action.ActionNetworkQosPolicies())
	carapace.Gen(network_qos_rule_deleteCmd).PositionalAnyCompletion(action.ActionNetworkQosPolicies())
	carapace.Gen(network_qos_rule_setCmd).PositionalCompletion(action.ActionNetworkQosPolicies())
	carapace.Gen(network_qos_rule_showCmd).PositionalCompletion(action.ActionNetworkQosPolicies())
	carapace.Gen(network_rbac_deleteCmd).PositionalAnyCompletion(action.ActionNetworkRbacPolicies())
	carapace.Gen(network_rbac_setCmd).PositionalCompletion(action.ActionNetworkRbacPolicies())
	carapace.Gen(network_rbac_showCmd).PositionalCompletion(action.ActionNetworkRbacPolicies())
	carapace.Gen(network_segment_deleteCmd).PositionalAnyCompletion(action.ActionNetworkSegments())
	carapace.Gen(network_segment_range_deleteCmd).PositionalAnyCompletion(action.ActionNetworkSegmentRanges())
	carapace.Gen(network_segment_range_setCmd).PositionalCompletion(action.ActionNetworkSegmentRanges())
	carapace.Gen(network_segment_range_showCmd).PositionalCompletion(action.ActionNetworkSegmentRanges())
	carapace.Gen(network_segment_setCmd).PositionalCompletion(action.ActionNetworkSegments())
	carapace.Gen(network_segment_showCmd).PositionalCompletion(action.ActionNetworkSegments())
	carapace.Gen(network_setCmd).PositionalCompletion(action.ActionNetworks())
	carapace.Gen(network_showCmd).PositionalCompletion(action.ActionNetworks())
	carapace.Gen(network_trunk_deleteCmd).PositionalAnyCompletion(action.ActionNetworkTrunks())
	carapace.Gen(network_trunk_setCmd).PositionalCompletion(action.ActionNetworkTrunks())
	carapace.Gen(network_trunk_showCmd).PositionalCompletion(action.ActionNetworkTrunks())
	carapace.Gen(network_trunk_subport_addCmd).PositionalCompletion(action.ActionNetworkTrunks(), action.ActionPorts())
	carapace.Gen(network_trunk_subport_listCmd).PositionalCompletion(action.ActionNetworkTrunks())
	carapace.Gen(network_trunk_subport_removeCmd).PositionalCompletion(action.ActionNetworkTrunks(), action.ActionPorts())
	carapace.Gen(network_trunk_unsetCmd).PositionalCompletion(action.ActionNetworkTrunks())
	carapace.Gen(network_unsetCmd).PositionalCompletion(action.ActionNetworks())
	carapace.Gen(port_deleteCmd).PositionalAnyCompletion(action.ActionPorts())
	carapace.Gen(port_setCmd).PositionalCompletion(action.ActionPorts())
	carapace.Gen(port_showCmd).PositionalCompletion(action.ActionPorts())
	carapace.Gen(port_unsetCmd).PositionalCompletion(action.ActionPorts())
	carapace.Gen(project_deleteCmd).PositionalAnyCompletion(action.ActionProjects())
	carapace.Gen(project_setCmd).PositionalCompletion(action.ActionProjects())
	carapace.Gen(project_showCmd).PositionalCompletion(action.ActionProjects())
	carapace.Gen(project_unsetCmd).PositionalCompletion(action.ActionProjects())
	carapace.Gen(quota_deleteCmd).PositionalCompletion(action.ActionProjects())
	carapace.Gen(quota_setCmd).PositionalCompletion(action.ActionProjects())
	carapace.Gen(quota_showCmd).PositionalCompletion(action.ActionProjects())
	carapace.Gen(region_deleteCmd).PositionalAnyCompletion(action.ActionRegions())
	carapace.Gen(region_setCmd).PositionalCompletion(action.ActionRegions())
	carapace.Gen(region_showCmd).PositionalCompletion(action.ActionRegions())
	carapace.Gen(registered_limit_deleteCmd).PositionalAnyCompletion(action.ActionRegisteredLimits())
	carapace.Gen(registered_limit_setCmd).PositionalCompletion(action.ActionRegisteredLimits())
	carapace.Gen(registered_limit_showCmd).PositionalCompletion(action.ActionRegisteredLimits())
	carapace.Gen(role_deleteCmd).PositionalAnyCompletion(action.ActionRoles())
	carapace.Gen(role_setCmd).PositionalCompletion(action.ActionRoles())
	carapace.Gen(role_showCmd).PositionalCompletion(action.ActionRoles())
	carapace.Gen(router_deleteCmd).PositionalAnyCompletion(action.ActionRouters())
	carapace.Gen(router_setCmd).PositionalCompletion(action.ActionRouters())
	carapace.Gen(router_showCmd).PositionalCompletion(action.ActionRouters())
	carapace.Gen(router_unsetCmd).PositionalCompletion(action.ActionRouters())
	carapace.Gen(security_group_deleteCmd).PositionalAnyCompletion(action.ActionSecurityGroups())
	carapace.Gen(security_group_rule_deleteCmd).PositionalAnyCompletion(action.ActionSecurityGroupRules())
	carapace.Gen(security_group_rule_showCmd).PositionalCompletion(action.ActionSecurityGroupRules())
	carapace.Gen(security_group_setCmd).PositionalCompletion(action.ActionSecurityGroups())
	carapace.Gen(security_group_showCmd).PositionalCompletion(action.ActionSecurityGroups())
	carapace.Gen(security_group_unsetCmd).PositionalCompletion(action.ActionSecurityGroups())
	carapace.Gen(server_add_floating_ipCmd).PositionalCompletion(action.ActionServers(), action.ActionFloatingIPs())
	carapace.Gen(server_add_networkCmd).PositionalCompletion(action.ActionServers(), action.ActionNetworks())
	carapace.Gen(server_add_portCmd).PositionalCompletion(action.ActionServers(), action.ActionPorts())
	carapace.Gen(server_add_security_groupCmd).PositionalCompletion(action.ActionServers(), action.ActionSecurityGroups())
	carapace.Gen(server_add_shareCmd).PositionalCompletion(action.ActionServers(), action.ActionShares())
	carapace.Gen(server_add_volumeCmd).PositionalCompletion(action.ActionServers(), action.ActionVolumes())
	carapace.Gen(server_backup_createCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_deleteCmd).PositionalAnyCompletion(action.ActionServers())
	carapace.Gen(server_dump_createCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_evacuateCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_event_listCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_event_showCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_group_deleteCmd).PositionalAnyCompletion(action.ActionServerGroups())
	carapace.Gen(server_group_showCmd).PositionalCompletion(action.ActionServerGroups())
	carapace.Gen(server_image_createCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_lockCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_migrateCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_migrate_confirmCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_migrate_revertCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_migration_abortCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_migration_force_completeCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_migration_showCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_pauseCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_rebootCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_rebuildCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_remove_floating_ipCmd).PositionalCompletion(action.ActionServers(), action.ActionFloatingIPs())
	carapace.Gen(server_remove_networkCmd).PositionalCompletion(action.ActionServers(), action.ActionNetworks())
	carapace.Gen(server_remove_portCmd).PositionalCompletion(action.ActionServers(), action.ActionPorts())
	carapace.Gen(server_remove_security_groupCmd).PositionalCompletion(action.ActionServers(), action.ActionSecurityGroups())
	carapace.Gen(server_remove_shareCmd).PositionalCompletion(action.ActionServers(), action.ActionShares())
	carapace.Gen(server_remove_volumeCmd).PositionalCompletion(action.ActionServers(), action.ActionVolumes())
	carapace.Gen(server_rescueCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_resizeCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_resize_confirmCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_resize_revertCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_restoreCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_resumeCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_setCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_share_listCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_share_showCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_shelveCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_showCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_sshCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_startCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_stopCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_suspendCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_unlockCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_unpauseCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_unrescueCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_unsetCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_unshelveCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_volume_listCmd).PositionalCompletion(action.ActionServers())
	carapace.Gen(server_volume_setCmd).PositionalCompletion(action.ActionServers(), action.ActionVolumes())
	carapace.Gen(service_deleteCmd).PositionalAnyCompletion(action.ActionServices())
	carapace.Gen(service_setCmd).PositionalCompletion(action.ActionServices())
	carapace.Gen(service_showCmd).PositionalCompletion(action.ActionServices())
	carapace.Gen(share_access_createCmd).PositionalCompletion(action.ActionShares())
	carapace.Gen(share_access_deleteCmd).PositionalCompletion(action.ActionShares(), action.ActionShareAccessRules())
	carapace.Gen(share_access_listCmd).PositionalCompletion(action.ActionShares())
	carapace.Gen(share_backup_deleteCmd).PositionalAnyCompletion(action.ActionShareBackups())
	carapace.Gen(share_backup_restoreCmd).PositionalCompletion(action.ActionShareBackups())
	carapace.Gen(share_backup_setCmd).PositionalCompletion(action.ActionShareBackups())
	carapace.Gen(share_backup_showCmd).PositionalCompletion(action.ActionShareBackups())
	carapace.Gen(share_backup_unsetCmd).PositionalCompletion(action.ActionShareBackups())
	carapace.Gen(share_deleteCmd).PositionalAnyCompletion(action.ActionShares())
	carapace.Gen(share_group_deleteCmd).PositionalAnyCompletion(action.ActionShareGroups())
	carapace.Gen(share_group_setCmd).PositionalCompletion(action.ActionShareGroups())
	carapace.Gen(share_group_showCmd).PositionalCompletion(action.ActionShareGroups())
	carapace.Gen(share_group_snapshot_deleteCmd).PositionalAnyCompletion(action.ActionShareGroupSnapshots())
	carapace.Gen(share_group_snapshot_setCmd).PositionalCompletion(action.ActionShareGroupSnapshots())
	carapace.Gen(share_group_snapshot_showCmd).PositionalCompletion(action.ActionShareGroupSnapshots())
	carapace.Gen(share_group_snapshot_unsetCmd).PositionalCompletion(action.ActionShareGroupSnapshots())
	carapace.Gen(share_group_type_access_createCmd).PositionalCompletion(action.ActionShareGroupTypes())
	carapace.Gen(share_group_type_access_deleteCmd).PositionalCompletion(action.ActionShareGroupTypes())
	carapace.Gen(share_group_type_access_listCmd).PositionalCompletion(action.ActionShareGroupTypes())
	carapace.Gen(share_group_type_deleteCmd).PositionalAnyCompletion(action.ActionShareGroupTypes())
	carapace.Gen(share_group_type_setCmd).PositionalCompletion(action.ActionShareGroupTypes())
	carapace.Gen(share_group_type_showCmd).PositionalCompletion(action.ActionShareGroupTypes())
	carapace.Gen(share_group_type_unsetCmd).PositionalCompletion(action.ActionShareGroupTypes())
	carapace.Gen(share_group_unsetCmd).PositionalCompletion(action.ActionShareGroups())
	carapace.Gen(share_instance_deleteCmd).PositionalAnyCompletion(action.ActionShareInstances())
	carapace.Gen(share_instance_setCmd).PositionalCompletion(action.ActionShareInstances())
	carapace.Gen(share_instance_showCmd).PositionalCompletion(action.ActionShareInstances())
	carapace.Gen(share_lock_deleteCmd).PositionalAnyCompletion(action.ActionShareLocks())
	carapace.Gen(share_lock_setCmd).PositionalCompletion(action.ActionShareLocks())
	carapace.Gen(share_lock_showCmd).PositionalCompletion(action.ActionShareLocks())
	carapace.Gen(share_lock_unsetCmd).PositionalCompletion(action.ActionShareLocks())
	carapace.Gen(share_message_deleteCmd).PositionalAnyCompletion(action.ActionShareMessages())
	carapace.Gen(share_message_showCmd).PositionalCompletion(action.ActionShareMessages())
	carapace.Gen(share_network_deleteCmd).PositionalAnyCompletion(action.ActionShareNetworks())
	carapace.Gen(share_network_setCmd).PositionalCompletion(action.ActionShareNetworks())
	carapace.Gen(share_network_showCmd).PositionalCompletion(action.ActionShareNetworks())
	carapace.Gen(share_network_subnet_createCmd).PositionalCompletion(action.ActionShareNetworks())
	carapace.Gen(share_network_unsetCmd).PositionalCompletion(action.ActionShareNetworks())
	carapace.Gen(share_qos_type_deleteCmd).PositionalAnyCompletion(action.ActionShareQosTypes())
	carapace.Gen(share_qos_type_setCmd).PositionalCompletion(action.ActionShareQosTypes())
	carapace.Gen(share_qos_type_showCmd).PositionalCompletion(action.ActionShareQosTypes())
	carapace.Gen(share_qos_type_unsetCmd).PositionalCompletion(action.ActionShareQosTypes())
	carapace.Gen(share_replica_deleteCmd).PositionalAnyCompletion(action.ActionShareReplicas())
	carapace.Gen(share_replica_setCmd).PositionalCompletion(action.ActionShareReplicas())
	carapace.Gen(share_replica_showCmd).PositionalCompletion(action.ActionShareReplicas())
	carapace.Gen(share_replica_unsetCmd).PositionalCompletion(action.ActionShareReplicas())
	carapace.Gen(share_security_service_deleteCmd).PositionalAnyCompletion(action.ActionShareSecurityServices())
	carapace.Gen(share_security_service_setCmd).PositionalCompletion(action.ActionShareSecurityServices())
	carapace.Gen(share_security_service_showCmd).PositionalCompletion(action.ActionShareSecurityServices())
	carapace.Gen(share_security_service_unsetCmd).PositionalCompletion(action.ActionShareSecurityServices())
	carapace.Gen(share_server_abandonCmd).PositionalCompletion(action.ActionShareServers())
	carapace.Gen(share_server_deleteCmd).PositionalAnyCompletion(action.ActionShareServers())
	carapace.Gen(share_server_setCmd).PositionalCompletion(action.ActionShareServers())
	carapace.Gen(share_server_showCmd).PositionalCompletion(action.ActionShareServers())
	carapace.Gen(share_setCmd).PositionalCompletion(action.ActionShares())
	carapace.Gen(share_showCmd).PositionalCompletion(action.ActionShares())
	carapace.Gen(share_snapshot_access_createCmd).PositionalCompletion(action.ActionShareSnapshots())
	carapace.Gen(share_snapshot_access_deleteCmd).PositionalCompletion(action.ActionShareSnapshots())
	carapace.Gen(share_snapshot_access_listCmd).PositionalCompletion(action.ActionShareSnapshots())
	carapace.Gen(share_snapshot_deleteCmd).PositionalAnyCompletion(action.ActionShareSnapshots())
	carapace.Gen(share_snapshot_setCmd).PositionalCompletion(action.ActionShareSnapshots())
	carapace.Gen(share_snapshot_showCmd).PositionalCompletion(action.ActionShareSnapshots())
	carapace.Gen(share_snapshot_unsetCmd).PositionalCompletion(action.ActionShareSnapshots())
	carapace.Gen(share_transfer_acceptCmd).PositionalCompletion(action.ActionShareTransfers())
	carapace.Gen(share_transfer_deleteCmd).PositionalAnyCompletion(action.ActionShareTransfers())
	carapace.Gen(share_transfer_showCmd).PositionalCompletion(action.ActionShareTransfers())
	carapace.Gen(share_type_access_createCmd).PositionalCompletion(action.ActionShareTypes())
	carapace.Gen(share_type_access_deleteCmd).PositionalCompletion(action.ActionShareTypes())
	carapace.Gen(share_type_access_listCmd).PositionalCompletion(action.ActionShareTypes())
	carapace.Gen(share_type_deleteCmd).PositionalAnyCompletion(action.ActionShareTypes())
	carapace.Gen(share_type_setCmd).PositionalCompletion(action.ActionShareTypes())
	carapace.Gen(share_type_showCmd).PositionalCompletion(action.ActionShareTypes())
	carapace.Gen(share_type_unsetCmd).PositionalCompletion(action.ActionShareTypes())
	carapace.Gen(share_unsetCmd).PositionalCompletion(action.ActionShares())
	carapace.Gen(subnet_deleteCmd).PositionalAnyCompletion(action.ActionSubnets())
	carapace.Gen(subnet_pool_add_networkCmd).PositionalCompletion(action.ActionNetworks(), action.ActionSubnetPools())
	carapace.Gen(subnet_pool_deleteCmd).PositionalAnyCompletion(action.ActionSubnetPools())
	carapace.Gen(subnet_pool_setCmd).PositionalCompletion(action.ActionSubnetPools())
	carapace.Gen(subnet_pool_showCmd).PositionalCompletion(action.ActionSubnetPools())
	carapace.Gen(subnet_pool_unsetCmd).PositionalCompletion(action.ActionSubnetPools())
	carapace.Gen(subnet_setCmd).PositionalCompletion(action.ActionSubnets())
	carapace.Gen(subnet_showCmd).PositionalCompletion(action.ActionSubnets())
	carapace.Gen(subnet_unsetCmd).PositionalCompletion(action.ActionSubnets())
	carapace.Gen(trust_deleteCmd).PositionalAnyCompletion(action.ActionTrusts())
	carapace.Gen(trust_showCmd).PositionalCompletion(action.ActionTrusts())
	carapace.Gen(user_deleteCmd).PositionalAnyCompletion(action.ActionUsers())
	carapace.Gen(user_setCmd).PositionalCompletion(action.ActionUsers())
	carapace.Gen(user_showCmd).PositionalCompletion(action.ActionUsers())
	carapace.Gen(volume_attachment_createCmd).PositionalCompletion(action.ActionVolumes())
	carapace.Gen(volume_attachment_deleteCmd).PositionalAnyCompletion(action.ActionVolumeAttachments())
	carapace.Gen(volume_attachment_setCmd).PositionalCompletion(action.ActionVolumeAttachments())
	carapace.Gen(volume_attachment_showCmd).PositionalCompletion(action.ActionVolumeAttachments())
	carapace.Gen(volume_backup_deleteCmd).PositionalAnyCompletion(action.ActionVolumeBackups())
	carapace.Gen(volume_backup_restoreCmd).PositionalCompletion(action.ActionVolumeBackups())
	carapace.Gen(volume_backup_setCmd).PositionalCompletion(action.ActionVolumeBackups())
	carapace.Gen(volume_backup_showCmd).PositionalCompletion(action.ActionVolumeBackups())
	carapace.Gen(volume_backup_unsetCmd).PositionalCompletion(action.ActionVolumeBackups())
	carapace.Gen(volume_deleteCmd).PositionalAnyCompletion(action.ActionVolumes())
	carapace.Gen(volume_group_add_volumeCmd).PositionalCompletion(action.ActionVolumeGroups(), action.ActionVolumes())
	carapace.Gen(volume_group_deleteCmd).PositionalAnyCompletion(action.ActionVolumeGroups())
	carapace.Gen(volume_group_remove_volumeCmd).PositionalCompletion(action.ActionVolumeGroups(), action.ActionVolumes())
	carapace.Gen(volume_group_setCmd).PositionalCompletion(action.ActionVolumeGroups())
	carapace.Gen(volume_group_showCmd).PositionalCompletion(action.ActionVolumeGroups())
	carapace.Gen(volume_group_snapshot_deleteCmd).PositionalAnyCompletion(action.ActionVolumeGroupSnapshots())
	carapace.Gen(volume_group_snapshot_showCmd).PositionalCompletion(action.ActionVolumeGroupSnapshots())
	carapace.Gen(volume_group_type_deleteCmd).PositionalAnyCompletion(action.ActionVolumeGroupTypes())
	carapace.Gen(volume_group_type_setCmd).PositionalCompletion(action.ActionVolumeGroupTypes())
	carapace.Gen(volume_group_type_showCmd).PositionalCompletion(action.ActionVolumeGroupTypes())
	carapace.Gen(volume_qos_deleteCmd).PositionalAnyCompletion(action.ActionVolumeQosSpecs())
	carapace.Gen(volume_qos_setCmd).PositionalCompletion(action.ActionVolumeQosSpecs())
	carapace.Gen(volume_qos_showCmd).PositionalCompletion(action.ActionVolumeQosSpecs())
	carapace.Gen(volume_qos_unsetCmd).PositionalCompletion(action.ActionVolumeQosSpecs())
	carapace.Gen(volume_setCmd).PositionalCompletion(action.ActionVolumes())
	carapace.Gen(volume_showCmd).PositionalCompletion(action.ActionVolumes())
	carapace.Gen(volume_snapshot_deleteCmd).PositionalAnyCompletion(action.ActionVolumeSnapshots())
	carapace.Gen(volume_snapshot_setCmd).PositionalCompletion(action.ActionVolumeSnapshots())
	carapace.Gen(volume_snapshot_showCmd).PositionalCompletion(action.ActionVolumeSnapshots())
	carapace.Gen(volume_snapshot_unsetCmd).PositionalCompletion(action.ActionVolumeSnapshots())
	carapace.Gen(volume_transfer_request_acceptCmd).PositionalCompletion(action.ActionVolumeTransfers())
	carapace.Gen(volume_transfer_request_deleteCmd).PositionalAnyCompletion(action.ActionVolumeTransfers())
	carapace.Gen(volume_transfer_request_showCmd).PositionalCompletion(action.ActionVolumeTransfers())
	carapace.Gen(volume_type_deleteCmd).PositionalAnyCompletion(action.ActionVolumeTypes())
	carapace.Gen(volume_type_setCmd).PositionalCompletion(action.ActionVolumeTypes())
	carapace.Gen(volume_type_showCmd).PositionalCompletion(action.ActionVolumeTypes())
	carapace.Gen(volume_type_unsetCmd).PositionalCompletion(action.ActionVolumeTypes())
	carapace.Gen(volume_unsetCmd).PositionalCompletion(action.ActionVolumes())

	carapace.Gen(access_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(address_group_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(address_group_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(address_scope_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(address_scope_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(aggregate_listCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
	})
	carapace.Gen(application_credential_createCmd).FlagCompletion(carapace.ActionMap{
		"role": action.ActionRoles(),
	})
	carapace.Gen(application_credential_listCmd).FlagCompletion(carapace.ActionMap{
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(availability_zone_listCmd).FlagCompletion(carapace.ActionMap{
		"network": action.ActionNetworks(),
		"volume":  action.ActionVolumes(),
	})
	carapace.Gen(bgp_peer_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgp_speaker_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgpvpn_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgpvpn_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgpvpn_network_association_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgpvpn_port_association_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgpvpn_router_association_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(consistency_group_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"volume-type":       action.ActionVolumeTypes(),
	})
	carapace.Gen(consistency_group_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"consistency-group": action.ActionConsistencyGroups(),
	})
	carapace.Gen(consistency_group_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"consistency-group": action.ActionConsistencyGroups(),
	})
	carapace.Gen(credential_createCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
	carapace.Gen(credential_listCmd).FlagCompletion(carapace.ActionMap{
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(credential_setCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
		"user":    action.ActionUsers(),
	})
	carapace.Gen(ec2_credentials_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(ec2_credentials_deleteCmd).FlagCompletion(carapace.ActionMap{
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(ec2_credentials_listCmd).FlagCompletion(carapace.ActionMap{
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(ec2_credentials_showCmd).FlagCompletion(carapace.ActionMap{
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(endpoint_add_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(endpoint_createCmd).FlagCompletion(carapace.ActionMap{
		"region": action.ActionRegions(),
	})
	carapace.Gen(endpoint_group_add_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(endpoint_group_listCmd).FlagCompletion(carapace.ActionMap{
		"domain":  action.ActionDomains(),
		"project": action.ActionProjects(),
	})
	carapace.Gen(endpoint_group_remove_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(endpoint_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"region":         action.ActionRegions(),
	})
	carapace.Gen(endpoint_remove_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(endpoint_setCmd).FlagCompletion(carapace.ActionMap{
		"region": action.ActionRegions(),
	})
	carapace.Gen(extension_listCmd).FlagCompletion(carapace.ActionMap{
		"network": action.ActionNetworks(),
		"volume":  action.ActionVolumes(),
	})
	carapace.Gen(firewall_group_createCmd).FlagCompletion(carapace.ActionMap{
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(firewall_group_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(firewall_group_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(firewall_group_setCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(firewall_group_unsetCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(flavor_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(flavor_setCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(flavor_unsetCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(floating_ip_createCmd).FlagCompletion(carapace.ActionMap{
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"qos-policy":     action.ActionNetworkQosPolicies(),
		"router":         action.ActionRouters(),
		"subnet":         action.ActionSubnets(),
	})
	carapace.Gen(floating_ip_listCmd).FlagCompletion(carapace.ActionMap{
		"network":        action.ActionNetworks(),
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"router":         action.ActionRouters(),
	})
	carapace.Gen(floating_ip_port_forwarding_createCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(floating_ip_port_forwarding_listCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(floating_ip_port_forwarding_setCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(floating_ip_setCmd).FlagCompletion(carapace.ActionMap{
		"port":       action.ActionPorts(),
		"qos-policy": action.ActionNetworkQosPolicies(),
		"router":     action.ActionRouters(),
	})
	carapace.Gen(floating_ip_unsetCmd).FlagCompletion(carapace.ActionMap{
		"port":       action.ActionPorts(),
		"qos-policy": action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(group_add_userCmd).FlagCompletion(carapace.ActionMap{
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(group_contains_userCmd).FlagCompletion(carapace.ActionMap{
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(group_createCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(group_deleteCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(group_listCmd).FlagCompletion(carapace.ActionMap{
		"domain":      action.ActionDomains(),
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(group_remove_userCmd).FlagCompletion(carapace.ActionMap{
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(group_setCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(group_showCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(identity_provider_createCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(image_add_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(image_createCmd).FlagCompletion(carapace.ActionMap{
		"file":           carapace.ActionFiles(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sign-key-path":  carapace.ActionFiles(),
		"volume":         action.ActionVolumes(),
	})
	carapace.Gen(image_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(image_member_getCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(image_member_listCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(image_remove_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(image_saveCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
	carapace.Gen(image_setCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(image_stageCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
	carapace.Gen(ip_availability_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(keypair_createCmd).FlagCompletion(carapace.ActionMap{
		"public-key":  carapace.ActionFiles(),
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(keypair_deleteCmd).FlagCompletion(carapace.ActionMap{
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(keypair_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(keypair_showCmd).FlagCompletion(carapace.ActionMap{
		"public-key":  carapace.ActionFiles(),
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(limit_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"region":         action.ActionRegions(),
	})
	carapace.Gen(limit_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"region":         action.ActionRegions(),
	})
	carapace.Gen(limits_showCmd).FlagCompletion(carapace.ActionMap{
		"domain":         action.ActionDomains(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(local_ip_association_createCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(local_ip_association_listCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(local_ip_createCmd).FlagCompletion(carapace.ActionMap{
		"network":        action.ActionNetworks(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(local_ip_listCmd).FlagCompletion(carapace.ActionMap{
		"network":        action.ActionNetworks(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_agent_listCmd).FlagCompletion(carapace.ActionMap{
		"network": action.ActionNetworks(),
		"router":  action.ActionRouters(),
	})
	carapace.Gen(network_auto_allocated_topology_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_auto_allocated_topology_deleteCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone-hint": action.ActionAvailabilityZones(),
		"project":                action.ActionProjects(),
		"project-domain":         action.ActionDomains(),
		"qos-policy":             action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(network_flavor_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_l3_conntrack_helper_createCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(network_l3_conntrack_helper_listCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(network_l3_conntrack_helper_setCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(network_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_meter_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_meter_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_qos_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_qos_policy_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_rbac_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_segment_createCmd).FlagCompletion(carapace.ActionMap{
		"network": action.ActionNetworks(),
	})
	carapace.Gen(network_segment_listCmd).FlagCompletion(carapace.ActionMap{
		"network": action.ActionNetworks(),
	})
	carapace.Gen(network_segment_range_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_setCmd).FlagCompletion(carapace.ActionMap{
		"qos-policy": action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(network_trunk_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(object_saveCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
	carapace.Gen(port_createCmd).FlagCompletion(carapace.ActionMap{
		"network":        action.ActionNetworks(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"qos-policy":     action.ActionNetworkQosPolicies(),
		"security-group": action.ActionSecurityGroups(),
	})
	carapace.Gen(port_listCmd).FlagCompletion(carapace.ActionMap{
		"network":        action.ActionNetworks(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"router":         action.ActionRouters(),
		"security-group": action.ActionSecurityGroups(),
		"server":         action.ActionServers(),
	})
	carapace.Gen(port_setCmd).FlagCompletion(carapace.ActionMap{
		"qos-policy":     action.ActionNetworkQosPolicies(),
		"security-group": action.ActionSecurityGroups(),
	})
	carapace.Gen(port_unsetCmd).FlagCompletion(carapace.ActionMap{
		"qos-policy":     action.ActionNetworkQosPolicies(),
		"security-group": action.ActionSecurityGroups(),
	})
	carapace.Gen(project_cleanupCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(project_createCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(project_deleteCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(project_listCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
		"user":   action.ActionUsers(),
	})
	carapace.Gen(project_setCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(project_showCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(quota_deleteCmd).FlagCompletion(carapace.ActionMap{
		"network": action.ActionNetworks(),
		"volume":  action.ActionVolumes(),
	})
	carapace.Gen(quota_listCmd).FlagCompletion(carapace.ActionMap{
		"network": action.ActionNetworks(),
		"volume":  action.ActionVolumes(),
	})
	carapace.Gen(quota_setCmd).FlagCompletion(carapace.ActionMap{
		"volume-type": action.ActionVolumeTypes(),
	})
	carapace.Gen(quota_showCmd).FlagCompletion(carapace.ActionMap{
		"network": action.ActionNetworks(),
		"volume":  action.ActionVolumes(),
	})
	carapace.Gen(registered_limit_createCmd).FlagCompletion(carapace.ActionMap{
		"region": action.ActionRegions(),
	})
	carapace.Gen(registered_limit_listCmd).FlagCompletion(carapace.ActionMap{
		"region": action.ActionRegions(),
	})
	carapace.Gen(registered_limit_setCmd).FlagCompletion(carapace.ActionMap{
		"region": action.ActionRegions(),
	})
	carapace.Gen(request_token_authorizeCmd).FlagCompletion(carapace.ActionMap{
		"role": action.ActionRoles(),
	})
	carapace.Gen(request_token_createCmd).FlagCompletion(carapace.ActionMap{
		"domain":  action.ActionDomains(),
		"project": action.ActionProjects(),
	})
	carapace.Gen(role_addCmd).FlagCompletion(carapace.ActionMap{
		"domain":         action.ActionDomains(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(role_assignment_listCmd).FlagCompletion(carapace.ActionMap{
		"domain":         action.ActionDomains(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"role":           action.ActionRoles(),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(role_createCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(role_deleteCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(role_listCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(role_removeCmd).FlagCompletion(carapace.ActionMap{
		"domain":         action.ActionDomains(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(role_setCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(role_showCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log-file":       carapace.ActionFiles(),
		"os-cacert":      carapace.ActionFiles(),
		"os-cert":        carapace.ActionFiles(),
		"os-cloud":       action.ActionClouds(),
		"os-key":         carapace.ActionFiles(),
		"os-region-name": action.ActionRegions(),
	})
	carapace.Gen(router_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone-hint": action.ActionAvailabilityZones(),
		"flavor":                 action.ActionFlavors(),
		"project":                action.ActionProjects(),
		"project-domain":         action.ActionDomains(),
		"qos-policy":             action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(router_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(router_ndp_proxy_createCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(router_ndp_proxy_listCmd).FlagCompletion(carapace.ActionMap{
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"router":         action.ActionRouters(),
	})
	carapace.Gen(router_setCmd).FlagCompletion(carapace.ActionMap{
		"qos-policy": action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(router_unsetCmd).FlagCompletion(carapace.ActionMap{
		"qos-policy": action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(security_group_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(security_group_default_statefulness_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(security_group_default_statefulness_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(security_group_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(security_group_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(security_group_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(server_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"block-device":      carapace.ActionFiles(),
		"file":              carapace.ActionFiles(),
		"flavor":            action.ActionFlavors(),
		"image":             action.ActionImages(),
		"key-name":          action.ActionKeypairs(),
		"network":           action.ActionNetworks(),
		"port":              action.ActionPorts(),
		"security-group":    action.ActionSecurityGroups(),
		"server-group":      action.ActionServerGroups(),
		"user-data":         carapace.ActionFiles(),
		"volume":            action.ActionVolumes(),
	})
	carapace.Gen(server_listCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"flavor":            action.ActionFlavors(),
		"image":             action.ActionImages(),
		"key-name":          action.ActionKeypairs(),
		"project":           action.ActionProjects(),
		"project-domain":    action.ActionDomains(),
		"user":              action.ActionUsers(),
		"user-domain":       action.ActionDomains(),
	})
	carapace.Gen(server_migration_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"server":         action.ActionServers(),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(server_rebuildCmd).FlagCompletion(carapace.ActionMap{
		"image":     action.ActionImages(),
		"key-name":  action.ActionKeypairs(),
		"user-data": carapace.ActionFiles(),
	})
	carapace.Gen(server_rescueCmd).FlagCompletion(carapace.ActionMap{
		"image": action.ActionImages(),
	})
	carapace.Gen(server_resizeCmd).FlagCompletion(carapace.ActionMap{
		"flavor": action.ActionFlavors(),
	})
	carapace.Gen(server_sshCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(server_unshelveCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
	})
	carapace.Gen(share_adoptCmd).FlagCompletion(carapace.ActionMap{
		"share-type": action.ActionShareTypes(),
	})
	carapace.Gen(share_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"share-group":       action.ActionShareGroups(),
		"share-network":     action.ActionShareNetworks(),
		"share-type":        action.ActionShareTypes(),
	})
	carapace.Gen(share_deleteCmd).FlagCompletion(carapace.ActionMap{
		"share-group": action.ActionShareGroups(),
	})
	carapace.Gen(share_group_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"share-network":     action.ActionShareNetworks(),
	})
	carapace.Gen(share_group_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"share-network":  action.ActionShareNetworks(),
		"share-server":   action.ActionShareServers(),
	})
	carapace.Gen(share_group_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"share-group": action.ActionShareGroups(),
	})
	carapace.Gen(share_group_type_access_createCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(share_group_type_access_deleteCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(share_listCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"project":           action.ActionProjects(),
		"project-domain":    action.ActionDomains(),
		"share-group":       action.ActionShareGroups(),
		"share-network":     action.ActionShareNetworks(),
		"share-server":      action.ActionShareServers(),
		"share-type":        action.ActionShareTypes(),
		"user":              action.ActionUsers(),
		"user-domain":       action.ActionDomains(),
	})
	carapace.Gen(share_lock_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(share_network_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
	})
	carapace.Gen(share_network_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(share_network_subnet_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
	})
	carapace.Gen(share_pool_listCmd).FlagCompletion(carapace.ActionMap{
		"share-type": action.ActionShareTypes(),
	})
	carapace.Gen(share_quota_deleteCmd).FlagCompletion(carapace.ActionMap{
		"share-type": action.ActionShareTypes(),
		"user":       action.ActionUsers(),
	})
	carapace.Gen(share_quota_setCmd).FlagCompletion(carapace.ActionMap{
		"share-type": action.ActionShareTypes(),
		"user":       action.ActionUsers(),
	})
	carapace.Gen(share_quota_showCmd).FlagCompletion(carapace.ActionMap{
		"share-type": action.ActionShareTypes(),
		"user":       action.ActionUsers(),
	})
	carapace.Gen(share_replica_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"share-network":     action.ActionShareNetworks(),
	})
	carapace.Gen(share_security_service_createCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
		"server": action.ActionServers(),
		"user":   action.ActionUsers(),
	})
	carapace.Gen(share_security_service_listCmd).FlagCompletion(carapace.ActionMap{
		"domain":        action.ActionDomains(),
		"server":        action.ActionServers(),
		"share-network": action.ActionShareNetworks(),
		"user":          action.ActionUsers(),
	})
	carapace.Gen(share_security_service_setCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
		"server": action.ActionServers(),
		"user":   action.ActionUsers(),
	})
	carapace.Gen(share_security_service_unsetCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
		"server": action.ActionServers(),
		"user":   action.ActionUsers(),
	})
	carapace.Gen(share_server_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"share-network":  action.ActionShareNetworks(),
	})
	carapace.Gen(subnet_createCmd).FlagCompletion(carapace.ActionMap{
		"network":         action.ActionNetworks(),
		"network-segment": action.ActionNetworkSegments(),
		"project":         action.ActionProjects(),
		"project-domain":  action.ActionDomains(),
		"subnet-pool":     action.ActionSubnetPools(),
	})
	carapace.Gen(subnet_listCmd).FlagCompletion(carapace.ActionMap{
		"network":        action.ActionNetworks(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"subnet-pool":    action.ActionSubnetPools(),
	})
	carapace.Gen(subnet_pool_createCmd).FlagCompletion(carapace.ActionMap{
		"address-scope":  action.ActionAddressScopes(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(subnet_pool_listCmd).FlagCompletion(carapace.ActionMap{
		"address-scope":  action.ActionAddressScopes(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(subnet_pool_setCmd).FlagCompletion(carapace.ActionMap{
		"address-scope": action.ActionAddressScopes(),
	})
	carapace.Gen(subnet_setCmd).FlagCompletion(carapace.ActionMap{
		"network-segment": action.ActionNetworkSegments(),
	})
	carapace.Gen(tap_flow_createCmd).FlagCompletion(carapace.ActionMap{
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(tap_flow_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(tap_mirror_createCmd).FlagCompletion(carapace.ActionMap{
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(tap_mirror_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(tap_service_createCmd).FlagCompletion(carapace.ActionMap{
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(tap_service_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(trust_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"role":           action.ActionRoles(),
	})
	carapace.Gen(usage_showCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(user_createCmd).FlagCompletion(carapace.ActionMap{
		"domain":         action.ActionDomains(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(user_deleteCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(user_listCmd).FlagCompletion(carapace.ActionMap{
		"domain":  action.ActionDomains(),
		"project": action.ActionProjects(),
	})
	carapace.Gen(user_setCmd).FlagCompletion(carapace.ActionMap{
		"domain":         action.ActionDomains(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(user_showCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(volume_attachment_createCmd).FlagCompletion(carapace.ActionMap{
		"server": action.ActionServers(),
	})
	carapace.Gen(volume_attachment_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"server-id":      action.ActionServers(),
	})
	carapace.Gen(volume_backup_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
	})
	carapace.Gen(volume_backup_listCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
		"volume":  action.ActionVolumes(),
	})
	carapace.Gen(volume_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"consistency-group": action.ActionConsistencyGroups(),
		"image":             action.ActionImages(),
	})
	carapace.Gen(volume_group_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"volume-type":       action.ActionVolumeTypes(),
	})
	carapace.Gen(volume_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(volume_message_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(volume_qos_disassociateCmd).FlagCompletion(carapace.ActionMap{
		"volume-type": action.ActionVolumeTypes(),
	})
	carapace.Gen(volume_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"volume": action.ActionVolumes(),
	})
	carapace.Gen(volume_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"volume":         action.ActionVolumes(),
	})
	carapace.Gen(volume_type_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"project":           action.ActionProjects(),
		"project-domain":    action.ActionDomains(),
	})
	carapace.Gen(volume_type_listCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
	})
	carapace.Gen(volume_type_setCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"project":           action.ActionProjects(),
		"project-domain":    action.ActionDomains(),
	})
	carapace.Gen(volume_type_unsetCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(vpn_endpoint_group_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(vpn_ike_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(vpn_ipsec_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(vpn_ipsec_site_connection_createCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(vpn_service_createCmd).FlagCompletion(carapace.ActionMap{
		"flavor":         action.ActionFlavors(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"router":         action.ActionRouters(),
		"subnet":         action.ActionSubnets(),
	})
	carapace.Gen(vpn_service_setCmd).FlagCompletion(carapace.ActionMap{
		"flavor": action.ActionFlavors(),
		"subnet": action.ActionSubnets(),
	})
}
