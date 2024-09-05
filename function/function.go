package function

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func LoadJSONConfig(filePath string, config interface{}) error {

	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	if err := json.Unmarshal(file, config); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

func LoadYAMLConfig(filePath string, config interface{}) error {

	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read YAML file: %w", err)
	}

	if err := yaml.Unmarshal(file, config); err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return nil
}

func LoadENVConfig(filePath string) error {

	if err := godotenv.Load(filePath); err != nil {
		return fmt.Errorf("failed to load ENV file: %w", err)
	}
	return nil
}
