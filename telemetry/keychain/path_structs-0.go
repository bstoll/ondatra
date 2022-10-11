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

// KeychainPath represents the /openconfig-keychain/keychains/keychain YANG schema element.
type KeychainPath struct {
	*ygot.NodePath
}

// KeychainPathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain YANG schema element.
type KeychainPathAny struct {
	*ygot.NodePath
}

// Keychain_NamePath represents the /openconfig-keychain/keychains/keychain/state/name YANG schema element.
type Keychain_NamePath struct {
	*ygot.NodePath
}

// Keychain_NamePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/state/name YANG schema element.
type Keychain_NamePathAny struct {
	*ygot.NodePath
}

// Keychain_TolerancePath represents the /openconfig-keychain/keychains/keychain/state/tolerance YANG schema element.
type Keychain_TolerancePath struct {
	*ygot.NodePath
}

// Keychain_TolerancePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/state/tolerance YANG schema element.
type Keychain_TolerancePathAny struct {
	*ygot.NodePath
}

// KeyAny (list): List of configured keys for the keychain.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "keys/key"
// Path from root: "/keychains/keychain/keys/key"
// KeyId (wildcarded): uint64
func (n *KeychainPath) KeyAny() *Keychain_KeyPathAny {
	return &Keychain_KeyPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"keys", "key"},
			map[string]interface{}{"key-id": "*"},
			n,
		),
	}
}

// KeyAny (list): List of configured keys for the keychain.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "keys/key"
// Path from root: "/keychains/keychain/keys/key"
// KeyId (wildcarded): uint64
func (n *KeychainPathAny) KeyAny() *Keychain_KeyPathAny {
	return &Keychain_KeyPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"keys", "key"},
			map[string]interface{}{"key-id": "*"},
			n,
		),
	}
}

// Key (list): List of configured keys for the keychain.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "keys/key"
// Path from root: "/keychains/keychain/keys/key"
// KeyId: uint64
func (n *KeychainPath) Key(KeyId uint64) *Keychain_KeyPath {
	return &Keychain_KeyPath{
		NodePath: ygot.NewNodePath(
			[]string{"keys", "key"},
			map[string]interface{}{"key-id": KeyId},
			n,
		),
	}
}

// Key (list): List of configured keys for the keychain.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "keys/key"
// Path from root: "/keychains/keychain/keys/key"
// KeyId: uint64
func (n *KeychainPathAny) Key(KeyId uint64) *Keychain_KeyPathAny {
	return &Keychain_KeyPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"keys", "key"},
			map[string]interface{}{"key-id": KeyId},
			n,
		),
	}
}

// Name (leaf): Keychain name.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/name"
// Path from root: "/keychains/keychain/state/name"
func (n *KeychainPath) Name() *Keychain_NamePath {
	return &Keychain_NamePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): Keychain name.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/name"
// Path from root: "/keychains/keychain/state/name"
func (n *KeychainPathAny) Name() *Keychain_NamePathAny {
	return &Keychain_NamePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// Tolerance (leaf): Tolerance (overlap time) that a receive key should be accepted.  May be
// expressed as range in seconds, or using the FOREVER value to indicate
// that the key does not expire.  The default value should be 0, i.e., no
// tolerance.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/tolerance"
// Path from root: "/keychains/keychain/state/tolerance"
func (n *KeychainPath) Tolerance() *Keychain_TolerancePath {
	return &Keychain_TolerancePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "tolerance"},
			map[string]interface{}{},
			n,
		),
	}
}

// Tolerance (leaf): Tolerance (overlap time) that a receive key should be accepted.  May be
// expressed as range in seconds, or using the FOREVER value to indicate
// that the key does not expire.  The default value should be 0, i.e., no
// tolerance.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/tolerance"
// Path from root: "/keychains/keychain/state/tolerance"
func (n *KeychainPathAny) Tolerance() *Keychain_TolerancePathAny {
	return &Keychain_TolerancePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "tolerance"},
			map[string]interface{}{},
			n,
		),
	}
}

// Keychain_KeyPath represents the /openconfig-keychain/keychains/keychain/keys/key YANG schema element.
type Keychain_KeyPath struct {
	*ygot.NodePath
}

// Keychain_KeyPathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key YANG schema element.
type Keychain_KeyPathAny struct {
	*ygot.NodePath
}

// Keychain_Key_CryptoAlgorithmPath represents the /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm YANG schema element.
type Keychain_Key_CryptoAlgorithmPath struct {
	*ygot.NodePath
}

// Keychain_Key_CryptoAlgorithmPathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm YANG schema element.
type Keychain_Key_CryptoAlgorithmPathAny struct {
	*ygot.NodePath
}

// Keychain_Key_KeyIdPath represents the /openconfig-keychain/keychains/keychain/keys/key/state/key-id YANG schema element.
type Keychain_Key_KeyIdPath struct {
	*ygot.NodePath
}

// Keychain_Key_KeyIdPathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/state/key-id YANG schema element.
type Keychain_Key_KeyIdPathAny struct {
	*ygot.NodePath
}

// Keychain_Key_SecretKeyPath represents the /openconfig-keychain/keychains/keychain/keys/key/state/secret-key YANG schema element.
type Keychain_Key_SecretKeyPath struct {
	*ygot.NodePath
}

