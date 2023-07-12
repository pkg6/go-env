package goenv

import (
	"encoding/json"
	"os"
)

var (
	knownEnvVars   = map[string]string{}
	defaultEnvFile = []string{".env"}
	iEnvs          = map[string]IEnv{}
)

// Load file
func Load(files ...string) (map[string]string, []error) {
	var errs []error
	if len(files) == 0 {
		files = defaultEnvFile
	}
	errorMap := make(map[string]string)
	classify := fileClassify(files)
	for s, classifyFiles := range classify {
		if i, ok := iEnvs[s]; ok {
			errs = append(errs, i.Load(classifyFiles...))
			errorMap = mapMerge(errorMap, i.SetEnv())
		}
	}
	return errorMap, errs
}

// SetEnvMap set env
func SetEnvMap(envs map[string]string) (failedEnvs map[string]string) {
	for key, value := range envs {
		if os.Getenv(key) != "" {
			continue
		}
		knownEnvVars[key] = value
		if err := os.Setenv(key, value); err != nil {
			failedEnvs[key] = value
		}
	}
	return failedEnvs
}

// GetDefault default val
func GetDefault(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func JsonUnmarshal(v any) error {
	marshal, err := json.Marshal(knownEnvVars)
	if err != nil {
		return err
	}
	return json.Unmarshal(marshal, v)
}

// Clear envs
func Clear() {
	for key := range knownEnvVars {
		_ = os.Unsetenv(key)
	}
	os.Clearenv()
}
