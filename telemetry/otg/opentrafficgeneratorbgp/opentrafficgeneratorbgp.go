/*
Package opentrafficgeneratorbgp is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by /usr/local/google/home/gdennis/go/pkg/mod/github.com/openconfig/ygot@v0.20.1/genutil/names.go
using the following YANG input files:
	- models-yang/models/isis/open-traffic-generator-isis.yang
	- models-yang/models/types/open-traffic-generator-types.yang
	- models-yang/models/flow/open-traffic-generator-flow.yang
	- models-yang/models/discovery/open-traffic-generator-discovery.yang
	- models-yang/models/interface/open-traffic-generator-port.yang
	- models-yang/models/bgp/open-traffic-generator-bgp.yang
Imported modules were sourced from:
	- models-yang/models/...
*/
package opentrafficgeneratorbgp

import (
	"github.com/openconfig/ygot/ygot"
)

// BgpPeerPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer YANG schema element.
type BgpPeerPath struct {
	*ygot.NodePath
}

// BgpPeerPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer YANG schema element.
type BgpPeerPathAny struct {
	*ygot.NodePath
}

// BgpPeer_NamePath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name YANG schema element.
type BgpPeer_NamePath struct {
	*ygot.NodePath
}

// BgpPeer_NamePathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name YANG schema element.
type BgpPeer_NamePathAny struct {
	*ygot.NodePath
}

// BgpPeer_SessionStatePath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state YANG schema element.
type BgpPeer_SessionStatePath struct {
	*ygot.NodePath
}

// BgpPeer_SessionStatePathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state YANG schema element.
type BgpPeer_SessionStatePathAny struct {
	*ygot.NodePath
}

// Counters (container): Counters of an individual BGP peer.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "state/counters"
// Path from root: "/bgp-peers/bgp-peer/state/counters"
func (n *BgpPeerPath) Counters() *BgpPeer_CountersPath {
	return &BgpPeer_CountersPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "counters"},
			map[string]interface{}{},
			n,
		),
	}
}

// Counters (container): Counters of an individual BGP peer.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "state/counters"
// Path from root: "/bgp-peers/bgp-peer/state/counters"
func (n *BgpPeerPathAny) Counters() *BgpPeer_CountersPathAny {
	return &BgpPeer_CountersPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "counters"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): An arbitary name of the BGP peer determined by the ATE
// configuration.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "state/name"
// Path from root: "/bgp-peers/bgp-peer/state/name"
func (n *BgpPeerPath) Name() *BgpPeer_NamePath {
	return &BgpPeer_NamePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): An arbitary name of the BGP peer determined by the ATE
// configuration.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "state/name"
// Path from root: "/bgp-peers/bgp-peer/state/name"
func (n *BgpPeerPathAny) Name() *BgpPeer_NamePathAny {
	return &BgpPeer_NamePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// SessionState (leaf): Operational state of the BGP peer
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "state/session-state"
// Path from root: "/bgp-peers/bgp-peer/state/session-state"
func (n *BgpPeerPath) SessionState() *BgpPeer_SessionStatePath {
	return &BgpPeer_SessionStatePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "session-state"},
			map[string]interface{}{},
			n,
		),
	}
}

// SessionState (leaf): Operational state of the BGP peer
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "state/session-state"
// Path from root: "/bgp-peers/bgp-peer/state/session-state"
func (n *BgpPeerPathAny) SessionState() *BgpPeer_SessionStatePathAny {
	return &BgpPeer_SessionStatePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "session-state"},
			map[string]interface{}{},
			n,
		),
	}
}

// BgpPeer_CountersPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters YANG schema element.
type BgpPeer_CountersPath struct {
	*ygot.NodePath
}

