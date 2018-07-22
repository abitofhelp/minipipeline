// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package stage provides the building blocks for making a pipeline.
package stage

// Interface ISend defines the methods that must be implemented for sending through a channel
// to the next stage in a pipeline.
type ISend interface {
	SendCounter() uint64
	Send() error
}
