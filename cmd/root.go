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

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ynse",
	Short: "a CLI for importing swedish bank statements to You Need A Budget (YNAB)",
	Long: `A CLI for importing swedish bank statements to You Need A Budget (YNAB).
See https://github.com/forsington/ynse for examples`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("access-token", "a", "YNAB-ACCESS-TOKEN-HERE", "your YNAB Personal Access Token")
	_ = viper.BindPFlag("accessToken", rootCmd.PersistentFlags().Lookup("access-token"))
}
