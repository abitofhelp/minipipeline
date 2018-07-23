// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package stage provides the building blocks for making a pipeline.
package stage

import (
	"fmt"
	"github.com/abitofhelp/minipipeline/message"
	"sync/atomic"
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

	// Field inputChannel is the previous stage's channel (OutputChannel) to this stage,
	// or nil if there isn't a previous stage.
	inputChannel <-chan message.IMessage

	// Field outputChannel is the InputChannel from the next stage, or nil if there isn't one.
	outputChannel chan<- message.IMessage
}

// Method Stage() returns the kind of stage.
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

// Method InputChannel returns the input channel for this stage.
func (s Stage) InputChannel() <-chan message.IMessage {
	return s.inputChannel
}

// Method SetInputChannel sets the input channel for this stage.
func (s *Stage) SetInputChannel(inputChannel <-chan message.IMessage) {
	s.inputChannel = inputChannel
}

// Method OutputChannel returns the output channel for this stage.
func (s Stage) OutputChannel() chan<- message.IMessage {
	return s.outputChannel
}

// Method SetOutputChannel sets the output channel for this stage.
func (s *Stage) SetOutputChannel(outputChannel chan<- message.IMessage) {
	s.outputChannel = outputChannel
}

// Method Send sends a message through the output channel.
func (s *Stage) Send(message message.IMessage) {
	fmt.Printf("OUT: %s: %s\n", s.stage.String(), message.Payload())

	s.outputChannel <- message
	s.incSendCounter()
}

// Method Receive receives a message from the input channel.
func (s *Stage) Receive() message.IMessage {
	message := <-s.inputChannel
	s.incReceiveCounter()

	fmt.Printf("IN: %s: %s\n", s.stage.String(), message.Payload())

	return message
}

// Function New creates a new instance of a stage.
func New(stage Stages, inputChannel <-chan message.IMessage, outputChannel chan<- message.IMessage) Stage {

	// Create the first stage in the pipeline: IntakeStep.
	step := Stage{
		stage:          stage,
		receiveCounter: 0,
		sendCounter:    0,
		inputChannel:   inputChannel,
		outputChannel:  outputChannel,
	}

	return step
}
