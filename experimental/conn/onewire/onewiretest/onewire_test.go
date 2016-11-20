// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package onewiretest

import (
	"testing"

	"github.com/google/periph/experimental/conn/onewire"
)

// TestDevTx tests the onewire.Dev implementation using the Playback bus impl.
func TestDevTx(t *testing.T) {
	p := Playback{
		Ops: []IO{
			{
				Write: []byte{10, 11},
				Read:  []byte{12, 13},
				Pull:  onewire.WeakPullup,
			},
			{
				Write: []byte{20, 21},
				Read:  []byte{22, 23},
				Pull:  onewire.StrongPullup,
			},
		},
	}
	d := onewire.Dev{Bus: &p}
	buf := []byte{0, 0}

	// Test Tx.
	err := d.Tx([]byte{10, 11}, buf)
	if err != nil {
		t.Fatal(err)
	}
	if buf[0] != 12 || buf[1] != 13 {
		t.Errorf("expected 12 & 13, got %d %d", buf[0], buf[1])
	}

	// Test TxPower.
	err = d.TxPower([]byte{20, 21}, buf)
	if err != nil {
		t.Fatal(err)
	}
	if buf[0] != 22 || buf[1] != 23 {
		t.Errorf("expected 12 & 13, got %d %d", buf[0], buf[1])
	}
}