// BgpPeer_CountersPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters YANG schema element.
type BgpPeer_CountersPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_FlapsPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps YANG schema element.
type BgpPeer_Counters_FlapsPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_FlapsPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps YANG schema element.
type BgpPeer_Counters_FlapsPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InKeepalivesPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives YANG schema element.
type BgpPeer_Counters_InKeepalivesPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InKeepalivesPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives YANG schema element.
type BgpPeer_Counters_InKeepalivesPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InNotificationsPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications YANG schema element.
type BgpPeer_Counters_InNotificationsPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InNotificationsPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications YANG schema element.
type BgpPeer_Counters_InNotificationsPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InOpensPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens YANG schema element.
type BgpPeer_Counters_InOpensPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InOpensPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens YANG schema element.
type BgpPeer_Counters_InOpensPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InRouteWithdrawPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw YANG schema element.
type BgpPeer_Counters_InRouteWithdrawPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InRouteWithdrawPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw YANG schema element.
type BgpPeer_Counters_InRouteWithdrawPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InRoutesPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes YANG schema element.
type BgpPeer_Counters_InRoutesPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InRoutesPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes YANG schema element.
type BgpPeer_Counters_InRoutesPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InUpdatesPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates YANG schema element.
type BgpPeer_Counters_InUpdatesPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_InUpdatesPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates YANG schema element.
type BgpPeer_Counters_InUpdatesPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutKeepalivesPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives YANG schema element.
type BgpPeer_Counters_OutKeepalivesPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutKeepalivesPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives YANG schema element.
type BgpPeer_Counters_OutKeepalivesPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutNotificationsPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications YANG schema element.
type BgpPeer_Counters_OutNotificationsPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutNotificationsPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications YANG schema element.
type BgpPeer_Counters_OutNotificationsPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutOpensPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens YANG schema element.
type BgpPeer_Counters_OutOpensPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutOpensPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens YANG schema element.
type BgpPeer_Counters_OutOpensPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutRouteWithdrawPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw YANG schema element.
type BgpPeer_Counters_OutRouteWithdrawPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutRouteWithdrawPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw YANG schema element.
type BgpPeer_Counters_OutRouteWithdrawPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutRoutesPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes YANG schema element.
type BgpPeer_Counters_OutRoutesPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutRoutesPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes YANG schema element.
type BgpPeer_Counters_OutRoutesPathAny struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutUpdatesPath represents the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates YANG schema element.
type BgpPeer_Counters_OutUpdatesPath struct {
	*ygot.NodePath
}

// BgpPeer_Counters_OutUpdatesPathAny represents the wildcard version of the /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates YANG schema element.
type BgpPeer_Counters_OutUpdatesPathAny struct {
	*ygot.NodePath
}

// Flaps (leaf): The total number of times the BGP session went from an
// ESTABLISHED state to an IDLE state.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "flaps"
// Path from root: "/bgp-peers/bgp-peer/state/counters/flaps"
func (n *BgpPeer_CountersPath) Flaps() *BgpPeer_Counters_FlapsPath {
	return &BgpPeer_Counters_FlapsPath{
		NodePath: ygot.NewNodePath(
			[]string{"flaps"},
			map[string]interface{}{},
			n,
		),
	}
}

// Flaps (leaf): The total number of times the BGP session went from an
// ESTABLISHED state to an IDLE state.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "flaps"
// Path from root: "/bgp-peers/bgp-peer/state/counters/flaps"
func (n *BgpPeer_CountersPathAny) Flaps() *BgpPeer_Counters_FlapsPathAny {
	return &BgpPeer_Counters_FlapsPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"flaps"},
			map[string]interface{}{},
			n,
		),
	}
}

// InKeepalives (leaf): The total number of KEEPALIVE messages received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-keepalives"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-keepalives"
func (n *BgpPeer_CountersPath) InKeepalives() *BgpPeer_Counters_InKeepalivesPath {
	return &BgpPeer_Counters_InKeepalivesPath{
		NodePath: ygot.NewNodePath(
			[]string{"in-keepalives"},
			map[string]interface{}{},
			n,
		),
	}
}

// InKeepalives (leaf): The total number of KEEPALIVE messages received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-keepalives"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-keepalives"
func (n *BgpPeer_CountersPathAny) InKeepalives() *BgpPeer_Counters_InKeepalivesPathAny {
	return &BgpPeer_Counters_InKeepalivesPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"in-keepalives"},
			map[string]interface{}{},
			n,
		),
	}
}

