/*
Package sampling is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by /usr/local/google/home/gdennis/go/pkg/mod/github.com/openconfig/ygot@v0.23.1/genutil/names.go
using the following YANG input files:
  - gnmi-collector-metadata.yang
  - gnsi/authz/gnsi-authz.yang
  - gnsi/cert/gnsi-cert.yang
  - gnsi/console/gnsi-console.yang
  - gnsi/pathz/gnsi-pathz.yang
  - gnsi/ssh/gnsi-ssh.yang
  - public/release/models/acl/openconfig-acl.yang
  - public/release/models/acl/openconfig-packet-match.yang
  - public/release/models/aft/openconfig-aft.yang
  - public/release/models/aft/openconfig-aft-network-instance.yang
  - public/release/models/ate/openconfig-ate-flow.yang
  - public/release/models/ate/openconfig-ate-intf.yang
  - public/release/models/bfd/openconfig-bfd.yang
  - public/release/models/bgp/openconfig-bgp-policy.yang
  - public/release/models/bgp/openconfig-bgp-types.yang
  - public/release/models/interfaces/openconfig-if-aggregate.yang
  - public/release/models/interfaces/openconfig-if-ethernet.yang
  - public/release/models/interfaces/openconfig-if-ethernet-ext.yang
  - public/release/models/interfaces/openconfig-if-ip-ext.yang
  - public/release/models/interfaces/openconfig-if-ip.yang
  - public/release/models/interfaces/openconfig-interfaces.yang
  - public/release/models/isis/openconfig-isis.yang
  - public/release/models/lacp/openconfig-lacp.yang
  - public/release/models/lldp/openconfig-lldp-types.yang
  - public/release/models/lldp/openconfig-lldp.yang
  - public/release/models/local-routing/openconfig-local-routing.yang
  - public/release/models/mpls/openconfig-mpls-types.yang
  - public/release/models/multicast/openconfig-pim.yang
  - public/release/models/network-instance/openconfig-network-instance.yang
  - public/release/models/openconfig-extensions.yang
  - public/release/models/optical-transport/openconfig-transport-types.yang
  - public/release/models/ospf/openconfig-ospfv2.yang
  - public/release/models/p4rt/openconfig-p4rt.yang
  - public/release/models/platform/openconfig-platform-cpu.yang
  - public/release/models/platform/openconfig-platform-fan.yang
  - public/release/models/platform/openconfig-platform-integrated-circuit.yang
  - public/release/models/platform/openconfig-platform-software.yang
  - public/release/models/platform/openconfig-platform-transceiver.yang
  - public/release/models/platform/openconfig-platform.yang
  - public/release/models/policy-forwarding/openconfig-policy-forwarding.yang
  - public/release/models/policy/openconfig-policy-types.yang
  - public/release/models/qos/openconfig-qos-elements.yang
  - public/release/models/qos/openconfig-qos-interfaces.yang
  - public/release/models/qos/openconfig-qos-types.yang
  - public/release/models/qos/openconfig-qos.yang
  - public/release/models/rib/openconfig-rib-bgp.yang
  - public/release/models/sampling/openconfig-sampling-sflow.yang
  - public/release/models/segment-routing/openconfig-segment-routing-types.yang
  - public/release/models/system/openconfig-system.yang
  - public/release/models/types/openconfig-inet-types.yang
  - public/release/models/types/openconfig-types.yang
  - public/release/models/types/openconfig-yang-types.yang
  - public/release/models/vlan/openconfig-vlan.yang
  - public/third_party/ietf/iana-if-type.yang
  - public/third_party/ietf/ietf-inet-types.yang
  - public/third_party/ietf/ietf-interfaces.yang
  - public/third_party/ietf/ietf-yang-types.yang

Imported modules were sourced from:
  - public/release/models/...
  - public/third_party/ietf/...
*/
package sampling

import (
	"github.com/openconfig/ygot/ygot"
)

