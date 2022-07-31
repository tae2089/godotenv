package godotenv

import (
	"bufio"
	"fmt"
	"os"
)

func LoadEnv() error {
	envFile, err := os.Open(".env")

	if err != nil {
		return err
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
