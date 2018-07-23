// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package pipeline provides the ability to construct and run a pipeline.
package pipeline

import (
	"fmt"
	"github.com/abitofhelp/minipipeline/message"
	"github.com/abitofhelp/minipipeline/stage"
)

type Builder struct {
	Pipeline
}

func (pb *Builder) Intake() *Builder {
	// Construct an instance of the Intake stage.
	step := stage.New(stage.Intake, nil, nil)
	pb.Pipeline.stages.PushBack(step)
	return pb
}

func (pb *Builder) Analysis() *Builder {
	// Construct an instance of the Analysis stage.
	step := stage.New(stage.Analysis, nil, nil)
	pb.Pipeline.stages.PushBack(step)
	return pb
}

func (pb *Builder) Transformation() *Builder {
	// Construct an instance of the Transformation stage.
	step := stage.New(stage.Transformation, nil, nil)
	pb.Pipeline.stages.PushBack(step)
	return pb
}

func (pb *Builder) Validation() *Builder {
	// Construct an instance of the Validation stage.
	step := stage.New(stage.Validation, nil, nil)
	pb.Pipeline.stages.PushBack(step)
	return pb
}

func (pb *Builder) Delivery() *Builder {
	// Construct an instance of the Delivery stage.
	step := stage.New(stage.Delivery, nil, nil)
	pb.Pipeline.stages.PushBack(step)
	return pb
}

// Function Build will create the pipeline from the stages that were
// assembled through the fluent interface.  It will create the channels
// that link the stages together for messages to flow through the
// pipeline.
// On success, it returns (pipeline instance, nil) or (nil, an error)
func (pb *Builder) Build() (*Pipeline, error) {

	var inputChan chan message.IMessage = nil
	var outputChan chan message.IMessage = nil

	// Create the channels for each stage in the configuration
	// that has been made through the fluent interface.
	// Traverse the stages from beginning to end, and
	// configure the channels for each stage.
	for e := pb.stages.Front(); e != nil; e = e.Next() {

		// Convert the list's node's value to IStage.
		step, ok := e.Value.(stage.Stage)
		if !ok {
			return nil, fmt.Errorf("failed to convert a list node value to type step.IStage: %v", e.Value)
		}

		// At the start of the list, we need to create and set
		// an input channel, since there isn't a node before it.
		if step.InputChannel() == nil {
			inputChan = make(chan message.IMessage, 100)
			step.SetInputChannel(inputChan)
		}

		// Set the output channel for this stage in the pipeline.
		outputChan = make(chan message.IMessage, 100)
		step.SetOutputChannel(outputChan)

		// If there is a step following this one,
		// the output channel from this step is the input
		// channel for the next step in the pipeline.
		next := e.Next()
		if next != nil {
			// Convert the next list's node's value to IStage.
			step, ok = next.Value.(stage.Stage)
			if !ok {
				return nil, fmt.Errorf("failed to convert a list node value to type step.IStage: %v", next.Value)
			}

			// Assign the outputChan of the current step to the inputChan of the next step.
			step.SetInputChannel(outputChan)
		}
	}

	return &pb.Pipeline, nil
}
