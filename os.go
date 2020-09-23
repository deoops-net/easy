package easy

import "path/filepath"

func RootDir() string {
	root, err := filepath.Abs("./")
	if err != nil {
		return "."
	}
	return root
}
