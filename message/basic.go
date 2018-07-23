// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

import (
	"time"
)

// Type Basic implements the IBasic interface to provide the minimum
// set of information a message transiting the pipeline.
type Basic struct {

	// The payload of the message in a string, so an implementation
	// can return a simple string, a base64 string, or even a JSON object.
	payload string

	// The Time in UTC when the message departed a stage in the pipeline.
	departedUtc time.Time
}

// Method Payload returns the payload of the message in a string, so an implementation
// can return a simple string, a base64 string, or even a JSON object.
func (b Basic) Payload() string {
	return b.payload
}

// Method DepartedUtc returns the Time in UTC when the message departed a
// stage in the pipeline.
func (b Basic) DepartedUtc() time.Time {
	return b.departedUtc
}

// Function NewBasic creates a new, initialized message.
// Parameter payload is the text to carry through the pipeline.
// Parameter arrivedUtc is the time in UTC when the message arrived in the stage.
// Parameter departedUtc is the time in UTC when the message departed from the stage.
// Returns nil on error.
func NewBasic(payload string, departedUtc time.Time) *Basic {

	msg := &Basic{
		payload:     payload,
		departedUtc: departedUtc,
	}

	return msg
}
