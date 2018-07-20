// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package intake implements a pipeline step in which items in an array are ingested and
// passed along to the next step in the pipeline.
package intake

import (
	"fmt"
	"github.com/abitofhelp/minipipeline/step"
	"github.com/karrick/godirwalk"
	"os"
	. "path/filepath"
	"sync"
	"sync/atomic"
)

const (
	// The name of this step.
	KStepName = "Intake"

	// The length of the send channel.
	KSendChannelSize = 1000
)

// Type Intake implements the IStep interface
// and scans a directory for file system paths,
// which are passed along to the pipeline for
// further processing.
type Intake struct {
	step.IFirst

	// Field name is the name of the step.
	name string

	// Field items is an array of items to process.
	inputItems []string

	// Field receiveCounter is the number of elements that will be
	// sent through the send channel.
	receiveCounter uint64

	// Field sendCounter is the number of elements that have been
	// sent through send channel.
	sendCounter uint64

	// Field sendChannel contains the elements being passed along to
	// the next step in the pipeline.
	sendChannel chan<- string
}

// Method Name returns the name of the step.
func (i Intake) Name() string {
	return i.name
}

// Method InputItems returns an array of items that will be processed by the pipeline.
func (i Intake) InputItems() []string {
	return i.inputItems
}

// Method ReceiveCounter returns the number of elements that have been
// processed as input, atomically.
func (i Intake) ReceiveCounter() uint64 {
	return atomic.LoadUint64(&i.receiveCounter)
}

// Method SendCounter returns the number of elements that have been
// sent through the send channel to the next step in the pipeline, atomically.
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

// Method setChannelOut sets the output channel for the step.
func (i *Intake) Send() error {
	// ToDo
	return nil
}

// Method Execute walks the source directory paths, and adds a file system path for each
// regular file that is found to the output channel.
func (i *Intake) Execute() error {
	return i.loadFilePathsToSendChannel()
}

// Recursively walk a file system hierarchy to locate files, and pass the paths into the pipeline for processing.
// Parameter pathToDirectory is the path to a folder containing files to process.
// Parameter pathsChannel is the unidirectional channel being used to feed the paths to the pipeline.

func (i *Intake) loadFilePathsToSendChannel() error {

	// Variable wg is main's WaitGroup, which detects when all of the
	// goroutines that were launched have completed.
	var wg sync.WaitGroup

	// Determine the source directory paths to process.
	for _, directoryPath := range i.inputItems {

		fmt.Printf("\nDirectoryPath: %s\n", directoryPath)

		godirwalk.Walk(directoryPath, &godirwalk.Options{

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
						i.incReceiveCounter()
						fmt.Printf("F: %s\n", path)
						// When the channel is full, we will block here...
						i.sendChannel <- Join(path, de.Name())
						i.incSendCounter()
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

// Function New creates a new instance of the Intake step.
func New(directoryPaths []string, sendChannel chan<- string) *Intake {

	// Create the first step in the pipeline: IntakeStep.
	intakeStep := &Intake{name: KStepName, inputItems: directoryPaths, receiveCounter: 0, sendCounter: 0, sendChannel: sendChannel}

	return intakeStep
}
