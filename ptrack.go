// Package ptrack provides functionality for tracking projects, specifically by identifying Git repositories.
//
// This package includes functions for scanning a directory and identifying Git repositories within it.
// It also provides helper functions for setting up test directory structures.
//
// The core function, ScanDirectory, recursively searches a given directory for
// directories that contain a ".git" subdirectory, indicating a Git repository.
// The test suite ensures the correct identification of projects in various scenarios,
// including nested directories and empty directories.
// The setupTestDir helper function facilitates creating temporary directory structures for testing.
// These tests cover different scenarios: finding projects at the root, in subdirectories,
// finding multiple projects, and finding no projects in empty directories.
//
// The package uses the testify/assert package for making assertions in the tests.
// The sort package is used to sort the found projects for reliable comparison due to potentially unordered results from the underlying file system operations.
//
// The ptrack package is designed to be a lightweight and efficient solution for
// identifying Git projects within a file system.
package ptrack
