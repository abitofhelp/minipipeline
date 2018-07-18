// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package main is the entry point for the application
// and is responsible for configuring the environment.
package main

import (
	"fmt"
	. "github.com/abitofhelp/MiniPipeline/step/intake"
	"os"
	"sync"
)

// Function main is the entry point for the application.
func main() {

	var wg sync.WaitGroup

	// Create the first step in the pipeline
	//intakeStep, err := NewFactory().Create("intake")

	directoryPaths := []string{"/tmp/", "/home/mjgardner/Downloads"}
	intakeStep := New(directoryPaths, &wg)
	if intakeStep == nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to create an intake step.")
		os.Exit(-1)
	}

	// Perform the first step...
	go intakeStep.Execute(&wg)

	processed := 1
	// Simulate the next step pulling from the output channel...
	// Spin off a goroutine to process each file in the channel
	for path := range intakeStep.ChannelOut() {
		go func() {
			fmt.Printf("\n(%d) Processed: %s", processed, path)
			processed++
			wg.Done()
		}()
	}

	// Wait for all goroutines to complete.
	wg.Wait()

	fmt.Println("\nTotal Found: %d", intakeStep.Counter())
	fmt.Printf("\nTotal Processed: %d", processed)

}
