package filesystem

import "os"

func CheckIfPathExists(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {

		return false, err
	}

	return true, nil
}
