/*
Package otg is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was true
in this case).

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
package otg

import (
	"fmt"
	"reflect"

	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Interface represents the /open-traffic-generator-discovery/interfaces/interface YANG schema element.
type Interface struct {
	Ipv4Neighbor map[string]*Interface_Ipv4Neighbor `path:"ipv4-neighbors/ipv4-neighbor" module:"open-traffic-generator-discovery/open-traffic-generator-discovery"`
	Ipv6Neighbor map[string]*Interface_Ipv6Neighbor `path:"ipv6-neighbors/ipv6-neighbor" module:"open-traffic-generator-discovery/open-traffic-generator-discovery"`
	Name         *string                            `path:"state/name|name" module:"open-traffic-generator-discovery/open-traffic-generator-discovery|open-traffic-generator-discovery" shadow-path:"name" shadow-module:"open-traffic-generator-discovery"`
}

// IsYANGGoStruct ensures that Interface implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Interface) IsYANGGoStruct() {}

// NewIpv4Neighbor creates a new entry in the Ipv4Neighbor list of the
// Interface struct. The keys of the list are populated from the input
// arguments.
func (t *Interface) NewIpv4Neighbor(Ipv4Address string) (*Interface_Ipv4Neighbor, error) {

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.Ipv4Neighbor == nil {
		t.Ipv4Neighbor = make(map[string]*Interface_Ipv4Neighbor)
	}

	key := Ipv4Address

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.Ipv4Neighbor[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list Ipv4Neighbor", key)
	}

	t.Ipv4Neighbor[key] = &Interface_Ipv4Neighbor{
		Ipv4Address: &Ipv4Address,
	}

	return t.Ipv4Neighbor[key], nil
}

// RenameIpv4Neighbor renames an entry in the list Ipv4Neighbor within
// the Interface struct. The entry with key oldK is renamed to newK updating
// the key within the value.
func (t *Interface) RenameIpv4Neighbor(oldK, newK string) error {
	if _, ok := t.Ipv4Neighbor[newK]; ok {
		return fmt.Errorf("key %v already exists in Ipv4Neighbor", newK)
	}

	e, ok := t.Ipv4Neighbor[oldK]
	if !ok {
		return fmt.Errorf("key %v not found in Ipv4Neighbor", oldK)
	}
	e.Ipv4Address = &newK

	t.Ipv4Neighbor[newK] = e
	delete(t.Ipv4Neighbor, oldK)
	return nil
}

// GetOrCreateIpv4Neighbor retrieves the value with the specified keys from
// the receiver Interface. If the entry does not exist, then it is created.
// It returns the existing or new list member.
func (t *Interface) GetOrCreateIpv4Neighbor(Ipv4Address string) *Interface_Ipv4Neighbor {

	key := Ipv4Address

	if v, ok := t.Ipv4Neighbor[key]; ok {
		return v
	}
	// Panic if we receive an error, since we should have retrieved an existing
	// list member. This allows chaining of GetOrCreate methods.
	v, err := t.NewIpv4Neighbor(Ipv4Address)
	if err != nil {
		panic(fmt.Sprintf("GetOrCreateIpv4Neighbor got unexpected error: %v", err))
	}
	return v
}

// GetIpv4Neighbor retrieves the value with the specified key from
// the Ipv4Neighbor map field of Interface. If the receiver is nil, or
// the specified key is not present in the list, nil is returned such that Get*
// methods may be safely chained.
func (t *Interface) GetIpv4Neighbor(Ipv4Address string) *Interface_Ipv4Neighbor {

	if t == nil {
		return nil
	}

	key := Ipv4Address

	if lm, ok := t.Ipv4Neighbor[key]; ok {
		return lm
	}
	return nil
}

// DeleteIpv4Neighbor deletes the value with the specified keys from
// the receiver Interface. If there is no such element, the function
// is a no-op.
func (t *Interface) DeleteIpv4Neighbor(Ipv4Address string) {
	key := Ipv4Address

	delete(t.Ipv4Neighbor, key)
}

// AppendIpv4Neighbor appends the supplied Interface_Ipv4Neighbor struct to the
// list Ipv4Neighbor of Interface. If the key value(s) specified in
// the supplied Interface_Ipv4Neighbor already exist in the list, an error is
// returned.
func (t *Interface) AppendIpv4Neighbor(v *Interface_Ipv4Neighbor) error {
	if v.Ipv4Address == nil {
		return fmt.Errorf("invalid nil key received for Ipv4Address")
	}

	key := *v.Ipv4Address

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.Ipv4Neighbor == nil {
		t.Ipv4Neighbor = make(map[string]*Interface_Ipv4Neighbor)
	}

	if _, ok := t.Ipv4Neighbor[key]; ok {
		return fmt.Errorf("duplicate key for list Ipv4Neighbor %v", key)
	}

	t.Ipv4Neighbor[key] = v
	return nil
}

// NewIpv6Neighbor creates a new entry in the Ipv6Neighbor list of the
// Interface struct. The keys of the list are populated from the input
// arguments.
func (t *Interface) NewIpv6Neighbor(Ipv6Address string) (*Interface_Ipv6Neighbor, error) {

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.Ipv6Neighbor == nil {
		t.Ipv6Neighbor = make(map[string]*Interface_Ipv6Neighbor)
	}

	key := Ipv6Address

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.Ipv6Neighbor[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list Ipv6Neighbor", key)
	}

	t.Ipv6Neighbor[key] = &Interface_Ipv6Neighbor{
		Ipv6Address: &Ipv6Address,
	}

	return t.Ipv6Neighbor[key], nil
}

// RenameIpv6Neighbor renames an entry in the list Ipv6Neighbor within
// the Interface struct. The entry with key oldK is renamed to newK updating
// the key within the value.
func (t *Interface) RenameIpv6Neighbor(oldK, newK string) error {
	if _, ok := t.Ipv6Neighbor[newK]; ok {
		return fmt.Errorf("key %v already exists in Ipv6Neighbor", newK)
	}

	e, ok := t.Ipv6Neighbor[oldK]
	if !ok {
		return fmt.Errorf("key %v not found in Ipv6Neighbor", oldK)
	}
	e.Ipv6Address = &newK

	t.Ipv6Neighbor[newK] = e
	delete(t.Ipv6Neighbor, oldK)
	return nil
}

// GetOrCreateIpv6Neighbor retrieves the value with the specified keys from
// the receiver Interface. If the entry does not exist, then it is created.
// It returns the existing or new list member.
func (t *Interface) GetOrCreateIpv6Neighbor(Ipv6Address string) *Interface_Ipv6Neighbor {

	key := Ipv6Address

	if v, ok := t.Ipv6Neighbor[key]; ok {
		return v
	}
	// Panic if we receive an error, since we should have retrieved an existing
	// list member. This allows chaining of GetOrCreate methods.
	v, err := t.NewIpv6Neighbor(Ipv6Address)
	if err != nil {
		panic(fmt.Sprintf("GetOrCreateIpv6Neighbor got unexpected error: %v", err))
	}
	return v
}

// GetIpv6Neighbor retrieves the value with the specified key from
// the Ipv6Neighbor map field of Interface. If the receiver is nil, or
// the specified key is not present in the list, nil is returned such that Get*
// methods may be safely chained.
func (t *Interface) GetIpv6Neighbor(Ipv6Address string) *Interface_Ipv6Neighbor {

	if t == nil {
		return nil
	}

	key := Ipv6Address

	if lm, ok := t.Ipv6Neighbor[key]; ok {
		return lm
	}
	return nil
}

// DeleteIpv6Neighbor deletes the value with the specified keys from
// the receiver Interface. If there is no such element, the function
// is a no-op.
func (t *Interface) DeleteIpv6Neighbor(Ipv6Address string) {
	key := Ipv6Address

	delete(t.Ipv6Neighbor, key)
}

// AppendIpv6Neighbor appends the supplied Interface_Ipv6Neighbor struct to the
// list Ipv6Neighbor of Interface. If the key value(s) specified in
// the supplied Interface_Ipv6Neighbor already exist in the list, an error is
// returned.
func (t *Interface) AppendIpv6Neighbor(v *Interface_Ipv6Neighbor) error {
	if v.Ipv6Address == nil {
		return fmt.Errorf("invalid nil key received for Ipv6Address")
	}

	key := *v.Ipv6Address

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.Ipv6Neighbor == nil {
		t.Ipv6Neighbor = make(map[string]*Interface_Ipv6Neighbor)
	}

	if _, ok := t.Ipv6Neighbor[key]; ok {
		return fmt.Errorf("duplicate key for list Ipv6Neighbor %v", key)
	}

	t.Ipv6Neighbor[key] = v
	return nil
}

// GetName retrieves the value of the leaf Name from the Interface
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Name is set, it can
// safely use t.GetName() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Name == nil' before retrieving the leaf's value.
func (t *Interface) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}
	return *t.Name
}

// PopulateDefaults recursively populates unset leaf fields in the Interface
// with default values as specified in the YANG schema, instantiating any nil
// container fields.
func (t *Interface) PopulateDefaults() {
	if t == nil {
		return
	}
	ygot.BuildEmptyTree(t)
	for _, e := range t.Ipv4Neighbor {
		e.PopulateDefaults()
	}
	for _, e := range t.Ipv6Neighbor {
		e.PopulateDefaults()
	}
}

// ΛListKeyMap returns the keys of the Interface struct, which is a YANG list entry.
func (t *Interface) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Name == nil {
		return nil, fmt.Errorf("nil value for key Name")
	}

	return map[string]interface{}{
		"name": *t.Name,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Interface) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Interface"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Interface) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Interface) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of Interface.
func (*Interface) ΛBelongingModule() string {
	return "open-traffic-generator-discovery"
}

// Interface_Ipv4Neighbor represents the /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor YANG schema element.
type Interface_Ipv4Neighbor struct {
	Ipv4Address      *string `path:"state/ipv4-address|ipv4-address" module:"open-traffic-generator-discovery/open-traffic-generator-discovery|open-traffic-generator-discovery" shadow-path:"ipv4-address" shadow-module:"open-traffic-generator-discovery"`
	LinkLayerAddress *string `path:"state/link-layer-address" module:"open-traffic-generator-discovery/open-traffic-generator-discovery"`
}

// IsYANGGoStruct ensures that Interface_Ipv4Neighbor implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Interface_Ipv4Neighbor) IsYANGGoStruct() {}

// GetIpv4Address retrieves the value of the leaf Ipv4Address from the Interface_Ipv4Neighbor
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Ipv4Address is set, it can
// safely use t.GetIpv4Address() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Ipv4Address == nil' before retrieving the leaf's value.
func (t *Interface_Ipv4Neighbor) GetIpv4Address() string {
	if t == nil || t.Ipv4Address == nil {
		return ""
	}
	return *t.Ipv4Address
}

// GetLinkLayerAddress retrieves the value of the leaf LinkLayerAddress from the Interface_Ipv4Neighbor
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if LinkLayerAddress is set, it can
// safely use t.GetLinkLayerAddress() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.LinkLayerAddress == nil' before retrieving the leaf's value.
func (t *Interface_Ipv4Neighbor) GetLinkLayerAddress() string {
	if t == nil || t.LinkLayerAddress == nil {
		return ""
	}
	return *t.LinkLayerAddress
}

// PopulateDefaults recursively populates unset leaf fields in the Interface_Ipv4Neighbor
// with default values as specified in the YANG schema, instantiating any nil
// container fields.
func (t *Interface_Ipv4Neighbor) PopulateDefaults() {
	if t == nil {
		return
	}
	ygot.BuildEmptyTree(t)
}

// ΛListKeyMap returns the keys of the Interface_Ipv4Neighbor struct, which is a YANG list entry.
func (t *Interface_Ipv4Neighbor) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Ipv4Address == nil {
		return nil, fmt.Errorf("nil value for key Ipv4Address")
	}

	return map[string]interface{}{
		"ipv4-address": *t.Ipv4Address,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Interface_Ipv4Neighbor) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Interface_Ipv4Neighbor"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Interface_Ipv4Neighbor) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Interface_Ipv4Neighbor) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of Interface_Ipv4Neighbor.
func (*Interface_Ipv4Neighbor) ΛBelongingModule() string {
	return "open-traffic-generator-discovery"
}

// Interface_Ipv6Neighbor represents the /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor YANG schema element.
type Interface_Ipv6Neighbor struct {
	Ipv6Address      *string `path:"state/ipv6-address|ipv6-address" module:"open-traffic-generator-discovery/open-traffic-generator-discovery|open-traffic-generator-discovery" shadow-path:"ipv6-address" shadow-module:"open-traffic-generator-discovery"`
	LinkLayerAddress *string `path:"state/link-layer-address" module:"open-traffic-generator-discovery/open-traffic-generator-discovery"`
}

// IsYANGGoStruct ensures that Interface_Ipv6Neighbor implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Interface_Ipv6Neighbor) IsYANGGoStruct() {}

// GetIpv6Address retrieves the value of the leaf Ipv6Address from the Interface_Ipv6Neighbor
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Ipv6Address is set, it can
// safely use t.GetIpv6Address() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Ipv6Address == nil' before retrieving the leaf's value.
func (t *Interface_Ipv6Neighbor) GetIpv6Address() string {
	if t == nil || t.Ipv6Address == nil {
		return ""
	}
	return *t.Ipv6Address
}

// GetLinkLayerAddress retrieves the value of the leaf LinkLayerAddress from the Interface_Ipv6Neighbor
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if LinkLayerAddress is set, it can
// safely use t.GetLinkLayerAddress() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.LinkLayerAddress == nil' before retrieving the leaf's value.
func (t *Interface_Ipv6Neighbor) GetLinkLayerAddress() string {
	if t == nil || t.LinkLayerAddress == nil {
		return ""
	}
	return *t.LinkLayerAddress
}

// PopulateDefaults recursively populates unset leaf fields in the Interface_Ipv6Neighbor
// with default values as specified in the YANG schema, instantiating any nil
// container fields.
func (t *Interface_Ipv6Neighbor) PopulateDefaults() {
	if t == nil {
		return
	}
	ygot.BuildEmptyTree(t)
}

// ΛListKeyMap returns the keys of the Interface_Ipv6Neighbor struct, which is a YANG list entry.
func (t *Interface_Ipv6Neighbor) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Ipv6Address == nil {
		return nil, fmt.Errorf("nil value for key Ipv6Address")
	}

	return map[string]interface{}{
		"ipv6-address": *t.Ipv6Address,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Interface_Ipv6Neighbor) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Interface_Ipv6Neighbor"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Interface_Ipv6Neighbor) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Interface_Ipv6Neighbor) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of Interface_Ipv6Neighbor.
func (*Interface_Ipv6Neighbor) ΛBelongingModule() string {
	return "open-traffic-generator-discovery"
}

// IsisRouter represents the /open-traffic-generator-isis/isis-routers/isis-router YANG schema element.
type IsisRouter struct {
	Counters *IsisRouter_Counters `path:"state/counters" module:"open-traffic-generator-isis/open-traffic-generator-isis"`
	Name     *string              `path:"state/name|name" module:"open-traffic-generator-isis/open-traffic-generator-isis|open-traffic-generator-isis" shadow-path:"name" shadow-module:"open-traffic-generator-isis"`
}

// IsYANGGoStruct ensures that IsisRouter implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*IsisRouter) IsYANGGoStruct() {}

// GetOrCreateCounters retrieves the value of the Counters field
// or returns the existing field if it already exists.
func (t *IsisRouter) GetOrCreateCounters() *IsisRouter_Counters {
	if t.Counters != nil {
		return t.Counters
	}
	t.Counters = &IsisRouter_Counters{}
	return t.Counters
}

// GetCounters returns the value of the Counters struct pointer
// from IsisRouter. If the receiver or the field Counters is nil, nil
// is returned such that the Get* methods can be safely chained.
func (t *IsisRouter) GetCounters() *IsisRouter_Counters {
	if t != nil && t.Counters != nil {
		return t.Counters
	}
	return nil
}

// GetName retrieves the value of the leaf Name from the IsisRouter
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Name is set, it can
// safely use t.GetName() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Name == nil' before retrieving the leaf's value.
func (t *IsisRouter) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}
	return *t.Name
}

// PopulateDefaults recursively populates unset leaf fields in the IsisRouter
// with default values as specified in the YANG schema, instantiating any nil
// container fields.
func (t *IsisRouter) PopulateDefaults() {
	if t == nil {
		return
	}
	ygot.BuildEmptyTree(t)
	t.Counters.PopulateDefaults()
}

// ΛListKeyMap returns the keys of the IsisRouter struct, which is a YANG list entry.
func (t *IsisRouter) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Name == nil {
		return nil, fmt.Errorf("nil value for key Name")
	}

	return map[string]interface{}{
		"name": *t.Name,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *IsisRouter) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["IsisRouter"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *IsisRouter) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *IsisRouter) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of IsisRouter.
func (*IsisRouter) ΛBelongingModule() string {
	return "open-traffic-generator-isis"
}

// IsisRouter_Counters represents the /open-traffic-generator-isis/isis-routers/isis-router/state/counters YANG schema element.
type IsisRouter_Counters struct {
	Level1 *IsisRouter_Counters_Level1 `path:"level1" module:"open-traffic-generator-isis"`
	Level2 *IsisRouter_Counters_Level2 `path:"level2" module:"open-traffic-generator-isis"`
}

// IsYANGGoStruct ensures that IsisRouter_Counters implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*IsisRouter_Counters) IsYANGGoStruct() {}

// GetOrCreateLevel1 retrieves the value of the Level1 field
// or returns the existing field if it already exists.
func (t *IsisRouter_Counters) GetOrCreateLevel1() *IsisRouter_Counters_Level1 {
	if t.Level1 != nil {
		return t.Level1
	}
	t.Level1 = &IsisRouter_Counters_Level1{}
	return t.Level1
}

// GetOrCreateLevel2 retrieves the value of the Level2 field
// or returns the existing field if it already exists.
func (t *IsisRouter_Counters) GetOrCreateLevel2() *IsisRouter_Counters_Level2 {
	if t.Level2 != nil {
		return t.Level2
	}
	t.Level2 = &IsisRouter_Counters_Level2{}
	return t.Level2
}

// GetLevel1 returns the value of the Level1 struct pointer
// from IsisRouter_Counters. If the receiver or the field Level1 is nil, nil
// is returned such that the Get* methods can be safely chained.
func (t *IsisRouter_Counters) GetLevel1() *IsisRouter_Counters_Level1 {
	if t != nil && t.Level1 != nil {
		return t.Level1
	}
	return nil
}

// GetLevel2 returns the value of the Level2 struct pointer
// from IsisRouter_Counters. If the receiver or the field Level2 is nil, nil
// is returned such that the Get* methods can be safely chained.
func (t *IsisRouter_Counters) GetLevel2() *IsisRouter_Counters_Level2 {
	if t != nil && t.Level2 != nil {
		return t.Level2
	}
	return nil
}

// PopulateDefaults recursively populates unset leaf fields in the IsisRouter_Counters
// with default values as specified in the YANG schema, instantiating any nil
// container fields.
func (t *IsisRouter_Counters) PopulateDefaults() {
	if t == nil {
		return
	}
	ygot.BuildEmptyTree(t)
	t.Level1.PopulateDefaults()
	t.Level2.PopulateDefaults()
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *IsisRouter_Counters) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["IsisRouter_Counters"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *IsisRouter_Counters) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *IsisRouter_Counters) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of IsisRouter_Counters.
func (*IsisRouter_Counters) ΛBelongingModule() string {
	return "open-traffic-generator-isis"
}
