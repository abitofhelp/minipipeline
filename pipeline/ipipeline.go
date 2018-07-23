// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package pipeline provides the ability to construct and run a pipeline.
package pipeline

// The type IPipeline is an interface consisting of each of the possible stages
// in the pipeline.  These stages will be part of a fluent interface, so it is
// a simple matter to change the order of the stages, remove some, or introduce
// new ones.
type IPipeline interface {
	Intake()
	Analysis() *IPipeline
	Transformation() *IPipeline
	Validation() *IPipeline
	Persistence() *IPipeline

	Builder() *IPipeline
}
