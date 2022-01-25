/*
Package device is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by /usr/local/google/home/gdennis/go/pkg/mod/github.com/openconfig/ygot@v0.14.0/genutil/names.go
using the following YANG input files:
	- gnmi-collector-metadata.yang
	- public/release/models/acl/openconfig-acl.yang
	- public/release/models/acl/openconfig-packet-match.yang
	- public/release/models/aft/openconfig-aft.yang
	- public/release/models/ate/openconfig-ate-flow.yang
	- public/release/models/ate/openconfig-ate-intf.yang
	- public/release/models/bfd/openconfig-bfd.yang
	- public/release/models/bgp/openconfig-bgp-policy.yang
	- public/release/models/bgp/openconfig-bgp-types.yang
	- public/release/models/interfaces/openconfig-if-aggregate.yang
	- public/release/models/interfaces/openconfig-if-ethernet.yang
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
	- public/release/models/platform/openconfig-platform-cpu.yang
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
package device

import (
	"github.com/openconfig/ondatra/config/acl"
	"github.com/openconfig/ondatra/config/gnmicollectormetadata"
	"github.com/openconfig/ondatra/config/interfaces"
	"github.com/openconfig/ondatra/config/lacp"
	"github.com/openconfig/ondatra/config/lldp"
	"github.com/openconfig/ondatra/config/localrouting"
	"github.com/openconfig/ondatra/config/networkinstance"
	"github.com/openconfig/ondatra/config/platform"
	"github.com/openconfig/ondatra/config/qos"
	"github.com/openconfig/ondatra/config/routingpolicy"
	"github.com/openconfig/ondatra/config/system"
	"github.com/openconfig/ygot/ygot"
)

// DevicePath represents the /device YANG schema element.
type DevicePath struct {
	*ygot.DeviceRootBase
}

// DeviceRoot returns a new path object from which YANG paths can be constructed.
func DeviceRoot(id string) *DevicePath {
	return &DevicePath{ygot.NewDeviceRootBase(id)}
}

// Acl (container): Top level enclosing container for ACL model config
// and operational state data
// ----------------------------------------
// Defining module: "openconfig-acl"
// Instantiating module: "openconfig-acl"
// Path from parent: "acl"
// Path from root: "/acl"
func (n *DevicePath) Acl() *acl.AclPath {
	return &acl.AclPath{
		NodePath: ygot.NewNodePath(
			[]string{"acl"},
			map[string]interface{}{},
			n,
		),
	}
}

// ComponentAny (list): List of components, keyed by component name.
// ----------------------------------------
// Defining module: "openconfig-platform"
// Instantiating module: "openconfig-platform"
// Path from parent: "components/component"
// Path from root: "/components/component"
// Name (wildcarded): string
func (n *DevicePath) ComponentAny() *platform.ComponentPathAny {
	return &platform.ComponentPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"components", "component"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Component (list): List of components, keyed by component name.
// ----------------------------------------
// Defining module: "openconfig-platform"
// Instantiating module: "openconfig-platform"
// Path from parent: "components/component"
// Path from root: "/components/component"
// Name: string
func (n *DevicePath) Component(Name string) *platform.ComponentPath {
	return &platform.ComponentPath{
		NodePath: ygot.NewNodePath(
			[]string{"components", "component"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// InterfaceAny (list): The list of named interfaces on the device.
// ----------------------------------------
// Defining module: "openconfig-interfaces"
// Instantiating module: "openconfig-interfaces"
// Path from parent: "interfaces/interface"
// Path from root: "/interfaces/interface"
// Name (wildcarded): string
func (n *DevicePath) InterfaceAny() *interfaces.InterfacePathAny {
	return &interfaces.InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Interface (list): The list of named interfaces on the device.
// ----------------------------------------
// Defining module: "openconfig-interfaces"
// Instantiating module: "openconfig-interfaces"
// Path from parent: "interfaces/interface"
// Path from root: "/interfaces/interface"
// Name: string
func (n *DevicePath) Interface(Name string) *interfaces.InterfacePath {
	return &interfaces.InterfacePath{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Lacp (container): Configuration and operational state data for LACP protocol
// operation on the aggregate interface
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "lacp"
// Path from root: "/lacp"
func (n *DevicePath) Lacp() *lacp.LacpPath {
	return &lacp.LacpPath{
		NodePath: ygot.NewNodePath(
			[]string{"lacp"},
			map[string]interface{}{},
			n,
		),
	}
}

// Lldp (container): Top-level container for LLDP configuration and state data
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "lldp"
// Path from root: "/lldp"
func (n *DevicePath) Lldp() *lldp.LldpPath {
	return &lldp.LldpPath{
		NodePath: ygot.NewNodePath(
			[]string{"lldp"},
			map[string]interface{}{},
			n,
		),
	}
}

// LocalRoutes (container): Top-level container for local routes
// ----------------------------------------
// Defining module: "openconfig-local-routing"
// Instantiating module: "openconfig-local-routing"
// Path from parent: "local-routes"
// Path from root: "/local-routes"
func (n *DevicePath) LocalRoutes() *localrouting.LocalRoutesPath {
	return &localrouting.LocalRoutesPath{
		NodePath: ygot.NewNodePath(
			[]string{"local-routes"},
			map[string]interface{}{},
			n,
		),
	}
}

// Meta (container):
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "meta"
// Path from root: "/meta"
func (n *DevicePath) Meta() *gnmicollectormetadata.MetaPath {
	return &gnmicollectormetadata.MetaPath{
		NodePath: ygot.NewNodePath(
			[]string{"meta"},
			map[string]interface{}{},
			n,
		),
	}
}

// NetworkInstanceAny (list): Network instances configured on the local system
// ----------------------------------------
// Defining module: "openconfig-network-instance"
// Instantiating module: "openconfig-network-instance"
// Path from parent: "network-instances/network-instance"
// Path from root: "/network-instances/network-instance"
// Name (wildcarded): string
func (n *DevicePath) NetworkInstanceAny() *networkinstance.NetworkInstancePathAny {
	return &networkinstance.NetworkInstancePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"network-instances", "network-instance"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// NetworkInstance (list): Network instances configured on the local system
// ----------------------------------------
// Defining module: "openconfig-network-instance"
// Instantiating module: "openconfig-network-instance"
// Path from parent: "network-instances/network-instance"
// Path from root: "/network-instances/network-instance"
// Name: string
func (n *DevicePath) NetworkInstance(Name string) *networkinstance.NetworkInstancePath {
	return &networkinstance.NetworkInstancePath{
		NodePath: ygot.NewNodePath(
			[]string{"network-instances", "network-instance"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Qos (container): Top-level container for QoS data
// ----------------------------------------
// Defining module: "openconfig-qos"
// Instantiating module: "openconfig-qos"
// Path from parent: "qos"
// Path from root: "/qos"
func (n *DevicePath) Qos() *qos.QosPath {
	return &qos.QosPath{
		NodePath: ygot.NewNodePath(
			[]string{"qos"},
			map[string]interface{}{},
			n,
		),
	}
}

// RoutingPolicy (container): Top-level container for all routing policy configuration
// ----------------------------------------
// Defining module: "openconfig-routing-policy"
// Instantiating module: "openconfig-routing-policy"
// Path from parent: "routing-policy"
// Path from root: "/routing-policy"
func (n *DevicePath) RoutingPolicy() *routingpolicy.RoutingPolicyPath {
	return &routingpolicy.RoutingPolicyPath{
		NodePath: ygot.NewNodePath(
			[]string{"routing-policy"},
			map[string]interface{}{},
			n,
		),
	}
}

// System (container): Enclosing container for system-related configuration and
// operational state data
// ----------------------------------------
// Defining module: "openconfig-system"
// Instantiating module: "openconfig-system"
// Path from parent: "system"
// Path from root: "/system"
func (n *DevicePath) System() *system.SystemPath {
	return &system.SystemPath{
		NodePath: ygot.NewNodePath(
			[]string{"system"},
			map[string]interface{}{},
			n,
		),
	}
}
