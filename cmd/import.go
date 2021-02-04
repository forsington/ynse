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
	"os"
	"strings"

	"github.com/forsington/ynse/bank"
	"github.com/forsington/ynse/budget"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		dryRun := viper.GetBool("dryRun")
		verbose := viper.GetBool("verbose")

		// call import
		imp := bank.New(bank.ImplementedParsers)
		transactions, err := imp.Import(filename, dir, bankName)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return
		}

		printPendingTransactions(transactions, verbose)

		if dryRun {
			fmt.Printf("would have created %d transactions, exiting dry run", len(transactions))
			return
		}

		fmt.Printf("%d transactions prepared, please verify that they are correct\n", len(transactions))
		isConfirmed := askForConfirmation("import them to YNAB?")
		if !isConfirmed {
			fmt.Println("exiting...")
			return
		}

		fmt.Printf("connecting to YNAB...\n")
		// push to budget & account
		budget := budget.New(budget.NewRepo(apiKey))
		trans, err := budget.Push(budgetID, accountID, transactions, allowDuplicates)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return
		}

		if len(trans) == 0 {
			fmt.Println("no new transactions to import, exiting...")
		} else {
			fmt.Printf("imported %d transactions to YNAB\n", len(trans))
		}
	},
}

func init() {
	importCmd.PersistentFlags().StringP("budget-id", "b", "", "your YNAB Budget ID")
	_ = viper.BindPFlag("budgetID", importCmd.PersistentFlags().Lookup("budget-id"))

	importCmd.PersistentFlags().String("account-id", "", "your YNAB Account ID")
	_ = viper.BindPFlag("accountID", importCmd.PersistentFlags().Lookup("account-id"))

	importCmd.PersistentFlags().StringP("filename", "f", "", "path to file")
	_ = viper.BindPFlag("filename", importCmd.PersistentFlags().Lookup("filename"))

	importCmd.PersistentFlags().StringP("dir", "d", "", "path to directory")
	_ = viper.BindPFlag("dir", importCmd.PersistentFlags().Lookup("dir"))

	importCmd.PersistentFlags().String("bank", "", "bank for the file to import")
	_ = viper.BindPFlag("bank", importCmd.PersistentFlags().Lookup("bank"))

	importCmd.PersistentFlags().Bool("allow-duplicates", false, "skip fuzzy check for existing transaction duplication")
	_ = viper.BindPFlag("allowDuplicates", importCmd.PersistentFlags().Lookup("allow-duplicates"))

	importCmd.PersistentFlags().Bool("dry-run", false, "dry run, doesn't create transactions in YNAB")
	_ = viper.BindPFlag("dryRun", importCmd.PersistentFlags().Lookup("dry-run"))

	importCmd.PersistentFlags().BoolP("verbose", "v", false, "dry run, doesn't create transactions in YNAB")
	_ = viper.BindPFlag("verbose", importCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.AddCommand(importCmd)
}

func askForConfirmation(question string) bool {
	var s string

	fmt.Printf("%s (y/N): ", question)
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}

func printPendingTransactions(transactions []*budget.Transaction, verbose bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("PREPARED TRANSACTIONS")
	t.AppendHeader(table.Row{"Date", "Payee", "Amount"})
	t.SetStyle(table.StyleLight)

	for i, transaction := range transactions {
		if i > 2 && !verbose {
			t.AppendRow(table.Row{"...", "...", "..."})
			break
		}
		t.AppendRow(table.Row{transaction.Date.Format("2006-01-02"), transaction.PayeeName, transaction.AmountPretty(budget.CurrencySEK)})
	}
	t.AppendSeparator()
	t.Render()
	fmt.Printf("to display all transactions, re-run with -v flag\n\n")
}