// Sampling_Sflow_CollectorPath represents the /openconfig-sampling/sampling/sflow/collectors/collector YANG schema element.
type Sampling_Sflow_CollectorPath struct {
	*ygot.NodePath
}

// Sampling_Sflow_CollectorPathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/collectors/collector YANG schema element.
type Sampling_Sflow_CollectorPathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Collector_AddressPath represents the /openconfig-sampling/sampling/sflow/collectors/collector/state/address YANG schema element.
type Sampling_Sflow_Collector_AddressPath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Collector_AddressPathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/collectors/collector/state/address YANG schema element.
type Sampling_Sflow_Collector_AddressPathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Collector_NetworkInstancePath represents the /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance YANG schema element.
type Sampling_Sflow_Collector_NetworkInstancePath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Collector_NetworkInstancePathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance YANG schema element.
type Sampling_Sflow_Collector_NetworkInstancePathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Collector_PacketsSentPath represents the /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent YANG schema element.
type Sampling_Sflow_Collector_PacketsSentPath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Collector_PacketsSentPathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent YANG schema element.
type Sampling_Sflow_Collector_PacketsSentPathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Collector_PortPath represents the /openconfig-sampling/sampling/sflow/collectors/collector/state/port YANG schema element.
type Sampling_Sflow_Collector_PortPath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Collector_PortPathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/collectors/collector/state/port YANG schema element.
type Sampling_Sflow_Collector_PortPathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Collector_SourceAddressPath represents the /openconfig-sampling/sampling/sflow/collectors/collector/state/source-address YANG schema element.
type Sampling_Sflow_Collector_SourceAddressPath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Collector_SourceAddressPathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/collectors/collector/state/source-address YANG schema element.
type Sampling_Sflow_Collector_SourceAddressPathAny struct {
	*ygot.NodePath
}

// Address (leaf): IPv4/IPv6 address of the sFlow collector.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/address"
// Path from root: "/sampling/sflow/collectors/collector/state/address"
func (n *Sampling_Sflow_CollectorPath) Address() *Sampling_Sflow_Collector_AddressPath {
	return &Sampling_Sflow_Collector_AddressPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "address"},
			map[string]interface{}{},
			n,
		),
	}
}

// Address (leaf): IPv4/IPv6 address of the sFlow collector.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/address"
// Path from root: "/sampling/sflow/collectors/collector/state/address"
func (n *Sampling_Sflow_CollectorPathAny) Address() *Sampling_Sflow_Collector_AddressPathAny {
	return &Sampling_Sflow_Collector_AddressPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "address"},
			map[string]interface{}{},
			n,
		),
	}
}

// NetworkInstance (leaf): Reference to the network instance used to reach the
// sFlow collector.  If uspecified, the collector destination
// is reachable in the default network instance.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/network-instance"
// Path from root: "/sampling/sflow/collectors/collector/state/network-instance"
func (n *Sampling_Sflow_CollectorPath) NetworkInstance() *Sampling_Sflow_Collector_NetworkInstancePath {
	return &Sampling_Sflow_Collector_NetworkInstancePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "network-instance"},
			map[string]interface{}{},
			n,
		),
	}
}

// NetworkInstance (leaf): Reference to the network instance used to reach the
// sFlow collector.  If uspecified, the collector destination
// is reachable in the default network instance.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/network-instance"
// Path from root: "/sampling/sflow/collectors/collector/state/network-instance"
func (n *Sampling_Sflow_CollectorPathAny) NetworkInstance() *Sampling_Sflow_Collector_NetworkInstancePathAny {
	return &Sampling_Sflow_Collector_NetworkInstancePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "network-instance"},
			map[string]interface{}{},
			n,
		),
	}
}

// PacketsSent (leaf): The total number of packets sampled and sent to the
// collector.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/packets-sent"
// Path from root: "/sampling/sflow/collectors/collector/state/packets-sent"
func (n *Sampling_Sflow_CollectorPath) PacketsSent() *Sampling_Sflow_Collector_PacketsSentPath {
	return &Sampling_Sflow_Collector_PacketsSentPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "packets-sent"},
			map[string]interface{}{},
			n,
		),
	}
}

