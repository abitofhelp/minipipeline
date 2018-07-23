// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package intake implements a pipeline stage in which items in an array are ingested and
// passed along to the next stage in the pipeline.
package analysis

//
//import (
//	"fmt"
//			//. "path/filepath"
//	"github.com/abitofhelp/minipipeline/message"
//	msg "github.com/abitofhelp/minipipeline/message"
//		"sync/atomic"
//	"time"
//)
//
//const (
//	// The length of the send channel.
//	KSendChannelSize = 1000
//)
//
//// Type Analysis implements the IStaged interface
//// and determines the checksum for the file path in the message.
//type Analysis struct {
//
//	// Field message is the pipeline message for this stage.
//	message message.Analysis
//
//	// Field receiveCounter is the number of messages that will be
//	// sent through the send channel.
//	receiveCounter uint64
//
//	// Field sendCounter is the number of messages that have been
//	// sent through send channel.
//	sendCounter uint64
//
//	// Field sendChannel contains the messages being passed along to
//	// the next stage in the pipeline.
//	sendChannel chan<- message.Analysis
//}
//
//// Method Message() returns a message that will be processed by the pipeline for this stage.
//func (i Analysis) Message() message.Analysis {
//	return i.message
//}
//
//// Method ReceiveCounter returns the number of messages that have been
//// processed as input, atomically.
//func (i Analysis) ReceiveCounter() uint64 {
//	return atomic.LoadUint64(&i.receiveCounter)
//}
//
//// Method SendCounter returns the number of messages that have been
//// sent through the send channel to the next stage in the pipeline, atomically.
//func (i Analysis) SendCounter() uint64 {
//	return atomic.LoadUint64(&i.sendCounter)
//}
//
//// Method setSendCounter sets the value of the sendCounter, atomically.
//func (i *Analysis) setSendCounter(value uint64) {
//	atomic.StoreUint64(&i.sendCounter, value)
//}
//
//// Method setReceiveCounter sets the value of the receiveCounter, atomically.
//func (i *Analysis) setReceiveCounter(value uint64) {
//	atomic.StoreUint64(&i.receiveCounter, value)
//}
//
//// Method incSendCounter increments the value of the sendCounter, atomically.
//func (i *Analysis) incSendCounter() {
//	atomic.AddUint64(&i.sendCounter, 1)
//}
//
//// Method incReceiveCounter increments the value of the receiveCounter, atomically.
//func (i *Analysis) incReceiveCounter() {
//	atomic.AddUint64(&i.receiveCounter, 1)
//}
//
//// Method Send sets sends a message through the send channel.
//func (i *Analysis) Send(payload string) error {
//	fmt.Printf("F: %s\n", payload)
//
//	// Ready to pass the intake message to the next stage in the pipeline.  Set its departure time.
//	i.message.SetPayload(payload)
//	i.message.SetDepartedUtc(time.Now().UTC())
//	i.sendChannel <- i.message
//	i.incSendCounter()
//	return nil
//}
//
//// Method Execute walks the source directory paths, and adds a file system path for each
//// regular file that is found to the output channel.
//func (i *Analysis) Execute() error {
//
//	// Ready to process the input directories...  For this stage, this is when the message arrived.
//	i.message.SetArrivedUtc(time.Now().UTC())
//
//}
//
//// Function New creates a new instance of the stage.
//func New(sendChannel chan<- msg.Intake) *Analysis {
//
//	// Create the message for this stage in the pipeline...
//	message := message.NewIntake("")
//
//	// Create the first stage in the pipeline: IntakeStep.
//	//intakeStep := &Intake{name: KStepName, directories: directories, message: *message, receiveCounter: 0, sendCounter: 0, sendChannel: sendChannel}
//
//	return nil
//}
