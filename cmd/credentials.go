package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var credentialsCreateCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Create DefectDojo stored credentials",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("credentials create called")
	},
}

var credentialsDeleteCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Delete DefectDojo stored credential",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("credentials delete called")
	},
}

var credentialsListCmd = &cobra.Command{
	Use:   "credentials",
	Short: "List DefectDojo stored credentials",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list credentials called")
	},
}

func init() {
	createCmd.AddCommand(credentialsCreateCmd)
  listCmd.AddCommand(credentialsListCmd)
  deleteCmd.AddCommand(credentialsDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// credentialsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// credentialsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
