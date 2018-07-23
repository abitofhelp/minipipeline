// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package intake implements a pipeline stage in which items in an array are ingested and
// passed along to the next stage in the pipeline.
package intake

import (
	"fmt"
	"github.com/abitofhelp/minipipeline/message"
	"github.com/karrick/godirwalk"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// The name of this stage.
	KStepName = "Intake"

	// The length of the send channel.
	KSendChannelSize = 1000
)

// Type Intake implements the IStep interface
// and scans a directory for file system paths,
// which are passed along to the pipeline for
// further processing.
type Intake struct {
	// Field name is the name of the stage.
	name string

	// Field directories is an array of directory paths to process.
	directories []string

	// Field message is the pipeline message for this stage.
	message message.Intake

	// Field receiveCounter is the number of messages that will be
	// sent through the send channel.
	receiveCounter uint64

	// Field sendCounter is the number of messages that have been
	// sent through send channel.
	sendCounter uint64

	prevChannel <-chan message.Intake

	nextChannel chan<- message.IStaged
}

// Method Name returns the name of the stage.
func (i Intake) Name() string {
	return i.name
}

// Method Directories() returns an array of directory paths that will be processed by the pipeline.
func (i Intake) Directories() []string {
	return i.directories
}

// Method Message() returns a message that will be processed by the pipeline for this stage.
func (i Intake) Message() message.Intake {
	return i.message
}

// Method ReceiveCounter returns the number of messages that have been
// processed as input, atomically.
func (i Intake) ReceiveCounter() uint64 {
	return atomic.LoadUint64(&i.receiveCounter)
}

// Method SendCounter returns the number of messages that have been
// sent through the send channel to the next stage in the pipeline, atomically.
func (i Intake) SendCounter() uint64 {
	return atomic.LoadUint64(&i.sendCounter)
}

// Method setSendCounter sets the value of the sendCounter, atomically.
func (i *Intake) setSendCounter(value uint64) {
	atomic.StoreUint64(&i.sendCounter, value)
}

// Method setReceiveCounter sets the value of the receiveCounter, atomically.
func (i *Intake) setReceiveCounter(value uint64) {
	atomic.StoreUint64(&i.receiveCounter, value)
}

// Method incSendCounter increments the value of the sendCounter, atomically.
func (i *Intake) incSendCounter() {
	atomic.AddUint64(&i.sendCounter, 1)
}

// Method incReceiveCounter increments the value of the receiveCounter, atomically.
func (i *Intake) incReceiveCounter() {
	atomic.AddUint64(&i.receiveCounter, 1)
}

// Method Send sets sends a message through the send channel.
func (i *Intake) Send(payload string) error {
	fmt.Printf("F: %s\n", payload)

	// Ready to pass the intake message to the next stage in the pipeline.  Set its departure time.
	i.message.SetPayload(payload)
	i.message.SetDepartedUtc(time.Now().UTC())
	i.sendChannel <- i.message
	i.incSendCounter()
	return nil
}

// Method Execute walks the source directory paths, and adds a file system path for each
// regular file that is found to the output channel.
func (i *Intake) Execute() error {

	// Ready to process the input directories...  For this stage, this is when the message arrived.
	i.message.SetArrivedUtc(time.Now().UTC())

	return i.loadFilePathsToSendChannel()
}

// Recursively walk a file system hierarchy to locate files, and pass the paths into the pipeline for processing.
// Parameter pathToDirectory is the path to a folder containing files to process.
// Parameter pathsChannel is the unidirectional channel being used to feed the paths to the pipeline.

func (i *Intake) loadFilePathsToSendChannel() error {

	// Field wg is main's WaitGroup, which detects when all of the
	// goroutines that were launched have completed.
	var wg sync.WaitGroup

	// Determine the messages to process.
	for _, dir := range i.directories {

		fmt.Printf("\\Directory: %s\n", dir)

		godirwalk.Walk(dir, &godirwalk.Options{

			FollowSymbolicLinks: false,
			Unsorted:            true,

			Callback: func(path string, de *godirwalk.Dirent) error {
				// Increment the producer's WaitGroup counter for the goroutine
				// launching the Consumer().
				wg.Add(1)

				// Launch the goroutine, which will block until the producer
				// sends file system directory paths into the channel.
				go func(path string) {
					// Decrement the consumer's WaitGroup counter just before the goroutine exits.
					defer wg.Done()

					if de.IsRegular() {

						// Increment the received counter...  We are receiving a path to a file rather than a message.
						i.incReceiveCounter()

						// Send the item to the next stage in the pipeline...
						i.Send(path)
					}
				}(path)

				// Signal no errors...
				return nil
			},

			ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
				// Your program may want to log the error somehow.
				fmt.Fprintf(os.Stderr, ">>>>>> ERROR: %s\n", err)

				// On error, we will skip the current file system node and continue
				// walking the file system hierarchy of remaining nodes.
				return godirwalk.SkipNode
			},
		})
	}

	// Wait here until the producer and consumer have completed their work,
	// which will be signaled the channel being closed and by wg's internal
	// goroutine counter being zero.
	wg.Wait()

	return nil
}

// ToDo: Virtual/abstract method...
//func CreateMessage() error {}
//Join(path, de.Name())
//}

// Function New creates a new instance of the Intake stage.
func New(directories []string, prevChannel <-chan message.Intake, nextChannel chan<- message.IStaged) *Intake {

	// Create the message for this stage in the pipeline...
	message := message.NewIntake("")

	// Create the first stage in the pipeline: IntakeStep.
	intakeStep := &Intake{name: KStepName, directories: directories, message: *message, receiveCounter: 0, sendCounter: 0, prevChannel: prevChannel, nextChannel: nextChannel}

	return intakeStep
}
