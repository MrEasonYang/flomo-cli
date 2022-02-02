package application

import "os"

// Common func to parse config path with the user home dir.
func ParseFilePath(path string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir + string(os.PathSeparator) + path
}
