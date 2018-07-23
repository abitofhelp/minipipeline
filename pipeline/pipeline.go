// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package pipeline provides the ability to construct and run a pipeline.
package pipeline

import (
	"container/list"
	"github.com/abitofhelp/minipipeline/message"
	"github.com/abitofhelp/minipipeline/stage"
)

type PipelineBuilder struct {
	Pipeline
}

func (pb *PipelineBuilder) Intake(inputChannel <-chan message.IStage, outputChannel chan<- message.IStage) *PipelineBuilder {
	// Construct an instance of the Intake stage.
	stage := stage.New(stage.Intake, inputChannel, outputChannel)
	pb.Pipeline.stages.PushFront(stage)
	return pb
}

func (pb *PipelineBuilder) Analysis(inputChannel <-chan message.IStage, outputChannel chan<- message.IStage) *PipelineBuilder {
	// Construct an instance of the Analysis stage.
	stage := stage.New(stage.Analysis, inputChannel, outputChannel)
	pb.Pipeline.stages.PushFront(stage)
	return pb
}

func (pb *PipelineBuilder) Transformation(inputChannel <-chan message.IStage, outputChannel chan<- message.IStage) *PipelineBuilder {
	// Construct an instance of the Transformation stage.
	stage := stage.New(stage.Transformation, inputChannel, outputChannel)
	pb.Pipeline.stages.PushFront(stage)
	return pb
}

func (pb *PipelineBuilder) Validation(inputChannel <-chan message.IStage, outputChannel chan<- message.IStage) *PipelineBuilder {
	// Construct an instance of the Validation stage.
	stage := stage.New(stage.Validation, inputChannel, outputChannel)
	pb.Pipeline.stages.PushFront(stage)
	return pb
}

func (pb *PipelineBuilder) Delivery(inputChannel <-chan message.IStage, outputChannel chan<- message.IStage) *PipelineBuilder {
	// Construct an instance of the Delivery stage.
	stage := stage.New(stage.Delivery, inputChannel, outputChannel)
	pb.Pipeline.stages.PushFront(stage)
	return pb
}

func (pb *PipelineBuilder) Build() Pipeline {
	return pb.Pipeline
}

// The type Pipeline implements the IPipeline interface and to make it possible
// to easily construct and rearrange a pipeline using a fluent interface.  Basically,
// it is going to be a bi-directional linked list.
type Pipeline struct {
	stages list.List
}

// Function New creates a new instance of a pipeline.
func New() Pipeline {

	intakeInput := make(<-chan message.IStage, 100)
	intakeOutput := make(chan<- message.IStage, 100)

	// Create the first stage in the pipeline: IntakeStep.
	pb := &PipelineBuilder{}
	pipeline := pb.
		Intake(intakeInput, intakeOutput).
		Analysis(nil, nil).
		Transformation(nil, nil).
		Validation(nil, nil).
		Delivery(nil, nil).
		Build()

	return pipeline
}
