package go_samples

import "os"

func readFile(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

func statFile(filename string) (os.FileInfo, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return nil, err
	} else {
		return stat, nil
	}
}

func writeFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	} else {
		return nil
	}
}
