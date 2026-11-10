package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
)

type ConfigTheme struct {
	TitleBG string `json:"title_bg"`
	TitleFG string `json:"title_fg"` // TODO: add the rest of styles

	ListItemBG string `json:"list_item_bg"`
	ListItemFG string `json:"list_item_fg"`

	SelectedListItemBG string `json:"selected_list_item_bg"`
	SelectedListItemFG string `json:"selected_list_item_fg"`
}

type Config struct {
	Theme *ConfigTheme `json:"theme"`
}

type LoadedConfig struct {
	theme theme
}

var defaultConfig = LoadedConfig{kanagawaDragon()}

func loadConfig() (LoadedConfig, error) {
	cfg, err := os.UserHomeDir()
	if err != nil {
		return LoadedConfig{}, err
	}

	p := filepath.Join(cfg, ".config", "mina", "config.json")
	configFile, err := os.Open(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return defaultConfig, nil
		}
		return LoadedConfig{}, err
	}
	defer configFile.Close()

	var config Config
	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return defaultConfig, nil
		}
		return LoadedConfig{}, err
	}

	loadedConfig := LoadedConfig{}
	if config.Theme == nil {
		loadedConfig.theme = kanagawaDragon()
	} else {
		loadedConfig.theme = loadTheme(*config.Theme)
	}

	return loadedConfig, nil
}
