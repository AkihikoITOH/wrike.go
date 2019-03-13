// Copyright 2016 Martin Angers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpcmw

import (
	"bytes"
	"fmt"
)

// Logger defines the Log method that is used to log structured
// data, in tuples of alternating keys/values. The go-kit logger
// satisfies this interface (github.com/go-kit/kit/log).
type Logger interface {
	Log(...interface{}) error
}

// PrintfLogger is an adapter to use Printf-style functions as a
// Logger in the middlewares that accept one. For example,
// the stdlib's log.Printf function can be used via this adapter.
type PrintfLogger func(string, ...interface{})

// Log implements Logger for the PrintfLogger function adapter.
func (fn PrintfLogger) Log(args ...interface{}) error {
	var buf bytes.Buffer
	for i := 0; i < len(args)-1; i += 2 {
		if i > 0 {
			buf.WriteByte(' ')
		}
		verb := "%v" //catch-all formatter
		val := args[i+1]

		// use quoted string formatter if possible
		switch val.(type) {
		case string:
			verb = "%q"
		case fmt.Stringer:
			verb = "%q"
		}

		fmt.Fprintf(&buf, "%s="+verb, args[i], args[i+1])
	}
	if buf.Len() > 0 {
		fn(buf.String())
	}
	return nil
}
