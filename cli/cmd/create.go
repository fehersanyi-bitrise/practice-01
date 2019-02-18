// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type command struct {
	ID      int    `json:"id"`
	Command string `json:"command"`
	Flag    string `json:"flag"`
	Log     string `json:"log"`
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "runs the command, stores in database",
	Long:  `With this command, you can run a terminal command in the agent container and store it in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		samplecommand := command{}
		samplecommand.Command, err = cmd.Flags().GetString("command")
		if err != nil {
			log.Printf("Erros parsing command flag: %s", err)
		}
		samplecommand.Flag, err = cmd.Flags().GetString("flag")
		if err != nil {
			log.Printf("Erros parsing flag flag, or not given: %s", err)
		}
		message, err := json.Marshal(samplecommand)
		if err != nil {
			log.Printf("Erros marshalling return command: %s", err)
		}
		req, err := http.NewRequest("POST", "http://localhost:3030/API/create", bytes.NewBuffer(message))
		if err != nil {
			log.Println(err)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}

		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Println(err)
			}
		}()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%s", body)
	},
}

func init() {
	createCmd.Flags().StringP("command", "c", "", "the command you would run")
	createCmd.Flags().StringP("flag", "f", "", "command flag or path")
	rootCmd.AddCommand(createCmd)
}
