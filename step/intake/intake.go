// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package step provides the building blocks for making a pipeline.
package intake

import (
	"fmt"
	"github.com/abitofhelp/MiniPipeline/step"
	"github.com/karrick/godirwalk"
	"os"
	. "path/filepath"
	"sync"
)

const (
	// The name of this step.
	kStepName = "Intake"

	// The length of the input channel.
	kIntakeChannelInSize = 100

	// The length of the output channel.
	kIntakeChannelOutSize = 1000
)

// Type Intake implements the IStep interface
// and scans a directory for file system paths,
// which are passed along to the pipeline for
// further processing.
type Intake struct {
	step.IStep

	// Field name is the name of the step.
	name string

	// Field directoryPaths is an array of directory paths from which
	// files will be grabbed.
	directoryPaths []string

	// Field counter is the number of file system paths that have been
	// passed along to the next step through a channel.
	counter uint64

	// Field channelIn contains the directories that will be scanned for
	// files that will be processed by the pipeline.
	channelIn chan string

	// Field channelOut contains the file system paths to files that
	// will be processed by the pipeline.
	channelOut chan string
}

// Method Name returns the name of the step.
func (i Intake) Name() string {
	return i.name
}

// Method setName sets the name of the step.
func (i *Intake) setName(name string) {
	i.name = name
}

// Method DirectoryPaths returns an array of directory paths from which
// files will be grabbed.
func (i Intake) DirectoryPaths() []string {
	return i.directoryPaths
}

// Method setDirectoryPaths sets an array of directory paths from which
//// files will be grabbed.
func (i *Intake) setDirectoryPaths(directoryPaths []string) {
	i.directoryPaths = directoryPaths
}

// Method Counter returns the number of file paths that
// were passed along to the next step though ChannelOut.
func (i Intake) Counter() uint64 {
	return i.counter
}

// Method setCounter sets the number of file paths that
// were passed along to the next step though ChannelOut.
func (i *Intake) setCounter(counter uint64) {
	i.counter = counter
}

// Method incrementCounter increases the counter by one.
func (i *Intake) incrementCounter() {
	i.counter++
}

// Method decrementCounter decreases the counter by one.
func (i *Intake) decrementCounter() {
	i.counter--
}

// Method ChannelIn returns the input channel for the step.
func (i Intake) ChannelIn() chan string {
	return i.channelIn
}

// Method setChannelIn sets the input channel fgr the step.
func (i *Intake) setChannelIn(channelIn chan string) {
	i.channelIn = channelIn
}

// Method ChannelOut returns the output channel for the step.
func (i Intake) ChannelOut() chan string {
	return i.channelOut
}

// Method setChannelOut sets the output channel for the step.
func (i *Intake) setChannelOut(channelOut chan string) {
	i.channelOut = channelOut
}

// Method Execute walks the source directory paths, and adds a file system path for each
// regular file that is found to the output channel.
func (i Intake) Execute(wg *sync.WaitGroup) error {

	// Recursively scan the path for files to process...
	go i.loadPathsToChannel(wg)

	return nil
}

// Recursively walk a file system hierarchy to locate files, and pass the paths into the pipeline for processing.
// Parameter pathToDirectory is the path to a folder containing files to process.
// Parameter pathsChannel is the unidirectional channel being used to feed the paths to the pipeline.
// Parameter commandChannel is used to start the pipeline's processing.
func (i Intake) loadPathsToChannel(wg *sync.WaitGroup) {
	defer wg.Done()

	// Determine the source directory paths to process.
	for directoryPath := range i.ChannelIn() {
		go func() {
			fmt.Printf("\nSource Directory: %s", directoryPath)

			godirwalk.Walk(directoryPath, &godirwalk.Options{

				FollowSymbolicLinks: false,
				Unsorted:            true,

				Callback: func(path string, de *godirwalk.Dirent) error {
					if de.IsRegular() {
						wg.Add(1)
						fmt.Printf("\n(%d) Found: %s", i.Counter(), path)
						i.ChannelOut() <- Join(path, de.Name())
						i.incrementCounter()
					}

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
		}()
	}
}

// Function New creates a new instance of the Intake step.
func New(directoryPaths []string, wg *sync.WaitGroup) *Intake {

	// Create the input and output channels.
	input := make(chan string, kIntakeChannelInSize)
	output := make(chan string, kIntakeChannelOutSize)

	// Create the first step in the pipeline, which
	// is responsible for determining all of the file
	// system paths from source directories.
	step := &Intake{name: kStepName, directoryPaths: directoryPaths, counter: 0, channelIn: input, channelOut: output}

	// Load the source directories into the input channel...
	for _, directoryPath := range directoryPaths {
		fmt.Printf("\nSource Directory: %s", directoryPath)
		step.channelIn <- directoryPath
	}

	return step
}
