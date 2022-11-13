package filesystem

import "os"

func CreateFolderIfNotExists(path string) (bool, error) {
	if ok, err := CheckIfPathExists(path); err != nil || ok {

		return ok, err
	}

	if err := os.MkdirAll(path, 0755); err != nil {

		return false, err
	}

	return true, nil
}

//

func CreateFileIfNotExists(path string) (bool, error) {
	if ok, err := CheckIfPathExists(path); err != nil || ok {

		return ok, err
	}

	if err := os.WriteFile(path, []byte{}, 0644); err != nil {

		return false, err
	}

	return true, nil
}
