/*
Package gnmicollectormetadata is a generated package which contains definitions
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
  - public/release/models/interfaces/openconfig-if-sdn-ext.yang
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
package gnmicollectormetadata

import (
	"github.com/openconfig/ygot/ygot"
)

// MetaPath represents the /gnmi-collector-metadata/meta YANG schema element.
type MetaPath struct {
	*ygot.NodePath
}

// MetaPathAny represents the wildcard version of the /gnmi-collector-metadata/meta YANG schema element.
type MetaPathAny struct {
	*ygot.NodePath
}

// Meta_ConnectErrorPath represents the /gnmi-collector-metadata/meta/connectError YANG schema element.
type Meta_ConnectErrorPath struct {
	*ygot.NodePath
}

// Meta_ConnectErrorPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/connectError YANG schema element.
type Meta_ConnectErrorPathAny struct {
	*ygot.NodePath
}

// Meta_ConnectedPath represents the /gnmi-collector-metadata/meta/connected YANG schema element.
type Meta_ConnectedPath struct {
	*ygot.NodePath
}

// Meta_ConnectedPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/connected YANG schema element.
type Meta_ConnectedPathAny struct {
	*ygot.NodePath
}

// Meta_ConnectedAddressPath represents the /gnmi-collector-metadata/meta/connectedAddress YANG schema element.
type Meta_ConnectedAddressPath struct {
	*ygot.NodePath
}

// Meta_ConnectedAddressPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/connectedAddress YANG schema element.
type Meta_ConnectedAddressPathAny struct {
	*ygot.NodePath
}

// Meta_LatencyAvgPath represents the /gnmi-collector-metadata/meta/latencyAvg YANG schema element.
type Meta_LatencyAvgPath struct {
	*ygot.NodePath
}

// Meta_LatencyAvgPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/latencyAvg YANG schema element.
type Meta_LatencyAvgPathAny struct {
	*ygot.NodePath
}

// Meta_LatencyMaxPath represents the /gnmi-collector-metadata/meta/latencyMax YANG schema element.
type Meta_LatencyMaxPath struct {
	*ygot.NodePath
}

// Meta_LatencyMaxPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/latencyMax YANG schema element.
type Meta_LatencyMaxPathAny struct {
	*ygot.NodePath
}

// Meta_LatencyMinPath represents the /gnmi-collector-metadata/meta/latencyMin YANG schema element.
type Meta_LatencyMinPath struct {
	*ygot.NodePath
}

// Meta_LatencyMinPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/latencyMin YANG schema element.
type Meta_LatencyMinPathAny struct {
	*ygot.NodePath
}

// Meta_LatestTimestampPath represents the /gnmi-collector-metadata/meta/latestTimestamp YANG schema element.
type Meta_LatestTimestampPath struct {
	*ygot.NodePath
}

// Meta_LatestTimestampPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/latestTimestamp YANG schema element.
type Meta_LatestTimestampPathAny struct {
	*ygot.NodePath
}

// Meta_SyncPath represents the /gnmi-collector-metadata/meta/sync YANG schema element.
type Meta_SyncPath struct {
	*ygot.NodePath
}

// Meta_SyncPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/sync YANG schema element.
type Meta_SyncPathAny struct {
	*ygot.NodePath
}

// Meta_TargetLeavesPath represents the /gnmi-collector-metadata/meta/targetLeaves YANG schema element.
type Meta_TargetLeavesPath struct {
	*ygot.NodePath
}

// Meta_TargetLeavesPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/targetLeaves YANG schema element.
type Meta_TargetLeavesPathAny struct {
	*ygot.NodePath
}

// Meta_TargetLeavesAddedPath represents the /gnmi-collector-metadata/meta/targetLeavesAdded YANG schema element.
type Meta_TargetLeavesAddedPath struct {
	*ygot.NodePath
}

// Meta_TargetLeavesAddedPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/targetLeavesAdded YANG schema element.
type Meta_TargetLeavesAddedPathAny struct {
	*ygot.NodePath
}

// Meta_TargetLeavesDeletedPath represents the /gnmi-collector-metadata/meta/targetLeavesDeleted YANG schema element.
type Meta_TargetLeavesDeletedPath struct {
	*ygot.NodePath
}

// Meta_TargetLeavesDeletedPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/targetLeavesDeleted YANG schema element.
type Meta_TargetLeavesDeletedPathAny struct {
	*ygot.NodePath
}

// Meta_TargetLeavesEmptyPath represents the /gnmi-collector-metadata/meta/targetLeavesEmpty YANG schema element.
type Meta_TargetLeavesEmptyPath struct {
	*ygot.NodePath
}

// Meta_TargetLeavesEmptyPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/targetLeavesEmpty YANG schema element.
type Meta_TargetLeavesEmptyPathAny struct {
	*ygot.NodePath
}

// Meta_TargetLeavesStalePath represents the /gnmi-collector-metadata/meta/targetLeavesStale YANG schema element.
type Meta_TargetLeavesStalePath struct {
	*ygot.NodePath
}

// Meta_TargetLeavesStalePathAny represents the wildcard version of the /gnmi-collector-metadata/meta/targetLeavesStale YANG schema element.
type Meta_TargetLeavesStalePathAny struct {
	*ygot.NodePath
}

// Meta_TargetLeavesSuppressedPath represents the /gnmi-collector-metadata/meta/targetLeavesSuppressed YANG schema element.
type Meta_TargetLeavesSuppressedPath struct {
	*ygot.NodePath
}

// Meta_TargetLeavesSuppressedPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/targetLeavesSuppressed YANG schema element.
type Meta_TargetLeavesSuppressedPathAny struct {
	*ygot.NodePath
}

// Meta_TargetLeavesUpdatedPath represents the /gnmi-collector-metadata/meta/targetLeavesUpdated YANG schema element.
type Meta_TargetLeavesUpdatedPath struct {
	*ygot.NodePath
}

// Meta_TargetLeavesUpdatedPathAny represents the wildcard version of the /gnmi-collector-metadata/meta/targetLeavesUpdated YANG schema element.
type Meta_TargetLeavesUpdatedPathAny struct {
	*ygot.NodePath
}

// Meta_TargetSizePath represents the /gnmi-collector-metadata/meta/targetSize YANG schema element.
type Meta_TargetSizePath struct {
	*ygot.NodePath
}

// Meta_TargetSizePathAny represents the wildcard version of the /gnmi-collector-metadata/meta/targetSize YANG schema element.
type Meta_TargetSizePathAny struct {
	*ygot.NodePath
}

// ConnectError (leaf): connectError is the error related to connection failure.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "connectError"
// Path from root: "/meta/connectError"
func (n *MetaPath) ConnectError() *Meta_ConnectErrorPath {
	return &Meta_ConnectErrorPath{
		NodePath: ygot.NewNodePath(
			[]string{"connectError"},
			map[string]interface{}{},
			n,
		),
	}
}

// ConnectError (leaf): connectError is the error related to connection failure.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "connectError"
// Path from root: "/meta/connectError"
func (n *MetaPathAny) ConnectError() *Meta_ConnectErrorPathAny {
	return &Meta_ConnectErrorPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"connectError"},
			map[string]interface{}{},
			n,
		),
	}
}

// Connected (leaf): connected reports whether the client has an active gRPC session with
// the target device; it requires at least 1 update delivered over the
// connection before being set to true.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "connected"
// Path from root: "/meta/connected"
func (n *MetaPath) Connected() *Meta_ConnectedPath {
	return &Meta_ConnectedPath{
		NodePath: ygot.NewNodePath(
			[]string{"connected"},
			map[string]interface{}{},
			n,
		),
	}
}

// Connected (leaf): connected reports whether the client has an active gRPC session with
// the target device; it requires at least 1 update delivered over the
// connection before being set to true.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "connected"
// Path from root: "/meta/connected"
func (n *MetaPathAny) Connected() *Meta_ConnectedPathAny {
	return &Meta_ConnectedPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"connected"},
			map[string]interface{}{},
			n,
		),
	}
}

// ConnectedAddress (leaf): connectedAddress denotes the last-hop IP address of a connected target
// in IP:Port format (e.g., '10.1.1.1:12345',
// '[123:123:123:123::1]:12345').
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "connectedAddress"
// Path from root: "/meta/connectedAddress"
func (n *MetaPath) ConnectedAddress() *Meta_ConnectedAddressPath {
	return &Meta_ConnectedAddressPath{
		NodePath: ygot.NewNodePath(
			[]string{"connectedAddress"},
			map[string]interface{}{},
			n,
		),
	}
}

// ConnectedAddress (leaf): connectedAddress denotes the last-hop IP address of a connected target
// in IP:Port format (e.g., '10.1.1.1:12345',
// '[123:123:123:123::1]:12345').
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "connectedAddress"
// Path from root: "/meta/connectedAddress"
func (n *MetaPathAny) ConnectedAddress() *Meta_ConnectedAddressPathAny {
	return &Meta_ConnectedAddressPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"connectedAddress"},
			map[string]interface{}{},
			n,
		),
	}
}

// LatencyAvg (leaf): latencyAvg is the average latency in nanoseconds between target
// timestamp and cache reception - latency being calculated by (timestamp
// of arrival) - (timestamp in update). It is reported per update
// window based on the commandline flag for metadata updates.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latencyAvg"
// Path from root: "/meta/latencyAvg"
func (n *MetaPath) LatencyAvg() *Meta_LatencyAvgPath {
	return &Meta_LatencyAvgPath{
		NodePath: ygot.NewNodePath(
			[]string{"latencyAvg"},
			map[string]interface{}{},
			n,
		),
	}
}

// LatencyAvg (leaf): latencyAvg is the average latency in nanoseconds between target
// timestamp and cache reception - latency being calculated by (timestamp
// of arrival) - (timestamp in update). It is reported per update
// window based on the commandline flag for metadata updates.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latencyAvg"
// Path from root: "/meta/latencyAvg"
func (n *MetaPathAny) LatencyAvg() *Meta_LatencyAvgPathAny {
	return &Meta_LatencyAvgPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"latencyAvg"},
			map[string]interface{}{},
			n,
		),
	}
}

// LatencyMax (leaf): latencyMax is the maximum latency in nanoseconds between target
// timestamp and cache reception - latency being calculated by (timestamp
// of arrival) - (timestamp in update). It is reported per update
// window based on the commandline flag for metadata updates.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latencyMax"
// Path from root: "/meta/latencyMax"
func (n *MetaPath) LatencyMax() *Meta_LatencyMaxPath {
	return &Meta_LatencyMaxPath{
		NodePath: ygot.NewNodePath(
			[]string{"latencyMax"},
			map[string]interface{}{},
			n,
		),
	}
}

// LatencyMax (leaf): latencyMax is the maximum latency in nanoseconds between target
// timestamp and cache reception - latency being calculated by (timestamp
// of arrival) - (timestamp in update). It is reported per update
// window based on the commandline flag for metadata updates.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latencyMax"
// Path from root: "/meta/latencyMax"
func (n *MetaPathAny) LatencyMax() *Meta_LatencyMaxPathAny {
	return &Meta_LatencyMaxPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"latencyMax"},
			map[string]interface{}{},
			n,
		),
	}
}

// LatencyMin (leaf): latencyMin is the minimum latency in nanoseconds - latency being
// calculated by (timestamp of arrival) - (timestamp in update). It is
// reported per update window based on the commandline flag for metadata
// updates.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latencyMin"
// Path from root: "/meta/latencyMin"
func (n *MetaPath) LatencyMin() *Meta_LatencyMinPath {
	return &Meta_LatencyMinPath{
		NodePath: ygot.NewNodePath(
			[]string{"latencyMin"},
			map[string]interface{}{},
			n,
		),
	}
}

// LatencyMin (leaf): latencyMin is the minimum latency in nanoseconds - latency being
// calculated by (timestamp of arrival) - (timestamp in update). It is
// reported per update window based on the commandline flag for metadata
// updates.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latencyMin"
// Path from root: "/meta/latencyMin"
func (n *MetaPathAny) LatencyMin() *Meta_LatencyMinPathAny {
	return &Meta_LatencyMinPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"latencyMin"},
			map[string]interface{}{},
			n,
		),
	}
}

// LatestTimestamp (leaf): latestTimestamp is the latest timestamp in nanoseconds since Epoch time
// of the latest update received from the target. This value is updated
// periodically so it may lag behind the actual target updates.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latestTimestamp"
// Path from root: "/meta/latestTimestamp"
func (n *MetaPath) LatestTimestamp() *Meta_LatestTimestampPath {
	return &Meta_LatestTimestampPath{
		NodePath: ygot.NewNodePath(
			[]string{"latestTimestamp"},
			map[string]interface{}{},
			n,
		),
	}
}

// LatestTimestamp (leaf): latestTimestamp is the latest timestamp in nanoseconds since Epoch time
// of the latest update received from the target. This value is updated
// periodically so it may lag behind the actual target updates.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latestTimestamp"
// Path from root: "/meta/latestTimestamp"
func (n *MetaPathAny) LatestTimestamp() *Meta_LatestTimestampPathAny {
	return &Meta_LatestTimestampPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"latestTimestamp"},
			map[string]interface{}{},
			n,
		),
	}
}

// Sync (leaf): sync indicates that at least one copy of the target's entire tree has
// been received, as indicated by the sync_response field in a gNMI
// SubscribeResponse message
// (https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-specification.md#3514-the-subscriberesponse-message).
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "sync"
// Path from root: "/meta/sync"
func (n *MetaPath) Sync() *Meta_SyncPath {
	return &Meta_SyncPath{
		NodePath: ygot.NewNodePath(
			[]string{"sync"},
			map[string]interface{}{},
			n,
		),
	}
}

// Sync (leaf): sync indicates that at least one copy of the target's entire tree has
// been received, as indicated by the sync_response field in a gNMI
// SubscribeResponse message
// (https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-specification.md#3514-the-subscriberesponse-message).
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "sync"
// Path from root: "/meta/sync"
func (n *MetaPathAny) Sync() *Meta_SyncPathAny {
	return &Meta_SyncPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"sync"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeaves (leaf): targetLeaves is the total number of leaves available for the target.
// Note that this does not include any intermediate nodes.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeaves"
// Path from root: "/meta/targetLeaves"
func (n *MetaPath) TargetLeaves() *Meta_TargetLeavesPath {
	return &Meta_TargetLeavesPath{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeaves"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeaves (leaf): targetLeaves is the total number of leaves available for the target.
// Note that this does not include any intermediate nodes.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeaves"
// Path from root: "/meta/targetLeaves"
func (n *MetaPathAny) TargetLeaves() *Meta_TargetLeavesPathAny {
	return &Meta_TargetLeavesPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeaves"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesAdded (leaf): targetLeavesAdded is the total number of leaves that have been added.
// This number may be larger than meta/targetLeaves due to deletes.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesAdded"
// Path from root: "/meta/targetLeavesAdded"
func (n *MetaPath) TargetLeavesAdded() *Meta_TargetLeavesAddedPath {
	return &Meta_TargetLeavesAddedPath{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesAdded"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesAdded (leaf): targetLeavesAdded is the total number of leaves that have been added.
// This number may be larger than meta/targetLeaves due to deletes.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesAdded"
// Path from root: "/meta/targetLeavesAdded"
func (n *MetaPathAny) TargetLeavesAdded() *Meta_TargetLeavesAddedPathAny {
	return &Meta_TargetLeavesAddedPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesAdded"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesDeleted (leaf): targetLeavesDeleted is the total number of leaves that have been
// deleted.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesDeleted"
// Path from root: "/meta/targetLeavesDeleted"
func (n *MetaPath) TargetLeavesDeleted() *Meta_TargetLeavesDeletedPath {
	return &Meta_TargetLeavesDeletedPath{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesDeleted"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesDeleted (leaf): targetLeavesDeleted is the total number of leaves that have been
// deleted.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesDeleted"
// Path from root: "/meta/targetLeavesDeleted"
func (n *MetaPathAny) TargetLeavesDeleted() *Meta_TargetLeavesDeletedPathAny {
	return &Meta_TargetLeavesDeletedPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesDeleted"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesEmpty (leaf): targetLeavesEmpty is the total number of empty notifications received
// from the target.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesEmpty"
// Path from root: "/meta/targetLeavesEmpty"
func (n *MetaPath) TargetLeavesEmpty() *Meta_TargetLeavesEmptyPath {
	return &Meta_TargetLeavesEmptyPath{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesEmpty"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesEmpty (leaf): targetLeavesEmpty is the total number of empty notifications received
// from the target.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesEmpty"
// Path from root: "/meta/targetLeavesEmpty"
func (n *MetaPathAny) TargetLeavesEmpty() *Meta_TargetLeavesEmptyPathAny {
	return &Meta_TargetLeavesEmptyPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesEmpty"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesStale (leaf): targetLeavesStale is the total number of leaf updates that were received
// with a timestamp older than the latest timestamp reported for the target.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesStale"
// Path from root: "/meta/targetLeavesStale"
func (n *MetaPath) TargetLeavesStale() *Meta_TargetLeavesStalePath {
	return &Meta_TargetLeavesStalePath{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesStale"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesStale (leaf): targetLeavesStale is the total number of leaf updates that were received
// with a timestamp older than the latest timestamp reported for the target.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesStale"
// Path from root: "/meta/targetLeavesStale"
func (n *MetaPathAny) TargetLeavesStale() *Meta_TargetLeavesStalePathAny {
	return &Meta_TargetLeavesStalePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesStale"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesSuppressed (leaf): targetLeavesSuppressed is the total number of leaf updates that were not
// forwarded to subscribers because the value had not changed.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesSuppressed"
// Path from root: "/meta/targetLeavesSuppressed"
func (n *MetaPath) TargetLeavesSuppressed() *Meta_TargetLeavesSuppressedPath {
	return &Meta_TargetLeavesSuppressedPath{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesSuppressed"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesSuppressed (leaf): targetLeavesSuppressed is the total number of leaf updates that were not
// forwarded to subscribers because the value had not changed.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesSuppressed"
// Path from root: "/meta/targetLeavesSuppressed"
func (n *MetaPathAny) TargetLeavesSuppressed() *Meta_TargetLeavesSuppressedPathAny {
	return &Meta_TargetLeavesSuppressedPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesSuppressed"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesUpdated (leaf): targetLeavesUpdated is the total number of leaf updates that have been
// received from the target by the collector
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesUpdated"
// Path from root: "/meta/targetLeavesUpdated"
func (n *MetaPath) TargetLeavesUpdated() *Meta_TargetLeavesUpdatedPath {
	return &Meta_TargetLeavesUpdatedPath{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesUpdated"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetLeavesUpdated (leaf): targetLeavesUpdated is the total number of leaf updates that have been
// received from the target by the collector
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetLeavesUpdated"
// Path from root: "/meta/targetLeavesUpdated"
func (n *MetaPathAny) TargetLeavesUpdated() *Meta_TargetLeavesUpdatedPathAny {
	return &Meta_TargetLeavesUpdatedPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"targetLeavesUpdated"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetSize (leaf): targetSize is the total number of bytes used to store all values. This
// count excludes all indexing overhead. This value is updated periodically
// and may not be up to date at all times.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetSize"
// Path from root: "/meta/targetSize"
func (n *MetaPath) TargetSize() *Meta_TargetSizePath {
	return &Meta_TargetSizePath{
		NodePath: ygot.NewNodePath(
			[]string{"targetSize"},
			map[string]interface{}{},
			n,
		),
	}
}

// TargetSize (leaf): targetSize is the total number of bytes used to store all values. This
// count excludes all indexing overhead. This value is updated periodically
// and may not be up to date at all times.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "targetSize"
// Path from root: "/meta/targetSize"
func (n *MetaPathAny) TargetSize() *Meta_TargetSizePathAny {
	return &Meta_TargetSizePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"targetSize"},
			map[string]interface{}{},
			n,
		),
	}
}

// WindowAny (list): latency statistics for a time window.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latency/window"
// Path from root: "/meta/latency/window"
// Size (wildcarded): string
func (n *MetaPath) WindowAny() *Meta_WindowPathAny {
	return &Meta_WindowPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"latency", "window"},
			map[string]interface{}{"size": "*"},
			n,
		),
	}
}

// WindowAny (list): latency statistics for a time window.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latency/window"
// Path from root: "/meta/latency/window"
// Size (wildcarded): string
func (n *MetaPathAny) WindowAny() *Meta_WindowPathAny {
	return &Meta_WindowPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"latency", "window"},
			map[string]interface{}{"size": "*"},
			n,
		),
	}
}

// Window (list): latency statistics for a time window.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latency/window"
// Path from root: "/meta/latency/window"
// Size: string
func (n *MetaPath) Window(Size string) *Meta_WindowPath {
	return &Meta_WindowPath{
		NodePath: ygot.NewNodePath(
			[]string{"latency", "window"},
			map[string]interface{}{"size": Size},
			n,
		),
	}
}

// Window (list): latency statistics for a time window.
// ----------------------------------------
// Defining module: "gnmi-collector-metadata"
// Instantiating module: "gnmi-collector-metadata"
// Path from parent: "latency/window"
// Path from root: "/meta/latency/window"
// Size: string
func (n *MetaPathAny) Window(Size string) *Meta_WindowPathAny {
	return &Meta_WindowPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"latency", "window"},
			map[string]interface{}{"size": Size},
			n,
		),
	}
}
