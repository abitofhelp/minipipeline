// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package pipeline provides the ability to construct and run a pipeline.
package pipeline

import "github.com/abitofhelp/minipipeline/stage"

// The type IPipeline is an interface consisting of each of the possible stages
// in the pipeline.  These stages will be part of a fluent interface, so it is
// a simple matter to change the order of the stages, remove some, or introduce
// new ones.
type IPipeline interface {
	// Function FirstStage locates the first stage in the pipeline.
	// Returns the (stage instance, nil) on success; Otherwise, (nil, error).
	FirstStage() (first *stage.Stage, err error)

	// Function LastStage locates the last stage in the pipeline.
	// Returns the (stage instance, nil) on success; Otherwise, (nil, error).
	LastStage() (last *stage.Stage, err error)

	// Function FindStage will locate a stage of interest in the pipeline.
	// Parameter stageOfInterest is the kind of stage to seek in the pipeline.
	// Returns (stage instance, nil) on success, otherwise (nil, error)
	FindStage(stageOfInterest stage.Stages) (found *stage.Stage, err error)
}
