package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var endpointsCreateCmd = &cobra.Command{
	Use:   "endpoints",
	Short: "Create an engagement endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create endpoints called")
	},
}

var endpointsDeleteCmd = &cobra.Command{
	Use:   "endpoints",
	Short: "Delete an engagement endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete endpoints called")
	},
}

var endpointsListCmd = &cobra.Command{
	Use:   "endpoints",
	Short: "List endpoints of an engagement",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list endpoints called")
	},
}

func init() {
  createCmd.AddCommand(endpointsCreateCmd)
  deleteCmd.AddCommand(endpointsDeleteCmd)
	listCmd.AddCommand(endpointsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// endpointsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// endpointsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
