// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package stage

//
//import (
//	"encoding/hex"
//	"fmt"
//	. "github.com/minio/highwayhash.git"
//	"io"
//	"os"
//)
//
//type FilePath struct {
//	IMessage
//
//	checksum string
//}
//
//func New() {
//	filepath := &FilePath{checksum: ""}
//}
//
//// E6115D76AC0A76FF95ADD7EF043C786A480086403E2175CBE24BFDD2E4579825
//
//// Parameter path is the path to the file from which the checksum will be created.
//// Parameter key is a 64 character, hexidecimal key.
//func Checksum(path string, hexKey string) (string, error) {
//	// DecodeString returns the bytes represented by the hexadecimal string.
//	key, err := hex.DecodeString(hexKey)
//	if err != nil {
//		fmt.Printf("Cannot decode hex key: %v", err) // add error handling
//		return
//	}
//
//	// Open the file that will be used to determine the checksum.
//	file, err := os.Open(path)
//	if err != nil {
//		fmt.Printf("Failed to open the file: %v", err) // add error handling
//		return
//	}
//	defer file.Close()
//
//	// Create an instance of the HighwayHash'er.
//	hasher, err := New(key)
//	if err != nil {
//		fmt.Printf("Failed to create HighwayHash instance: %v", err) // add error handling
//		return
//	}
//
//	// Stream the file through the hasher.
//	if _, err = io.Copy(hasher, file); err != nil {
//		fmt.Printf("Failed to read from file: %v", err) // add error handling
//		return
//	}
//
//	// Determine the checksum.
//	checksum := hasher.Sum(nil)
//
//	fingerprint := hex.EncodeToString(checksum)
//	if fingerprint == "" {
//		fmt.Printf("Cannot decode the checksum bytes to a hex string.") // add error handling
//		return
//	}
//}
