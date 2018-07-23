// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

import (
	"time"
)

// Interface IBasic defines the minimum methods that must be implemented for
// a message flowing through the pipeline.
type IBasic interface {
	// Method Payload returns the payload of the message in a string, so an implementation
	// can return a simple string, a base64 string, or even a JSON object.
	Payload() string

	// Method DepartedUtc returns the Time in UTC when the message departed a
	// stage in the pipeline.
	DepartedUtc() time.Time
}

// Function ContainsNil checks any number of items, which are in the IBasic family, for nil.
// Returns true if any of the items are nil; Otherwise, false.
func ContainsNil(items ...IBasic) bool {
	for _, item := range items {
		if item == nil {
			return true
		}
	}
	return false
}
