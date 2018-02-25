package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type globalConfig struct {
	CredentialsFile string `toml:"credentials_file"`
}

func getDefaultConfig() globalConfig {
	return globalConfig{
		CredentialsFile: "service_account.json",
	}
}

func (c *globalConfig) save(file string) error {

	dir := filepath.Dir(file)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(c)
	if err != nil {
		return err
	}

	return nil
}
