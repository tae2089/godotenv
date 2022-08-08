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
	for scanner.Scan() {
		split(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func split(envText string) {
	envs := strings.Split(envText, "=")
	fmt.Println("key: ", envs[0], " value: ", envs[1])
}
