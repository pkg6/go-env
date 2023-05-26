package goenv

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

type JsonEnv struct {
	fileAbstract
}

func (f *JsonEnv) EnvMap() (map[string]string, error) {
	return f.filesEnvMap(f.FileMap)
}

func (f *JsonEnv) FileMap(file string) (map[string]string, error) {
	fileData, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}
	envs := map[string]string{}
	keyPairs := make(map[string]any)
	err = json.Unmarshal(fileData, &keyPairs)
	if err != nil {
		return nil, err
	}
	return f.mapKeyPairs(envs, keyPairs), nil
}

func (f *JsonEnv) mapKeyPairs(data map[string]string, keyPairs map[string]any) map[string]string {
	for key, value := range keyPairs {
		switch reflect.TypeOf(value).String() {
		case "bool":
			data[key] = strconv.FormatBool(value.(bool))
		case "float64":
			data[key] = strconv.FormatFloat(value.(float64), 'f', -1, 64)
		case "int64":
			data[key] = strconv.FormatInt(value.(int64), 10)
		case "string":
			data[key] = value.(string)
		default:
			extra := f.mapKeyPairs(make(map[string]string), value.(map[string]any))
			for subKey, subVal := range extra {
				data[key+"_"+subKey] = subVal
			}
		}
	}
	return data
}

func (f *JsonEnv) SetEnv() (failedEnvs map[string]string) {
	envMap, _ := f.EnvMap()
	return SetEnvMap(envMap)
}
