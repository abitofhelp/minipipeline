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

	// The Time in UTC when the message arrived in a stage of the pipeline
	//// stage in the pipeline.
	arrivedUtc time.Time

	// The Time in UTC when the message departed a stage in the pipeline.
	departedUtc time.Time
}

// Method Payload returns the payload of the message in a string, so an implementation
// can return a simple string, a base64 string, or even a JSON object.
func (b Basic) Payload() string {
	return b.payload
}

// Method SetPayload sets the payload in the instance.
func (b *Basic) SetPayload(payload string) {
	b.payload = payload
}

// Method ArrivedUtc returns the Time in UTC when the message arrived in a
// stage of the pipeline.
func (b Basic) ArrivedUtc() time.Time {
	return b.arrivedUtc
}

// Method SetArrivedUtc sets the arrived in UTC in the instance.
func (b *Basic) SetArrivedUtc(arrivedUtc time.Time) {
	// Todo: Add validation that it is in UTC.
	b.arrivedUtc = arrivedUtc
}

// Method DepartedUtc returns the Time in UTC when the message departed a
// stage in the pipeline.
func (b Basic) DepartedUtc() time.Time {
	return b.departedUtc
}

// Method SetDepartedUtc sets the departed in UTC in the instance.
func (b *Basic) SetDepartedUtc(departedUtc time.Time) {
	// Todo: Add validation that it is in UTC.
	b.departedUtc = departedUtc
}

// Function Elapsed determines the time that a message spent in a stage of a pipeline.
func (b Basic) Elapsed() time.Duration {
	return b.departedUtc.Sub(b.arrivedUtc)
}

// Function NewBasic creates a new, initialized message.
// Parameter payload is the text to carry through the pipeline.
// Returns nil on error.
func NewBasic(payload string) *Basic {
	return NewBasicDetails(payload, time.Time{}, time.Time{})
}

// Function NewBasic creates a new, initialized message.
// Parameter payload is the text to carry through the pipeline.
// Parameter arrivedUtc is the time in UTC when the message arrived in the stage.
// Parameter departedUtc is the time in UTC when the message departed from the stage.
// Returns nil on error.
func NewBasicDetails(payload string, arrivedUtc time.Time, departedUtc time.Time) *Basic {

	msg := &Basic{
		payload:     payload,
		arrivedUtc:  arrivedUtc,
		departedUtc: departedUtc,
	}

	return msg
}
