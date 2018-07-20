// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package step provides the building blocks for making a pipeline.
package step

// Interface ISend defines the methods that must be implemented for sending through a channel
// to the next step in a pipeline.
type ISend interface {
	SendCounter() uint64
	Send() error
}
