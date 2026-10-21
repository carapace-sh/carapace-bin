package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
)

func init() {
	carapace.Gen(accelerator_arq_bindCmd).PositionalCompletion(action.ActionAcceleratorArqs())
	carapace.Gen(accelerator_arq_createCmd).PositionalCompletion(action.ActionAcceleratorDeviceProfiles())
	carapace.Gen(accelerator_arq_deleteCmd).PositionalAnyCompletion(action.ActionAcceleratorArqs())
	carapace.Gen(accelerator_arq_showCmd).PositionalCompletion(action.ActionAcceleratorArqs())
	carapace.Gen(accelerator_arq_unbindCmd).PositionalCompletion(action.ActionAcceleratorArqs())
	carapace.Gen(accelerator_deployable_programCmd).PositionalCompletion(action.ActionAcceleratorDeployables(), action.ActionImages())
	carapace.Gen(accelerator_deployable_showCmd).PositionalCompletion(action.ActionAcceleratorDeployables())
	carapace.Gen(accelerator_device_attribute_createCmd).PositionalCompletion(action.ActionAcceleratorDeployables())
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
	carapace.Gen(aggregate_add_hostCmd).PositionalCompletion(action.ActionAggregates(), action.ActionComputeHosts())
	carapace.Gen(aggregate_cache_imageCmd).PositionalCompletion(action.ActionAggregates(), action.ActionImages())
	carapace.Gen(aggregate_deleteCmd).PositionalAnyCompletion(action.ActionAggregates())
	carapace.Gen(aggregate_remove_hostCmd).PositionalCompletion(action.ActionAggregates(), action.ActionComputeHosts())
	carapace.Gen(aggregate_setCmd).PositionalCompletion(action.ActionAggregates())
	carapace.Gen(aggregate_showCmd).PositionalCompletion(action.ActionAggregates())
	carapace.Gen(aggregate_unsetCmd).PositionalCompletion(action.ActionAggregates())
	carapace.Gen(application_credential_deleteCmd).PositionalAnyCompletion(action.ActionApplicationCredentials())
	carapace.Gen(application_credential_showCmd).PositionalCompletion(action.ActionApplicationCredentials())
	carapace.Gen(bgp_dragent_add_speakerCmd).PositionalCompletion(action.ActionBgpDragents(), action.ActionBgpSpeakers())
	carapace.Gen(bgp_dragent_remove_speakerCmd).PositionalCompletion(action.ActionBgpDragents(), action.ActionBgpSpeakers())
	carapace.Gen(bgp_peer_deleteCmd).PositionalAnyCompletion(action.ActionBgpPeers())
	carapace.Gen(bgp_peer_setCmd).PositionalCompletion(action.ActionBgpPeers())
	carapace.Gen(bgp_peer_showCmd).PositionalCompletion(action.ActionBgpPeers())
	carapace.Gen(bgp_speaker_add_networkCmd).PositionalCompletion(action.ActionBgpSpeakers(), action.ActionNetworks())
	carapace.Gen(bgp_speaker_add_peerCmd).PositionalCompletion(action.ActionBgpSpeakers(), action.ActionBgpPeers())
	carapace.Gen(bgp_speaker_deleteCmd).PositionalAnyCompletion(action.ActionBgpSpeakers())
	carapace.Gen(bgp_speaker_remove_networkCmd).PositionalCompletion(action.ActionBgpSpeakers(), action.ActionNetworks())
	carapace.Gen(bgp_speaker_remove_peerCmd).PositionalCompletion(action.ActionBgpSpeakers(), action.ActionBgpPeers())
	carapace.Gen(bgp_speaker_setCmd).PositionalCompletion(action.ActionBgpSpeakers())
	carapace.Gen(bgp_speaker_showCmd).PositionalCompletion(action.ActionBgpSpeakers())
	carapace.Gen(bgpvpn_deleteCmd).PositionalAnyCompletion(action.ActionBgpvpns())
	carapace.Gen(bgpvpn_setCmd).PositionalCompletion(action.ActionBgpvpns())
	carapace.Gen(bgpvpn_showCmd).PositionalCompletion(action.ActionBgpvpns())
	carapace.Gen(bgpvpn_unsetCmd).PositionalCompletion(action.ActionBgpvpns())
	carapace.Gen(cached_image_deleteCmd).PositionalAnyCompletion(action.ActionCachedImages())
	carapace.Gen(compute_agent_deleteCmd).PositionalAnyCompletion(action.ActionComputeAgents())
	carapace.Gen(compute_agent_setCmd).PositionalCompletion(action.ActionComputeAgents())
	carapace.Gen(consistency_group_add_volumeCmd).PositionalCompletion(action.ActionConsistencyGroups(), action.ActionVolumes())
	carapace.Gen(consistency_group_deleteCmd).PositionalAnyCompletion(action.ActionConsistencyGroups())
	carapace.Gen(consistency_group_remove_volumeCmd).PositionalCompletion(action.ActionConsistencyGroups(), action.ActionVolumes())
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
	carapace.Gen(firewall_group_policy_add_ruleCmd).PositionalCompletion(action.ActionFirewallPolicies(), action.ActionFirewallRules())
	carapace.Gen(firewall_group_policy_deleteCmd).PositionalAnyCompletion(action.ActionFirewallPolicies())
	carapace.Gen(firewall_group_policy_remove_ruleCmd).PositionalCompletion(action.ActionFirewallPolicies(), action.ActionFirewallRules())
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
	carapace.Gen(floating_ip_port_forwarding_createCmd).PositionalCompletion(action.ActionFloatingIPs())
	carapace.Gen(floating_ip_port_forwarding_deleteCmd).PositionalCompletion(action.ActionFloatingIPs(), action.ActionFloatingIPPortForwardings())
	carapace.Gen(floating_ip_port_forwarding_listCmd).PositionalCompletion(action.ActionFloatingIPs())
	carapace.Gen(floating_ip_port_forwarding_setCmd).PositionalCompletion(action.ActionFloatingIPs(), action.ActionFloatingIPPortForwardings())
	carapace.Gen(floating_ip_port_forwarding_showCmd).PositionalCompletion(action.ActionFloatingIPs(), action.ActionFloatingIPPortForwardings())
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
	carapace.Gen(image_add_projectCmd).PositionalCompletion(action.ActionImages(), action.ActionProjectIDs())
	carapace.Gen(image_deleteCmd).PositionalAnyCompletion(action.ActionImages())
	carapace.Gen(image_importCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_member_getCmd).PositionalCompletion(action.ActionImages(), action.ActionProjectIDs())
	carapace.Gen(image_member_listCmd).PositionalCompletion(action.ActionImages())
	carapace.Gen(image_metadef_namespace_deleteCmd).PositionalAnyCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_namespace_setCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_namespace_showCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_namespace_unsetCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_object_createCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_object_deleteCmd).PositionalCompletion(action.ActionImageMetadefNamespaces(), action.ActionImageMetadefObjects())
	carapace.Gen(image_metadef_object_listCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_object_property_showCmd).PositionalCompletion(action.ActionImageMetadefNamespaces(), action.ActionImageMetadefObjects())
	carapace.Gen(image_metadef_object_showCmd).PositionalCompletion(action.ActionImageMetadefNamespaces(), action.ActionImageMetadefObjects())
	carapace.Gen(image_metadef_object_updateCmd).PositionalCompletion(action.ActionImageMetadefNamespaces(), action.ActionImageMetadefObjects())
	carapace.Gen(image_metadef_property_createCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_property_deleteCmd).PositionalCompletion(action.ActionImageMetadefNamespaces(), action.ActionImageMetadefProperties())
	carapace.Gen(image_metadef_property_listCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_metadef_property_setCmd).PositionalCompletion(action.ActionImageMetadefNamespaces(), action.ActionImageMetadefProperties())
	carapace.Gen(image_metadef_property_showCmd).PositionalCompletion(action.ActionImageMetadefNamespaces(), action.ActionImageMetadefProperties())
	carapace.Gen(image_metadef_resource_type_association_createCmd).PositionalCompletion(action.ActionImageMetadefNamespaces(), action.ActionImageMetadefResourceTypes())
	carapace.Gen(image_metadef_resource_type_association_deleteCmd).PositionalCompletion(action.ActionImageMetadefNamespaces(), action.ActionImageMetadefResourceTypes())
	carapace.Gen(image_metadef_resource_type_association_listCmd).PositionalCompletion(action.ActionImageMetadefNamespaces())
	carapace.Gen(image_remove_projectCmd).PositionalCompletion(action.ActionImages(), action.ActionProjects())
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
	carapace.Gen(local_ip_association_createCmd).PositionalCompletion(action.ActionLocalIPs())
	carapace.Gen(local_ip_association_deleteCmd).PositionalCompletion(action.ActionLocalIPs(), action.ActionPorts())
	carapace.Gen(local_ip_association_listCmd).PositionalCompletion(action.ActionLocalIPs())
	carapace.Gen(local_ip_deleteCmd).PositionalAnyCompletion(action.ActionLocalIPs())
	carapace.Gen(local_ip_setCmd).PositionalCompletion(action.ActionLocalIPs())
	carapace.Gen(local_ip_showCmd).PositionalCompletion(action.ActionLocalIPs())
	carapace.Gen(mapping_deleteCmd).PositionalAnyCompletion(action.ActionMappings())
	carapace.Gen(mapping_setCmd).PositionalCompletion(action.ActionMappings())
	carapace.Gen(mapping_showCmd).PositionalCompletion(action.ActionMappings())
	carapace.Gen(network_agent_add_networkCmd).PositionalCompletion(action.ActionNetworkAgents(), action.ActionNetworks())
	carapace.Gen(network_agent_add_routerCmd).PositionalCompletion(action.ActionNetworkAgents(), action.ActionRouters())
	carapace.Gen(network_agent_deleteCmd).PositionalAnyCompletion(action.ActionNetworkAgents())
	carapace.Gen(network_agent_remove_networkCmd).PositionalCompletion(action.ActionNetworkAgents(), action.ActionNetworks())
	carapace.Gen(network_agent_remove_routerCmd).PositionalCompletion(action.ActionNetworkAgents(), action.ActionRouters())
	carapace.Gen(network_agent_router_setCmd).PositionalCompletion(action.ActionNetworkAgents(), action.ActionRouters())
	carapace.Gen(network_agent_setCmd).PositionalCompletion(action.ActionNetworkAgents())
	carapace.Gen(network_agent_showCmd).PositionalCompletion(action.ActionNetworkAgents())
	carapace.Gen(network_deleteCmd).PositionalAnyCompletion(action.ActionNetworks())
	carapace.Gen(network_flavor_add_profileCmd).PositionalCompletion(action.ActionNetworkFlavors(), action.ActionNetworkFlavorProfiles())
	carapace.Gen(network_flavor_deleteCmd).PositionalAnyCompletion(action.ActionNetworkFlavors())
	carapace.Gen(network_flavor_remove_profileCmd).PositionalCompletion(action.ActionNetworkFlavors(), action.ActionNetworkFlavorProfiles())
	carapace.Gen(network_flavor_setCmd).PositionalCompletion(action.ActionNetworkFlavors())
	carapace.Gen(network_flavor_showCmd).PositionalCompletion(action.ActionNetworkFlavors())
	carapace.Gen(network_meter_deleteCmd).PositionalAnyCompletion(action.ActionNetworkMeters())
	carapace.Gen(network_meter_rule_createCmd).PositionalCompletion(action.ActionNetworkMeters())
	carapace.Gen(network_meter_showCmd).PositionalCompletion(action.ActionNetworkMeters())
	carapace.Gen(network_onboard_subnetCmd).PositionalCompletion(action.ActionNetworks(), action.ActionSubnetPools())
	carapace.Gen(network_qos_policy_deleteCmd).PositionalAnyCompletion(action.ActionNetworkQosPolicies())
	carapace.Gen(network_qos_policy_setCmd).PositionalCompletion(action.ActionNetworkQosPolicies())
	carapace.Gen(network_qos_policy_showCmd).PositionalCompletion(action.ActionNetworkQosPolicies())
	carapace.Gen(network_qos_rule_deleteCmd).PositionalCompletion(action.ActionNetworkQosPolicies(), action.ActionNetworkQosRules())
	carapace.Gen(network_qos_rule_setCmd).PositionalCompletion(action.ActionNetworkQosPolicies(), action.ActionNetworkQosRules())
	carapace.Gen(network_qos_rule_showCmd).PositionalCompletion(action.ActionNetworkQosPolicies(), action.ActionNetworkQosRules())
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
	carapace.Gen(object_createCmd).PositionalCompletion(action.ActionContainers(), carapace.ActionFiles())
	carapace.Gen(object_deleteCmd).PositionalCompletion(action.ActionContainers(), action.ActionObjects())
	carapace.Gen(object_listCmd).PositionalCompletion(action.ActionContainers())
	carapace.Gen(object_saveCmd).PositionalCompletion(carapace.ActionFiles(), action.ActionContainers(), action.ActionObjectsFromArg(1))
	carapace.Gen(object_setCmd).PositionalCompletion(action.ActionContainers(), action.ActionObjects())
	carapace.Gen(object_showCmd).PositionalCompletion(action.ActionContainers(), action.ActionObjects())
	carapace.Gen(object_unsetCmd).PositionalCompletion(action.ActionContainers(), action.ActionObjects())
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
	carapace.Gen(role_addCmd).PositionalCompletion(action.ActionRoles())
	carapace.Gen(role_deleteCmd).PositionalAnyCompletion(action.ActionRoles())
	carapace.Gen(role_removeCmd).PositionalCompletion(action.ActionRoles())
	carapace.Gen(role_setCmd).PositionalCompletion(action.ActionRoles())
	carapace.Gen(role_showCmd).PositionalCompletion(action.ActionRoles())
	carapace.Gen(router_add_portCmd).PositionalCompletion(action.ActionRouters(), action.ActionPorts())
	carapace.Gen(router_add_subnetCmd).PositionalCompletion(action.ActionRouters(), action.ActionSubnets())
	carapace.Gen(router_deleteCmd).PositionalAnyCompletion(action.ActionRouters())
	carapace.Gen(router_remove_portCmd).PositionalCompletion(action.ActionRouters(), action.ActionPorts())
	carapace.Gen(router_remove_subnetCmd).PositionalCompletion(action.ActionRouters(), action.ActionSubnets())
	carapace.Gen(router_setCmd).PositionalCompletion(action.ActionRouters())
	carapace.Gen(router_showCmd).PositionalCompletion(action.ActionRouters())
	carapace.Gen(router_unsetCmd).PositionalCompletion(action.ActionRouters())
	carapace.Gen(security_group_deleteCmd).PositionalAnyCompletion(action.ActionSecurityGroups())
	carapace.Gen(security_group_rule_createCmd).PositionalCompletion(action.ActionSecurityGroups())
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
	carapace.Gen(server_migration_abortCmd).PositionalCompletion(action.ActionServers(), action.ActionServerMigrations())
	carapace.Gen(server_migration_force_completeCmd).PositionalCompletion(action.ActionServers(), action.ActionServerMigrations())
	carapace.Gen(server_migration_showCmd).PositionalCompletion(action.ActionServers(), action.ActionServerMigrations())
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
	carapace.Gen(server_share_showCmd).PositionalCompletion(action.ActionServers(), action.ActionShares())
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
	carapace.Gen(volume_backup_createCmd).PositionalCompletion(action.ActionVolumes())
	carapace.Gen(volume_backup_deleteCmd).PositionalAnyCompletion(action.ActionVolumeBackups())
	carapace.Gen(volume_backup_restoreCmd).PositionalCompletion(action.ActionVolumeBackups())
	carapace.Gen(volume_backup_setCmd).PositionalCompletion(action.ActionVolumeBackups())
	carapace.Gen(volume_backup_showCmd).PositionalCompletion(action.ActionVolumeBackups())
	carapace.Gen(volume_backup_unsetCmd).PositionalCompletion(action.ActionVolumeBackups())
	carapace.Gen(volume_deleteCmd).PositionalAnyCompletion(action.ActionVolumes())
	carapace.Gen(volume_group_add_volumeCmd).PositionalCompletion(action.ActionVolumeGroups(), action.ActionVolumes())
	carapace.Gen(volume_group_deleteCmd).PositionalAnyCompletion(action.ActionVolumeGroups())
	carapace.Gen(volume_group_failoverCmd).PositionalCompletion(action.ActionVolumeGroups())
	carapace.Gen(volume_group_remove_volumeCmd).PositionalCompletion(action.ActionVolumeGroups(), action.ActionVolumes())
	carapace.Gen(volume_group_setCmd).PositionalCompletion(action.ActionVolumeGroups())
	carapace.Gen(volume_group_showCmd).PositionalCompletion(action.ActionVolumeGroups())
	carapace.Gen(volume_group_snapshot_createCmd).PositionalCompletion(action.ActionVolumeGroups())
	carapace.Gen(volume_group_snapshot_deleteCmd).PositionalAnyCompletion(action.ActionVolumeGroupSnapshots())
	carapace.Gen(volume_group_snapshot_showCmd).PositionalCompletion(action.ActionVolumeGroupSnapshots())
	carapace.Gen(volume_group_type_deleteCmd).PositionalAnyCompletion(action.ActionVolumeGroupTypes())
	carapace.Gen(volume_group_type_setCmd).PositionalCompletion(action.ActionVolumeGroupTypes())
	carapace.Gen(volume_group_type_showCmd).PositionalCompletion(action.ActionVolumeGroupTypes())
	carapace.Gen(volume_host_setCmd).PositionalCompletion(action.ActionVolumeHosts())
	carapace.Gen(volume_qos_associateCmd).PositionalCompletion(action.ActionVolumeQosSpecs(), action.ActionVolumeTypes())
	carapace.Gen(volume_qos_deleteCmd).PositionalAnyCompletion(action.ActionVolumeQosSpecs())
	carapace.Gen(volume_qos_disassociateCmd).PositionalCompletion(action.ActionVolumeQosSpecs(), action.ActionVolumeTypes())
	carapace.Gen(volume_qos_setCmd).PositionalCompletion(action.ActionVolumeQosSpecs())
	carapace.Gen(volume_qos_showCmd).PositionalCompletion(action.ActionVolumeQosSpecs())
	carapace.Gen(volume_qos_unsetCmd).PositionalCompletion(action.ActionVolumeQosSpecs())
	carapace.Gen(volume_service_setCmd).PositionalCompletion(action.ActionVolumeHosts())
	carapace.Gen(volume_setCmd).PositionalCompletion(action.ActionVolumes())
	carapace.Gen(volume_showCmd).PositionalCompletion(action.ActionVolumes())
	carapace.Gen(volume_snapshot_createCmd).PositionalCompletion(carapace.ActionValues(), action.ActionVolumes())
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

	carapace.Gen(accelerator_arq_bindCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "arq", "list"),
	})
	carapace.Gen(accelerator_arq_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "arq", "list"),
	})
	carapace.Gen(accelerator_arq_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("accelerator", "arq", "list"),
		"sort-column": action.ActionColumns("accelerator", "arq", "list"),
	})
	carapace.Gen(accelerator_arq_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "arq", "list"),
	})
	carapace.Gen(accelerator_arq_unbindCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "arq", "list"),
	})
	carapace.Gen(accelerator_deployable_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("accelerator", "deployable", "list"),
		"sort-column": action.ActionColumns("accelerator", "deployable", "list"),
	})
	carapace.Gen(accelerator_deployable_programCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "deployable", "list"),
	})
	carapace.Gen(accelerator_deployable_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "deployable", "list"),
	})
	carapace.Gen(accelerator_device_attribute_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "device", "attribute", "list"),
	})
	carapace.Gen(accelerator_device_attribute_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("accelerator", "device", "attribute", "list"),
		"sort-column": action.ActionColumns("accelerator", "device", "attribute", "list"),
	})
	carapace.Gen(accelerator_device_attribute_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "device", "attribute", "list"),
	})
	carapace.Gen(accelerator_device_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("accelerator", "device", "list"),
		"sort-column": action.ActionColumns("accelerator", "device", "list"),
	})
	carapace.Gen(accelerator_device_profile_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "device", "profile", "list"),
	})
	carapace.Gen(accelerator_device_profile_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("accelerator", "device", "profile", "list"),
		"sort-column": action.ActionColumns("accelerator", "device", "profile", "list"),
	})
	carapace.Gen(accelerator_device_profile_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "device", "profile", "list"),
	})
	carapace.Gen(accelerator_device_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("accelerator", "device", "list"),
	})
	carapace.Gen(access_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("access", "rule", "list"),
		"sort-column": action.ActionColumns("access", "rule", "list"),
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(access_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("access", "rule", "list"),
	})
	carapace.Gen(address_group_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("address", "group", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(address_group_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("address", "group", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("address", "group", "list"),
	})
	carapace.Gen(address_group_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("address", "group", "list"),
	})
	carapace.Gen(address_scope_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("address", "scope", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(address_scope_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("address", "scope", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("address", "scope", "list"),
	})
	carapace.Gen(address_scope_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("address", "scope", "list"),
	})
	carapace.Gen(aggregate_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("aggregate", "list"),
	})
	carapace.Gen(aggregate_listCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("aggregate", "list"),
		"sort-column":       action.ActionColumns("aggregate", "list"),
	})
	carapace.Gen(aggregate_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("aggregate", "list"),
	})
	carapace.Gen(application_credential_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("application", "credential", "list"),
		"role":   action.ActionRoles(),
	})
	carapace.Gen(application_credential_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("application", "credential", "list"),
		"sort-column": action.ActionColumns("application", "credential", "list"),
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(application_credential_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("application", "credential", "list"),
	})
	carapace.Gen(availability_zone_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("availability", "zone", "list"),
		"network":     action.ActionNetworks(),
		"sort-column": action.ActionColumns("availability", "zone", "list"),
		"volume":      action.ActionVolumes(),
	})
	carapace.Gen(bgp_dragent_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("bgp", "dragent", "list"),
		"sort-column": action.ActionColumns("bgp", "dragent", "list"),
	})
	carapace.Gen(bgp_peer_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("bgp", "peer", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgp_peer_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("bgp", "peer", "list"),
		"sort-column": action.ActionColumns("bgp", "peer", "list"),
	})
	carapace.Gen(bgp_peer_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("bgp", "peer", "list"),
	})
	carapace.Gen(bgp_speaker_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("bgp", "speaker", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgp_speaker_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("bgp", "speaker", "list"),
		"sort-column": action.ActionColumns("bgp", "speaker", "list"),
	})
	carapace.Gen(bgp_speaker_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("bgp", "speaker", "list"),
	})
	carapace.Gen(bgpvpn_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("bgpvpn", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgpvpn_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("bgpvpn", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("bgpvpn", "list"),
	})
	carapace.Gen(bgpvpn_network_association_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("bgpvpn", "network", "association", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgpvpn_network_association_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("bgpvpn", "network", "association", "list"),
		"sort-column": action.ActionColumns("bgpvpn", "network", "association", "list"),
	})
	carapace.Gen(bgpvpn_network_association_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("bgpvpn", "network", "association", "list"),
	})
	carapace.Gen(bgpvpn_port_association_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("bgpvpn", "port", "association", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgpvpn_port_association_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("bgpvpn", "port", "association", "list"),
		"sort-column": action.ActionColumns("bgpvpn", "port", "association", "list"),
	})
	carapace.Gen(bgpvpn_port_association_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("bgpvpn", "port", "association", "list"),
	})
	carapace.Gen(bgpvpn_router_association_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("bgpvpn", "router", "association", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(bgpvpn_router_association_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("bgpvpn", "router", "association", "list"),
		"sort-column": action.ActionColumns("bgpvpn", "router", "association", "list"),
	})
	carapace.Gen(bgpvpn_router_association_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("bgpvpn", "router", "association", "list"),
	})
	carapace.Gen(bgpvpn_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("bgpvpn", "list"),
	})
	carapace.Gen(block_storage_cluster_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("block", "storage", "cluster", "list"),
		"sort-column": action.ActionColumns("block", "storage", "cluster", "list"),
	})
	carapace.Gen(block_storage_cluster_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("block", "storage", "cluster", "list"),
	})
	carapace.Gen(block_storage_log_level_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("block", "storage", "log", "level", "list"),
		"sort-column": action.ActionColumns("block", "storage", "log", "level", "list"),
	})
	carapace.Gen(block_storage_resource_filter_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("block", "storage", "resource", "filter", "list"),
		"sort-column": action.ActionColumns("block", "storage", "resource", "filter", "list"),
	})
	carapace.Gen(block_storage_resource_filter_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("block", "storage", "resource", "filter", "list"),
	})
	carapace.Gen(block_storage_snapshot_manageable_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("block", "storage", "snapshot", "manageable", "list"),
		"sort-column": action.ActionColumns("block", "storage", "snapshot", "manageable", "list"),
	})
	carapace.Gen(block_storage_volume_manageable_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("block", "storage", "volume", "manageable", "list"),
		"sort-column": action.ActionColumns("block", "storage", "volume", "manageable", "list"),
	})
	carapace.Gen(cached_image_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("cached", "image", "list"),
		"sort-column": action.ActionColumns("cached", "image", "list"),
	})
	carapace.Gen(catalog_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("catalog", "list"),
		"sort-column": action.ActionColumns("catalog", "list"),
	})
	carapace.Gen(catalog_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("catalog", "list"),
	})
	carapace.Gen(command_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("command", "list"),
		"sort-column": action.ActionColumns("command", "list"),
	})
	carapace.Gen(compute_agent_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("compute", "agent", "list"),
	})
	carapace.Gen(compute_agent_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("compute", "agent", "list"),
		"sort-column": action.ActionColumns("compute", "agent", "list"),
	})
	carapace.Gen(compute_service_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("compute", "service", "list"),
		"sort-column": action.ActionColumns("compute", "service", "list"),
	})
	carapace.Gen(consistency_group_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("consistency", "group", "list"),
		"snapshot":          action.ActionConsistencyGroupSnapshots(),
		"source":            action.ActionConsistencyGroups(),
		"volume-type":       action.ActionVolumeTypes(),
	})
	carapace.Gen(consistency_group_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("consistency", "group", "list"),
		"sort-column": action.ActionColumns("consistency", "group", "list"),
	})
	carapace.Gen(consistency_group_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("consistency", "group", "list"),
	})
	carapace.Gen(consistency_group_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"column":            action.ActionColumns("consistency", "group", "snapshot", "list"),
		"consistency-group": action.ActionConsistencyGroups(),
	})
	carapace.Gen(consistency_group_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"column":            action.ActionColumns("consistency", "group", "snapshot", "list"),
		"consistency-group": action.ActionConsistencyGroups(),
		"sort-column":       action.ActionColumns("consistency", "group", "snapshot", "list"),
	})
	carapace.Gen(consistency_group_snapshot_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("consistency", "group", "snapshot", "list"),
	})
	carapace.Gen(consumer_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("consumer", "list"),
	})
	carapace.Gen(consumer_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("consumer", "list"),
		"sort-column": action.ActionColumns("consumer", "list"),
	})
	carapace.Gen(consumer_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("consumer", "list"),
	})
	carapace.Gen(container_createCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("container", "list"),
		"sort-column": action.ActionColumns("container", "list"),
	})
	carapace.Gen(container_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("container", "list"),
		"sort-column": action.ActionColumns("container", "list"),
	})
	carapace.Gen(container_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("container", "list"),
	})
	carapace.Gen(credential_createCmd).FlagCompletion(carapace.ActionMap{
		"column":  action.ActionColumns("credential", "list"),
		"project": action.ActionProjects(),
	})
	carapace.Gen(credential_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("credential", "list"),
		"sort-column": action.ActionColumns("credential", "list"),
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(credential_setCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
		"user":    action.ActionUsers(),
	})
	carapace.Gen(credential_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("credential", "list"),
	})
	carapace.Gen(default_security_group_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("default", "security", "group", "rule", "list"),
	})
	carapace.Gen(default_security_group_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("default", "security", "group", "rule", "list"),
		"sort-column": action.ActionColumns("default", "security", "group", "rule", "list"),
	})
	carapace.Gen(default_security_group_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("default", "security", "group", "rule", "list"),
	})
	carapace.Gen(domain_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("domain", "list"),
	})
	carapace.Gen(domain_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("domain", "list"),
		"sort-column": action.ActionColumns("domain", "list"),
	})
	carapace.Gen(domain_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("domain", "list"),
	})
	carapace.Gen(ec2_credentials_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("ec2", "credentials", "list"),
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
		"column":      action.ActionColumns("ec2", "credentials", "list"),
		"sort-column": action.ActionColumns("ec2", "credentials", "list"),
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(ec2_credentials_showCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("ec2", "credentials", "list"),
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(endpoint_add_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(endpoint_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("endpoint", "list"),
		"region": action.ActionRegions(),
	})
	carapace.Gen(endpoint_group_add_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(endpoint_group_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("endpoint", "group", "list"),
	})
	carapace.Gen(endpoint_group_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("endpoint", "group", "list"),
		"domain":      action.ActionDomains(),
		"project":     action.ActionProjects(),
		"sort-column": action.ActionColumns("endpoint", "group", "list"),
	})
	carapace.Gen(endpoint_group_remove_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(endpoint_group_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("endpoint", "group", "list"),
	})
	carapace.Gen(endpoint_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("endpoint", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"region":         action.ActionRegions(),
		"sort-column":    action.ActionColumns("endpoint", "list"),
	})
	carapace.Gen(endpoint_remove_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(endpoint_setCmd).FlagCompletion(carapace.ActionMap{
		"region": action.ActionRegions(),
	})
	carapace.Gen(endpoint_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("endpoint", "list"),
	})
	carapace.Gen(extension_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("extension", "list"),
		"network":     action.ActionNetworks(),
		"sort-column": action.ActionColumns("extension", "list"),
		"volume":      action.ActionVolumes(),
	})
	carapace.Gen(extension_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("extension", "list"),
	})
	carapace.Gen(federation_domain_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("federation", "domain", "list"),
		"sort-column": action.ActionColumns("federation", "domain", "list"),
	})
	carapace.Gen(federation_project_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("federation", "project", "list"),
		"sort-column": action.ActionColumns("federation", "project", "list"),
	})
	carapace.Gen(federation_protocol_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("federation", "protocol", "list"),
	})
	carapace.Gen(federation_protocol_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("federation", "protocol", "list"),
		"sort-column": action.ActionColumns("federation", "protocol", "list"),
	})
	carapace.Gen(federation_protocol_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("federation", "protocol", "list"),
	})
	carapace.Gen(firewall_group_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("firewall", "group", "list"),
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(firewall_group_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("firewall", "group", "list"),
		"sort-column": action.ActionColumns("firewall", "group", "list"),
	})
	carapace.Gen(firewall_group_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("firewall", "group", "policy", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(firewall_group_policy_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("firewall", "group", "policy", "list"),
		"sort-column": action.ActionColumns("firewall", "group", "policy", "list"),
	})
	carapace.Gen(firewall_group_policy_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("firewall", "group", "policy", "list"),
	})
	carapace.Gen(firewall_group_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("firewall", "group", "rule", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(firewall_group_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("firewall", "group", "rule", "list"),
		"sort-column": action.ActionColumns("firewall", "group", "rule", "list"),
	})
	carapace.Gen(firewall_group_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("firewall", "group", "rule", "list"),
	})
	carapace.Gen(firewall_group_setCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(firewall_group_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("firewall", "group", "list"),
	})
	carapace.Gen(firewall_group_unsetCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(flavor_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("flavor", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(flavor_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("flavor", "list"),
		"sort-column": action.ActionColumns("flavor", "list"),
	})
	carapace.Gen(flavor_setCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(flavor_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("flavor", "list"),
	})
	carapace.Gen(flavor_unsetCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(floating_ip_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("floating", "ip", "list"),
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"qos-policy":     action.ActionNetworkQosPolicies(),
		"router":         action.ActionRouters(),
		"subnet":         action.ActionSubnets(),
	})
	carapace.Gen(floating_ip_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("floating", "ip", "list"),
		"network":        action.ActionNetworks(),
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"router":         action.ActionRouters(),
		"sort-column":    action.ActionColumns("floating", "ip", "list"),
	})
	carapace.Gen(floating_ip_port_forwarding_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("floating", "ip", "port", "forwarding", "list"),
		"port":   action.ActionPorts(),
	})
	carapace.Gen(floating_ip_port_forwarding_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("floating", "ip", "port", "forwarding", "list"),
		"port":        action.ActionPorts(),
		"sort-column": action.ActionColumns("floating", "ip", "port", "forwarding", "list"),
	})
	carapace.Gen(floating_ip_port_forwarding_setCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(floating_ip_port_forwarding_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("floating", "ip", "port", "forwarding", "list"),
	})
	carapace.Gen(floating_ip_setCmd).FlagCompletion(carapace.ActionMap{
		"port":       action.ActionPorts(),
		"qos-policy": action.ActionNetworkQosPolicies(),
		"router":     action.ActionRouters(),
	})
	carapace.Gen(floating_ip_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("floating", "ip", "list"),
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
		"column": action.ActionColumns("group", "list"),
		"domain": action.ActionDomains(),
	})
	carapace.Gen(group_deleteCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(group_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("group", "list"),
		"domain":      action.ActionDomains(),
		"sort-column": action.ActionColumns("group", "list"),
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
		"column": action.ActionColumns("group", "list"),
		"domain": action.ActionDomains(),
	})
	carapace.Gen(host_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("host", "list"),
		"sort-column": action.ActionColumns("host", "list"),
	})
	carapace.Gen(host_showCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("host", "list"),
		"sort-column": action.ActionColumns("host", "list"),
	})
	carapace.Gen(hypervisor_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("hypervisor", "list"),
		"sort-column": action.ActionColumns("hypervisor", "list"),
	})
	carapace.Gen(hypervisor_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("hypervisor", "list"),
	})
	carapace.Gen(identity_provider_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("identity", "provider", "list"),
		"domain": action.ActionDomains(),
	})
	carapace.Gen(identity_provider_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("identity", "provider", "list"),
		"sort-column": action.ActionColumns("identity", "provider", "list"),
	})
	carapace.Gen(identity_provider_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("identity", "provider", "list"),
	})
	carapace.Gen(image_add_projectCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(image_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("image", "list"),
		"file":           carapace.ActionFiles(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sign-key-path":  carapace.ActionFiles(),
		"volume":         action.ActionVolumes(),
	})
	carapace.Gen(image_importCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("image", "list"),
	})
	carapace.Gen(image_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("image", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("image", "list"),
	})
	carapace.Gen(image_member_getCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("image", "member", "list"),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(image_member_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("image", "member", "list"),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("image", "member", "list"),
	})
	carapace.Gen(image_metadef_namespace_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("image", "metadef", "namespace", "list"),
	})
	carapace.Gen(image_metadef_namespace_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("image", "metadef", "namespace", "list"),
		"sort-column": action.ActionColumns("image", "metadef", "namespace", "list"),
	})
	carapace.Gen(image_metadef_namespace_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("image", "metadef", "namespace", "list"),
	})
	carapace.Gen(image_metadef_object_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("image", "metadef", "object", "list"),
	})
	carapace.Gen(image_metadef_object_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("image", "metadef", "object", "list"),
		"sort-column": action.ActionColumns("image", "metadef", "object", "list"),
	})
	carapace.Gen(image_metadef_object_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("image", "metadef", "object", "list"),
	})
	carapace.Gen(image_metadef_property_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("image", "metadef", "property", "list"),
	})
	carapace.Gen(image_metadef_property_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("image", "metadef", "property", "list"),
		"sort-column": action.ActionColumns("image", "metadef", "property", "list"),
	})
	carapace.Gen(image_metadef_property_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("image", "metadef", "property", "list"),
	})
	carapace.Gen(image_metadef_resource_type_association_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("image", "metadef", "resource", "type", "association", "list"),
	})
	carapace.Gen(image_metadef_resource_type_association_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("image", "metadef", "resource", "type", "association", "list"),
		"sort-column": action.ActionColumns("image", "metadef", "resource", "type", "association", "list"),
	})
	carapace.Gen(image_metadef_resource_type_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("image", "metadef", "resource", "type", "list"),
		"sort-column": action.ActionColumns("image", "metadef", "resource", "type", "list"),
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
	carapace.Gen(image_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("image", "list"),
	})
	carapace.Gen(image_stageCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
	carapace.Gen(image_stores_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("image", "stores", "list"),
		"sort-column": action.ActionColumns("image", "stores", "list"),
	})
	carapace.Gen(image_task_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("image", "task", "list"),
		"sort-column": action.ActionColumns("image", "task", "list"),
	})
	carapace.Gen(image_task_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("image", "task", "list"),
	})
	carapace.Gen(implied_role_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("implied", "role", "list"),
	})
	carapace.Gen(implied_role_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("implied", "role", "list"),
		"sort-column": action.ActionColumns("implied", "role", "list"),
	})
	carapace.Gen(ip_availability_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("ip", "availability", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("ip", "availability", "list"),
	})
	carapace.Gen(ip_availability_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("ip", "availability", "list"),
	})
	carapace.Gen(keypair_createCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("keypair", "list"),
		"public-key":  carapace.ActionFiles(),
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(keypair_deleteCmd).FlagCompletion(carapace.ActionMap{
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(keypair_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("keypair", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("keypair", "list"),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(keypair_showCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("keypair", "list"),
		"public-key":  carapace.ActionFiles(),
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})
	carapace.Gen(limit_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("limit", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"region":         action.ActionRegions(),
	})
	carapace.Gen(limit_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("limit", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"region":         action.ActionRegions(),
		"sort-column":    action.ActionColumns("limit", "list"),
	})
	carapace.Gen(limit_setCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("limit", "list"),
	})
	carapace.Gen(limit_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("limit", "list"),
	})
	carapace.Gen(limits_showCmd).FlagCompletion(carapace.ActionMap{
		"domain":         action.ActionDomains(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(local_ip_association_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("local", "ip", "association", "list"),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(local_ip_association_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("local", "ip", "association", "list"),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("local", "ip", "association", "list"),
	})
	carapace.Gen(local_ip_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("local", "ip", "list"),
		"network":        action.ActionNetworks(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(local_ip_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("local", "ip", "list"),
		"network":        action.ActionNetworks(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("local", "ip", "list"),
	})
	carapace.Gen(local_ip_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("local", "ip", "list"),
	})
	carapace.Gen(mapping_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("mapping", "list"),
	})
	carapace.Gen(mapping_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("mapping", "list"),
		"sort-column": action.ActionColumns("mapping", "list"),
	})
	carapace.Gen(mapping_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("mapping", "list"),
	})
	carapace.Gen(module_listCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("module", "list"),
	})
	carapace.Gen(network_agent_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "agent", "list"),
		"network":     action.ActionNetworks(),
		"router":      action.ActionRouters(),
		"sort-column": action.ActionColumns("network", "agent", "list"),
	})
	carapace.Gen(network_agent_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "agent", "list"),
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
		"column":                 action.ActionColumns("network", "list"),
		"project":                action.ActionProjects(),
		"project-domain":         action.ActionDomains(),
		"qos-policy":             action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(network_flavor_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("network", "flavor", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_flavor_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "flavor", "list"),
		"sort-column": action.ActionColumns("network", "flavor", "list"),
	})
	carapace.Gen(network_flavor_profile_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "flavor", "profile", "list"),
	})
	carapace.Gen(network_flavor_profile_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "flavor", "profile", "list"),
		"sort-column": action.ActionColumns("network", "flavor", "profile", "list"),
	})
	carapace.Gen(network_flavor_profile_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "flavor", "profile", "list"),
	})
	carapace.Gen(network_flavor_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "flavor", "list"),
	})
	carapace.Gen(network_l3_conntrack_helper_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "l3", "conntrack", "helper", "list"),
		"port":   action.ActionPorts(),
	})
	carapace.Gen(network_l3_conntrack_helper_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "l3", "conntrack", "helper", "list"),
		"port":        action.ActionPorts(),
		"sort-column": action.ActionColumns("network", "l3", "conntrack", "helper", "list"),
	})
	carapace.Gen(network_l3_conntrack_helper_setCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(network_l3_conntrack_helper_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "l3", "conntrack", "helper", "list"),
	})
	carapace.Gen(network_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("network", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("network", "list"),
	})
	carapace.Gen(network_meter_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("network", "meter", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_meter_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "meter", "list"),
		"sort-column": action.ActionColumns("network", "meter", "list"),
	})
	carapace.Gen(network_meter_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("network", "meter", "rule", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_meter_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "meter", "rule", "list"),
		"sort-column": action.ActionColumns("network", "meter", "rule", "list"),
	})
	carapace.Gen(network_meter_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "meter", "rule", "list"),
	})
	carapace.Gen(network_meter_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "meter", "list"),
	})
	carapace.Gen(network_qos_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("network", "qos", "policy", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_qos_policy_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("network", "qos", "policy", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("network", "qos", "policy", "list"),
	})
	carapace.Gen(network_qos_policy_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "qos", "policy", "list"),
	})
	carapace.Gen(network_qos_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "qos", "rule", "list"),
	})
	carapace.Gen(network_qos_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "qos", "rule", "list"),
		"sort-column": action.ActionColumns("network", "qos", "rule", "list"),
	})
	carapace.Gen(network_qos_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "qos", "rule", "list"),
	})
	carapace.Gen(network_qos_rule_type_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "qos", "rule", "type", "list"),
		"sort-column": action.ActionColumns("network", "qos", "rule", "type", "list"),
	})
	carapace.Gen(network_qos_rule_type_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "qos", "rule", "type", "list"),
	})
	carapace.Gen(network_rbac_createCmd).FlagCompletion(carapace.ActionMap{
		"column":                action.ActionColumns("network", "rbac", "list"),
		"project":               action.ActionProjects(),
		"project-domain":        action.ActionDomains(),
		"target-project":        action.ActionProjects(),
		"target-project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_rbac_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("network", "rbac", "list"),
		"sort-column":    action.ActionColumns("network", "rbac", "list"),
		"target-project": action.ActionProjects(),
	})
	carapace.Gen(network_rbac_setCmd).FlagCompletion(carapace.ActionMap{
		"target-project":        action.ActionProjects(),
		"target-project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_rbac_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "rbac", "list"),
	})
	carapace.Gen(network_segment_createCmd).FlagCompletion(carapace.ActionMap{
		"column":  action.ActionColumns("network", "segment", "list"),
		"network": action.ActionNetworks(),
	})
	carapace.Gen(network_segment_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "segment", "list"),
		"network":     action.ActionNetworks(),
		"sort-column": action.ActionColumns("network", "segment", "list"),
	})
	carapace.Gen(network_segment_range_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("network", "segment", "range", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_segment_range_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "segment", "range", "list"),
		"sort-column": action.ActionColumns("network", "segment", "range", "list"),
	})
	carapace.Gen(network_segment_range_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "segment", "range", "list"),
	})
	carapace.Gen(network_segment_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "segment", "list"),
	})
	carapace.Gen(network_service_provider_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "service", "provider", "list"),
		"sort-column": action.ActionColumns("network", "service", "provider", "list"),
	})
	carapace.Gen(network_setCmd).FlagCompletion(carapace.ActionMap{
		"qos-policy": action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(network_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "list"),
	})
	carapace.Gen(network_subport_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "subport", "list"),
		"sort-column": action.ActionColumns("network", "subport", "list"),
	})
	carapace.Gen(network_trunk_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("network", "trunk", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(network_trunk_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "trunk", "list"),
		"sort-column": action.ActionColumns("network", "trunk", "list"),
	})
	carapace.Gen(network_trunk_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("network", "trunk", "list"),
	})
	carapace.Gen(network_trunk_subport_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("network", "trunk", "subport", "list"),
		"sort-column": action.ActionColumns("network", "trunk", "subport", "list"),
	})
	carapace.Gen(object_createCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("object", "list"),
		"sort-column": action.ActionColumns("object", "list"),
	})
	carapace.Gen(object_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("object", "list"),
		"sort-column": action.ActionColumns("object", "list"),
	})
	carapace.Gen(object_saveCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
	carapace.Gen(object_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("object", "list"),
	})
	carapace.Gen(policy_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("policy", "list"),
	})
	carapace.Gen(policy_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("policy", "list"),
		"sort-column": action.ActionColumns("policy", "list"),
	})
	carapace.Gen(policy_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("policy", "list"),
	})
	carapace.Gen(port_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("port", "list"),
		"network":        action.ActionNetworks(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"qos-policy":     action.ActionNetworkQosPolicies(),
		"security-group": action.ActionSecurityGroups(),
	})
	carapace.Gen(port_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("port", "list"),
		"network":        action.ActionNetworks(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"router":         action.ActionRouters(),
		"security-group": action.ActionSecurityGroups(),
		"server":         action.ActionServers(),
		"sort-column":    action.ActionColumns("port", "list"),
	})
	carapace.Gen(port_setCmd).FlagCompletion(carapace.ActionMap{
		"qos-policy":     action.ActionNetworkQosPolicies(),
		"security-group": action.ActionSecurityGroups(),
	})
	carapace.Gen(port_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("port", "list"),
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
		"column": action.ActionColumns("project", "list"),
		"domain": action.ActionDomains(),
	})
	carapace.Gen(project_deleteCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(project_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("project", "list"),
		"domain":      action.ActionDomains(),
		"sort-column": action.ActionColumns("project", "list"),
		"user":        action.ActionUsers(),
	})
	carapace.Gen(project_setCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(project_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("project", "list"),
		"domain": action.ActionDomains(),
	})
	carapace.Gen(quota_deleteCmd).FlagCompletion(carapace.ActionMap{
		"network": action.ActionNetworks(),
		"volume":  action.ActionVolumes(),
	})
	carapace.Gen(quota_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("quota", "list"),
		"network":     action.ActionNetworks(),
		"sort-column": action.ActionColumns("quota", "list"),
		"volume":      action.ActionVolumes(),
	})
	carapace.Gen(quota_setCmd).FlagCompletion(carapace.ActionMap{
		"volume-type": action.ActionVolumeTypes(),
	})
	carapace.Gen(quota_showCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("quota", "list"),
		"network":     action.ActionNetworks(),
		"sort-column": action.ActionColumns("quota", "list"),
		"volume":      action.ActionVolumes(),
	})
	carapace.Gen(region_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("region", "list"),
	})
	carapace.Gen(region_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("region", "list"),
		"sort-column": action.ActionColumns("region", "list"),
	})
	carapace.Gen(region_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("region", "list"),
	})
	carapace.Gen(registered_limit_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("registered", "limit", "list"),
		"region": action.ActionRegions(),
	})
	carapace.Gen(registered_limit_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("registered", "limit", "list"),
		"region":      action.ActionRegions(),
		"sort-column": action.ActionColumns("registered", "limit", "list"),
	})
	carapace.Gen(registered_limit_setCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("registered", "limit", "list"),
		"region": action.ActionRegions(),
	})
	carapace.Gen(registered_limit_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("registered", "limit", "list"),
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
		"column":         action.ActionColumns("role", "assignment", "list"),
		"domain":         action.ActionDomains(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"role":           action.ActionRoles(),
		"sort-column":    action.ActionColumns("role", "assignment", "list"),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(role_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("role", "list"),
		"domain": action.ActionDomains(),
	})
	carapace.Gen(role_deleteCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(role_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("role", "list"),
		"domain":      action.ActionDomains(),
		"sort-column": action.ActionColumns("role", "list"),
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
		"column": action.ActionColumns("role", "list"),
		"domain": action.ActionDomains(),
	})
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log-file":               carapace.ActionFiles(),
		"os-auth-url":            action.ActionCloudAuthURLs(),
		"os-cacert":              carapace.ActionFiles(),
		"os-cert":                carapace.ActionFiles(),
		"os-cloud":               action.ActionClouds(),
		"os-default-domain":      action.ActionCloudDomainIDs(),
		"os-default-domain-id":   action.ActionCloudDomainIDs(),
		"os-default-domain-name": action.ActionCloudDomainNames(),
		"os-domain-id":           action.ActionCloudDomainIDs(),
		"os-domain-name":         action.ActionCloudDomainNames(),
		"os-key":                 carapace.ActionFiles(),
		"os-project-domain-id":   action.ActionCloudDomainIDs(),
		"os-project-domain-name": action.ActionCloudDomainNames(),
		"os-project-id":          action.ActionCloudProjectIDs(),
		"os-project-name":        action.ActionCloudProjectNames(),
		"os-region-name":         action.ActionCloudRegions(),
		"os-user-domain-id":      action.ActionCloudDomainIDs(),
		"os-user-domain-name":    action.ActionCloudDomainNames(),
		"os-user-id":             action.ActionCloudUserIDs(),
		"os-username":            action.ActionCloudUsernames(),
	})
	carapace.Gen(router_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone-hint": action.ActionAvailabilityZones(),
		"column":                 action.ActionColumns("router", "list"),
		"flavor":                 action.ActionFlavors(),
		"project":                action.ActionProjects(),
		"project-domain":         action.ActionDomains(),
		"qos-policy":             action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(router_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("router", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("router", "list"),
	})
	carapace.Gen(router_ndp_proxy_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("router", "ndp", "proxy", "list"),
		"port":   action.ActionPorts(),
	})
	carapace.Gen(router_ndp_proxy_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("router", "ndp", "proxy", "list"),
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"router":         action.ActionRouters(),
		"sort-column":    action.ActionColumns("router", "ndp", "proxy", "list"),
	})
	carapace.Gen(router_ndp_proxy_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("router", "ndp", "proxy", "list"),
	})
	carapace.Gen(router_setCmd).FlagCompletion(carapace.ActionMap{
		"qos-policy": action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(router_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("router", "list"),
	})
	carapace.Gen(router_unsetCmd).FlagCompletion(carapace.ActionMap{
		"qos-policy": action.ActionNetworkQosPolicies(),
	})
	carapace.Gen(security_group_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("security", "group", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(security_group_default_statefulness_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("security", "group", "default", "statefulness", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(security_group_default_statefulness_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("security", "group", "default", "statefulness", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("security", "group", "default", "statefulness", "list"),
	})
	carapace.Gen(security_group_default_statefulness_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("security", "group", "default", "statefulness", "list"),
	})
	carapace.Gen(security_group_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("security", "group", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("security", "group", "list"),
	})
	carapace.Gen(security_group_rule_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("security", "group", "rule", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(security_group_rule_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("security", "group", "rule", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("security", "group", "rule", "list"),
	})
	carapace.Gen(security_group_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("security", "group", "rule", "list"),
	})
	carapace.Gen(security_group_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("security", "group", "list"),
	})
	carapace.Gen(server_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"block-device":      carapace.ActionFiles(),
		"column":            action.ActionColumns("server", "list"),
		"file":              carapace.ActionFiles(),
		"flavor":            action.ActionFlavors(),
		"image":             action.ActionImages(),
		"key-name":          action.ActionKeypairs(),
		"network":           action.ActionNetworks(),
		"port":              action.ActionPorts(),
		"security-group":    action.ActionSecurityGroups(),
		"server-group":      action.ActionServerGroups(),
		"snapshot":          action.ActionImages(),
		"user-data":         carapace.ActionFiles(),
		"volume":            action.ActionVolumes(),
	})
	carapace.Gen(server_evacuateCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("server", "list"),
	})
	carapace.Gen(server_event_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("server", "event", "list"),
		"sort-column": action.ActionColumns("server", "event", "list"),
	})
	carapace.Gen(server_event_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("server", "event", "list"),
	})
	carapace.Gen(server_group_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("server", "group", "list"),
	})
	carapace.Gen(server_group_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("server", "group", "list"),
		"sort-column": action.ActionColumns("server", "group", "list"),
	})
	carapace.Gen(server_group_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("server", "group", "list"),
	})
	carapace.Gen(server_listCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("server", "list"),
		"flavor":            action.ActionFlavors(),
		"image":             action.ActionImages(),
		"key-name":          action.ActionKeypairs(),
		"project":           action.ActionProjects(),
		"project-domain":    action.ActionDomains(),
		"sort-column":       action.ActionColumns("server", "list"),
		"user":              action.ActionUsers(),
		"user-domain":       action.ActionDomains(),
	})
	carapace.Gen(server_migration_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("server", "migration", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"server":         action.ActionServers(),
		"sort-column":    action.ActionColumns("server", "migration", "list"),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(server_migration_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("server", "migration", "list"),
	})
	carapace.Gen(server_rebuildCmd).FlagCompletion(carapace.ActionMap{
		"column":    action.ActionColumns("server", "list"),
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
	carapace.Gen(server_share_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("server", "share", "list"),
		"sort-column": action.ActionColumns("server", "share", "list"),
	})
	carapace.Gen(server_share_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("server", "share", "list"),
	})
	carapace.Gen(server_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("server", "list"),
	})
	carapace.Gen(server_sshCmd).FlagCompletion(carapace.ActionMap{
		"port": action.ActionPorts(),
	})
	carapace.Gen(server_unshelveCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
	})
	carapace.Gen(server_volume_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("server", "volume", "list"),
		"sort-column": action.ActionColumns("server", "volume", "list"),
	})
	carapace.Gen(service_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("service", "list"),
	})
	carapace.Gen(service_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("service", "list"),
		"sort-column": action.ActionColumns("service", "list"),
	})
	carapace.Gen(service_provider_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("service", "provider", "list"),
	})
	carapace.Gen(service_provider_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("service", "provider", "list"),
		"sort-column": action.ActionColumns("service", "provider", "list"),
	})
	carapace.Gen(service_provider_setCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("service", "provider", "list"),
	})
	carapace.Gen(service_provider_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("service", "provider", "list"),
	})
	carapace.Gen(service_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("service", "list"),
	})
	carapace.Gen(share_access_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "access", "list"),
	})
	carapace.Gen(share_access_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "access", "list"),
		"sort-column": action.ActionColumns("share", "access", "list"),
	})
	carapace.Gen(share_access_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "access", "list"),
	})
	carapace.Gen(share_adoptCmd).FlagCompletion(carapace.ActionMap{
		"column":     action.ActionColumns("share", "list"),
		"share-type": action.ActionShareTypes(),
	})
	carapace.Gen(share_availability_zone_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "availability", "zone", "list"),
		"sort-column": action.ActionColumns("share", "availability", "zone", "list"),
	})
	carapace.Gen(share_backup_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "backup", "list"),
	})
	carapace.Gen(share_backup_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "backup", "list"),
		"share":       action.ActionShares(),
		"sort-column": action.ActionColumns("share", "backup", "list"),
	})
	carapace.Gen(share_backup_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "backup", "list"),
	})
	carapace.Gen(share_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("share", "list"),
		"share-group":       action.ActionShareGroups(),
		"share-network":     action.ActionShareNetworks(),
		"share-type":        action.ActionShareTypes(),
	})
	carapace.Gen(share_deleteCmd).FlagCompletion(carapace.ActionMap{
		"share-group": action.ActionShareGroups(),
	})
	carapace.Gen(share_export_location_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "export", "location", "list"),
		"sort-column": action.ActionColumns("share", "export", "location", "list"),
	})
	carapace.Gen(share_export_location_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "export", "location", "list"),
	})
	carapace.Gen(share_group_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("share", "group", "list"),
		"share-network":     action.ActionShareNetworks(),
	})
	carapace.Gen(share_group_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("share", "group", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"share-network":  action.ActionShareNetworks(),
		"share-server":   action.ActionShareServers(),
		"snapshot":       action.ActionShareGroupSnapshots(),
		"sort-column":    action.ActionColumns("share", "group", "list"),
	})
	carapace.Gen(share_group_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "group", "list"),
	})
	carapace.Gen(share_group_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "group", "snapshot", "list"),
	})
	carapace.Gen(share_group_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "group", "snapshot", "list"),
		"share-group": action.ActionShareGroups(),
		"sort-column": action.ActionColumns("share", "group", "snapshot", "list"),
	})
	carapace.Gen(share_group_snapshot_members_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "group", "snapshot", "members", "list"),
		"sort-column": action.ActionColumns("share", "group", "snapshot", "members", "list"),
	})
	carapace.Gen(share_group_snapshot_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "group", "snapshot", "list"),
	})
	carapace.Gen(share_group_type_access_createCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(share_group_type_access_deleteCmd).FlagCompletion(carapace.ActionMap{
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(share_group_type_access_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "group", "type", "access", "list"),
		"sort-column": action.ActionColumns("share", "group", "type", "access", "list"),
	})
	carapace.Gen(share_group_type_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "group", "type", "list"),
	})
	carapace.Gen(share_group_type_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "group", "type", "list"),
		"sort-column": action.ActionColumns("share", "group", "type", "list"),
	})
	carapace.Gen(share_group_type_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "group", "type", "list"),
	})
	carapace.Gen(share_instance_export_location_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "instance", "export", "location", "list"),
		"sort-column": action.ActionColumns("share", "instance", "export", "location", "list"),
	})
	carapace.Gen(share_instance_export_location_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "instance", "export", "location", "list"),
	})
	carapace.Gen(share_instance_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "instance", "list"),
		"share":       action.ActionShares(),
		"sort-column": action.ActionColumns("share", "instance", "list"),
	})
	carapace.Gen(share_instance_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "instance", "list"),
	})
	carapace.Gen(share_listCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("share", "list"),
		"project":           action.ActionProjects(),
		"project-domain":    action.ActionDomains(),
		"share-group":       action.ActionShareGroups(),
		"share-network":     action.ActionShareNetworks(),
		"share-server":      action.ActionShareServers(),
		"share-type":        action.ActionShareTypes(),
		"snapshot":          action.ActionShareSnapshots(),
		"sort-column":       action.ActionColumns("share", "list"),
		"user":              action.ActionUsers(),
		"user-domain":       action.ActionDomains(),
	})
	carapace.Gen(share_lock_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "lock", "list"),
	})
	carapace.Gen(share_lock_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("share", "lock", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("share", "lock", "list"),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(share_lock_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "lock", "list"),
	})
	carapace.Gen(share_message_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "message", "list"),
		"sort-column": action.ActionColumns("share", "message", "list"),
	})
	carapace.Gen(share_message_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "message", "list"),
	})
	carapace.Gen(share_network_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("share", "network", "list"),
	})
	carapace.Gen(share_network_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("share", "network", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("share", "network", "list"),
	})
	carapace.Gen(share_network_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "network", "list"),
	})
	carapace.Gen(share_network_subnet_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
	})
	carapace.Gen(share_pool_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "pool", "list"),
		"share-type":  action.ActionShareTypes(),
		"sort-column": action.ActionColumns("share", "pool", "list"),
	})
	carapace.Gen(share_qos_type_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "qos", "type", "list"),
	})
	carapace.Gen(share_qos_type_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "qos", "type", "list"),
		"sort-column": action.ActionColumns("share", "qos", "type", "list"),
	})
	carapace.Gen(share_qos_type_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "qos", "type", "list"),
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
		"column":            action.ActionColumns("share", "replica", "list"),
		"share-network":     action.ActionShareNetworks(),
	})
	carapace.Gen(share_replica_export_location_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "replica", "export", "location", "list"),
		"sort-column": action.ActionColumns("share", "replica", "export", "location", "list"),
	})
	carapace.Gen(share_replica_export_location_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "replica", "export", "location", "list"),
	})
	carapace.Gen(share_replica_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "replica", "list"),
		"share":       action.ActionShares(),
		"sort-column": action.ActionColumns("share", "replica", "list"),
	})
	carapace.Gen(share_replica_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "replica", "list"),
	})
	carapace.Gen(share_security_service_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "security", "service", "list"),
		"domain": action.ActionDomains(),
		"server": action.ActionServers(),
		"user":   action.ActionUsers(),
	})
	carapace.Gen(share_security_service_listCmd).FlagCompletion(carapace.ActionMap{
		"column":        action.ActionColumns("share", "security", "service", "list"),
		"domain":        action.ActionDomains(),
		"server":        action.ActionServers(),
		"share-network": action.ActionShareNetworks(),
		"sort-column":   action.ActionColumns("share", "security", "service", "list"),
		"user":          action.ActionUsers(),
	})
	carapace.Gen(share_security_service_setCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
		"server": action.ActionServers(),
		"user":   action.ActionUsers(),
	})
	carapace.Gen(share_security_service_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "security", "service", "list"),
	})
	carapace.Gen(share_security_service_unsetCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
		"server": action.ActionServers(),
		"user":   action.ActionUsers(),
	})
	carapace.Gen(share_server_adoptCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "server", "list"),
	})
	carapace.Gen(share_server_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("share", "server", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"share-network":  action.ActionShareNetworks(),
		"sort-column":    action.ActionColumns("share", "server", "list"),
	})
	carapace.Gen(share_server_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "server", "list"),
	})
	carapace.Gen(share_service_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "service", "list"),
		"sort-column": action.ActionColumns("share", "service", "list"),
	})
	carapace.Gen(share_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "list"),
	})
	carapace.Gen(share_snapshot_access_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "snapshot", "access", "list"),
	})
	carapace.Gen(share_snapshot_access_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "snapshot", "access", "list"),
		"sort-column": action.ActionColumns("share", "snapshot", "access", "list"),
	})
	carapace.Gen(share_snapshot_adoptCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "snapshot", "list"),
	})
	carapace.Gen(share_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "snapshot", "list"),
	})
	carapace.Gen(share_snapshot_export_location_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "snapshot", "export", "location", "list"),
		"sort-column": action.ActionColumns("share", "snapshot", "export", "location", "list"),
	})
	carapace.Gen(share_snapshot_export_location_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "snapshot", "export", "location", "list"),
	})
	carapace.Gen(share_snapshot_instance_export_location_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "snapshot", "instance", "export", "location", "list"),
		"sort-column": action.ActionColumns("share", "snapshot", "instance", "export", "location", "list"),
	})
	carapace.Gen(share_snapshot_instance_export_location_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "snapshot", "instance", "export", "location", "list"),
	})
	carapace.Gen(share_snapshot_instance_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "snapshot", "instance", "list"),
		"snapshot":    action.ActionShareSnapshotIDs(),
		"sort-column": action.ActionColumns("share", "snapshot", "instance", "list"),
	})
	carapace.Gen(share_snapshot_instance_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "snapshot", "instance", "list"),
	})
	carapace.Gen(share_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "snapshot", "list"),
		"share":       action.ActionShares(),
		"sort-column": action.ActionColumns("share", "snapshot", "list"),
	})
	carapace.Gen(share_snapshot_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "snapshot", "list"),
	})
	carapace.Gen(share_transfer_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "transfer", "list"),
	})
	carapace.Gen(share_transfer_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "transfer", "list"),
		"sort-column": action.ActionColumns("share", "transfer", "list"),
	})
	carapace.Gen(share_transfer_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "transfer", "list"),
	})
	carapace.Gen(share_type_access_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "type", "access", "list"),
		"sort-column": action.ActionColumns("share", "type", "access", "list"),
	})
	carapace.Gen(share_type_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "type", "list"),
	})
	carapace.Gen(share_type_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("share", "type", "list"),
		"sort-column": action.ActionColumns("share", "type", "list"),
	})
	carapace.Gen(share_type_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("share", "type", "list"),
	})
	carapace.Gen(subnet_createCmd).FlagCompletion(carapace.ActionMap{
		"column":          action.ActionColumns("subnet", "list"),
		"network":         action.ActionNetworks(),
		"network-segment": action.ActionNetworkSegments(),
		"project":         action.ActionProjects(),
		"project-domain":  action.ActionDomains(),
		"subnet-pool":     action.ActionSubnetPools(),
	})
	carapace.Gen(subnet_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("subnet", "list"),
		"network":        action.ActionNetworks(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("subnet", "list"),
		"subnet-pool":    action.ActionSubnetPools(),
	})
	carapace.Gen(subnet_pool_createCmd).FlagCompletion(carapace.ActionMap{
		"address-scope":  action.ActionAddressScopes(),
		"column":         action.ActionColumns("subnet", "pool", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(subnet_pool_listCmd).FlagCompletion(carapace.ActionMap{
		"address-scope":  action.ActionAddressScopes(),
		"column":         action.ActionColumns("subnet", "pool", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("subnet", "pool", "list"),
	})
	carapace.Gen(subnet_pool_setCmd).FlagCompletion(carapace.ActionMap{
		"address-scope": action.ActionAddressScopes(),
	})
	carapace.Gen(subnet_pool_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("subnet", "pool", "list"),
	})
	carapace.Gen(subnet_setCmd).FlagCompletion(carapace.ActionMap{
		"network-segment": action.ActionNetworkSegments(),
	})
	carapace.Gen(subnet_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("subnet", "list"),
	})
	carapace.Gen(tap_flow_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("tap", "flow", "list"),
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(tap_flow_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("tap", "flow", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("tap", "flow", "list"),
	})
	carapace.Gen(tap_flow_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("tap", "flow", "list"),
	})
	carapace.Gen(tap_flow_updateCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("tap", "flow", "list"),
	})
	carapace.Gen(tap_mirror_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("tap", "mirror", "list"),
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(tap_mirror_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("tap", "mirror", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("tap", "mirror", "list"),
	})
	carapace.Gen(tap_mirror_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("tap", "mirror", "list"),
	})
	carapace.Gen(tap_mirror_updateCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("tap", "mirror", "list"),
	})
	carapace.Gen(tap_service_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("tap", "service", "list"),
		"port":           action.ActionPorts(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(tap_service_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("tap", "service", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("tap", "service", "list"),
	})
	carapace.Gen(tap_service_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("tap", "service", "list"),
	})
	carapace.Gen(tap_service_updateCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("tap", "service", "list"),
	})
	carapace.Gen(trust_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("trust", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"role":           action.ActionRoles(),
	})
	carapace.Gen(trust_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("trust", "list"),
		"sort-column": action.ActionColumns("trust", "list"),
	})
	carapace.Gen(trust_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("trust", "list"),
	})
	carapace.Gen(usage_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("usage", "list"),
		"sort-column": action.ActionColumns("usage", "list"),
	})
	carapace.Gen(usage_showCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("usage", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(user_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("user", "list"),
		"domain":         action.ActionDomains(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(user_deleteCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(),
	})
	carapace.Gen(user_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("user", "list"),
		"domain":      action.ActionDomains(),
		"project":     action.ActionProjects(),
		"sort-column": action.ActionColumns("user", "list"),
	})
	carapace.Gen(user_setCmd).FlagCompletion(carapace.ActionMap{
		"domain":         action.ActionDomains(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(user_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("user", "list"),
		"domain": action.ActionDomains(),
	})
	carapace.Gen(volume_attachment_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "attachment", "list"),
		"server": action.ActionServers(),
	})
	carapace.Gen(volume_attachment_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("volume", "attachment", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"server-id":      action.ActionServerIDs(),
		"sort-column":    action.ActionColumns("volume", "attachment", "list"),
		"volume-id":      action.ActionVolumeIDs(),
	})
	carapace.Gen(volume_attachment_setCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "attachment", "list"),
	})
	carapace.Gen(volume_attachment_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "attachment", "list"),
	})
	carapace.Gen(volume_backend_pool_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("volume", "backend", "pool", "list"),
		"sort-column": action.ActionColumns("volume", "backend", "pool", "list"),
	})
	carapace.Gen(volume_backup_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("volume", "backup", "list"),
		"snapshot":          action.ActionVolumeSnapshots(),
	})
	carapace.Gen(volume_backup_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("volume", "backup", "list"),
		"project":     action.ActionProjects(),
		"sort-column": action.ActionColumns("volume", "backup", "list"),
		"volume":      action.ActionVolumes(),
	})
	carapace.Gen(volume_backup_restoreCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "backup", "list"),
	})
	carapace.Gen(volume_backup_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "backup", "list"),
	})
	carapace.Gen(volume_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("volume", "list"),
		"consistency-group": action.ActionConsistencyGroups(),
		"image":             action.ActionImages(),
		"snapshot":          action.ActionVolumeSnapshots(),
	})
	carapace.Gen(volume_group_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("volume", "group", "list"),
		"volume-type":       action.ActionVolumeTypes(),
	})
	carapace.Gen(volume_group_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("volume", "group", "list"),
		"sort-column": action.ActionColumns("volume", "group", "list"),
	})
	carapace.Gen(volume_group_setCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "group", "list"),
	})
	carapace.Gen(volume_group_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "group", "list"),
	})
	carapace.Gen(volume_group_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "group", "snapshot", "list"),
	})
	carapace.Gen(volume_group_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("volume", "group", "snapshot", "list"),
		"sort-column": action.ActionColumns("volume", "group", "snapshot", "list"),
	})
	carapace.Gen(volume_group_snapshot_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "group", "snapshot", "list"),
	})
	carapace.Gen(volume_group_type_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "group", "type", "list"),
	})
	carapace.Gen(volume_group_type_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("volume", "group", "type", "list"),
		"sort-column": action.ActionColumns("volume", "group", "type", "list"),
	})
	carapace.Gen(volume_group_type_setCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "group", "type", "list"),
	})
	carapace.Gen(volume_group_type_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "group", "type", "list"),
	})
	carapace.Gen(volume_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("volume", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("volume", "list"),
		"user":           action.ActionUsers(),
		"user-domain":    action.ActionDomains(),
	})
	carapace.Gen(volume_message_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("volume", "message", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("volume", "message", "list"),
	})
	carapace.Gen(volume_message_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "message", "list"),
	})
	carapace.Gen(volume_qos_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "qos", "list"),
	})
	carapace.Gen(volume_qos_disassociateCmd).FlagCompletion(carapace.ActionMap{
		"volume-type": action.ActionVolumeTypes(),
	})
	carapace.Gen(volume_qos_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("volume", "qos", "list"),
		"sort-column": action.ActionColumns("volume", "qos", "list"),
	})
	carapace.Gen(volume_qos_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "qos", "list"),
	})
	carapace.Gen(volume_service_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("volume", "service", "list"),
		"sort-column": action.ActionColumns("volume", "service", "list"),
	})
	carapace.Gen(volume_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "list"),
	})
	carapace.Gen(volume_snapshot_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "snapshot", "list"),
		"volume": action.ActionVolumes(),
	})
	carapace.Gen(volume_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("volume", "snapshot", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"sort-column":    action.ActionColumns("volume", "snapshot", "list"),
		"volume":         action.ActionVolumes(),
	})
	carapace.Gen(volume_snapshot_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "snapshot", "list"),
	})
	carapace.Gen(volume_summaryCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "list"),
	})
	carapace.Gen(volume_transfer_request_acceptCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "transfer", "request", "list"),
	})
	carapace.Gen(volume_transfer_request_createCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "transfer", "request", "list"),
	})
	carapace.Gen(volume_transfer_request_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("volume", "transfer", "request", "list"),
		"sort-column": action.ActionColumns("volume", "transfer", "request", "list"),
	})
	carapace.Gen(volume_transfer_request_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "transfer", "request", "list"),
	})
	carapace.Gen(volume_type_createCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("volume", "type", "list"),
		"project":           action.ActionProjects(),
		"project-domain":    action.ActionDomains(),
	})
	carapace.Gen(volume_type_listCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"column":            action.ActionColumns("volume", "type", "list"),
		"sort-column":       action.ActionColumns("volume", "type", "list"),
	})
	carapace.Gen(volume_type_setCmd).FlagCompletion(carapace.ActionMap{
		"availability-zone": action.ActionAvailabilityZones(),
		"project":           action.ActionProjects(),
		"project-domain":    action.ActionDomains(),
	})
	carapace.Gen(volume_type_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("volume", "type", "list"),
	})
	carapace.Gen(volume_type_unsetCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(vpn_endpoint_group_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("vpn", "endpoint", "group", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(vpn_endpoint_group_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("vpn", "endpoint", "group", "list"),
		"sort-column": action.ActionColumns("vpn", "endpoint", "group", "list"),
	})
	carapace.Gen(vpn_endpoint_group_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("vpn", "endpoint", "group", "list"),
	})
	carapace.Gen(vpn_ike_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("vpn", "ike", "policy", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(vpn_ike_policy_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("vpn", "ike", "policy", "list"),
		"sort-column": action.ActionColumns("vpn", "ike", "policy", "list"),
	})
	carapace.Gen(vpn_ike_policy_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("vpn", "ike", "policy", "list"),
	})
	carapace.Gen(vpn_ipsec_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("vpn", "ipsec", "policy", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(vpn_ipsec_policy_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("vpn", "ipsec", "policy", "list"),
		"sort-column": action.ActionColumns("vpn", "ipsec", "policy", "list"),
	})
	carapace.Gen(vpn_ipsec_policy_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("vpn", "ipsec", "policy", "list"),
	})
	carapace.Gen(vpn_ipsec_site_connection_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("vpn", "ipsec", "site", "connection", "list"),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
	})
	carapace.Gen(vpn_ipsec_site_connection_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("vpn", "ipsec", "site", "connection", "list"),
		"sort-column": action.ActionColumns("vpn", "ipsec", "site", "connection", "list"),
	})
	carapace.Gen(vpn_ipsec_site_connection_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("vpn", "ipsec", "site", "connection", "list"),
	})
	carapace.Gen(vpn_service_createCmd).FlagCompletion(carapace.ActionMap{
		"column":         action.ActionColumns("vpn", "service", "list"),
		"flavor":         action.ActionFlavors(),
		"project":        action.ActionProjects(),
		"project-domain": action.ActionDomains(),
		"router":         action.ActionRouters(),
		"subnet":         action.ActionSubnets(),
	})
	carapace.Gen(vpn_service_listCmd).FlagCompletion(carapace.ActionMap{
		"column":      action.ActionColumns("vpn", "service", "list"),
		"sort-column": action.ActionColumns("vpn", "service", "list"),
	})
	carapace.Gen(vpn_service_setCmd).FlagCompletion(carapace.ActionMap{
		"flavor": action.ActionFlavors(),
		"subnet": action.ActionSubnets(),
	})
	carapace.Gen(vpn_service_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("vpn", "service", "list"),
	})
}
