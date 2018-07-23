// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

import (
	"time"
)

// Interface IMessage defines the methods that must be implemented for
// a message flowing through the pipeline.
type IMessage interface {
	// Method Payload returns the payload of the message in a string, so an implementation
	// can return a simple string, a base64 string, or even a JSON object.
	Payload() string

	// Method EnteredUtc returns the Time in UTC when the message entered the pipeline.
	EnteredUtc() time.Time

	// Method SetEnteredUtc sets the Time in UTC when the message entered the pipeline.
	SetEnteredUtc(exitedUtc time.Time)

	// Method ExitedUtc returns the Time in UTC when the message exited the pipeline.
	ExitedUtc() time.Time

	// Method SetExitedUtc sets the Time in UTC when the message exited the pipeline.
	SetExitedUtc(exitedUtc time.Time)

	// Method Elapsed returns the amount of time the message took to transit the pipeline.
	Elapsed() time.Duration
}
