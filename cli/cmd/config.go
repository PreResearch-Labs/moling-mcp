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

package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gojue/moling/services"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show current configuration",
	Long: `Show the current service configuration. You can refer to the configuration file to modify the configuration.
`,
	RunE: ConfigCommandFunc,
}

// bashCommandFunc executes the "bash" command.
func ConfigCommandFunc(command *cobra.Command, args []string) error {
	loger := initLogger(fmt.Sprintf("%s/logs/moling_debug.log", mlConfig.DataPath))
	mlConfig.SetLogger(loger)
	loger.Info().Msg("Start to show config")
	ctx := context.WithValue(context.Background(), services.MoLingConfigKey, mlConfig)
	ctx = context.WithValue(ctx, services.MoLingLoggerKey, loger)
	bf := bytes.Buffer{}
	bf.WriteString("\n{\n")
	first := true
	for _, nsv := range services.ServiceList() {
		srv, err := nsv(ctx, args)
		if err != nil {
			return err
		}
		if !first {
			bf.WriteString(",\n")
		}
		bf.WriteString(fmt.Sprintf("\t\"%s\":\n", srv.Name()))
		bf.WriteString(fmt.Sprintf("\t%s\n", srv.Config()))
		first = false
	}
	bf.WriteString("}\n")

	// 解析原始 JSON 字符串
	var data interface{}
	err := json.Unmarshal(bf.Bytes(), &data)
	if err != nil {
		return fmt.Errorf("Error unmarshaling JSON: %v\n", err)
	}

	// 格式化 JSON
	formattedJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshaling JSON: %v\n", err)

	}

	loger.Printf("Config: \n%s", formattedJson)
	return nil
}

func init() {
	rootCmd.AddCommand(configCmd)
}
