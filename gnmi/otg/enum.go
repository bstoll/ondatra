/*
Package otg is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was true
in this case).

This package was generated by ygnmi version: v0.7.1: (ygot: v0.25.4)
using the following YANG input files:
  - models-yang/models/bgp/open-traffic-generator-bgp.yang
  - models-yang/models/discovery/open-traffic-generator-discovery.yang
  - models-yang/models/flow/open-traffic-generator-flow.yang
  - models-yang/models/interface/open-traffic-generator-port.yang
  - models-yang/models/isis/open-traffic-generator-isis.yang
  - models-yang/models/lacp/open-traffic-generator-lacp.yang
  - models-yang/models/lag/open-traffic-generator-lag.yang
  - models-yang/models/lldp/open-traffic-generator-lldp.yang
  - models-yang/models/types/open-traffic-generator-types.yang

Imported modules were sourced from:
  - models-yang/models/...
*/
package otg

import (
	"github.com/openconfig/ygot/ygot"
)

// E_BgpPeer_SessionState is a derived int64 type which is used to represent
// the enumerated node BgpPeer_SessionState. An additional value named
// BgpPeer_SessionState_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_BgpPeer_SessionState int64

// IsYANGGoEnum ensures that BgpPeer_SessionState implements the yang.GoEnum
// interface. This ensures that BgpPeer_SessionState can be identified as a
// mapped type for a YANG enumeration.
func (E_BgpPeer_SessionState) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  BgpPeer_SessionState.
func (E_BgpPeer_SessionState) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_BgpPeer_SessionState.
func (e E_BgpPeer_SessionState) String() string {
	return ygot.EnumLogString(e, int64(e), "E_BgpPeer_SessionState")
}

const (
	// BgpPeer_SessionState_UNSET corresponds to the value UNSET of BgpPeer_SessionState
	BgpPeer_SessionState_UNSET E_BgpPeer_SessionState = 0
	// BgpPeer_SessionState_IDLE corresponds to the value IDLE of BgpPeer_SessionState
	BgpPeer_SessionState_IDLE E_BgpPeer_SessionState = 1
	// BgpPeer_SessionState_CONNECT corresponds to the value CONNECT of BgpPeer_SessionState
	BgpPeer_SessionState_CONNECT E_BgpPeer_SessionState = 2
	// BgpPeer_SessionState_ACTIVE corresponds to the value ACTIVE of BgpPeer_SessionState
	BgpPeer_SessionState_ACTIVE E_BgpPeer_SessionState = 3
	// BgpPeer_SessionState_OPEN_SENT corresponds to the value OPEN_SENT of BgpPeer_SessionState
	BgpPeer_SessionState_OPEN_SENT E_BgpPeer_SessionState = 4
	// BgpPeer_SessionState_OPEN_CONFIRM corresponds to the value OPEN_CONFIRM of BgpPeer_SessionState
	BgpPeer_SessionState_OPEN_CONFIRM E_BgpPeer_SessionState = 5
	// BgpPeer_SessionState_ESTABLISHED corresponds to the value ESTABLISHED of BgpPeer_SessionState
	BgpPeer_SessionState_ESTABLISHED E_BgpPeer_SessionState = 6
)

// E_Capabilities_Name is a derived int64 type which is used to represent
// the enumerated node Capabilities_Name. An additional value named
// Capabilities_Name_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Capabilities_Name int64

// IsYANGGoEnum ensures that Capabilities_Name implements the yang.GoEnum
// interface. This ensures that Capabilities_Name can be identified as a
// mapped type for a YANG enumeration.
func (E_Capabilities_Name) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Capabilities_Name.
func (E_Capabilities_Name) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_Capabilities_Name.
func (e E_Capabilities_Name) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Capabilities_Name")
}

const (
	// Capabilities_Name_UNSET corresponds to the value UNSET of Capabilities_Name
	Capabilities_Name_UNSET E_Capabilities_Name = 0
	// Capabilities_Name_MAC_BRIDGE corresponds to the value MAC_BRIDGE of Capabilities_Name
	Capabilities_Name_MAC_BRIDGE E_Capabilities_Name = 1
	// Capabilities_Name_TWO_PORT_MAC_RELAY corresponds to the value TWO_PORT_MAC_RELAY of Capabilities_Name
	Capabilities_Name_TWO_PORT_MAC_RELAY E_Capabilities_Name = 2
	// Capabilities_Name_REPEATER corresponds to the value REPEATER of Capabilities_Name
	Capabilities_Name_REPEATER E_Capabilities_Name = 3
	// Capabilities_Name_DOCSIS_CABLE_DEVICE corresponds to the value DOCSIS_CABLE_DEVICE of Capabilities_Name
	Capabilities_Name_DOCSIS_CABLE_DEVICE E_Capabilities_Name = 4
	// Capabilities_Name_S_VLAN corresponds to the value S_VLAN of Capabilities_Name
	Capabilities_Name_S_VLAN E_Capabilities_Name = 5
	// Capabilities_Name_TELEPHONE corresponds to the value TELEPHONE of Capabilities_Name
	Capabilities_Name_TELEPHONE E_Capabilities_Name = 6
	// Capabilities_Name_OTHER corresponds to the value OTHER of Capabilities_Name
	Capabilities_Name_OTHER E_Capabilities_Name = 7
	// Capabilities_Name_ROUTER corresponds to the value ROUTER of Capabilities_Name
	Capabilities_Name_ROUTER E_Capabilities_Name = 8
	// Capabilities_Name_C_VLAN corresponds to the value C_VLAN of Capabilities_Name
	Capabilities_Name_C_VLAN E_Capabilities_Name = 9
	// Capabilities_Name_STATION_ONLY corresponds to the value STATION_ONLY of Capabilities_Name
	Capabilities_Name_STATION_ONLY E_Capabilities_Name = 10
	// Capabilities_Name_WLAN_ACCESS_POINT corresponds to the value WLAN_ACCESS_POINT of Capabilities_Name
	Capabilities_Name_WLAN_ACCESS_POINT E_Capabilities_Name = 11
)

