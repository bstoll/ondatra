// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ondatra

import (
	"golang.org/x/net/context"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"flag"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"

	opb "github.com/openconfig/ondatra/proto"
)

func TestRunTests(t *testing.T) {
	emptyTB, err := os.CreateTemp(t.TempDir(), "*.textproto")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	if err := emptyTB.Close(); err != nil {
		t.Errorf("Failed to close temp file: %v", err)
	}
	flag.Set("testbed", emptyTB.Name())

	tests := []struct {
		desc                   string
		interrupt              bool
		reserveErr, releaseErr error
		wantErr                string
	}{{
		desc: "success",
	}, {
		desc:       "reserve error",
		reserveErr: errors.New("reserve error"),
		wantErr:    "reserve error",
	}, {
		desc:       "release error",
		releaseErr: errors.New("release error"),
		wantErr:    "release error",
	}, {
		desc:      "interrupted",
		interrupt: true,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			fakeBind := fakebind.Setup()
			fakeBind.ReserveFn = func(context.Context, *opb.Testbed, time.Duration, time.Duration, map[string]string) (*binding.Reservation, error) {
				return new(binding.Reservation), test.reserveErr
			}
			releaseCh := make(chan struct{}, 1)
			defer close(releaseCh)
			fakeBind.ReleaseFn = func(context.Context) error {
				releaseCh <- struct{}{}
				return test.releaseErr
			}
			runFn := func() int {
				if test.interrupt {
					sigc <- os.Interrupt
					// Wait for the release immediately to confirm it is triggered from
					// the interrupt and not from the  release at the end of runTests.
					<-releaseCh
				}
				return 0
			}

			gotErr := runTests(runFn, func() (binding.Binding, error) { return fakeBind, nil })
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("runTests() got err %v, want %q", gotErr, test.wantErr)
			}
			// If there was no error reserving and the test wasn't interrupted
			// wait for the deferred release at the end of runTests.
			if test.reserveErr == nil && !test.interrupt {
				<-releaseCh
			}
		})
	}
}

func TestGetters(t *testing.T) {
	fakebind.Setup().WithReservation(fakeRes)

	t.Run("DUT", func(t *testing.T) {
		id := "dut_arista"
		d := DUT(t, id)
		if got, want := d.ID(), id; got != want {
			t.Errorf("DUT id = %q, want %q", got, want)
		}
		if got, want := d.Name(), "pf01.xxx01"; got != want {
			t.Errorf("DUT name = %q, want %q", got, want)
		}
		if got, want := d.Vendor(), ARISTA; got != want {
			t.Errorf("DUT vendor = %v, want %v", got, want)
		}
		if got, want := d.Model(), "aristaModel"; got != want {
			t.Errorf("DUT model = %v, want %v", got, want)
		}
		if got, want := d.Version(), "aristaVersion"; got != want {
			t.Errorf("DUT version = %v, want %v", got, want)
		}
	})

	t.Run("DUTs", func(t *testing.T) {
		id := "dut_cisco"
		duts := DUTs(t)
		d, ok := duts[id]
		if !ok {
			t.Errorf("DUT id %q not present in DUTs map %q, must be present", id, duts)
		}
		if got, want := d.ID(), id; got != want {
			t.Errorf("DUT id = %q, want %q", got, want)
		}
	})

	t.Run("DUT failure", func(t *testing.T) {
		id := "gaga"
		got := testt.ExpectFatal(t, func(t testing.TB) {
			DUT(t, id)
		})
		if !strings.Contains(got, id) {
			t.Errorf("DUT(%q) failed with message %q, want %q", id, got, id)
		}
	})

	t.Run("ATE", func(t *testing.T) {
		id := "ate_ixia"
		a := ATE(t, id)
		if got, want := a.ID(), id; got != want {
			t.Errorf("ATE id = %q, want %q", got, want)
		}
		if got, want := a.Name(), "ix1"; got != want {
			t.Errorf("ATE name = %q, want %q", got, want)
		}
	})

	t.Run("ATEs", func(t *testing.T) {
		id := "ate_ixia"
		ates := ATEs(t)
		a, ok := ates[id]
		if !ok {
			t.Errorf("ATE id %q not present in ATEs map %q, must be present", id, ates)
		}
		if got, want := a.ID(), id; got != want {
			t.Errorf("ATE id = %q, want %q", got, want)
		}
	})

	t.Run("ATE failure", func(t *testing.T) {
		id := "gaga"
		got := testt.ExpectFatal(t, func(t testing.TB) {
			ATE(t, id)
		})
		if !strings.Contains(got, id) {
			t.Errorf("ATE(%q) failed with message %q, want %q", id, got, id)
		}
	})

	t.Run("Port", func(t *testing.T) {
		did, pid := "dut_juniper", "port1"
		p := DUT(t, did).Port(t, pid)
		if got, want := p.ID(), pid; got != want {
			t.Errorf("port id = %q, want %q", got, want)
		}
		if got, want := p.Device().ID(), did; got != want {
			t.Errorf("port device id = %q, want %q", got, want)
		}
		if got, want := p.Speed(), Speed10Gb; got != want {
			t.Errorf("port speed = %d, want %d", got, want)
		}
		if got, want := p.CardModel(), "EX9200-40T"; got != want {
			t.Errorf("card model = %q, want %q", got, want)
		}
		if got, want := p.PMD(), PMD100GBASEFR; got != want {
			t.Errorf("pmd = %q, want %q", got, want)
		}
	})

	t.Run("Port failure", func(t *testing.T) {
		d := DUT(t, "dut_arista")
		pid := "gaga"
		got := testt.ExpectFatal(t, func(t testing.TB) {
			d.Port(t, pid)
		})
		if !strings.Contains(got, pid) {
			t.Errorf("Port(%q) failed with message %q, want %q", pid, got, pid)
		}
	})

	t.Run("Get Ixia Port", func(t *testing.T) {
		iid, pid := "ate_ixia", "port2"
		p := ATE(t, iid).Port(t, pid)
		if got, want := p.ID(), pid; got != want {
			t.Errorf("port id = %q, want %q", got, want)
		}
		if got, want := p.Device().ID(), iid; got != want {
			t.Errorf("port device id = %q, want %q", got, want)
		}
		if got, want := p.Speed(), Speed100Gb; got != want {
			t.Errorf("port speed = %d, want %d", got, want)
		}
		if got, want := p.CardModel(), "NOVUS"; got != want {
			t.Errorf("card model = %q, want %q", got, want)
		}
		if got, want := p.PMD(), PMD100GBASEFR; got != want {
			t.Errorf("pmd = %q, want %q", got, want)
		}
	})
}
