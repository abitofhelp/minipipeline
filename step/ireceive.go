// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package step provides the building blocks for making a pipeline.
package step

// Interface IReceive defines the methods that must be implemented for receiving through a channel
// from the previous step in a pipeline.
type IReceive interface {
	ReceiveCounter() uint64
	Receive() error
}
