package config

import (
	"encoding/json"
	"os"

	"github.com/Graphyk/cours-go/internal/database"
	"github.com/spf13/cobra"
)

type Config struct {
	DB database.DBConfig `json:"db"`
}

func fromCLI(cmd *cobra.Command) (*Config, error) {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return nil, err
	}

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := json.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	return config, nil
}
