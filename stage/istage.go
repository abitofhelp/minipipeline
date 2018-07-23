// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package stage provides the building blocks for making a pipeline.
package stage

import (
	"github.com/abitofhelp/minipipeline/message"
)

// Interface IStage defines the methods that all stages in a pipeline must implement.
type IStage interface {
	// Method Stage() returns the kind of stage.
	Stage() Stages

	// Method ReceiveCounter returns the number of messages that have been
	// received by this stage.
	ReceiveCounter() uint64

	// Method SendCounter returns the number of messages that have been
	// sent through the send channel to the next stage in the pipeline.
	SendCounter() uint64

	// Method InputChannel returns the input channel for this stage.
	InputChannel() <-chan message.IMessage

	// Method SetInputChannel sets the input channel for this stage.
	SetInputChannel(inputChannel <-chan message.IMessage)

	// Method OutputChannel returns the output channel for this stage.
	OutputChannel() chan<- message.IMessage

	// Method SetOutputChannel sets the output channel for this stage.
	SetOutputChannel(outputChannel chan<- message.IMessage)

	// Method Send sends a message through the output channel.
	Send(message message.IMessage)

	// Method Receive receives a message from the input channel.
	Receive() message.IMessage

	// Method Execute performs whatever action is required in the stage.
	Execute()
}
