package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var engagementsCreateCmd = &cobra.Command{
	Use:   "engagements",
	Short: "Create an engagement",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create engagements called")
	},
}

var engagementsDeleteCmd = &cobra.Command{
	Use:   "engagements",
	Short: "Delete an engagement",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete engagements called")
	},
}

var engagementsListCmd = &cobra.Command{
	Use:   "engagements",
	Short: "List engagements",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list engagements called")
	},
}

func init() {
  createCmd.AddCommand(engagementsCreateCmd)
  deleteCmd.AddCommand(engagementsDeleteCmd)
	listCmd.AddCommand(engagementsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// engagementsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// engagementsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
