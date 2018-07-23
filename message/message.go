// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

import (
	"time"
)

// Type Message implements the IBasic interface to provide the minimum
// set of information a message transiting the pipeline.
type Message struct {

	// The payload of the message in a string, so an implementation
	// can return a simple string, a base64 string, or even a JSON object.
	payload string

	// The Time in UTC when the message entered the pipeline.
	enteredUtc time.Time

	// The Time in UTC when the message exited the pipeline.
	exitedUtc time.Time
}

// Method Payload returns the payload of the message in a string, so an implementation
// can return a simple string, a base64 string, or even a JSON object.
func (b Message) Payload() string {
	return b.payload
}

// Method EnteredUtc returns the Time in UTC when the message entered the pipeline.
func (b Message) EnteredUtc() time.Time {
	return b.enteredUtc
}

// Method SetEnteredUtc sets the Time in UTC when the message entered the pipeline.
func (b *Message) SetEnteredUtc(enteredUtc time.Time) {
	b.enteredUtc = enteredUtc
}

// Method ExitedUtc returns the Time in UTC when the message exited the pipeline.
func (b Message) ExitedUtc() time.Time {
	return b.exitedUtc
}

// Method SetExitedUtc sets the Time in UTC when the message exited the pipeline.
func (b *Message) SetExitedUtc(exitedUtc time.Time) {
	b.exitedUtc = exitedUtc
}

// Method Elapsed returns the amount of time the message took to transit the pipeline.
func (b Message) Elapsed() time.Duration {
	return b.enteredUtc.Sub(b.enteredUtc)
}

// Function NewBasic creates a new, initialized message.
// Parameter payload is the text to carry through the pipeline.
// Parameter enteredUtc is the Time in UTC when the message entered the pipeline.
// Parameter exitedUtc is the Time in UTC when the message exited the pipeline.
// Returns nil on error.
func New(payload string, enteredUtc time.Time, exitedUtc time.Time) *Message {

	msg := &Message{
		payload:    payload,
		enteredUtc: enteredUtc,
		exitedUtc:  exitedUtc,
	}

	return msg
}
