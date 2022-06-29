/*
Package port is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by /usr/local/google/home/alexmasi/go/pkg/mod/github.com/openconfig/ygot@v0.23.1/genutil/names.go
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
package port

import (
	"github.com/openconfig/ygot/ygot"
)

// PortPath represents the /open-traffic-generator-port/ports/port YANG schema element.
type PortPath struct {
	*ygot.NodePath
}

// PortPathAny represents the wildcard version of the /open-traffic-generator-port/ports/port YANG schema element.
type PortPathAny struct {
	*ygot.NodePath
}

// Port_InRatePath represents the /open-traffic-generator-port/ports/port/state/in-rate YANG schema element.
type Port_InRatePath struct {
	*ygot.NodePath
}

// Port_InRatePathAny represents the wildcard version of the /open-traffic-generator-port/ports/port/state/in-rate YANG schema element.
type Port_InRatePathAny struct {
	*ygot.NodePath
}

// Port_LinkPath represents the /open-traffic-generator-port/ports/port/state/link YANG schema element.
type Port_LinkPath struct {
	*ygot.NodePath
}

// Port_LinkPathAny represents the wildcard version of the /open-traffic-generator-port/ports/port/state/link YANG schema element.
type Port_LinkPathAny struct {
	*ygot.NodePath
}

// Port_NamePath represents the /open-traffic-generator-port/ports/port/state/name YANG schema element.
type Port_NamePath struct {
	*ygot.NodePath
}

// Port_NamePathAny represents the wildcard version of the /open-traffic-generator-port/ports/port/state/name YANG schema element.
type Port_NamePathAny struct {
	*ygot.NodePath
}

// Port_OutRatePath represents the /open-traffic-generator-port/ports/port/state/out-rate YANG schema element.
type Port_OutRatePath struct {
	*ygot.NodePath
}

// Port_OutRatePathAny represents the wildcard version of the /open-traffic-generator-port/ports/port/state/out-rate YANG schema element.
type Port_OutRatePathAny struct {
	*ygot.NodePath
}

// Counters (container): Counters of an OTG port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "state/counters"
// Path from root: "/ports/port/state/counters"
func (n *PortPath) Counters() *Port_CountersPath {
	return &Port_CountersPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "counters"},
			map[string]interface{}{},
			n,
		),
	}
}

// Counters (container): Counters of an OTG port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "state/counters"
// Path from root: "/ports/port/state/counters"
func (n *PortPathAny) Counters() *Port_CountersPathAny {
	return &Port_CountersPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "counters"},
			map[string]interface{}{},
			n,
		),
	}
}

// InRate (leaf): The current receive rate of an OTG port, measured in
// bits per second.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "state/in-rate"
// Path from root: "/ports/port/state/in-rate"
func (n *PortPath) InRate() *Port_InRatePath {
	return &Port_InRatePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "in-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// InRate (leaf): The current receive rate of an OTG port, measured in
// bits per second.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "state/in-rate"
// Path from root: "/ports/port/state/in-rate"
func (n *PortPathAny) InRate() *Port_InRatePathAny {
	return &Port_InRatePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "in-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// Link (leaf): The current state of an OTG port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "state/link"
// Path from root: "/ports/port/state/link"
func (n *PortPath) Link() *Port_LinkPath {
	return &Port_LinkPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "link"},
			map[string]interface{}{},
			n,
		),
	}
}

// Link (leaf): The current state of an OTG port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "state/link"
// Path from root: "/ports/port/state/link"
func (n *PortPathAny) Link() *Port_LinkPathAny {
	return &Port_LinkPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "link"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): An arbitary name of an OTG port determined by the OTG
// configuration.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "state/name"
// Path from root: "/ports/port/state/name"
func (n *PortPath) Name() *Port_NamePath {
	return &Port_NamePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): An arbitary name of an OTG port determined by the OTG
// configuration.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "state/name"
// Path from root: "/ports/port/state/name"
func (n *PortPathAny) Name() *Port_NamePathAny {
	return &Port_NamePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutRate (leaf): The current transmit rate of an OTG port, measured in
// bits per second.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "state/out-rate"
// Path from root: "/ports/port/state/out-rate"
func (n *PortPath) OutRate() *Port_OutRatePath {
	return &Port_OutRatePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "out-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutRate (leaf): The current transmit rate of an OTG port, measured in
// bits per second.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "state/out-rate"
// Path from root: "/ports/port/state/out-rate"
func (n *PortPathAny) OutRate() *Port_OutRatePathAny {
	return &Port_OutRatePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "out-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// Port_CountersPath represents the /open-traffic-generator-port/ports/port/state/counters YANG schema element.
type Port_CountersPath struct {
	*ygot.NodePath
}

// Port_CountersPathAny represents the wildcard version of the /open-traffic-generator-port/ports/port/state/counters YANG schema element.
type Port_CountersPathAny struct {
	*ygot.NodePath
}

// Port_Counters_InFramesPath represents the /open-traffic-generator-port/ports/port/state/counters/in-frames YANG schema element.
type Port_Counters_InFramesPath struct {
	*ygot.NodePath
}

// Port_Counters_InFramesPathAny represents the wildcard version of the /open-traffic-generator-port/ports/port/state/counters/in-frames YANG schema element.
type Port_Counters_InFramesPathAny struct {
	*ygot.NodePath
}

// Port_Counters_InOctetsPath represents the /open-traffic-generator-port/ports/port/state/counters/in-octets YANG schema element.
type Port_Counters_InOctetsPath struct {
	*ygot.NodePath
}

// Port_Counters_InOctetsPathAny represents the wildcard version of the /open-traffic-generator-port/ports/port/state/counters/in-octets YANG schema element.
type Port_Counters_InOctetsPathAny struct {
	*ygot.NodePath
}

// Port_Counters_OutFramesPath represents the /open-traffic-generator-port/ports/port/state/counters/out-frames YANG schema element.
type Port_Counters_OutFramesPath struct {
	*ygot.NodePath
}

// Port_Counters_OutFramesPathAny represents the wildcard version of the /open-traffic-generator-port/ports/port/state/counters/out-frames YANG schema element.
type Port_Counters_OutFramesPathAny struct {
	*ygot.NodePath
}

// Port_Counters_OutOctetsPath represents the /open-traffic-generator-port/ports/port/state/counters/out-octets YANG schema element.
type Port_Counters_OutOctetsPath struct {
	*ygot.NodePath
}

// Port_Counters_OutOctetsPathAny represents the wildcard version of the /open-traffic-generator-port/ports/port/state/counters/out-octets YANG schema element.
type Port_Counters_OutOctetsPathAny struct {
	*ygot.NodePath
}

// InFrames (leaf): The total number of packets received on the port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "in-frames"
// Path from root: "/ports/port/state/counters/in-frames"
func (n *Port_CountersPath) InFrames() *Port_Counters_InFramesPath {
	return &Port_Counters_InFramesPath{
		NodePath: ygot.NewNodePath(
			[]string{"in-frames"},
			map[string]interface{}{},
			n,
		),
	}
}

// InFrames (leaf): The total number of packets received on the port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "in-frames"
// Path from root: "/ports/port/state/counters/in-frames"
func (n *Port_CountersPathAny) InFrames() *Port_Counters_InFramesPathAny {
	return &Port_Counters_InFramesPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"in-frames"},
			map[string]interface{}{},
			n,
		),
	}
}

// InOctets (leaf): The total number of octets received on the port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "in-octets"
// Path from root: "/ports/port/state/counters/in-octets"
func (n *Port_CountersPath) InOctets() *Port_Counters_InOctetsPath {
	return &Port_Counters_InOctetsPath{
		NodePath: ygot.NewNodePath(
			[]string{"in-octets"},
			map[string]interface{}{},
			n,
		),
	}
}

// InOctets (leaf): The total number of octets received on the port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "in-octets"
// Path from root: "/ports/port/state/counters/in-octets"
func (n *Port_CountersPathAny) InOctets() *Port_Counters_InOctetsPathAny {
	return &Port_Counters_InOctetsPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"in-octets"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutFrames (leaf): The total number of packets transmitted on the port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "out-frames"
// Path from root: "/ports/port/state/counters/out-frames"
func (n *Port_CountersPath) OutFrames() *Port_Counters_OutFramesPath {
	return &Port_Counters_OutFramesPath{
		NodePath: ygot.NewNodePath(
			[]string{"out-frames"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutFrames (leaf): The total number of packets transmitted on the port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "out-frames"
// Path from root: "/ports/port/state/counters/out-frames"
func (n *Port_CountersPathAny) OutFrames() *Port_Counters_OutFramesPathAny {
	return &Port_Counters_OutFramesPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"out-frames"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutOctets (leaf): The total number of octets transmitted on the port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "out-octets"
// Path from root: "/ports/port/state/counters/out-octets"
func (n *Port_CountersPath) OutOctets() *Port_Counters_OutOctetsPath {
	return &Port_Counters_OutOctetsPath{
		NodePath: ygot.NewNodePath(
			[]string{"out-octets"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutOctets (leaf): The total number of octets transmitted on the port.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "out-octets"
// Path from root: "/ports/port/state/counters/out-octets"
func (n *Port_CountersPathAny) OutOctets() *Port_Counters_OutOctetsPathAny {
	return &Port_Counters_OutOctetsPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"out-octets"},
			map[string]interface{}{},
			n,
		),
	}
}
