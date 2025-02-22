// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package cgotlstest

// #include <pthread.h>
// extern void setTLS(int);
// extern int getTLS();
import "C"

import (
	"runtime"
	"testing"
)

func testTLS(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if val := C.getTLS(); val != 0 {
		t.Fatalf("at start, C.getTLS() = %#x, want 0", val)
	}

	const keyVal = 0x1234
	C.setTLS(keyVal)
	if val := C.getTLS(); val != keyVal {
		t.Fatalf("at end, C.getTLS() = %#x, want %#x", val, keyVal)
	}
}
