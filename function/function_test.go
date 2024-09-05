package function_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xadichamahkamova/config-helper/function"
)

type Config struct {
	AppName string `json:"app_name" yaml:"app_name"`
	Port    int    `json:"port" yaml:"port"`
}

func TestLoadJSONConfig(t *testing.T) {
	expectedConfig := Config{AppName: "TestApp", Port: 8080}

	err := function.LoadJSONConfig("../testdata/config.json", &expectedConfig)
	assert.NoError(t, err)
	assert.Equal(t, "TestApp", expectedConfig.AppName)
	assert.Equal(t, 8080, expectedConfig.Port)
}

func TestLoadYAMLConfig(t *testing.T) {
	expectedConfig := Config{AppName: "TestApp", Port: 8080}

	err := function.LoadYAMLConfig("../testdata/config.yaml", &expectedConfig)
	assert.NoError(t, err)
	assert.Equal(t, "TestApp", expectedConfig.AppName)
	assert.Equal(t, 8080, expectedConfig.Port)
}

func TestLoadENVConfig(t *testing.T) {
	err := function.LoadENVConfig("../testdata/.env")
	assert.NoError(t, err)

	assert.Equal(t, "TestApp", os.Getenv("APP_NAME"))
	assert.Equal(t, "8080", os.Getenv("PORT"))
}
