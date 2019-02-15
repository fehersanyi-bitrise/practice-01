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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "runs the command, stores in database",
	Long:  `With this command, you can run a terminal command in the agent container and store it in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		type command struct {
			ID      int    `json:"id"`
			Command string `json:"command"`
			Flag    string `json:"flag"`
			Log     string `json:"log"`
		}
		samplecommand := command{}
		samplecommand.Command = "pwd"
		message, _ := json.Marshal(samplecommand)
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
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