// PacketsSent (leaf): The total number of packets sampled and sent to the
// collector.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/packets-sent"
// Path from root: "/sampling/sflow/collectors/collector/state/packets-sent"
func (n *Sampling_Sflow_CollectorPathAny) PacketsSent() *Sampling_Sflow_Collector_PacketsSentPathAny {
	return &Sampling_Sflow_Collector_PacketsSentPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "packets-sent"},
			map[string]interface{}{},
			n,
		),
	}
}

// Port (leaf): UDP port number for the sFlow collector.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/port"
// Path from root: "/sampling/sflow/collectors/collector/state/port"
func (n *Sampling_Sflow_CollectorPath) Port() *Sampling_Sflow_Collector_PortPath {
	return &Sampling_Sflow_Collector_PortPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "port"},
			map[string]interface{}{},
			n,
		),
	}
}

// Port (leaf): UDP port number for the sFlow collector.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/port"
// Path from root: "/sampling/sflow/collectors/collector/state/port"
func (n *Sampling_Sflow_CollectorPathAny) Port() *Sampling_Sflow_Collector_PortPathAny {
	return &Sampling_Sflow_Collector_PortPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "port"},
			map[string]interface{}{},
			n,
		),
	}
}

// SourceAddress (leaf): Sets the source IPv4/IPv6 address for sFlow datagrams sent
// to sFlow collectors.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/source-address"
// Path from root: "/sampling/sflow/collectors/collector/state/source-address"
func (n *Sampling_Sflow_CollectorPath) SourceAddress() *Sampling_Sflow_Collector_SourceAddressPath {
	return &Sampling_Sflow_Collector_SourceAddressPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "source-address"},
			map[string]interface{}{},
			n,
		),
	}
}

// SourceAddress (leaf): Sets the source IPv4/IPv6 address for sFlow datagrams sent
// to sFlow collectors.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/source-address"
// Path from root: "/sampling/sflow/collectors/collector/state/source-address"
func (n *Sampling_Sflow_CollectorPathAny) SourceAddress() *Sampling_Sflow_Collector_SourceAddressPathAny {
	return &Sampling_Sflow_Collector_SourceAddressPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "source-address"},
			map[string]interface{}{},
			n,
		),
	}
}

// Sampling_Sflow_InterfacePath represents the /openconfig-sampling/sampling/sflow/interfaces/interface YANG schema element.
type Sampling_Sflow_InterfacePath struct {
	*ygot.NodePath
}

// Sampling_Sflow_InterfacePathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/interfaces/interface YANG schema element.
type Sampling_Sflow_InterfacePathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_EgressSamplingRatePath represents the /openconfig-sampling/sampling/sflow/interfaces/interface/state/egress-sampling-rate YANG schema element.
type Sampling_Sflow_Interface_EgressSamplingRatePath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_EgressSamplingRatePathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/interfaces/interface/state/egress-sampling-rate YANG schema element.
type Sampling_Sflow_Interface_EgressSamplingRatePathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_EnabledPath represents the /openconfig-sampling/sampling/sflow/interfaces/interface/state/enabled YANG schema element.
type Sampling_Sflow_Interface_EnabledPath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_EnabledPathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/interfaces/interface/state/enabled YANG schema element.
type Sampling_Sflow_Interface_EnabledPathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_IngressSamplingRatePath represents the /openconfig-sampling/sampling/sflow/interfaces/interface/state/ingress-sampling-rate YANG schema element.
type Sampling_Sflow_Interface_IngressSamplingRatePath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_IngressSamplingRatePathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/interfaces/interface/state/ingress-sampling-rate YANG schema element.
type Sampling_Sflow_Interface_IngressSamplingRatePathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_NamePath represents the /openconfig-sampling/sampling/sflow/interfaces/interface/state/name YANG schema element.
type Sampling_Sflow_Interface_NamePath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_NamePathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/interfaces/interface/state/name YANG schema element.
type Sampling_Sflow_Interface_NamePathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_PacketsSampledPath represents the /openconfig-sampling/sampling/sflow/interfaces/interface/state/packets-sampled YANG schema element.
type Sampling_Sflow_Interface_PacketsSampledPath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_PacketsSampledPathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/interfaces/interface/state/packets-sampled YANG schema element.
type Sampling_Sflow_Interface_PacketsSampledPathAny struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_PollingIntervalPath represents the /openconfig-sampling/sampling/sflow/interfaces/interface/state/polling-interval YANG schema element.
type Sampling_Sflow_Interface_PollingIntervalPath struct {
	*ygot.NodePath
}

