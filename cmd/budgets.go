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

	"github.com/forsington/ynse/budget"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// budgetsCmd represents the budgets command
var budgetsCmd = &cobra.Command{
	Use:   "budgets <budget-id>",
	Short: "list budgets and accounts",
	Long:  `Prints the available budgets and accounts for the given API Key.`,
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := viper.GetString("accessToken")
		srvc := budget.New(budget.NewRepo(accessToken))

		budgets, err := srvc.Get()
		if err != nil {
			fmt.Println("error in service", err.Error())
			return
		}

		fmt.Println("\nBudgets")
		for _, budget := range budgets {
			fmt.Printf("├──├ %s\n", budget.Name)
			fmt.Printf("   ├─── ID: %s\n", budget.ID)
			fmt.Printf("   ├──├ Accounts\n")
			for _, acc := range budget.Accounts {
				fmt.Printf("      ├──├ %s\n", acc.Name)
				fmt.Printf("         ├── ID: %s\n", acc.ID)

			}
			fmt.Printf("\n")

		}
	},
}

func init() {
	rootCmd.AddCommand(budgetsCmd)
}
