package ptrack

import (
	"fmt"
	"os"
	"path/filepath"
)

// ScanDirectory searches a directory recursively for Git repositories.
// It returns a slice of strings, where each string is the absolute path
// to a directory containing a .git directory.
// It returns an error if there was a problem walking the directory.
func ScanDirectory(root string) ([]string, error) {
	// projects will store the paths to the git project directories found.
	var projects []string
	// filepath.Walk recursively walks the root directory, calling the anonymous function for each file and directory.
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// If there's an error accessing the file or directory, return the error.
		if err != nil {
			return err // Propagate errors
		}
		// Check if the current item is a directory and its name is ".git".
		if info.IsDir() && info.Name() == ".git" {
			// If it's a .git directory, get the parent directory which is the project root.
			projectDir := filepath.Dir(path) // The parent directory of .git
			// Add the project directory to the projects slice.
			projects = append(projects, projectDir)
			// Skip further scanning of this directory because it is a .git directory itself,
			// we are only interested in the top level directory of a git project.
			return filepath.SkipDir // Don't scan inside the .git directory
		}
		// If not a .git directory, continue scanning.
		return nil
	})

	// If filepath.Walk returned an error, wrap and return it.
	if err != nil {
		return nil, fmt.Errorf("error walking directory %s: %w", root, err)
	}

	// Return the slice of project paths and nil error if no errors occurred.
	return projects, nil
}
