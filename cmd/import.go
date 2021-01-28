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
	"fmt"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import {-f file | -d dir} --bank shb --api-key key --budget-id abcd --account-id abcd",
	Short: "import a file or all files from a directory",
	Long:  `import transactions from bank statement files to a YNAB budget`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import called")

		apiKey := viper.GetString("apiKey")
		fmt.Println("apiKey", apiKey)
		// call import

		// push to budget & account

	},
}

func init() {

	rootCmd.AddCommand(importCmd)
}
