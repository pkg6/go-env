package goenv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileEnv struct {
	fileAbstract
}

func (f *FileEnv) EnvMap() (map[string]string, error) {
	return f.filesEnvMap(f.FileMap)
}

func (f *FileEnv) FileMap(file string) (map[string]string, error) {
	fileData, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(fileData), "\n")
	envs := map[string]string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		lineParts := strings.SplitN(line, "=", 2)
		if len(lineParts) != 2 {
			return nil, fmt.Errorf("failed to parse line, expected 2 parts got %d", len(lineParts))
		}
		key := strings.TrimSpace(lineParts[0])
		value := strings.TrimSpace(lineParts[1])
		envs[key] = value
	}
	return envs, nil
}

func (f *FileEnv) SetEnv() (failedEnvs map[string]string) {
	envMap, _ := f.EnvMap()
	return SetEnvMap(envMap)
}