// Sampling_Sflow_Interface_PollingIntervalPathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/interfaces/interface/state/polling-interval YANG schema element.
type Sampling_Sflow_Interface_PollingIntervalPathAny struct {
	*ygot.NodePath
}

// EgressSamplingRate (leaf): Sets the egress packet sampling rate.  The rate is expressed
// as an integer N, where the intended sampling rate is 1/N
// packets.  An implementation may implement the sampling rate as
// a statistical average, rather than a strict periodic sampling.
//
// The allowable sampling rate range is generally a property of
// the system, e.g., determined by the capability of the
// hardware.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/egress-sampling-rate"
// Path from root: "/sampling/sflow/interfaces/interface/state/egress-sampling-rate"
func (n *Sampling_Sflow_InterfacePath) EgressSamplingRate() *Sampling_Sflow_Interface_EgressSamplingRatePath {
	return &Sampling_Sflow_Interface_EgressSamplingRatePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "egress-sampling-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// EgressSamplingRate (leaf): Sets the egress packet sampling rate.  The rate is expressed
// as an integer N, where the intended sampling rate is 1/N
// packets.  An implementation may implement the sampling rate as
// a statistical average, rather than a strict periodic sampling.
//
// The allowable sampling rate range is generally a property of
// the system, e.g., determined by the capability of the
// hardware.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/egress-sampling-rate"
// Path from root: "/sampling/sflow/interfaces/interface/state/egress-sampling-rate"
func (n *Sampling_Sflow_InterfacePathAny) EgressSamplingRate() *Sampling_Sflow_Interface_EgressSamplingRatePathAny {
	return &Sampling_Sflow_Interface_EgressSamplingRatePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "egress-sampling-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// Enabled (leaf): Enables or disables sFlow on the interface.  If sFlow is
// globally disabled, this leaf is ignored.  If sFlow
// is globally enabled, this leaf may be used to disable it
// for a specific interface.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/enabled"
// Path from root: "/sampling/sflow/interfaces/interface/state/enabled"
func (n *Sampling_Sflow_InterfacePath) Enabled() *Sampling_Sflow_Interface_EnabledPath {
	return &Sampling_Sflow_Interface_EnabledPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "enabled"},
			map[string]interface{}{},
			n,
		),
	}
}

// Enabled (leaf): Enables or disables sFlow on the interface.  If sFlow is
// globally disabled, this leaf is ignored.  If sFlow
// is globally enabled, this leaf may be used to disable it
// for a specific interface.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/enabled"
// Path from root: "/sampling/sflow/interfaces/interface/state/enabled"
func (n *Sampling_Sflow_InterfacePathAny) Enabled() *Sampling_Sflow_Interface_EnabledPathAny {
	return &Sampling_Sflow_Interface_EnabledPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "enabled"},
			map[string]interface{}{},
			n,
		),
	}
}

