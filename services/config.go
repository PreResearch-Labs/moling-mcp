// Copyright 2025 CFC4N <cfc4n.cs@gmail.com>. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Repository: https://github.com/gojue/moling

package services

import "github.com/rs/zerolog"

// Config is an interface that defines a method for checking configuration validity.
type Config interface {
	// Check validates the configuration and returns an error if the configuration is invalid.
	Check() error
}

type MoLingConfig struct {
	ConfigFile   string   `json:"config_file"`   // The path to the configuration file.
	DataPath     string   `json:"data_path"`     // The data path for the server, used for storing files. MacOS: /Users/username/.moling, Linux: /home/username/.moling, Windows: C:\Users\username\.moling
	AllowDir     []string `json:"allow_dir"`     // The directories that are allowed to be accessed by the server.
	AllowCommand []string `json:"allow_command"` // The commands that are allowed to be executed by the server.
	Version      string   `json:"version"`       // The version of the MoLing server.
	logger       zerolog.Logger
}

func (cfg *MoLingConfig) Check() error {
	panic("not implemented yet") // TODO: Implement Check
}

func (cfg *MoLingConfig) Logger() zerolog.Logger {
	return cfg.logger
}

func (cfg *MoLingConfig) SetLogger(logger zerolog.Logger) {
	cfg.logger = logger
}
