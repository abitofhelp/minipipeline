// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package step provides the building blocks for making a pipeline.
package step

import "sync"

// Interface IStep defines the methods that all steps in a pipeline must implement.
type IStep interface {
	ChannelIn() interface{}
	setChannelIn(ch chan interface{})
	ChannelOut() interface{}
	setChannelOut(ch chan interface{})
	Execute(wg *sync.WaitGroup) error
}
