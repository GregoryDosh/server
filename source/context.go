// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package source

import (
	"context"
)

// key defines the key type for storing
// the source Service in the context.
const key = "source"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the source Service
// associated with this context.
func FromContext(c context.Context) Service {
	// get source value from context
	v := c.Value(key)
	if v == nil {
		return nil
	}

	// cast source value to expected Service type
	s, ok := v.(Service)
	if !ok {
		return nil
	}

	return s
}

// ToContext adds the source Service to this
// context if it supports the Setter interface.
func ToContext(c Setter, s Service) {
	c.Set(key, s)
}
