/*
Package keychain is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by /usr/local/google/home/gdennis/go/pkg/mod/github.com/openconfig/ygot@v0.25.2/genutil/names.go
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
  - public/release/models/optical-transport/openconfig-terminal-device.yang
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
package keychain

import (
	"github.com/openconfig/ygot/ygot"
)

// Keychain_Key_ReceiveLifetimePath represents the /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime YANG schema element.
type Keychain_Key_ReceiveLifetimePath struct {
	*ygot.NodePath
}

// Keychain_Key_ReceiveLifetimePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime YANG schema element.
type Keychain_Key_ReceiveLifetimePathAny struct {
	*ygot.NodePath
}

// Keychain_Key_ReceiveLifetime_EndTimePath represents the /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time YANG schema element.
type Keychain_Key_ReceiveLifetime_EndTimePath struct {
	*ygot.NodePath
}

// Keychain_Key_ReceiveLifetime_EndTimePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time YANG schema element.
type Keychain_Key_ReceiveLifetime_EndTimePathAny struct {
	*ygot.NodePath
}

// Keychain_Key_ReceiveLifetime_StartTimePath represents the /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/start-time YANG schema element.
type Keychain_Key_ReceiveLifetime_StartTimePath struct {
	*ygot.NodePath
}

// Keychain_Key_ReceiveLifetime_StartTimePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/start-time YANG schema element.
type Keychain_Key_ReceiveLifetime_StartTimePathAny struct {
	*ygot.NodePath
}

// EndTime (leaf): The time at which the key becomes invalid for use.
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
//
// Leaving this value unset, or setting it to 0, indicates that
// the key remains valid forever (no end time).
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/end-time"
// Path from root: "/keychains/keychain/keys/key/receive-lifetime/state/end-time"
func (n *Keychain_Key_ReceiveLifetimePath) EndTime() *Keychain_Key_ReceiveLifetime_EndTimePath {
	return &Keychain_Key_ReceiveLifetime_EndTimePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "end-time"},
			map[string]interface{}{},
			n,
		),
	}
}

// EndTime (leaf): The time at which the key becomes invalid for use.
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
//
// Leaving this value unset, or setting it to 0, indicates that
// the key remains valid forever (no end time).
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/end-time"
// Path from root: "/keychains/keychain/keys/key/receive-lifetime/state/end-time"
func (n *Keychain_Key_ReceiveLifetimePathAny) EndTime() *Keychain_Key_ReceiveLifetime_EndTimePathAny {
	return &Keychain_Key_ReceiveLifetime_EndTimePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "end-time"},
			map[string]interface{}{},
			n,
		),
	}
}

// StartTime (leaf): The time at which the key becomes valid for use.
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/start-time"
// Path from root: "/keychains/keychain/keys/key/receive-lifetime/state/start-time"
func (n *Keychain_Key_ReceiveLifetimePath) StartTime() *Keychain_Key_ReceiveLifetime_StartTimePath {
	return &Keychain_Key_ReceiveLifetime_StartTimePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "start-time"},
			map[string]interface{}{},
			n,
		),
	}
}

// StartTime (leaf): The time at which the key becomes valid for use.
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/start-time"
// Path from root: "/keychains/keychain/keys/key/receive-lifetime/state/start-time"
func (n *Keychain_Key_ReceiveLifetimePathAny) StartTime() *Keychain_Key_ReceiveLifetime_StartTimePathAny {
	return &Keychain_Key_ReceiveLifetime_StartTimePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "start-time"},
			map[string]interface{}{},
			n,
		),
	}
}

// Keychain_Key_SendLifetimePath represents the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime YANG schema element.
type Keychain_Key_SendLifetimePath struct {
	*ygot.NodePath
}

// Keychain_Key_SendLifetimePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime YANG schema element.
type Keychain_Key_SendLifetimePathAny struct {
	*ygot.NodePath
}

// Keychain_Key_SendLifetime_EndTimePath represents the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time YANG schema element.
type Keychain_Key_SendLifetime_EndTimePath struct {
	*ygot.NodePath
}

// Keychain_Key_SendLifetime_EndTimePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time YANG schema element.
type Keychain_Key_SendLifetime_EndTimePathAny struct {
	*ygot.NodePath
}

// Keychain_Key_SendLifetime_SendAndReceivePath represents the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive YANG schema element.
type Keychain_Key_SendLifetime_SendAndReceivePath struct {
	*ygot.NodePath
}

// Keychain_Key_SendLifetime_SendAndReceivePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive YANG schema element.
type Keychain_Key_SendLifetime_SendAndReceivePathAny struct {
	*ygot.NodePath
}

// Keychain_Key_SendLifetime_StartTimePath represents the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time YANG schema element.
type Keychain_Key_SendLifetime_StartTimePath struct {
	*ygot.NodePath
}

// Keychain_Key_SendLifetime_StartTimePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time YANG schema element.
type Keychain_Key_SendLifetime_StartTimePathAny struct {
	*ygot.NodePath
}

// EndTime (leaf): The time at which the key becomes invalid for use.
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
//
// Leaving this value unset, or setting it to 0, indicates that
// the key remains valid forever (no end time).
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/end-time"
// Path from root: "/keychains/keychain/keys/key/send-lifetime/state/end-time"
func (n *Keychain_Key_SendLifetimePath) EndTime() *Keychain_Key_SendLifetime_EndTimePath {
	return &Keychain_Key_SendLifetime_EndTimePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "end-time"},
			map[string]interface{}{},
			n,
		),
	}
}

// EndTime (leaf): The time at which the key becomes invalid for use.
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
//
// Leaving this value unset, or setting it to 0, indicates that
// the key remains valid forever (no end time).
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/end-time"
// Path from root: "/keychains/keychain/keys/key/send-lifetime/state/end-time"
func (n *Keychain_Key_SendLifetimePathAny) EndTime() *Keychain_Key_SendLifetime_EndTimePathAny {
	return &Keychain_Key_SendLifetime_EndTimePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "end-time"},
			map[string]interface{}{},
			n,
		),
	}
}

// SendAndReceive (leaf): When this is set to true (the default value), the specified
// send lifetime is also used in the receive direction.  When set
// to false, the device should use the specified receive-lifetime
// for the receive direction (asymmetric mode).  If send-and-receive
// is false, and the device does not support asymmetric configuration,
// the config should be rejected as unsupported.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/send-and-receive"
// Path from root: "/keychains/keychain/keys/key/send-lifetime/state/send-and-receive"
func (n *Keychain_Key_SendLifetimePath) SendAndReceive() *Keychain_Key_SendLifetime_SendAndReceivePath {
	return &Keychain_Key_SendLifetime_SendAndReceivePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "send-and-receive"},
			map[string]interface{}{},
			n,
		),
	}
}

// SendAndReceive (leaf): When this is set to true (the default value), the specified
// send lifetime is also used in the receive direction.  When set
// to false, the device should use the specified receive-lifetime
// for the receive direction (asymmetric mode).  If send-and-receive
// is false, and the device does not support asymmetric configuration,
// the config should be rejected as unsupported.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/send-and-receive"
// Path from root: "/keychains/keychain/keys/key/send-lifetime/state/send-and-receive"
func (n *Keychain_Key_SendLifetimePathAny) SendAndReceive() *Keychain_Key_SendLifetime_SendAndReceivePathAny {
	return &Keychain_Key_SendLifetime_SendAndReceivePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "send-and-receive"},
			map[string]interface{}{},
			n,
		),
	}
}

// StartTime (leaf): The time at which the key becomes valid for use.
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/start-time"
// Path from root: "/keychains/keychain/keys/key/send-lifetime/state/start-time"
func (n *Keychain_Key_SendLifetimePath) StartTime() *Keychain_Key_SendLifetime_StartTimePath {
	return &Keychain_Key_SendLifetime_StartTimePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "start-time"},
			map[string]interface{}{},
			n,
		),
	}
}

// StartTime (leaf): The time at which the key becomes valid for use.
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/start-time"
// Path from root: "/keychains/keychain/keys/key/send-lifetime/state/start-time"
func (n *Keychain_Key_SendLifetimePathAny) StartTime() *Keychain_Key_SendLifetime_StartTimePathAny {
	return &Keychain_Key_SendLifetime_StartTimePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "start-time"},
			map[string]interface{}{},
			n,
		),
	}
}