// E_ExtendedIpv4Reachability_Prefix_RedistributionType is a derived int64 type which is used to represent
// the enumerated node ExtendedIpv4Reachability_Prefix_RedistributionType. An additional value named
// ExtendedIpv4Reachability_Prefix_RedistributionType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_ExtendedIpv4Reachability_Prefix_RedistributionType int64

// IsYANGGoEnum ensures that ExtendedIpv4Reachability_Prefix_RedistributionType implements the yang.GoEnum
// interface. This ensures that ExtendedIpv4Reachability_Prefix_RedistributionType can be identified as a
// mapped type for a YANG enumeration.
func (E_ExtendedIpv4Reachability_Prefix_RedistributionType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  ExtendedIpv4Reachability_Prefix_RedistributionType.
func (E_ExtendedIpv4Reachability_Prefix_RedistributionType) ΛMap() map[string]map[int64]ygot.EnumDefinition {
	return ΛEnum
}

// String returns a logging-friendly string for E_ExtendedIpv4Reachability_Prefix_RedistributionType.
func (e E_ExtendedIpv4Reachability_Prefix_RedistributionType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_ExtendedIpv4Reachability_Prefix_RedistributionType")
}

const (
	// ExtendedIpv4Reachability_Prefix_RedistributionType_UNSET corresponds to the value UNSET of ExtendedIpv4Reachability_Prefix_RedistributionType
	ExtendedIpv4Reachability_Prefix_RedistributionType_UNSET E_ExtendedIpv4Reachability_Prefix_RedistributionType = 0
	// ExtendedIpv4Reachability_Prefix_RedistributionType_UP corresponds to the value UP of ExtendedIpv4Reachability_Prefix_RedistributionType
	ExtendedIpv4Reachability_Prefix_RedistributionType_UP E_ExtendedIpv4Reachability_Prefix_RedistributionType = 1
	// ExtendedIpv4Reachability_Prefix_RedistributionType_DOWN corresponds to the value DOWN of ExtendedIpv4Reachability_Prefix_RedistributionType
	ExtendedIpv4Reachability_Prefix_RedistributionType_DOWN E_ExtendedIpv4Reachability_Prefix_RedistributionType = 2
)

// E_Ipv4ExternalReachability_Prefix_OriginType is a derived int64 type which is used to represent
// the enumerated node Ipv4ExternalReachability_Prefix_OriginType. An additional value named
// Ipv4ExternalReachability_Prefix_OriginType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Ipv4ExternalReachability_Prefix_OriginType int64

// IsYANGGoEnum ensures that Ipv4ExternalReachability_Prefix_OriginType implements the yang.GoEnum
// interface. This ensures that Ipv4ExternalReachability_Prefix_OriginType can be identified as a
// mapped type for a YANG enumeration.
func (E_Ipv4ExternalReachability_Prefix_OriginType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Ipv4ExternalReachability_Prefix_OriginType.
func (E_Ipv4ExternalReachability_Prefix_OriginType) ΛMap() map[string]map[int64]ygot.EnumDefinition {
	return ΛEnum
}

// String returns a logging-friendly string for E_Ipv4ExternalReachability_Prefix_OriginType.
func (e E_Ipv4ExternalReachability_Prefix_OriginType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Ipv4ExternalReachability_Prefix_OriginType")
}

const (
	// Ipv4ExternalReachability_Prefix_OriginType_UNSET corresponds to the value UNSET of Ipv4ExternalReachability_Prefix_OriginType
	Ipv4ExternalReachability_Prefix_OriginType_UNSET E_Ipv4ExternalReachability_Prefix_OriginType = 0
	// Ipv4ExternalReachability_Prefix_OriginType_INTERNAL corresponds to the value INTERNAL of Ipv4ExternalReachability_Prefix_OriginType
	Ipv4ExternalReachability_Prefix_OriginType_INTERNAL E_Ipv4ExternalReachability_Prefix_OriginType = 1
	// Ipv4ExternalReachability_Prefix_OriginType_EXTERNAL corresponds to the value EXTERNAL of Ipv4ExternalReachability_Prefix_OriginType
	Ipv4ExternalReachability_Prefix_OriginType_EXTERNAL E_Ipv4ExternalReachability_Prefix_OriginType = 2
)

// E_Ipv4ExternalReachability_Prefix_RedistributionType is a derived int64 type which is used to represent
// the enumerated node Ipv4ExternalReachability_Prefix_RedistributionType. An additional value named
// Ipv4ExternalReachability_Prefix_RedistributionType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Ipv4ExternalReachability_Prefix_RedistributionType int64

// IsYANGGoEnum ensures that Ipv4ExternalReachability_Prefix_RedistributionType implements the yang.GoEnum
// interface. This ensures that Ipv4ExternalReachability_Prefix_RedistributionType can be identified as a
// mapped type for a YANG enumeration.
func (E_Ipv4ExternalReachability_Prefix_RedistributionType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Ipv4ExternalReachability_Prefix_RedistributionType.
func (E_Ipv4ExternalReachability_Prefix_RedistributionType) ΛMap() map[string]map[int64]ygot.EnumDefinition {
	return ΛEnum
}

// String returns a logging-friendly string for E_Ipv4ExternalReachability_Prefix_RedistributionType.
func (e E_Ipv4ExternalReachability_Prefix_RedistributionType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Ipv4ExternalReachability_Prefix_RedistributionType")
}

const (
	// Ipv4ExternalReachability_Prefix_RedistributionType_UNSET corresponds to the value UNSET of Ipv4ExternalReachability_Prefix_RedistributionType
	Ipv4ExternalReachability_Prefix_RedistributionType_UNSET E_Ipv4ExternalReachability_Prefix_RedistributionType = 0
	// Ipv4ExternalReachability_Prefix_RedistributionType_UP corresponds to the value UP of Ipv4ExternalReachability_Prefix_RedistributionType
	Ipv4ExternalReachability_Prefix_RedistributionType_UP E_Ipv4ExternalReachability_Prefix_RedistributionType = 1
	// Ipv4ExternalReachability_Prefix_RedistributionType_DOWN corresponds to the value DOWN of Ipv4ExternalReachability_Prefix_RedistributionType
	Ipv4ExternalReachability_Prefix_RedistributionType_DOWN E_Ipv4ExternalReachability_Prefix_RedistributionType = 2
)

// E_Ipv6Reachability_Prefix_OriginType is a derived int64 type which is used to represent
// the enumerated node Ipv6Reachability_Prefix_OriginType. An additional value named
// Ipv6Reachability_Prefix_OriginType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Ipv6Reachability_Prefix_OriginType int64

// IsYANGGoEnum ensures that Ipv6Reachability_Prefix_OriginType implements the yang.GoEnum
// interface. This ensures that Ipv6Reachability_Prefix_OriginType can be identified as a
// mapped type for a YANG enumeration.
func (E_Ipv6Reachability_Prefix_OriginType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Ipv6Reachability_Prefix_OriginType.
func (E_Ipv6Reachability_Prefix_OriginType) ΛMap() map[string]map[int64]ygot.EnumDefinition {
	return ΛEnum
}

// String returns a logging-friendly string for E_Ipv6Reachability_Prefix_OriginType.
func (e E_Ipv6Reachability_Prefix_OriginType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Ipv6Reachability_Prefix_OriginType")
}

const (
	// Ipv6Reachability_Prefix_OriginType_UNSET corresponds to the value UNSET of Ipv6Reachability_Prefix_OriginType
	Ipv6Reachability_Prefix_OriginType_UNSET E_Ipv6Reachability_Prefix_OriginType = 0
	// Ipv6Reachability_Prefix_OriginType_INTERNAL corresponds to the value INTERNAL of Ipv6Reachability_Prefix_OriginType
	Ipv6Reachability_Prefix_OriginType_INTERNAL E_Ipv6Reachability_Prefix_OriginType = 1
	// Ipv6Reachability_Prefix_OriginType_EXTERNAL corresponds to the value EXTERNAL of Ipv6Reachability_Prefix_OriginType
	Ipv6Reachability_Prefix_OriginType_EXTERNAL E_Ipv6Reachability_Prefix_OriginType = 2
)

// E_Ipv6Reachability_Prefix_RedistributionType is a derived int64 type which is used to represent
// the enumerated node Ipv6Reachability_Prefix_RedistributionType. An additional value named
// Ipv6Reachability_Prefix_RedistributionType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Ipv6Reachability_Prefix_RedistributionType int64

// IsYANGGoEnum ensures that Ipv6Reachability_Prefix_RedistributionType implements the yang.GoEnum
// interface. This ensures that Ipv6Reachability_Prefix_RedistributionType can be identified as a
// mapped type for a YANG enumeration.
func (E_Ipv6Reachability_Prefix_RedistributionType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Ipv6Reachability_Prefix_RedistributionType.
func (E_Ipv6Reachability_Prefix_RedistributionType) ΛMap() map[string]map[int64]ygot.EnumDefinition {
	return ΛEnum
}

// String returns a logging-friendly string for E_Ipv6Reachability_Prefix_RedistributionType.
func (e E_Ipv6Reachability_Prefix_RedistributionType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Ipv6Reachability_Prefix_RedistributionType")
}

const (
	// Ipv6Reachability_Prefix_RedistributionType_UNSET corresponds to the value UNSET of Ipv6Reachability_Prefix_RedistributionType
	Ipv6Reachability_Prefix_RedistributionType_UNSET E_Ipv6Reachability_Prefix_RedistributionType = 0
	// Ipv6Reachability_Prefix_RedistributionType_UP corresponds to the value UP of Ipv6Reachability_Prefix_RedistributionType
	Ipv6Reachability_Prefix_RedistributionType_UP E_Ipv6Reachability_Prefix_RedistributionType = 1
	// Ipv6Reachability_Prefix_RedistributionType_DOWN corresponds to the value DOWN of Ipv6Reachability_Prefix_RedistributionType
	Ipv6Reachability_Prefix_RedistributionType_DOWN E_Ipv6Reachability_Prefix_RedistributionType = 2
)

// E_Lacp_LacpActivityType is a derived int64 type which is used to represent
// the enumerated node Lacp_LacpActivityType. An additional value named
// Lacp_LacpActivityType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Lacp_LacpActivityType int64

// IsYANGGoEnum ensures that Lacp_LacpActivityType implements the yang.GoEnum
// interface. This ensures that Lacp_LacpActivityType can be identified as a
// mapped type for a YANG enumeration.
func (E_Lacp_LacpActivityType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Lacp_LacpActivityType.
func (E_Lacp_LacpActivityType) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_Lacp_LacpActivityType.
func (e E_Lacp_LacpActivityType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Lacp_LacpActivityType")
}

const (
	// Lacp_LacpActivityType_UNSET corresponds to the value UNSET of Lacp_LacpActivityType
	Lacp_LacpActivityType_UNSET E_Lacp_LacpActivityType = 0
	// Lacp_LacpActivityType_ACTIVE corresponds to the value ACTIVE of Lacp_LacpActivityType
	Lacp_LacpActivityType_ACTIVE E_Lacp_LacpActivityType = 1
	// Lacp_LacpActivityType_PASSIVE corresponds to the value PASSIVE of Lacp_LacpActivityType
	Lacp_LacpActivityType_PASSIVE E_Lacp_LacpActivityType = 2
)

// E_Lacp_LacpSynchronizationType is a derived int64 type which is used to represent
// the enumerated node Lacp_LacpSynchronizationType. An additional value named
// Lacp_LacpSynchronizationType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Lacp_LacpSynchronizationType int64

// IsYANGGoEnum ensures that Lacp_LacpSynchronizationType implements the yang.GoEnum
// interface. This ensures that Lacp_LacpSynchronizationType can be identified as a
// mapped type for a YANG enumeration.
func (E_Lacp_LacpSynchronizationType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Lacp_LacpSynchronizationType.
func (E_Lacp_LacpSynchronizationType) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_Lacp_LacpSynchronizationType.
func (e E_Lacp_LacpSynchronizationType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Lacp_LacpSynchronizationType")
}

const (
	// Lacp_LacpSynchronizationType_UNSET corresponds to the value UNSET of Lacp_LacpSynchronizationType
	Lacp_LacpSynchronizationType_UNSET E_Lacp_LacpSynchronizationType = 0
	// Lacp_LacpSynchronizationType_IN_SYNC corresponds to the value IN_SYNC of Lacp_LacpSynchronizationType
	Lacp_LacpSynchronizationType_IN_SYNC E_Lacp_LacpSynchronizationType = 1
	// Lacp_LacpSynchronizationType_OUT_SYNC corresponds to the value OUT_SYNC of Lacp_LacpSynchronizationType
	Lacp_LacpSynchronizationType_OUT_SYNC E_Lacp_LacpSynchronizationType = 2
)

// E_Lacp_LacpTimeoutType is a derived int64 type which is used to represent
// the enumerated node Lacp_LacpTimeoutType. An additional value named
// Lacp_LacpTimeoutType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Lacp_LacpTimeoutType int64

// IsYANGGoEnum ensures that Lacp_LacpTimeoutType implements the yang.GoEnum
// interface. This ensures that Lacp_LacpTimeoutType can be identified as a
// mapped type for a YANG enumeration.
func (E_Lacp_LacpTimeoutType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Lacp_LacpTimeoutType.
func (E_Lacp_LacpTimeoutType) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_Lacp_LacpTimeoutType.
func (e E_Lacp_LacpTimeoutType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Lacp_LacpTimeoutType")
}

const (
	// Lacp_LacpTimeoutType_UNSET corresponds to the value UNSET of Lacp_LacpTimeoutType
	Lacp_LacpTimeoutType_UNSET E_Lacp_LacpTimeoutType = 0
	// Lacp_LacpTimeoutType_LONG corresponds to the value LONG of Lacp_LacpTimeoutType
	Lacp_LacpTimeoutType_LONG E_Lacp_LacpTimeoutType = 1
	// Lacp_LacpTimeoutType_SHORT corresponds to the value SHORT of Lacp_LacpTimeoutType
	Lacp_LacpTimeoutType_SHORT E_Lacp_LacpTimeoutType = 2
)

// E_Lag_OperStatus is a derived int64 type which is used to represent
// the enumerated node Lag_OperStatus. An additional value named
// Lag_OperStatus_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Lag_OperStatus int64

// IsYANGGoEnum ensures that Lag_OperStatus implements the yang.GoEnum
// interface. This ensures that Lag_OperStatus can be identified as a
// mapped type for a YANG enumeration.
func (E_Lag_OperStatus) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Lag_OperStatus.
func (E_Lag_OperStatus) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_Lag_OperStatus.
func (e E_Lag_OperStatus) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Lag_OperStatus")
}

const (
	// Lag_OperStatus_UNSET corresponds to the value UNSET of Lag_OperStatus
	Lag_OperStatus_UNSET E_Lag_OperStatus = 0
	// Lag_OperStatus_UP corresponds to the value UP of Lag_OperStatus
	Lag_OperStatus_UP E_Lag_OperStatus = 1
	// Lag_OperStatus_DOWN corresponds to the value DOWN of Lag_OperStatus
	Lag_OperStatus_DOWN E_Lag_OperStatus = 2
)

// E_LldpNeighbor_ChassisIdType is a derived int64 type which is used to represent
// the enumerated node LldpNeighbor_ChassisIdType. An additional value named
// LldpNeighbor_ChassisIdType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_LldpNeighbor_ChassisIdType int64

// IsYANGGoEnum ensures that LldpNeighbor_ChassisIdType implements the yang.GoEnum
// interface. This ensures that LldpNeighbor_ChassisIdType can be identified as a
// mapped type for a YANG enumeration.
func (E_LldpNeighbor_ChassisIdType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  LldpNeighbor_ChassisIdType.
func (E_LldpNeighbor_ChassisIdType) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_LldpNeighbor_ChassisIdType.
func (e E_LldpNeighbor_ChassisIdType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_LldpNeighbor_ChassisIdType")
}

const (
	// LldpNeighbor_ChassisIdType_UNSET corresponds to the value UNSET of LldpNeighbor_ChassisIdType
	LldpNeighbor_ChassisIdType_UNSET E_LldpNeighbor_ChassisIdType = 0
	// LldpNeighbor_ChassisIdType_PORT_COMPONENT corresponds to the value PORT_COMPONENT of LldpNeighbor_ChassisIdType
	LldpNeighbor_ChassisIdType_PORT_COMPONENT E_LldpNeighbor_ChassisIdType = 1
	// LldpNeighbor_ChassisIdType_NETWORK_ADDRESS corresponds to the value NETWORK_ADDRESS of LldpNeighbor_ChassisIdType
	LldpNeighbor_ChassisIdType_NETWORK_ADDRESS E_LldpNeighbor_ChassisIdType = 2
	// LldpNeighbor_ChassisIdType_CHASSIS_COMPONENT corresponds to the value CHASSIS_COMPONENT of LldpNeighbor_ChassisIdType
	LldpNeighbor_ChassisIdType_CHASSIS_COMPONENT E_LldpNeighbor_ChassisIdType = 3
	// LldpNeighbor_ChassisIdType_MAC_ADDRESS corresponds to the value MAC_ADDRESS of LldpNeighbor_ChassisIdType
	LldpNeighbor_ChassisIdType_MAC_ADDRESS E_LldpNeighbor_ChassisIdType = 4
	// LldpNeighbor_ChassisIdType_INTERFACE_NAME corresponds to the value INTERFACE_NAME of LldpNeighbor_ChassisIdType
	LldpNeighbor_ChassisIdType_INTERFACE_NAME E_LldpNeighbor_ChassisIdType = 5
	// LldpNeighbor_ChassisIdType_LOCAL corresponds to the value LOCAL of LldpNeighbor_ChassisIdType
	LldpNeighbor_ChassisIdType_LOCAL E_LldpNeighbor_ChassisIdType = 6
	// LldpNeighbor_ChassisIdType_INTERFACE_ALIAS corresponds to the value INTERFACE_ALIAS of LldpNeighbor_ChassisIdType
	LldpNeighbor_ChassisIdType_INTERFACE_ALIAS E_LldpNeighbor_ChassisIdType = 7
)

// E_LldpNeighbor_PortIdType is a derived int64 type which is used to represent
// the enumerated node LldpNeighbor_PortIdType. An additional value named
// LldpNeighbor_PortIdType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_LldpNeighbor_PortIdType int64

// IsYANGGoEnum ensures that LldpNeighbor_PortIdType implements the yang.GoEnum
// interface. This ensures that LldpNeighbor_PortIdType can be identified as a
// mapped type for a YANG enumeration.
func (E_LldpNeighbor_PortIdType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  LldpNeighbor_PortIdType.
func (E_LldpNeighbor_PortIdType) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_LldpNeighbor_PortIdType.
func (e E_LldpNeighbor_PortIdType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_LldpNeighbor_PortIdType")
}

const (
	// LldpNeighbor_PortIdType_UNSET corresponds to the value UNSET of LldpNeighbor_PortIdType
	LldpNeighbor_PortIdType_UNSET E_LldpNeighbor_PortIdType = 0
	// LldpNeighbor_PortIdType_PORT_COMPONENT corresponds to the value PORT_COMPONENT of LldpNeighbor_PortIdType
	LldpNeighbor_PortIdType_PORT_COMPONENT E_LldpNeighbor_PortIdType = 1
	// LldpNeighbor_PortIdType_NETWORK_ADDRESS corresponds to the value NETWORK_ADDRESS of LldpNeighbor_PortIdType
	LldpNeighbor_PortIdType_NETWORK_ADDRESS E_LldpNeighbor_PortIdType = 2
	// LldpNeighbor_PortIdType_AGENT_CIRCUIT_ID corresponds to the value AGENT_CIRCUIT_ID of LldpNeighbor_PortIdType
	LldpNeighbor_PortIdType_AGENT_CIRCUIT_ID E_LldpNeighbor_PortIdType = 3
	// LldpNeighbor_PortIdType_MAC_ADDRESS corresponds to the value MAC_ADDRESS of LldpNeighbor_PortIdType
	LldpNeighbor_PortIdType_MAC_ADDRESS E_LldpNeighbor_PortIdType = 4
	// LldpNeighbor_PortIdType_INTERFACE_NAME corresponds to the value INTERFACE_NAME of LldpNeighbor_PortIdType
	LldpNeighbor_PortIdType_INTERFACE_NAME E_LldpNeighbor_PortIdType = 5
	// LldpNeighbor_PortIdType_LOCAL corresponds to the value LOCAL of LldpNeighbor_PortIdType
	LldpNeighbor_PortIdType_LOCAL E_LldpNeighbor_PortIdType = 6
	// LldpNeighbor_PortIdType_INTERFACE_ALIAS corresponds to the value INTERFACE_ALIAS of LldpNeighbor_PortIdType
	LldpNeighbor_PortIdType_INTERFACE_ALIAS E_LldpNeighbor_PortIdType = 7
)

// E_Lsps_Flags is a derived int64 type which is used to represent
// the enumerated node Lsps_Flags. An additional value named
// Lsps_Flags_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Lsps_Flags int64

// IsYANGGoEnum ensures that Lsps_Flags implements the yang.GoEnum
// interface. This ensures that Lsps_Flags can be identified as a
// mapped type for a YANG enumeration.
func (E_Lsps_Flags) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Lsps_Flags.
func (E_Lsps_Flags) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_Lsps_Flags.
func (e E_Lsps_Flags) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Lsps_Flags")
}

const (
	// Lsps_Flags_UNSET corresponds to the value UNSET of Lsps_Flags
	Lsps_Flags_UNSET E_Lsps_Flags = 0
	// Lsps_Flags_PARTITION_REPAIR corresponds to the value PARTITION_REPAIR of Lsps_Flags
	Lsps_Flags_PARTITION_REPAIR E_Lsps_Flags = 1
	// Lsps_Flags_ATTACHED_ERROR corresponds to the value ATTACHED_ERROR of Lsps_Flags
	Lsps_Flags_ATTACHED_ERROR E_Lsps_Flags = 2
	// Lsps_Flags_ATTACHED_EXPENSE corresponds to the value ATTACHED_EXPENSE of Lsps_Flags
	Lsps_Flags_ATTACHED_EXPENSE E_Lsps_Flags = 3
	// Lsps_Flags_ATTACHED_DELAY corresponds to the value ATTACHED_DELAY of Lsps_Flags
	Lsps_Flags_ATTACHED_DELAY E_Lsps_Flags = 4
	// Lsps_Flags_ATTACHED_DEFAULT corresponds to the value ATTACHED_DEFAULT of Lsps_Flags
	Lsps_Flags_ATTACHED_DEFAULT E_Lsps_Flags = 5
	// Lsps_Flags_OVERLOAD corresponds to the value OVERLOAD of Lsps_Flags
	Lsps_Flags_OVERLOAD E_Lsps_Flags = 6
)

// E_Lsps_PduType is a derived int64 type which is used to represent
// the enumerated node Lsps_PduType. An additional value named
// Lsps_PduType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Lsps_PduType int64

// IsYANGGoEnum ensures that Lsps_PduType implements the yang.GoEnum
// interface. This ensures that Lsps_PduType can be identified as a
// mapped type for a YANG enumeration.
func (E_Lsps_PduType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Lsps_PduType.
func (E_Lsps_PduType) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_Lsps_PduType.
func (e E_Lsps_PduType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Lsps_PduType")
}

const (
	// Lsps_PduType_UNSET corresponds to the value UNSET of Lsps_PduType
	Lsps_PduType_UNSET E_Lsps_PduType = 0
	// Lsps_PduType_LEVEL_1 corresponds to the value LEVEL_1 of Lsps_PduType
	Lsps_PduType_LEVEL_1 E_Lsps_PduType = 1
	// Lsps_PduType_LEVEL_2 corresponds to the value LEVEL_2 of Lsps_PduType
	Lsps_PduType_LEVEL_2 E_Lsps_PduType = 2
)

// E_Port_Link is a derived int64 type which is used to represent
// the enumerated node Port_Link. An additional value named
// Port_Link_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Port_Link int64

// IsYANGGoEnum ensures that Port_Link implements the yang.GoEnum
// interface. This ensures that Port_Link can be identified as a
// mapped type for a YANG enumeration.
func (E_Port_Link) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Port_Link.
func (E_Port_Link) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_Port_Link.
func (e E_Port_Link) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Port_Link")
}

const (
	// Port_Link_UNSET corresponds to the value UNSET of Port_Link
	Port_Link_UNSET E_Port_Link = 0
	// Port_Link_UP corresponds to the value UP of Port_Link
	Port_Link_UP E_Port_Link = 1
	// Port_Link_DOWN corresponds to the value DOWN of Port_Link
	Port_Link_DOWN E_Port_Link = 2
)

// E_State_CommunityType is a derived int64 type which is used to represent
// the enumerated node State_CommunityType. An additional value named
// State_CommunityType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_State_CommunityType int64

// IsYANGGoEnum ensures that State_CommunityType implements the yang.GoEnum
// interface. This ensures that State_CommunityType can be identified as a
// mapped type for a YANG enumeration.
func (E_State_CommunityType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  State_CommunityType.
func (E_State_CommunityType) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_State_CommunityType.
func (e E_State_CommunityType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_State_CommunityType")
}

const (
	// State_CommunityType_UNSET corresponds to the value UNSET of State_CommunityType
	State_CommunityType_UNSET E_State_CommunityType = 0
	// State_CommunityType_MANUAL_AS_NUMBER corresponds to the value MANUAL_AS_NUMBER of State_CommunityType
	State_CommunityType_MANUAL_AS_NUMBER E_State_CommunityType = 1
	// State_CommunityType_NO_EXPORT corresponds to the value NO_EXPORT of State_CommunityType
	State_CommunityType_NO_EXPORT E_State_CommunityType = 2
	// State_CommunityType_NO_ADVERTISE corresponds to the value NO_ADVERTISE of State_CommunityType
	State_CommunityType_NO_ADVERTISE E_State_CommunityType = 3
	// State_CommunityType_NO_EXPORT_SUBCONFED corresponds to the value NO_EXPORT_SUBCONFED of State_CommunityType
	State_CommunityType_NO_EXPORT_SUBCONFED E_State_CommunityType = 4
	// State_CommunityType_LLGR_STALE corresponds to the value LLGR_STALE of State_CommunityType
	State_CommunityType_LLGR_STALE E_State_CommunityType = 5
	// State_CommunityType_NO_LLGR corresponds to the value NO_LLGR of State_CommunityType
	State_CommunityType_NO_LLGR E_State_CommunityType = 6
)

// E_State_Flags is a derived int64 type which is used to represent
// the enumerated node State_Flags. An additional value named
// State_Flags_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_State_Flags int64

// IsYANGGoEnum ensures that State_Flags implements the yang.GoEnum
// interface. This ensures that State_Flags can be identified as a
// mapped type for a YANG enumeration.
func (E_State_Flags) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  State_Flags.
func (E_State_Flags) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_State_Flags.
func (e E_State_Flags) String() string {
	return ygot.EnumLogString(e, int64(e), "E_State_Flags")
}

const (
	// State_Flags_UNSET corresponds to the value UNSET of State_Flags
	State_Flags_UNSET E_State_Flags = 0
	// State_Flags_EXTERNAL_FLAG corresponds to the value EXTERNAL_FLAG of State_Flags
	State_Flags_EXTERNAL_FLAG E_State_Flags = 1
	// State_Flags_READVERTISEMENT_FLAG corresponds to the value READVERTISEMENT_FLAG of State_Flags
	State_Flags_READVERTISEMENT_FLAG E_State_Flags = 2
	// State_Flags_NODE_FLAG corresponds to the value NODE_FLAG of State_Flags
	State_Flags_NODE_FLAG E_State_Flags = 3
)

// E_State_SegmentType is a derived int64 type which is used to represent
// the enumerated node State_SegmentType. An additional value named
// State_SegmentType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_State_SegmentType int64

// IsYANGGoEnum ensures that State_SegmentType implements the yang.GoEnum
// interface. This ensures that State_SegmentType can be identified as a
// mapped type for a YANG enumeration.
func (E_State_SegmentType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  State_SegmentType.
func (E_State_SegmentType) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_State_SegmentType.
func (e E_State_SegmentType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_State_SegmentType")
}

const (
	// State_SegmentType_UNSET corresponds to the value UNSET of State_SegmentType
	State_SegmentType_UNSET E_State_SegmentType = 0
	// State_SegmentType_AS_SEQUENCE corresponds to the value AS_SEQUENCE of State_SegmentType
	State_SegmentType_AS_SEQUENCE E_State_SegmentType = 1
	// State_SegmentType_AS_SET corresponds to the value AS_SET of State_SegmentType
	State_SegmentType_AS_SET E_State_SegmentType = 2
	// State_SegmentType_AS_CONFED_SEQUENCE corresponds to the value AS_CONFED_SEQUENCE of State_SegmentType
	State_SegmentType_AS_CONFED_SEQUENCE E_State_SegmentType = 3
	// State_SegmentType_AS_CONFED_SET corresponds to the value AS_CONFED_SET of State_SegmentType
	State_SegmentType_AS_CONFED_SET E_State_SegmentType = 4
)

// E_UnicastIpv4Prefix_Origin is a derived int64 type which is used to represent
// the enumerated node UnicastIpv4Prefix_Origin. An additional value named
// UnicastIpv4Prefix_Origin_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_UnicastIpv4Prefix_Origin int64

// IsYANGGoEnum ensures that UnicastIpv4Prefix_Origin implements the yang.GoEnum
// interface. This ensures that UnicastIpv4Prefix_Origin can be identified as a
// mapped type for a YANG enumeration.
func (E_UnicastIpv4Prefix_Origin) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  UnicastIpv4Prefix_Origin.
func (E_UnicastIpv4Prefix_Origin) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_UnicastIpv4Prefix_Origin.
func (e E_UnicastIpv4Prefix_Origin) String() string {
	return ygot.EnumLogString(e, int64(e), "E_UnicastIpv4Prefix_Origin")
}

const (
	// UnicastIpv4Prefix_Origin_UNSET corresponds to the value UNSET of UnicastIpv4Prefix_Origin
	UnicastIpv4Prefix_Origin_UNSET E_UnicastIpv4Prefix_Origin = 0
	// UnicastIpv4Prefix_Origin_IGP corresponds to the value IGP of UnicastIpv4Prefix_Origin
	UnicastIpv4Prefix_Origin_IGP E_UnicastIpv4Prefix_Origin = 1
	// UnicastIpv4Prefix_Origin_EGP corresponds to the value EGP of UnicastIpv4Prefix_Origin
	UnicastIpv4Prefix_Origin_EGP E_UnicastIpv4Prefix_Origin = 2
	// UnicastIpv4Prefix_Origin_INCOMPLETE corresponds to the value INCOMPLETE of UnicastIpv4Prefix_Origin
	UnicastIpv4Prefix_Origin_INCOMPLETE E_UnicastIpv4Prefix_Origin = 3
)

// E_UnicastIpv6Prefix_Origin is a derived int64 type which is used to represent
// the enumerated node UnicastIpv6Prefix_Origin. An additional value named
// UnicastIpv6Prefix_Origin_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_UnicastIpv6Prefix_Origin int64

// IsYANGGoEnum ensures that UnicastIpv6Prefix_Origin implements the yang.GoEnum
// interface. This ensures that UnicastIpv6Prefix_Origin can be identified as a
// mapped type for a YANG enumeration.
func (E_UnicastIpv6Prefix_Origin) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  UnicastIpv6Prefix_Origin.
func (E_UnicastIpv6Prefix_Origin) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_UnicastIpv6Prefix_Origin.
func (e E_UnicastIpv6Prefix_Origin) String() string {
	return ygot.EnumLogString(e, int64(e), "E_UnicastIpv6Prefix_Origin")
}

const (
	// UnicastIpv6Prefix_Origin_UNSET corresponds to the value UNSET of UnicastIpv6Prefix_Origin
	UnicastIpv6Prefix_Origin_UNSET E_UnicastIpv6Prefix_Origin = 0
	// UnicastIpv6Prefix_Origin_IGP corresponds to the value IGP of UnicastIpv6Prefix_Origin
	UnicastIpv6Prefix_Origin_IGP E_UnicastIpv6Prefix_Origin = 1
	// UnicastIpv6Prefix_Origin_EGP corresponds to the value EGP of UnicastIpv6Prefix_Origin
	UnicastIpv6Prefix_Origin_EGP E_UnicastIpv6Prefix_Origin = 2
	// UnicastIpv6Prefix_Origin_INCOMPLETE corresponds to the value INCOMPLETE of UnicastIpv6Prefix_Origin
	UnicastIpv6Prefix_Origin_INCOMPLETE E_UnicastIpv6Prefix_Origin = 3
)
