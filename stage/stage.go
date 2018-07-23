// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package stage provides the building blocks for making a pipeline.
package stage

import (
	"fmt"
	"github.com/abitofhelp/minipipeline/message"
	"sync/atomic"
	"time"
)

type Stage struct {

	// Field stage indicates the kind of stage.
	stage Stages

	// Field receiveCounter is the number of messages that have been
	// received by this stage.
	receiveCounter uint64

	// Field sendCounter is the number of messages that have been
	// sent through the send channel to the next stage in the pipeline.
	sendCounter uint64

	// Field stageError contains an error if the stage has failed while processing a message.
	stageError error

	// Field previous points to the previous stage in the pipeline or nil if there isn't one.
	// It permits us to traverse backwards through the pipeline's stages.
	//previous *IStage

	// Field next points to the next stage in the pipeline or nil is there isn't one.
	// It permits us to traverse forward through the pipeline's stages.
	//next *IStage

	// Field inputChannel is the previous stage's channel (OutputChannel) to this stage,
	// or nil if there isn't a previous stage.
	inputChannel <-chan message.IStage

	// Field outputChannel is the InputChannel from the next stage, or nil if there isn't one.
	outputChannel chan<- message.IStage
}

// Method Message() returns the kind of stage.
func (s Stage) Stage() Stages {
	return s.stage
}

// Method ReceiveCounter returns the number of messages that have been
// processed as input, atomically.
func (s Stage) ReceiveCounter() uint64 {
	return atomic.LoadUint64(&s.receiveCounter)
}

// Method SendCounter returns the number of messages that have been
// sent through the send channel to the next stage in the pipeline, atomically.
func (s Stage) SendCounter() uint64 {
	return atomic.LoadUint64(&s.sendCounter)
}

// Method setSendCounter sets the value of the sendCounter, atomically.
func (s *Stage) setSendCounter(value uint64) {
	atomic.StoreUint64(&s.sendCounter, value)
}

// Method setReceiveCounter sets the value of the receiveCounter, atomically.
func (s *Stage) setReceiveCounter(value uint64) {
	atomic.StoreUint64(&s.receiveCounter, value)
}

// Method incSendCounter increments the value of the sendCounter, atomically.
func (s *Stage) incSendCounter() {
	atomic.AddUint64(&s.sendCounter, 1)
}

// Method incReceiveCounter increments the value of the receiveCounter, atomically.
func (s *Stage) incReceiveCounter() {
	atomic.AddUint64(&s.receiveCounter, 1)
}

// Method Message() returns an error if the stage has failed while processing a message;
// otherwise nil.
func (s Stage) StageError() error {
	return s.stageError
}

// Method Previous returns the previous stage in the pipeline or nil if there isn't one.
// It permits us to traverse backwards through the pipeline's stages.
//func (s *Stage) Previous() *IStage {
//	return s.previous
//}

// Method Next returns the next stage in the pipeline or nil is there isn't one.
// It permits us to traverse forward through the pipeline's stages.
//func (s *Stage) Next() *IStage {
//	return s.next
//}

// Method InputChannel returns the previous stage's channel (OutputChannel) to this stage,
//// or nil if there isn't a previous stage.
//func (s *Stage) InputChannel() <-chan message.IStage {
//	if s.Previous() != nil {
//		return s.Previous().OutputChannel()
//	}
//
//	return nil
//}
//
//// Method OutputChannel returns the InputChannel from the next stage, or nil if there isn't one.
//func (s *Stage) OutputChannel() chan<- message.IStage {
//	if s.next != nil {
//		//return s.Next().InputChannel()
//	}
//
//	return nil
//}

// Method Send sets sends a message through the send channel.
func (s *Stage) Send(payload string) error {
	fmt.Printf("F: %s\n", payload)

	msg := message.NewStage(s.Stage(), payload, time.Now().UTC())
	s.outputChannel <- msg
	s.incSendCounter()
	return nil
}

// Method Execute walks the source directory paths, and adds a file system path for each
// regular file that is found to the output channel.
func (s *Stage) Execute() error {
	// Todo: abstract method or use a closure?
	return nil //i.loadFilePathsToSendChannel()
}

// Function New creates a new instance of a stage.
func New(stage Stages, inputChannel <-chan message.IStage, outputChannel chan<- message.IStage) *Stage {

	// Create the first stage in the pipeline: IntakeStep.
	step := &Stage{
		stage:          stage,
		receiveCounter: 0,
		sendCounter:    0,
		stageError:     nil,
		inputChannel:   inputChannel,
		outputChannel:  outputChannel,
		//	previous:       previous,
		//next:           next,
	}

	return step
}
