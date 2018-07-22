// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package stage provides the building blocks for making a pipeline.
package stage

// Interface IFirst defines the methods that the first stage in a pipeline must implement.
type IFirst interface {
	IStep
	IReceive
	ISend
}
