package godotenv_test

import (
	"github.com/stretchr/testify/assert"
	"godotenv"
	"testing"
)

//
func TestSetupEnvNotFoundFile(t *testing.T) {
	t.Skip("env 파일이 없으면 에러나는 거 체크 확인")
	err := godotenv.LoadEnv()
	assert.Error(t, err, "open .env: no such file or directory")
}

func TestLoadEnv(t *testing.T) {
	err := godotenv.LoadEnv()
	assert.Equal(t, err, nil)
}