// InNotifications (leaf): The total number of NOTIFICATION messages received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-notifications"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-notifications"
func (n *BgpPeer_CountersPath) InNotifications() *BgpPeer_Counters_InNotificationsPath {
	return &BgpPeer_Counters_InNotificationsPath{
		NodePath: ygot.NewNodePath(
			[]string{"in-notifications"},
			map[string]interface{}{},
			n,
		),
	}
}

// InNotifications (leaf): The total number of NOTIFICATION messages received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-notifications"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-notifications"
func (n *BgpPeer_CountersPathAny) InNotifications() *BgpPeer_Counters_InNotificationsPathAny {
	return &BgpPeer_Counters_InNotificationsPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"in-notifications"},
			map[string]interface{}{},
			n,
		),
	}
}

// InOpens (leaf): The total number of OPEN messages received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-opens"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-opens"
func (n *BgpPeer_CountersPath) InOpens() *BgpPeer_Counters_InOpensPath {
	return &BgpPeer_Counters_InOpensPath{
		NodePath: ygot.NewNodePath(
			[]string{"in-opens"},
			map[string]interface{}{},
			n,
		),
	}
}

// InOpens (leaf): The total number of OPEN messages received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-opens"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-opens"
func (n *BgpPeer_CountersPathAny) InOpens() *BgpPeer_Counters_InOpensPathAny {
	return &BgpPeer_Counters_InOpensPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"in-opens"},
			map[string]interface{}{},
			n,
		),
	}
}

// InRouteWithdraw (leaf): The total number of route withdraws received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-route-withdraw"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-route-withdraw"
func (n *BgpPeer_CountersPath) InRouteWithdraw() *BgpPeer_Counters_InRouteWithdrawPath {
	return &BgpPeer_Counters_InRouteWithdrawPath{
		NodePath: ygot.NewNodePath(
			[]string{"in-route-withdraw"},
			map[string]interface{}{},
			n,
		),
	}
}

// InRouteWithdraw (leaf): The total number of route withdraws received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-route-withdraw"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-route-withdraw"
func (n *BgpPeer_CountersPathAny) InRouteWithdraw() *BgpPeer_Counters_InRouteWithdrawPathAny {
	return &BgpPeer_Counters_InRouteWithdrawPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"in-route-withdraw"},
			map[string]interface{}{},
			n,
		),
	}
}

// InRoutes (leaf): The total number of routes received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-routes"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-routes"
func (n *BgpPeer_CountersPath) InRoutes() *BgpPeer_Counters_InRoutesPath {
	return &BgpPeer_Counters_InRoutesPath{
		NodePath: ygot.NewNodePath(
			[]string{"in-routes"},
			map[string]interface{}{},
			n,
		),
	}
}

// InRoutes (leaf): The total number of routes received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-routes"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-routes"
func (n *BgpPeer_CountersPathAny) InRoutes() *BgpPeer_Counters_InRoutesPathAny {
	return &BgpPeer_Counters_InRoutesPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"in-routes"},
			map[string]interface{}{},
			n,
		),
	}
}

// InUpdates (leaf): The total number of UPDATE messages received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-updates"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-updates"
func (n *BgpPeer_CountersPath) InUpdates() *BgpPeer_Counters_InUpdatesPath {
	return &BgpPeer_Counters_InUpdatesPath{
		NodePath: ygot.NewNodePath(
			[]string{"in-updates"},
			map[string]interface{}{},
			n,
		),
	}
}

