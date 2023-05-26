package goenv

import (
	"os"
	"path/filepath"
	"strings"
)

func init() {
	AddIEnv(".env", &FileEnv{})
	AddIEnv(".json", &JsonEnv{})
}

func AddIEnv(ext string, env IEnv) {
	iEnvs[ext] = env
}

type IEnv interface {
	Load(files ...string) FileError
	EnvMap() (map[string]string, error)
	FileMap(file string) (map[string]string, error)
	SetEnv() (failedEnvs map[string]string)
}

type fileAbstract struct {
	Files []string
}

func (f *fileAbstract) Load(files ...string) FileError {
	err := FileError{}
	for _, file := range files {
		if fileExist(file) {
			f.Files = append(f.Files, file)
		} else {
			err.Files = append(err.Files, file)
		}
	}
	return err
}

type filesEnvMapFn func(file string) (map[string]string, error)

func (f fileAbstract) filesEnvMap(fileMapFn filesEnvMapFn) (map[string]string, error) {
	envs := map[string]string{}
	for _, file := range f.Files {
		fileMap, err := fileMapFn(file)
		if err != nil {
			return nil, err
		}
		envs = mapMerge(fileMap)
	}
	return envs, nil
}

type FileError struct {
	Files []string
}

func (f FileError) Error() string {
	if len(f.Files) > 0 {
		return strings.Join(f.Files, ",") + " files does not exist"
	}
	return ""
}

func fileClassify(files []string) map[string][]string {
	maps := make(map[string][]string)
	for _, f := range files {
		ext := filepath.Ext(f)
		if _, ok := maps[ext]; ok {
			maps[ext] = append(maps[ext], f)
		} else {
			maps[ext] = []string{f}
		}
	}
	return maps
}
func fileExist(path string) bool {
	abs, _ := filepath.Abs(path)
	if len(abs) == 0 {
		return false
	}
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func mapMerge(maps ...map[string]string) map[string]string {
	m := make(map[string]string)
	for _, sMap := range maps {
		for k, v := range sMap {
			m[k] = v
		}
	}
	return m
}
