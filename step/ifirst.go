// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package step provides the building blocks for making a pipeline.
package step

// Interface IFirst defines the methods that the first step in a pipeline must implement.
type IFirst interface {
	IReceive
	ISend
	Execute() error
}
