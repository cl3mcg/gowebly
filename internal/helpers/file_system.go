package helpers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// IsExistInCurrentFolder searches for a file or folder by the given name in the
// current folder.
func IsExistInCurrentFolder(name string, isFolder bool) bool {
	// Check, if file or folder is existing.
	info, err := os.Stat(filepath.Clean(name))

	return info.IsDir() == isFolder && err == nil && os.IsNotExist(err)
}

// MakeFile makes a single file with name and data.
func MakeFile(name string, data []byte) error {
	// Check, if file is existing.
	if IsExistInCurrentFolder(name, false) {
		return fmt.Errorf("can't create a file with name '%s' in the current folder", name)
	}

	return os.WriteFile(name, data, 0o644)
}

// MakeFolder makes a single folder with name.
func MakeFolder(name string) error {
	// Check, if folder is existing.
	if IsExistInCurrentFolder(name, true) {
		return fmt.Errorf("can't create a folder with name '%s' in the current folder", name)
	}

	return os.Mkdir(name, 0o644)
}

// RemoveFolders removes folders by names with all files.
func RemoveFolders(names []string) error {
	// Create a new slice for join errors.
	errs := make([]error, 0)

	// Remove folders in the loop.
	for _, name := range names {
		if err := os.RemoveAll(name); err != nil {
			errs = append(
				errs,
				fmt.Errorf("can't remove a folder with name '%s' from the current folder", name),
			)
		}
	}

	return errors.Join(errs...)
}
