// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package main is the entry point for the application
// and is responsible for configuring the environment.
package main

import (
	"fmt"
	"github.com/abitofhelp/minipipeline/pipeline"
)

// Function main is the entry point for the application and is responsible
// for configuring its environment.
func main() {

	pipe, err := pipeline.New()
	fmt.Println(pipe, err)

}

/*
	// Field wg is main's WaitGroup that is used to be able
	// to detect when all of the goroutines that are launched
	// have completed.
	var wg sync.WaitGroup

	// Create the channel for sending to the next stage.
	intakeSendChannel := make(chan message.Intake, intake.KSendChannelSize)

	// Increment the producer's WaitGroup counter for the goroutine
	// launching the Consumer().
	wg.Add(1)
	// Launch the goroutine, which will block until the producer
	// sends file system directory paths into the channel.
	go func() {
		// Decrement the consumer's WaitGroup counter just before the goroutine exits.
		defer wg.Done()
		Consumer(intakeSendChannel)
	}()

	paths := []string{"/tmp", "/home/mjgardner/Downloads"}

	intakeStep := intake.New(paths, intakeSendChannel)

	// Increment the producer's WaitGroup counter for the goroutine
	// launching the Producer().
	wg.Add(1)
	// Launch the goroutine, which will send file system directory
	// paths into the channel.
	go func() {
		// Decrement the producer's WaitGroup counter just before the goroutine exits.
		defer wg.Done()

		// The consumer is running in parallel to the producer.  Function main() will wait for
		// the consumer and producer goroutines to complete at its wg.Wait().  It is important to
		// know that if the producer does not close the channel, then the consumer will block
		// and wg.Wait() will never be reached in Main().  Closing the channel signals the consumer
		// that production has completed.  When the consumer empties the channel, wg.Wait() in
		// Main() will complete.  The channel will be closed just before the goroutine exits.
		defer close(intakeSendChannel)

		// Load the source directories into the channel so they can be consumed.
		intakeStep.Execute()
	}()

	// Wait here until the producer and consumer have completed their work,
	// which will be signaled the channel being closed and by wg's internal
	// goroutine counter being zero.
	wg.Wait()*/

// Adios!
//fmt.Println("All done!")
//}

/*

// Function Consumer receives file system directory paths from the channel.
// Parameter ch is a unidirectional channel from which directory paths are
// retrieved.
// It returns when all of the goroutines have completed processing the channel.
func Consumer(ch <-chan message.Intake) {
	// Field wg is the consumer's WaitGroup, which detects when all of the
	// goroutines that were launched have completed.
	var wg sync.WaitGroup

	// The consumer's wait group will block at wg.Wait(), which will be invoked just
	// before exiting from the function.  It will block until wg's internal counter is zero,
	// which happens when all of the goroutines that were launched have completed.
	defer wg.Wait()

	for msg := range ch {
		// Increment the consumer's WaitGroup counter for each goroutine that is launched.
		wg.Add(1)
		go func(path message.Intake) {
			// Decrement the consumers's WaitGroup counter after each goroutine completes its work.
			defer wg.Done()

			// The consumer's work is pretty simple...  Write the directory msg that was retrieved
			// from the channel to stdout.
			fmt.Printf("R: %s\n\tElapsed: %s\n", path.Payload(), path.Elapsed())
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
*/
