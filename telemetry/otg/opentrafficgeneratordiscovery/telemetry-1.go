package opentrafficgeneratordiscovery

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry/otg"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Ipv4Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Ipv4Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_Ipv4Neighbor_LinkLayerAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ipv4Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ipv4Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ipv4Neighbor_LinkLayerAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address with a ONCE subscription.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ipv4Neighbor_LinkLayerAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Ipv4Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ipv4Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ipv4Neighbor_LinkLayerAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Ipv4Neighbor_LinkLayerAddressPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address to the batch object.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ipv4Neighbor_LinkLayerAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_Ipv4Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ipv4Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ipv4Neighbor", structs[pre], queryPath, true, false)
			qv := convertInterface_Ipv4Neighbor_LinkLayerAddressPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Ipv4Neighbor_LinkLayerAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address to the batch object.
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ipv4Neighbor_LinkLayerAddressPath extracts the value of the leaf LinkLayerAddress from its parent oc.Interface_Ipv4Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Ipv4Neighbor_LinkLayerAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ipv4Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LinkLayerAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ipv6NeighborPath) Lookup(t testing.TB) *oc.QualifiedInterface_Ipv6Neighbor {
	t.Helper()
	goStruct := &oc.Interface_Ipv6Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Ipv6Neighbor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Ipv6Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ipv6NeighborPath) Get(t testing.TB) *oc.Interface_Ipv6Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ipv6NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Ipv6Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Ipv6Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ipv6Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ipv6Neighbor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Ipv6Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor with a ONCE subscription.
func (n *Interface_Ipv6NeighborPathAny) Get(t testing.TB) []*oc.Interface_Ipv6Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Ipv6Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ipv6NeighborPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Ipv6Neighbor {
	t.Helper()
	c := &oc.CollectionInterface_Ipv6Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Ipv6Neighbor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Ipv6Neighbor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Ipv6Neighbor)))
		return false
	})
	return c
}

func watch_Interface_Ipv6NeighborPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Ipv6Neighbor) bool) *oc.Interface_Ipv6NeighborWatcher {
	t.Helper()
	w := &oc.Interface_Ipv6NeighborWatcher{}
	gs := &oc.Interface_Ipv6Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ipv6Neighbor", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Ipv6Neighbor{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Ipv6Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ipv6NeighborPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Ipv6Neighbor) bool) *oc.Interface_Ipv6NeighborWatcher {
	t.Helper()
	return watch_Interface_Ipv6NeighborPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ipv6NeighborPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Ipv6Neighbor) *oc.QualifiedInterface_Ipv6Neighbor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Ipv6Neighbor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor to the batch object.
func (n *Interface_Ipv6NeighborPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ipv6NeighborPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Ipv6Neighbor {
	t.Helper()
	c := &oc.CollectionInterface_Ipv6Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Ipv6Neighbor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ipv6NeighborPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Ipv6Neighbor) bool) *oc.Interface_Ipv6NeighborWatcher {
	t.Helper()
	w := &oc.Interface_Ipv6NeighborWatcher{}
	structs := map[string]*oc.Interface_Ipv6Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ipv6Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ipv6Neighbor", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Ipv6Neighbor{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Ipv6Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ipv6NeighborPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Ipv6Neighbor) bool) *oc.Interface_Ipv6NeighborWatcher {
	t.Helper()
	return watch_Interface_Ipv6NeighborPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor to the batch object.
func (n *Interface_Ipv6NeighborPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Ipv6Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Ipv6Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_Ipv6Neighbor_Ipv6AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ipv6Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ipv6Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ipv6Neighbor_Ipv6AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address with a ONCE subscription.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ipv6Neighbor_Ipv6AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Ipv6Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ipv6Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ipv6Neighbor_Ipv6AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Ipv6Neighbor_Ipv6AddressPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address to the batch object.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ipv6Neighbor_Ipv6AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_Ipv6Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ipv6Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ipv6Neighbor", structs[pre], queryPath, true, false)
			qv := convertInterface_Ipv6Neighbor_Ipv6AddressPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Ipv6Neighbor_Ipv6AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address to the batch object.
func (n *Interface_Ipv6Neighbor_Ipv6AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ipv6Neighbor_Ipv6AddressPath extracts the value of the leaf Ipv6Address from its parent oc.Interface_Ipv6Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Ipv6Neighbor_Ipv6AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ipv6Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Ipv6Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
