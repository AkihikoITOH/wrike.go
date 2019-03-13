// Copyright 2016 Martin Angers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package httpcmw supports creating middleware chains for
// HTTP clients. It abstracts the *http.Client into a Doer
// interface and wraps middleware around the Do call.
package httpcmw

import "net/http"

// Doer defines the method required for an HTTP client to
// execute requests. An *http.Client satisfies this interface.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// DoerFunc is a function signature that satisfies the Doer
// interface.
type DoerFunc func(*http.Request) (*http.Response, error)

// Do satisfies the Doer interface for a DoerFunc by calling the
// function.
func (fn DoerFunc) Do(req *http.Request) (*http.Response, error) {
	return fn(req)
}

// Wrapper defines the Wrap method required to build a middleware-style
// chain of calls.
type Wrapper interface {
	Wrap(Doer) Doer
}

// WrapperFunc is a function type that satisfies the Wrapper interface.
type WrapperFunc func(Doer) Doer

// Wrap satisfies the Wrapper interface for a WrapperFunc by calling the function.
func (fn WrapperFunc) Wrap(d Doer) Doer {
	return fn(d)
}

// Wrap wraps the Doer d with the provided middleware ws.
// It returns a Doer that will call
// ws[0] -> ws[1] -> ... -> ws[n-1] -> d.
// Each Wrapper may stop the chain of calls by not
// calling the next Doer in the chain.
func Wrap(d Doer, ws ...Wrapper) Doer {
	for i := len(ws) - 1; i >= 0; i-- {
		d = ws[i].Wrap(d)
	}
	return d
}