// InUpdates (leaf): The total number of UPDATE messages received.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "in-updates"
// Path from root: "/bgp-peers/bgp-peer/state/counters/in-updates"
func (n *BgpPeer_CountersPathAny) InUpdates() *BgpPeer_Counters_InUpdatesPathAny {
	return &BgpPeer_Counters_InUpdatesPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"in-updates"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutKeepalives (leaf): The total number of KEEPALIVE messages sent.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-keepalives"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-keepalives"
func (n *BgpPeer_CountersPath) OutKeepalives() *BgpPeer_Counters_OutKeepalivesPath {
	return &BgpPeer_Counters_OutKeepalivesPath{
		NodePath: ygot.NewNodePath(
			[]string{"out-keepalives"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutKeepalives (leaf): The total number of KEEPALIVE messages sent.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-keepalives"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-keepalives"
func (n *BgpPeer_CountersPathAny) OutKeepalives() *BgpPeer_Counters_OutKeepalivesPathAny {
	return &BgpPeer_Counters_OutKeepalivesPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"out-keepalives"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutNotifications (leaf): The total number of NOTIFICATION messages sent.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-notifications"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-notifications"
func (n *BgpPeer_CountersPath) OutNotifications() *BgpPeer_Counters_OutNotificationsPath {
	return &BgpPeer_Counters_OutNotificationsPath{
		NodePath: ygot.NewNodePath(
			[]string{"out-notifications"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutNotifications (leaf): The total number of NOTIFICATION messages sent.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-notifications"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-notifications"
func (n *BgpPeer_CountersPathAny) OutNotifications() *BgpPeer_Counters_OutNotificationsPathAny {
	return &BgpPeer_Counters_OutNotificationsPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"out-notifications"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutOpens (leaf): The total number of OPEN messages sent.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-opens"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-opens"
func (n *BgpPeer_CountersPath) OutOpens() *BgpPeer_Counters_OutOpensPath {
	return &BgpPeer_Counters_OutOpensPath{
		NodePath: ygot.NewNodePath(
			[]string{"out-opens"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutOpens (leaf): The total number of OPEN messages sent.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-opens"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-opens"
func (n *BgpPeer_CountersPathAny) OutOpens() *BgpPeer_Counters_OutOpensPathAny {
	return &BgpPeer_Counters_OutOpensPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"out-opens"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutRouteWithdraw (leaf): The total number of route withdraws sent.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-route-withdraw"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-route-withdraw"
func (n *BgpPeer_CountersPath) OutRouteWithdraw() *BgpPeer_Counters_OutRouteWithdrawPath {
	return &BgpPeer_Counters_OutRouteWithdrawPath{
		NodePath: ygot.NewNodePath(
			[]string{"out-route-withdraw"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutRouteWithdraw (leaf): The total number of route withdraws sent.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-route-withdraw"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-route-withdraw"
func (n *BgpPeer_CountersPathAny) OutRouteWithdraw() *BgpPeer_Counters_OutRouteWithdrawPathAny {
	return &BgpPeer_Counters_OutRouteWithdrawPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"out-route-withdraw"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutRoutes (leaf): The total number of routes advertised.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-routes"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-routes"
func (n *BgpPeer_CountersPath) OutRoutes() *BgpPeer_Counters_OutRoutesPath {
	return &BgpPeer_Counters_OutRoutesPath{
		NodePath: ygot.NewNodePath(
			[]string{"out-routes"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutRoutes (leaf): The total number of routes advertised.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-routes"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-routes"
func (n *BgpPeer_CountersPathAny) OutRoutes() *BgpPeer_Counters_OutRoutesPathAny {
	return &BgpPeer_Counters_OutRoutesPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"out-routes"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutUpdates (leaf): The total number of UPDATE messages sent.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-updates"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-updates"
func (n *BgpPeer_CountersPath) OutUpdates() *BgpPeer_Counters_OutUpdatesPath {
	return &BgpPeer_Counters_OutUpdatesPath{
		NodePath: ygot.NewNodePath(
			[]string{"out-updates"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutUpdates (leaf): The total number of UPDATE messages sent.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "out-updates"
// Path from root: "/bgp-peers/bgp-peer/state/counters/out-updates"
func (n *BgpPeer_CountersPathAny) OutUpdates() *BgpPeer_Counters_OutUpdatesPathAny {
	return &BgpPeer_Counters_OutUpdatesPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"out-updates"},
			map[string]interface{}{},
			n,
		),
	}
}
