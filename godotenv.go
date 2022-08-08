package godotenv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnv() error {
	envFile, err := os.Open(".env")

	if err != nil {
		return err
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	fileEnvMap := map[string]string{}
	for scanner.Scan() {
		key, value := split(scanner.Text())
		fileEnvMap[key] = value
	}
	registerEnv(fileEnvMap, false)
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func split(envText string) (key, value string) {
	envs := strings.Split(envText, "=")
	fmt.Println("key: ", envs[0], " value: ", envs[1])
	return envs[0], envs[1]
}

func registerEnv(fileEnvMap map[string]string, overWrite bool) {
	currentSystemEnvMap := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentSystemEnvMap[key] = true
	}

	for key, value := range fileEnvMap {
		if !currentSystemEnvMap[key] || overWrite {
			os.Setenv(key, value)
		}
	}
}
