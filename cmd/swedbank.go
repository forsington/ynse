package cmd

/*
Copyright © 2021 HAMPUS FORS <h@f0.rs>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"log"

	"github.com/spf13/cobra"
)

// swedbankCmd represents the swedbank command
var swedbankCmd = &cobra.Command{
	Use:     "swedbank",
	Aliases: []string{"sw"},
	Short:   "[NOT IMPLEMENTED]",
	Long:    `[NOT IMPLEMENTED]`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("Swedbank bank statements are not supported yet")
	},
}

func init() {
	importCmd.AddCommand(swedbankCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// swedbankCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// swedbankCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
