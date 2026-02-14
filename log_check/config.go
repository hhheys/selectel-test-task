package logcheck

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	EnableEnglishLettersRule bool `yaml:"enable_english_letter_rule"`
	EnableLowercaseRule      bool `yaml:"enable_lowercase_rule"`
	EnableNoSpecialRule      bool `yaml:"enable_no_special_rule"`
	EnableSensitiveRule      bool `yaml:"enable_sensitive_rule"`

	SensitiveKeywords []string `yaml:"sensitive_keywords"`
	ForbiddenRunes    []string `yaml:"forbidden_runes"`
}

var config = Config{}

var configPath string = "../logcheck.yaml"

func LoadConfig() error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &config)
}
