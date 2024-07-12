package util

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"

	klog "k8s.io/klog/v2"
)

// Gzip compresses a file using gzip and writes the result to disk
func Gzip(input, output string) error {
	// Open input file
	inputFile, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer func() {
		closeErr := inputFile.Close()
		if err == nil {
			err = closeErr
		} else {
			klog.Errorf("failed to close input file: %v", err)
		}
	}()

	// Create the output .tgz file
	outputFile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() {
		closeErr := outputFile.Close()
		if err == nil {
			err = closeErr
		} else {
			klog.Errorf("failed to close output file: %v", err)
		}
	}()

	// Do the gzip
	gzipWriter := gzip.NewWriter(outputFile)
	defer func() {
		closeErr := gzipWriter.Close()
		if err == nil {
			err = closeErr
		} else {
			klog.Errorf("failed to close gzipWriter: %v", err)
		}
	}()
	if _, err := io.Copy(gzipWriter, inputFile); err != nil {
		return fmt.Errorf("failed to write compressed data: %w", err)
	}

	return nil
}
