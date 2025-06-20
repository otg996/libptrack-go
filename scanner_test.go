package ptrack

import (
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/test-go/testify/require"
)

// A helper to create a test directory structure.
func setupTestDir(t *testing.T, structure map[string]bool) string {
	t.Helper() // Marks this as a test helper function
	// Create a temporary directory for the test.
	tempDir := t.TempDir()

	// Iterate through the directory structure to create.
	for path, isGitDir := range structure {
		// Construct the full path for the item.
		fullPath := filepath.Join(tempDir, path)
		if isGitDir {
			// If it's a Git directory, create the .git directory.
			gitPath := filepath.Join(fullPath, ".git")

			// Create the .git directory and any necessary parent directories.
			err := os.MkdirAll(gitPath, 0755)
			if err != nil {
				// If creating the .git directory fails, fail the test.
				t.Fatalf("Failed to create .git dir at %s: %v", fullPath, err)
			}
			// } else {
			// Could add file creation here if needed, but not currently required by the test.
		}
	}

	// Clean up the temporary directory after the test.
	t.Cleanup(func() {
		// Remove the temporary directory and all its contents after the test.
		err := os.RemoveAll(tempDir) // t.Cleanup automatically calls this after the test
		if err != nil {
			t.Fatalf("Failed to remove temp dir %s: %v", tempDir, err)
		}
	})

	// Return the path to the temporary directory.
	return tempDir
}

// TestScanDirectory tests the ScanDirectory function.
func TestScanDirectory(t *testing.T) {
	t.Parallel()
	t.Run("Finds project in root", func(t *testing.T) {
		t.Parallel()
		// Arrange: Create a test directory with a .git directory at the root.
		testDir := setupTestDir(t, map[string]bool{
			".": true, // "." represents the root of the testDir
		})

		// Expected result: The root directory should be identified as a project.
		expected := []string{testDir}

		// Act: Call ScanDirectory to find projects.
		actual, err := ScanDirectory(testDir)

		// Require no error and that the actual results match the expected results.
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("Finds project in a subdirectory", func(t *testing.T) {
		t.Parallel()
		// Arrange: Create a test directory with a subdirectory containing a .git directory.
		testDir := setupTestDir(t, map[string]bool{
			"my-project": true,
		})
		// Expected result: The subdirectory "my-project" should be identified as a project.
		expected := []string{filepath.Join(testDir, "my-project")}

		// Act: Call ScanDirectory to find projects.
		actual, err := ScanDirectory(testDir)

		// Require no error and that the actual results match the expected results.
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("Finds multiple projects and finds nested .git", func(t *testing.T) {
		t.Parallel()
		// Arrange: Create a test directory with multiple projects, including a nested .git directory which should *not* be a project.
		testDir := setupTestDir(t, map[string]bool{
			"project-a":     true,
			"project-b":     true,
			"project-b/sub": true, // This one should be ignored
			"not-a-project": false,
		})
		// Expected result: The top-level project directories should be found. Subdirectories under projects should not.
		expected := []string{
			filepath.Join(testDir, "project-a"),
			filepath.Join(testDir, "project-b"),
			filepath.Join(testDir, "project-b", "sub"),
		}

		// Act: Call ScanDirectory to find projects.
		actual, err := ScanDirectory(testDir)

		// Assert: Sort the slices before comparison to handle potential order differences.
		sort.Strings(expected)
		sort.Strings(actual)

		// Require no error and that the actual results match the expected results.
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("Finds no projects in an empty directory", func(t *testing.T) {
		t.Parallel()
		// Arrange: Create an empty test directory.
		testDir := setupTestDir(t, map[string]bool{})
		// Expected result: An empty slice, as there are no projects.
		var expected []string // Expect an empty (nil) slice

		// Act: Call ScanDirectory to find projects.
		actual, err := ScanDirectory(testDir)

		// Require no error and that the actual results match the expected results.  Check the length is also 0
		require.NoError(t, err)
		require.Empty(t, actual)
		require.Equal(t, expected, actual)
	})
}

func TestScanDirectoryAgainstReferenceSuite(t *testing.T) {
	t.Parallel()
	const referenceSuitePath = "./spec/reference-test-suite"

	// PrepareSuite is an external function.
	testableSuiteDir, err := PrepareSuite(referenceSuitePath)
	if err != nil {
		t.Fatal("failed to create testing suite:", err.Error())
	}

	// Expected list of project paths.
	expected := []string{
		filepath.Join(testableSuiteDir, "deep/down/in/the/filesystem/a-deep-project"),
		filepath.Join(testableSuiteDir, "multiple/project-a"),
		filepath.Join(testableSuiteDir, "multiple/project-b"),
		filepath.Join(testableSuiteDir, "nested/single-nested-project"),
		filepath.Join(testableSuiteDir, "project-with-submodule"),
		filepath.Join(testableSuiteDir, "project-with-submodule/vendor/my-lib"),
		filepath.Join(testableSuiteDir, "root-level-project"),
	}
	sort.Strings(expected)

	// Call ScanDirectory to find projects within the prepared test suite.
	actual, err := ScanDirectory(testableSuiteDir)
	sort.Strings(actual)

	// Assert that no error occurred.
	require.NoError(t, err)
	// Assert that the actual results match the expected results.
	require.Equal(t, expected, actual)
}
