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
	"strconv"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id := command{}
		sv, err := strconv.Atoi(args[0])
		if err != nil {
			log.Printf("can't convert string into int %s", err)
		}
		id.ID = sv
		sendid, err := json.Marshal(id)
		if err != nil {
			log.Printf("Could not marshal argument: %s", err)
		}
		req, err := http.NewRequest("POST", "http://localhost:3030/API/get", bytes.NewBuffer(sendid))
		if err != nil {
			log.Printf("request failed: %s", err)
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
		err = json.Unmarshal(body, &id.Log)
		if err != nil {
			log.Printf("can't read body %s", err)
		}
		log.Print(id)
		log.Print(resp.Status)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
