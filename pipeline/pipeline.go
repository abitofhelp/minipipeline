// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package pipeline provides the ability to construct and run a pipeline.
package pipeline

import (
	"container/list"
	"fmt"
	"github.com/abitofhelp/minipipeline/stage"
)

// The type Pipeline implements the IPipeline interface and to make it possible
// to easily construct and rearrange a pipeline using a fluent interface.  Basically,
// it is going to be a bi-directional linked list.
type Pipeline struct {
	// Field stages is a list of the stages from the fluent interface.
	// The list is in the order that will be used for the pipeline.
	stages list.List
}

// Function FirstStage locates the first stage in the pipeline.
// Returns the (stage instance, nil) on success; Otherwise, (nil, error).
func (p Pipeline) FirstStage() (first *stage.Stage, err error) {
	e := p.stages.Front()

	// Convert the list's node's value to IStage.
	step, ok := e.Value.(*stage.Stage)
	if !ok {
		return nil, fmt.Errorf("failed to convert a list node value to type step.IStage: %v", e.Value)
	}

	return step, nil
}

// Function LastStage locates the last stage in the pipeline.
// Returns the (stage instance, nil) on success; Otherwise, (nil, error).
func (p Pipeline) LastStage() (last *stage.Stage, err error) {
	e := p.stages.Back()

	// Convert the list's node's value to IStage.
	step, ok := e.Value.(*stage.Stage)
	if !ok {
		return nil, fmt.Errorf("failed to convert a list node value to type step.IStage: %v", e.Value)
	}

	return step, nil
}

// Function FindStage will locate a stage of interest in the pipeline.
// Parameter stageOfInterest is the kind of stage to seek in the pipeline.
// Returns (stage instance, nil) on success, otherwise (nil, error)
func (p Pipeline) FindStage(stageOfInterest stage.Stages) (found *stage.Stage, err error) {

	// Traverse the stages in the pipeline to find the one that has been requested.
	for e := p.stages.Front(); e != nil; e = e.Next() {

		// Convert the list's node's value to IStage.
		step, ok := e.Value.(*stage.Stage)
		if !ok {
			return nil, fmt.Errorf("failed to convert a list node value to type step.IStage: %v", e.Value)
		}

		if step.Stage() == stageOfInterest {
			return step, nil
		}
	}

	return nil, fmt.Errorf("failed to find a stage %s in the pipeline", stageOfInterest.String())
}

// Function New creates a new instance of a pipeline.
func New() (*Pipeline, error) {
	// Use the Builder to create the pipeline using a fluent interface.
	// The order of the stages is the same as in the pipeline.
	// Build will take care of setting up the channels between the stages.
	pb := &Builder{}
	pipeline, err := pb.
		Intake().
		Analysis().
		Transformation().
		Validation().
		Delivery().
		Build()

	return pipeline, err
}
