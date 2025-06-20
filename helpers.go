package ptrack

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// PrepareSuite prepares a test suite by creating a temporary directory,
// copying the source suite into it, and renaming "git-dir" directories to ".git".
// It returns the path to the prepared suite or an error if something goes wrong.
func PrepareSuite(sourceSuitePath string) (string, error) {
	// 1. Create a new temporary directory for the test run.
	tempDir, err := os.MkdirTemp("", "ptrack-testsuite-")
	if err != nil {
		// If the temporary directory creation fails, return an error.
		return "", fmt.Errorf("failed to create temp dir: %w", err)
	}

	// 2. Perform a deep copy of the source suite into the temp directory.
	err = filepath.Walk(sourceSuitePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			// If there's an error during the walk, return it.
			return err
		}
		relPath, err := filepath.Rel(sourceSuitePath, path)
		if err != nil {
			// Handle errors during relative path calculation.
			return fmt.Errorf("failed to relative read path: %w", err)
		}
		if relPath == "." {
			// Skip the root directory itself.
			return nil
		}
		destPath := filepath.Join(tempDir, relPath)
		if info.IsDir() {
			// Create the directory in the destination.
			return os.MkdirAll(destPath, info.Mode())
		}
		data, err := os.ReadFile(path)
		if err != nil {
			// Handle errors during file reading.
			return fmt.Errorf("failed to read path: %w", err)
		}

		// Write the file to the destination.
		return os.WriteFile(destPath, data, info.Mode())
	})
	if err != nil {
		// If the copy fails, clean up the temp dir and return an error.
		nErr := os.RemoveAll(tempDir)
		if nErr != nil {
			return "", fmt.Errorf("failed to delete temporary directory: %w after failure to copy reference suite: %w", nErr, err)
		}

		return "", fmt.Errorf("failed to copy reference suite: %w", err)
	}

	// 3. Collect and then rename all "git-dir" directories.
	var dirsToRename []string
	err = filepath.Walk(tempDir, func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() && info.Name() == "git-dir" {
			// Collect all "git-dir" directories.
			dirsToRename = append(dirsToRename, path)
		}

		return nil
	})
	if err != nil {
		// Handle errors during walking the temporary directory to find "git-dir" directories.
		nErr := os.RemoveAll(tempDir)
		if nErr != nil {
			return "", fmt.Errorf("failed to delete temporary directory: %w after failure to find git-dird to rename: %w", nErr, err)
		}

		return "", fmt.Errorf("failed to find git-dirs to rename: %w", err)
	}

	for _, path := range dirsToRename {
		// Rename "git-dir" to ".git".
		newPath := filepath.Join(filepath.Dir(path), ".git")
		if err := os.Rename(path, newPath); err != nil {
			// If renaming fails, clean up the temp directory and return an error.
			nErr := os.RemoveAll(tempDir)
			if nErr != nil {
				return "", fmt.Errorf("failed to delete temporary directory: %w after failure to rename: %w", nErr, err)
			}

			return "", fmt.Errorf("failed to rename '%s': %w", path, err)
		}
	}

	// Return the path to the prepared suite.
	return tempDir, nil
}