// Keychain_Key_SecretKeyPathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/state/secret-key YANG schema element.
type Keychain_Key_SecretKeyPathAny struct {
	*ygot.NodePath
}

// CryptoAlgorithm (leaf): Cryptographic algorithm associated with the key.  Note that not all cryptographic
// algorithms are available in all contexts (e.g., across different protocols).
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/crypto-algorithm"
// Path from root: "/keychains/keychain/keys/key/state/crypto-algorithm"
func (n *Keychain_KeyPath) CryptoAlgorithm() *Keychain_Key_CryptoAlgorithmPath {
	return &Keychain_Key_CryptoAlgorithmPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "crypto-algorithm"},
			map[string]interface{}{},
			n,
		),
	}
}

// CryptoAlgorithm (leaf): Cryptographic algorithm associated with the key.  Note that not all cryptographic
// algorithms are available in all contexts (e.g., across different protocols).
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/crypto-algorithm"
// Path from root: "/keychains/keychain/keys/key/state/crypto-algorithm"
func (n *Keychain_KeyPathAny) CryptoAlgorithm() *Keychain_Key_CryptoAlgorithmPathAny {
	return &Keychain_Key_CryptoAlgorithmPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "crypto-algorithm"},
			map[string]interface{}{},
			n,
		),
	}
}

// KeyId (leaf): Identifier for the key within the keychain.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/key-id"
// Path from root: "/keychains/keychain/keys/key/state/key-id"
func (n *Keychain_KeyPath) KeyId() *Keychain_Key_KeyIdPath {
	return &Keychain_Key_KeyIdPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "key-id"},
			map[string]interface{}{},
			n,
		),
	}
}

// KeyId (leaf): Identifier for the key within the keychain.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/key-id"
// Path from root: "/keychains/keychain/keys/key/state/key-id"
func (n *Keychain_KeyPathAny) KeyId() *Keychain_Key_KeyIdPathAny {
	return &Keychain_Key_KeyIdPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "key-id"},
			map[string]interface{}{},
			n,
		),
	}
}

// ReceiveLifetime (container): Specify the validity lifetime of the key in the receive direction.
// Some platforms may only support symmetric send and receive lifetimes,
// in which case the receive-lifetime is typically not specified.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "receive-lifetime"
// Path from root: "/keychains/keychain/keys/key/receive-lifetime"
func (n *Keychain_KeyPath) ReceiveLifetime() *Keychain_Key_ReceiveLifetimePath {
	return &Keychain_Key_ReceiveLifetimePath{
		NodePath: ygot.NewNodePath(
			[]string{"receive-lifetime"},
			map[string]interface{}{},
			n,
		),
	}
}

// ReceiveLifetime (container): Specify the validity lifetime of the key in the receive direction.
// Some platforms may only support symmetric send and receive lifetimes,
// in which case the receive-lifetime is typically not specified.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "receive-lifetime"
// Path from root: "/keychains/keychain/keys/key/receive-lifetime"
func (n *Keychain_KeyPathAny) ReceiveLifetime() *Keychain_Key_ReceiveLifetimePathAny {
	return &Keychain_Key_ReceiveLifetimePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"receive-lifetime"},
			map[string]interface{}{},
			n,
		),
	}
}

// SecretKey (leaf): Authentication key supplied as an encrypted value.  The system should store and
// return the key in encrypted form.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/secret-key"
// Path from root: "/keychains/keychain/keys/key/state/secret-key"
func (n *Keychain_KeyPath) SecretKey() *Keychain_Key_SecretKeyPath {
	return &Keychain_Key_SecretKeyPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "secret-key"},
			map[string]interface{}{},
			n,
		),
	}
}

// SecretKey (leaf): Authentication key supplied as an encrypted value.  The system should store and
// return the key in encrypted form.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "state/secret-key"
// Path from root: "/keychains/keychain/keys/key/state/secret-key"
func (n *Keychain_KeyPathAny) SecretKey() *Keychain_Key_SecretKeyPathAny {
	return &Keychain_Key_SecretKeyPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "secret-key"},
			map[string]interface{}{},
			n,
		),
	}
}

// SendLifetime (container): Specifies the lifetime of the key for sending authentication
// information to the peer.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "send-lifetime"
// Path from root: "/keychains/keychain/keys/key/send-lifetime"
func (n *Keychain_KeyPath) SendLifetime() *Keychain_Key_SendLifetimePath {
	return &Keychain_Key_SendLifetimePath{
		NodePath: ygot.NewNodePath(
			[]string{"send-lifetime"},
			map[string]interface{}{},
			n,
		),
	}
}

// SendLifetime (container): Specifies the lifetime of the key for sending authentication
// information to the peer.
// ----------------------------------------
// Defining module: "openconfig-keychain"
// Instantiating module: "openconfig-keychain"
// Path from parent: "send-lifetime"
// Path from root: "/keychains/keychain/keys/key/send-lifetime"
func (n *Keychain_KeyPathAny) SendLifetime() *Keychain_Key_SendLifetimePathAny {
	return &Keychain_Key_SendLifetimePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"send-lifetime"},
			map[string]interface{}{},
			n,
		),
	}
}
