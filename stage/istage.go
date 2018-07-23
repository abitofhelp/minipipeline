// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package stage provides the building blocks for making a pipeline.
package stage

import (
	"github.com/abitofhelp/minipipeline/message"
)

// Interface IStage defines the methods that all stages in a pipeline must implement.
type IStage interface {
	// Method Message returns a message that will be processed by this stage.
	Message() message.IStage

	// Method ReceiveCounter returns the number of messages that have been
	// received by this stage.
	ReceiveCounter() uint64

	// Method SendCounter returns the number of messages that have been
	// sent through the send channel to the next stage in the pipeline.
	SendCounter() uint64

	// Method StageError returns an error if the stage has failed while processing a message.
	StageError() error

	// Method Previous returns the previous stage in the pipeline or nil if there isn't one.
	// It permits us to traverse backwards through the pipeline's stages.
	//Previous() *IStage

	// Method Next returns the next stage in the pipeline or nil is there isn't one.
	// It permits us to traverse forward through the pipeline's stages.
	//Next() *IStage

	// Method InputChannel returns the previous stage's channel (OutputChannel) to this stage,
	// or nil if there isn't a previous stage.
	InputChannel() <-chan *message.IStage

	// Method OutputChannel returns the InputChannel from the next stage, or nil if there isn't one.
	OutputChannel() chan<- *message.IStage

	// Method Send sends a message through the send channel to the next stage in the pipeline.
	Send(payload string)

	// Method Execute performs whatever action is required in the stage.
	Execute()
}
