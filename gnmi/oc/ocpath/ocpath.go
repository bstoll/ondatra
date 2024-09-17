/*
Package ocpath is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema.

This package was generated by ygnmi version: v0.11.1: (ygot: v0.29.18)
using the following YANG input files:
  - gnmi-collector-metadata.yang
  - public/release/models/acl/openconfig-acl.yang
  - public/release/models/acl/openconfig-packet-match.yang
  - public/release/models/aft/openconfig-aft-network-instance.yang
  - public/release/models/aft/openconfig-aft-summary.yang
  - public/release/models/aft/openconfig-aft.yang
  - public/release/models/ate/openconfig-ate-flow.yang
  - public/release/models/ate/openconfig-ate-intf.yang
  - public/release/models/bfd/openconfig-bfd.yang
  - public/release/models/bgp/openconfig-bgp-policy.yang
  - public/release/models/bgp/openconfig-bgp-types.yang
  - public/release/models/extensions/openconfig-metadata.yang
  - public/release/models/gnsi/openconfig-gnsi-acctz.yang
  - public/release/models/gnsi/openconfig-gnsi-authz.yang
  - public/release/models/gnsi/openconfig-gnsi-certz.yang
  - public/release/models/gnsi/openconfig-gnsi-credentialz.yang
  - public/release/models/gnsi/openconfig-gnsi-pathz.yang
  - public/release/models/gnsi/openconfig-gnsi.yang
  - public/release/models/gribi/openconfig-gribi.yang
  - public/release/models/interfaces/openconfig-if-aggregate.yang
  - public/release/models/interfaces/openconfig-if-ethernet-ext.yang
  - public/release/models/interfaces/openconfig-if-ethernet.yang
  - public/release/models/interfaces/openconfig-if-ip-ext.yang
  - public/release/models/interfaces/openconfig-if-ip.yang
  - public/release/models/interfaces/openconfig-if-sdn-ext.yang
  - public/release/models/interfaces/openconfig-interfaces.yang
  - public/release/models/isis/openconfig-isis-policy.yang
  - public/release/models/isis/openconfig-isis.yang
  - public/release/models/lacp/openconfig-lacp.yang
  - public/release/models/lldp/openconfig-lldp-types.yang
  - public/release/models/lldp/openconfig-lldp.yang
  - public/release/models/local-routing/openconfig-local-routing.yang
  - public/release/models/mpls/openconfig-mpls-types.yang
  - public/release/models/multicast/openconfig-pim.yang
  - public/release/models/network-instance/openconfig-network-instance.yang
  - public/release/models/openconfig-extensions.yang
  - public/release/models/optical-transport/openconfig-terminal-device.yang
  - public/release/models/optical-transport/openconfig-transport-types.yang
  - public/release/models/ospf/openconfig-ospf-policy.yang
  - public/release/models/ospf/openconfig-ospfv2.yang
  - public/release/models/p4rt/openconfig-p4rt.yang
  - public/release/models/platform/openconfig-platform-common.yang
  - public/release/models/platform/openconfig-platform-controller-card.yang
  - public/release/models/platform/openconfig-platform-cpu.yang
  - public/release/models/platform/openconfig-platform-ext.yang
  - public/release/models/platform/openconfig-platform-fabric.yang
  - public/release/models/platform/openconfig-platform-fan.yang
  - public/release/models/platform/openconfig-platform-integrated-circuit.yang
  - public/release/models/platform/openconfig-platform-linecard.yang
  - public/release/models/platform/openconfig-platform-pipeline-counters.yang
  - public/release/models/platform/openconfig-platform-psu.yang
  - public/release/models/platform/openconfig-platform-software.yang
  - public/release/models/platform/openconfig-platform-transceiver.yang
  - public/release/models/platform/openconfig-platform.yang
  - public/release/models/policy-forwarding/openconfig-policy-forwarding.yang
  - public/release/models/policy/openconfig-policy-types.yang
  - public/release/models/qos/openconfig-qos-elements.yang
  - public/release/models/qos/openconfig-qos-interfaces.yang
  - public/release/models/qos/openconfig-qos-types.yang
  - public/release/models/qos/openconfig-qos.yang
  - public/release/models/relay-agent/openconfig-relay-agent.yang
  - public/release/models/rib/openconfig-rib-bgp.yang
  - public/release/models/sampling/openconfig-sampling-sflow.yang
  - public/release/models/segment-routing/openconfig-segment-routing-types.yang
  - public/release/models/system/openconfig-system-bootz.yang
  - public/release/models/system/openconfig-system-controlplane.yang
  - public/release/models/system/openconfig-system-utilization.yang
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
package ocpath

import (
	oc "github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ondatra/gnmi/oc/acl"
	"github.com/openconfig/ondatra/gnmi/oc/ateflow"
	"github.com/openconfig/ondatra/gnmi/oc/definedsets"
	"github.com/openconfig/ondatra/gnmi/oc/gnmicollectormetadata"
	"github.com/openconfig/ondatra/gnmi/oc/interfaces"
	"github.com/openconfig/ondatra/gnmi/oc/keychain"
	"github.com/openconfig/ondatra/gnmi/oc/lacp"
	"github.com/openconfig/ondatra/gnmi/oc/lldp"
	"github.com/openconfig/ondatra/gnmi/oc/networkinstance"
	"github.com/openconfig/ondatra/gnmi/oc/platform"
	"github.com/openconfig/ondatra/gnmi/oc/qos"
	"github.com/openconfig/ondatra/gnmi/oc/relayagent"
	"github.com/openconfig/ondatra/gnmi/oc/routingpolicy"
	"github.com/openconfig/ondatra/gnmi/oc/sampling"
	"github.com/openconfig/ondatra/gnmi/oc/system"
	"github.com/openconfig/ondatra/gnmi/oc/terminaldevice"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// RootPath represents the /root YANG schema element.
type RootPath struct {
	*ygnmi.DeviceRootBase
}

// Root returns a root path object from which YANG paths can be constructed.
func Root() *RootPath {
	return &RootPath{ygnmi.NewDeviceRootBase()}
}

// Acl (container): Top level enclosing container for ACL model config
// and operational state data
//
//	Defining module:      "openconfig-acl"
//	Instantiating module: "openconfig-acl"
//	Path from parent:     "acl"
//	Path from root:       "/acl"
func (n *RootPath) Acl() *acl.AclPath {
	ps := &acl.AclPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"acl"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// ComponentAny (list): List of components, keyed by component name.
//
//	Defining module:      "openconfig-platform"
//	Instantiating module: "openconfig-platform"
//	Path from parent:     "components/component"
//	Path from root:       "/components/component"
func (n *RootPath) ComponentAny() *platform.ComponentPathAny {
	ps := &platform.ComponentPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"components", "component"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
	return ps
}

// Component (list): List of components, keyed by component name.
//
//	Defining module:      "openconfig-platform"
//	Instantiating module: "openconfig-platform"
//	Path from parent:     "components/component"
//	Path from root:       "/components/component"
//
//	Name: string
func (n *RootPath) Component(Name string) *platform.ComponentPath {
	ps := &platform.ComponentPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"components", "component"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
	return ps
}

// ComponentMap (list): List of components, keyed by component name.
//
//	Defining module:      "openconfig-platform"
//	Instantiating module: "openconfig-platform"
//	Path from parent:     "components/component"
//	Path from root:       "/components/component"
func (n *RootPath) ComponentMap() *platform.ComponentPathMap {
	ps := &platform.ComponentPathMap{
		NodePath: ygnmi.NewNodePath(
			[]string{"components"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// DefinedSets (container): Top level enclosing container for defined-set model
// config and operational state data.
//
//	Defining module:      "openconfig-defined-sets"
//	Instantiating module: "openconfig-defined-sets"
//	Path from parent:     "defined-sets"
//	Path from root:       "/defined-sets"
func (n *RootPath) DefinedSets() *definedsets.DefinedSetsPath {
	ps := &definedsets.DefinedSetsPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"defined-sets"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// FlowAny (list): A flow of packets between one or more internal and external sources
// and one or more internal and external destinations that the target
// is able to track and report statistics on. Each flow is identified by
// an arbitrary string identifier.
//
//	Defining module:      "openconfig-ate-flow"
//	Instantiating module: "openconfig-ate-flow"
//	Path from parent:     "flows/flow"
//	Path from root:       "/flows/flow"
func (n *RootPath) FlowAny() *ateflow.FlowPathAny {
	ps := &ateflow.FlowPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"flows", "flow"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
	return ps
}

// Flow (list): A flow of packets between one or more internal and external sources
// and one or more internal and external destinations that the target
// is able to track and report statistics on. Each flow is identified by
// an arbitrary string identifier.
//
//	Defining module:      "openconfig-ate-flow"
//	Instantiating module: "openconfig-ate-flow"
//	Path from parent:     "flows/flow"
//	Path from root:       "/flows/flow"
//
//	Name: string
func (n *RootPath) Flow(Name string) *ateflow.FlowPath {
	ps := &ateflow.FlowPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"flows", "flow"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
	return ps
}

// FlowMap (list): A flow of packets between one or more internal and external sources
// and one or more internal and external destinations that the target
// is able to track and report statistics on. Each flow is identified by
// an arbitrary string identifier.
//
//	Defining module:      "openconfig-ate-flow"
//	Instantiating module: "openconfig-ate-flow"
//	Path from parent:     "flows/flow"
//	Path from root:       "/flows/flow"
func (n *RootPath) FlowMap() *ateflow.FlowPathMap {
	ps := &ateflow.FlowPathMap{
		NodePath: ygnmi.NewNodePath(
			[]string{"flows"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// InterfaceAny (list): The list of named interfaces on the device.
//
//	Defining module:      "openconfig-interfaces"
//	Instantiating module: "openconfig-interfaces"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/interfaces/interface"
func (n *RootPath) InterfaceAny() *interfaces.InterfacePathAny {
	ps := &interfaces.InterfacePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
	return ps
}

// Interface (list): The list of named interfaces on the device.
//
//	Defining module:      "openconfig-interfaces"
//	Instantiating module: "openconfig-interfaces"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/interfaces/interface"
//
//	Name: string
func (n *RootPath) Interface(Name string) *interfaces.InterfacePath {
	ps := &interfaces.InterfacePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
	return ps
}

// InterfaceMap (list): The list of named interfaces on the device.
//
//	Defining module:      "openconfig-interfaces"
//	Instantiating module: "openconfig-interfaces"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/interfaces/interface"
func (n *RootPath) InterfaceMap() *interfaces.InterfacePathMap {
	ps := &interfaces.InterfacePathMap{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// KeychainAny (list): List of defined keychains.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keychains/keychain"
//	Path from root:       "/keychains/keychain"
func (n *RootPath) KeychainAny() *keychain.KeychainPathAny {
	ps := &keychain.KeychainPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"keychains", "keychain"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
	return ps
}

// Keychain (list): List of defined keychains.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keychains/keychain"
//	Path from root:       "/keychains/keychain"
//
//	Name: string
func (n *RootPath) Keychain(Name string) *keychain.KeychainPath {
	ps := &keychain.KeychainPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"keychains", "keychain"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
	return ps
}

// KeychainMap (list): List of defined keychains.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keychains/keychain"
//	Path from root:       "/keychains/keychain"
func (n *RootPath) KeychainMap() *keychain.KeychainPathMap {
	ps := &keychain.KeychainPathMap{
		NodePath: ygnmi.NewNodePath(
			[]string{"keychains"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// Lacp (container): Configuration and operational state data for LACP protocol
// operation on the aggregate interface
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp"
//	Path from root:       "/lacp"
func (n *RootPath) Lacp() *lacp.LacpPath {
	ps := &lacp.LacpPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// Lldp (container): Top-level container for LLDP configuration and state data
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "lldp"
//	Path from root:       "/lldp"
func (n *RootPath) Lldp() *lldp.LldpPath {
	ps := &lldp.LldpPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lldp"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// Meta (container):
//
//	Defining module:      "gnmi-collector-metadata"
//	Instantiating module: "gnmi-collector-metadata"
//	Path from parent:     "meta"
//	Path from root:       "/meta"
func (n *RootPath) Meta() *gnmicollectormetadata.MetaPath {
	ps := &gnmicollectormetadata.MetaPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"meta"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// NetworkInstanceAny (list): Network instances configured on the local system
//
// IPv4 and IPv6 forwarding are enabled by default within an L3
// network-instance and subsequently, routes can be populated
// into the network-instance using protocols that enable IPv4 and
// IPv6 configuration without explicitly enabling these.
//
//	Defining module:      "openconfig-network-instance"
//	Instantiating module: "openconfig-network-instance"
//	Path from parent:     "network-instances/network-instance"
//	Path from root:       "/network-instances/network-instance"
func (n *RootPath) NetworkInstanceAny() *networkinstance.NetworkInstancePathAny {
	ps := &networkinstance.NetworkInstancePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"network-instances", "network-instance"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
	return ps
}

// NetworkInstance (list): Network instances configured on the local system
//
// IPv4 and IPv6 forwarding are enabled by default within an L3
// network-instance and subsequently, routes can be populated
// into the network-instance using protocols that enable IPv4 and
// IPv6 configuration without explicitly enabling these.
//
//	Defining module:      "openconfig-network-instance"
//	Instantiating module: "openconfig-network-instance"
//	Path from parent:     "network-instances/network-instance"
//	Path from root:       "/network-instances/network-instance"
//
//	Name: string
func (n *RootPath) NetworkInstance(Name string) *networkinstance.NetworkInstancePath {
	ps := &networkinstance.NetworkInstancePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"network-instances", "network-instance"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
	return ps
}

// NetworkInstanceMap (list): Network instances configured on the local system
//
// IPv4 and IPv6 forwarding are enabled by default within an L3
// network-instance and subsequently, routes can be populated
// into the network-instance using protocols that enable IPv4 and
// IPv6 configuration without explicitly enabling these.
//
//	Defining module:      "openconfig-network-instance"
//	Instantiating module: "openconfig-network-instance"
//	Path from parent:     "network-instances/network-instance"
//	Path from root:       "/network-instances/network-instance"
func (n *RootPath) NetworkInstanceMap() *networkinstance.NetworkInstancePathMap {
	ps := &networkinstance.NetworkInstancePathMap{
		NodePath: ygnmi.NewNodePath(
			[]string{"network-instances"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// Qos (container): Top-level container for QoS data
//
//	Defining module:      "openconfig-qos"
//	Instantiating module: "openconfig-qos"
//	Path from parent:     "qos"
//	Path from root:       "/qos"
func (n *RootPath) Qos() *qos.QosPath {
	ps := &qos.QosPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"qos"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// RelayAgent (container): Top level container for relay-agent configuration and
// operational state data
//
//	Defining module:      "openconfig-relay-agent"
//	Instantiating module: "openconfig-relay-agent"
//	Path from parent:     "relay-agent"
//	Path from root:       "/relay-agent"
func (n *RootPath) RelayAgent() *relayagent.RelayAgentPath {
	ps := &relayagent.RelayAgentPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"relay-agent"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// RoutingPolicy (container): Top-level container for all routing policy configuration
//
//	Defining module:      "openconfig-routing-policy"
//	Instantiating module: "openconfig-routing-policy"
//	Path from parent:     "routing-policy"
//	Path from root:       "/routing-policy"
func (n *RootPath) RoutingPolicy() *routingpolicy.RoutingPolicyPath {
	ps := &routingpolicy.RoutingPolicyPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"routing-policy"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// Sampling (container): Top-level container for sampling-related configuration and
// operational state data
//
//	Defining module:      "openconfig-sampling"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "sampling"
//	Path from root:       "/sampling"
func (n *RootPath) Sampling() *sampling.SamplingPath {
	ps := &sampling.SamplingPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"sampling"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// System (container): Enclosing container for system-related configuration and
// operational state data
//
//	Defining module:      "openconfig-system"
//	Instantiating module: "openconfig-system"
//	Path from parent:     "system"
//	Path from root:       "/system"
func (n *RootPath) System() *system.SystemPath {
	ps := &system.SystemPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"system"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// TerminalDevice (container): Top-level container for the terminal device
//
//	Defining module:      "openconfig-terminal-device"
//	Instantiating module: "openconfig-terminal-device"
//	Path from parent:     "terminal-device"
//	Path from root:       "/terminal-device"
func (n *RootPath) TerminalDevice() *terminaldevice.TerminalDevicePath {
	ps := &terminaldevice.TerminalDevicePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"terminal-device"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// Batch contains a collection of paths.
// Use batch to call Lookup, Watch, etc. on multiple paths at once.
type Batch struct {
	paths []ygnmi.PathStruct
}

// AddPaths adds the paths to the batch.
func (b *Batch) AddPaths(paths ...ygnmi.PathStruct) *Batch {
	b.paths = append(b.paths, paths...)
	return b
}

// State returns a Query that can be used in gNMI operations.
// The returned query is immutable, adding paths does not modify existing queries.
func (b *Batch) State() ygnmi.SingletonQuery[*oc.Root] {
	queryPaths := make([]ygnmi.PathStruct, len(b.paths))
	copy(queryPaths, b.paths)
	return ygnmi.NewSingletonQuery[*oc.Root](
		"Root",
		true,
		false,
		false,
		false,
		true,
		false,
		ygnmi.NewDeviceRootBase(),
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		queryPaths,
		nil,
	)
}

// Config returns a Query that can be used in gNMI operations.
// The returned query is immutable, adding paths does not modify existing queries.
func (b *Batch) Config() ygnmi.SingletonQuery[*oc.Root] {
	queryPaths := make([]ygnmi.PathStruct, len(b.paths))
	copy(queryPaths, b.paths)
	return ygnmi.NewSingletonQuery[*oc.Root](
		"Root",
		false,
		true,
		false,
		false,
		true,
		false,
		ygnmi.NewDeviceRootBase(),
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		queryPaths,
		nil,
	)
}

func binarySliceToFloatSlice(in []oc.Binary) []float32 {
	converted := make([]float32, 0, len(in))
	for _, binary := range in {
		converted = append(converted, ygot.BinaryToFloat32(binary))
	}
	return converted
}

// State returns a Query that can be used in gNMI operations.
func (n *RootPath) State() ygnmi.SingletonQuery[*oc.Root] {
	return ygnmi.NewSingletonQuery[*oc.Root](
		"Root",
		true,
		false,
		false,
		false,
		true,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *RootPath) Config() ygnmi.ConfigQuery[*oc.Root] {
	return ygnmi.NewConfigQuery[*oc.Root](
		"Root",
		false,
		true,
		false,
		false,
		true,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}
