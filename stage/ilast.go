// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package stage provides the building blocks for making a pipeline.
package stage

// Interface ILast defines the methods that the last stage in a pipeline must implement.
type ILast interface {
	IStep
	IReceive
}
