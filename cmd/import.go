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
	"github.com/forsington/ynse/bank"
	"github.com/forsington/ynse/budget"
	"github.com/forsington/ynse/importer"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import {-f file | -d dir} --bank shb --api-key key --budget-id abcd --account-id abcd",
	Short: "import a file or all files from a directory",
	Long:  `import transactions from bank statement files to a YNAB budget`,
	Run: func(cmd *cobra.Command, args []string) {

		apiKey := viper.GetString("apiKey")
		accountID := viper.GetString("accountID")
		budgetID := viper.GetString("budgetID")
		filename := viper.GetString("filename")
		dir := viper.GetString("dir")
		bankName := viper.GetString("bank")
		allowDuplicates := viper.GetBool("allowDuplicates")

		// call import
		imp := importer.New(bank.ImplementedParsers)
		transactions, err := imp.Import(filename, dir, bankName)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		// push to budget & account
		budget := budget.New(budget.NewRepo(apiKey))
		trans, err := budget.Push(budgetID, accountID, transactions, allowDuplicates)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		if len(trans) == 0 {
			fmt.Println("no new transactions to import, exiting...")
		} else {
			fmt.Println("successfully imported", len(trans), "transactions")
		}
	},
}

func init() {
	importCmd.PersistentFlags().StringP("budget-id", "b", "", "your YNAB Budget ID")
	_ = viper.BindPFlag("budgetID", importCmd.PersistentFlags().Lookup("budget-id"))
	importCmd.MarkFlagRequired("budget-id")

	importCmd.PersistentFlags().StringP("account-id", "c", "", "your YNAB Account ID")
	_ = viper.BindPFlag("accountID", importCmd.PersistentFlags().Lookup("account-id"))

	importCmd.PersistentFlags().StringP("filename", "f", "", "path to file")
	_ = viper.BindPFlag("filename", importCmd.PersistentFlags().Lookup("filename"))

	importCmd.PersistentFlags().StringP("dir", "d", "", "path to directory")
	_ = viper.BindPFlag("dir", importCmd.PersistentFlags().Lookup("dir"))

	importCmd.PersistentFlags().StringP("bank", "k", "", "bank for the file to import")
	_ = viper.BindPFlag("bank", importCmd.PersistentFlags().Lookup("bank"))

	importCmd.PersistentFlags().Bool("allow-duplicates", false, "skip fuzzy check for existing transaction duplication")
	_ = viper.BindPFlag("allowDuplicates", importCmd.PersistentFlags().Lookup("allow-duplicates"))

	rootCmd.AddCommand(importCmd)
}
