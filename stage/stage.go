// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package stage provides the building blocks for making a pipeline.
package stage

import (
	"fmt"
	"github.com/abitofhelp/minipipeline/message"
	"sync"
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
	inputChannel chan message.IMessage

	// Field outputChannel is the InputChannel from the next stage, or nil if there isn't one.
	outputChannel chan message.IMessage
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
func (s Stage) InputChannel() chan message.IMessage {
	return s.inputChannel
}

// Method SetInputChannel sets the input channel for this stage.
func (s *Stage) SetInputChannel(inputChannel chan message.IMessage) {
	s.inputChannel = inputChannel
}

// Method OutputChannel returns the output channel for this stage.
func (s Stage) OutputChannel() chan message.IMessage {
	return s.outputChannel
}

// Method SetOutputChannel sets the output channel for this stage.
func (s *Stage) SetOutputChannel(outputChannel chan message.IMessage) {
	s.outputChannel = outputChannel
}

// Method Send sends a message through the output channel.
func (s *Stage) Send(message message.IMessage) {
	fmt.Printf("OUT: %s: %s\n", s.stage.String(), message.Payload())

	s.outputChannel <- message
	s.incSendCounter()
}

// Method Receive receives a message from the input channel.
func (s *Stage) Receive() {

	// Field wg is the consumer's WaitGroup, which detects when all of the
	// goroutines that were launched have completed.
	var wg sync.WaitGroup

	// The consumer's wait group will block at wg.Wait(), which will be invoked just
	// before exiting from the function.  It will block until wg's internal counter is zero,
	// which happens when all of the goroutines that were launched have completed.
	defer wg.Wait()

	for msg := range s.InputChannel() {
		// Increment the consumer's WaitGroup counter for each goroutine that is launched.
		wg.Add(1)
		go func(path message.IMessage) {
			// Decrement the consumers's WaitGroup counter after each goroutine completes its work.
			defer wg.Done()

			s.incReceiveCounter()
			fmt.Printf("IN: %s: %s\n", s.stage.String(), msg.Payload())

			// Todo
			// For now, we simply pass the message along to the next stage...
			s.Send(msg)

		}(msg)
		// The closure is only bound to that one variable, 'msg'.  There is a very good
		// chance that not adding 'msg' as a parameter to the closure, will result in seeing
		// the last element printed for every iteration instead of each value in sequence.
		// This is due to the high probability that goroutines will execute after the loop.
		//
		// By adding 'msg' as a parameter to the closure, 'msg' is evaluated at each iteration
		// and placed on the stack for the goroutine, so each slice element is available
		// to the goroutine when it is eventually executed.
	}
}

// Function New creates a new instance of a stage.
func New(stage Stages, inputChannel chan message.IMessage, outputChannel chan message.IMessage) *Stage {

	// Create the first stage in the pipeline: IntakeStep.
	step := &Stage{
		stage:          stage,
		receiveCounter: 0,
		sendCounter:    0,
		inputChannel:   inputChannel,
		outputChannel:  outputChannel,
	}

	return step
}
