// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package virtual implements a machine that isn't. (Work In Progress)
package virtual

import (
	"io"

	"github.com/cznic/xc"
)

var (
	// Testing amends things for tests.
	Testing bool

	dict = xc.Dict
)

// Exec runs the program in b and returns its exit status or an error, if any.
func Exec(b *Binary, args []string, stdin io.Reader, stdout, stderr io.Writer, heapSize, stackSize int) (exitStatus int, err error) {
	m, err := newMachine(b.Data, b.Text, b.BSS, heapSize, stdin, stdout, stderr)
	if err != nil {
		return 0, err
	}

	m.lines = b.Lines
	m.functions = b.Functions

	defer func() {
		if e := m.close(); e != nil && err == nil {
			err = e
		}
	}()

	t, err := m.newThread(stackSize)
	if err != nil {
		return 0, err
	}

	argv := make([]uintptr, len(args)+1)
	for i, v := range args {
		argv[i] = m.CString(v)
	}
	pargv := m.malloc(len(argv) * ptrSize)
	for i, v := range argv {
		t.writePtr(pargv+uintptr(i*ptrSize), v)
	}

	// void _start(int args, char **argv);
	t.rp = t.sp
	t.sp -= i32StackSz
	t.writeI32(t.sp, int32(len(args))) // argc
	t.sp -= ptrStackSz
	t.writePtr(t.sp, pargv) // argv
	t.sp -= ptrStackSz
	t.writePtr(t.sp, 0xcafebabe) // return address, not used
	return t.run(b.Code)
}