// IngressSamplingRate (leaf): Sets the ingress packet sampling rate.  The rate is expressed
// as an integer N, where the intended sampling rate is 1/N
// packets.  An implementation may implement the sampling rate as
// a statistical average, rather than a strict periodic sampling.
//
// The allowable sampling rate range is generally a property of
// the system, e.g., determined by the capability of the
// hardware.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/ingress-sampling-rate"
// Path from root: "/sampling/sflow/interfaces/interface/state/ingress-sampling-rate"
func (n *Sampling_Sflow_InterfacePath) IngressSamplingRate() *Sampling_Sflow_Interface_IngressSamplingRatePath {
	return &Sampling_Sflow_Interface_IngressSamplingRatePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "ingress-sampling-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// IngressSamplingRate (leaf): Sets the ingress packet sampling rate.  The rate is expressed
// as an integer N, where the intended sampling rate is 1/N
// packets.  An implementation may implement the sampling rate as
// a statistical average, rather than a strict periodic sampling.
//
// The allowable sampling rate range is generally a property of
// the system, e.g., determined by the capability of the
// hardware.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/ingress-sampling-rate"
// Path from root: "/sampling/sflow/interfaces/interface/state/ingress-sampling-rate"
func (n *Sampling_Sflow_InterfacePathAny) IngressSamplingRate() *Sampling_Sflow_Interface_IngressSamplingRatePathAny {
	return &Sampling_Sflow_Interface_IngressSamplingRatePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "ingress-sampling-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): Reference to the interface for sFlow configuration and
// state.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/name"
// Path from root: "/sampling/sflow/interfaces/interface/state/name"
func (n *Sampling_Sflow_InterfacePath) Name() *Sampling_Sflow_Interface_NamePath {
	return &Sampling_Sflow_Interface_NamePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): Reference to the interface for sFlow configuration and
// state.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/name"
// Path from root: "/sampling/sflow/interfaces/interface/state/name"
func (n *Sampling_Sflow_InterfacePathAny) Name() *Sampling_Sflow_Interface_NamePathAny {
	return &Sampling_Sflow_Interface_NamePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// PacketsSampled (leaf): Total number of packets sampled from the interface.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/packets-sampled"
// Path from root: "/sampling/sflow/interfaces/interface/state/packets-sampled"
func (n *Sampling_Sflow_InterfacePath) PacketsSampled() *Sampling_Sflow_Interface_PacketsSampledPath {
	return &Sampling_Sflow_Interface_PacketsSampledPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "packets-sampled"},
			map[string]interface{}{},
			n,
		),
	}
}

// PacketsSampled (leaf): Total number of packets sampled from the interface.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/packets-sampled"
// Path from root: "/sampling/sflow/interfaces/interface/state/packets-sampled"
func (n *Sampling_Sflow_InterfacePathAny) PacketsSampled() *Sampling_Sflow_Interface_PacketsSampledPathAny {
	return &Sampling_Sflow_Interface_PacketsSampledPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "packets-sampled"},
			map[string]interface{}{},
			n,
		),
	}
}

// PollingInterval (leaf): Sets the traffic sampling polling interval.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/polling-interval"
// Path from root: "/sampling/sflow/interfaces/interface/state/polling-interval"
func (n *Sampling_Sflow_InterfacePath) PollingInterval() *Sampling_Sflow_Interface_PollingIntervalPath {
	return &Sampling_Sflow_Interface_PollingIntervalPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "polling-interval"},
			map[string]interface{}{},
			n,
		),
	}
}

// PollingInterval (leaf): Sets the traffic sampling polling interval.
// ----------------------------------------
// Defining module: "openconfig-sampling-sflow"
// Instantiating module: "openconfig-sampling"
// Path from parent: "state/polling-interval"
// Path from root: "/sampling/sflow/interfaces/interface/state/polling-interval"
func (n *Sampling_Sflow_InterfacePathAny) PollingInterval() *Sampling_Sflow_Interface_PollingIntervalPathAny {
	return &Sampling_Sflow_Interface_PollingIntervalPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "polling-interval"},
			map[string]interface{}{},
			n,
		),
	}
}
