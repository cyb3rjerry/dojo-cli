package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var productsCreateCmd = &cobra.Command{
	Use:   "products",
	Short: "Create a product",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create products called")
	},
}

var productsDeleteCmd = &cobra.Command{
	Use:   "products",
	Short: "Delete a product",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete products called")
	},
}

var productsListCmd = &cobra.Command{
	Use:   "products",
	Short: "List products",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list products called")
	},
}

func init() {
  createCmd.AddCommand(productsCreateCmd)
  deleteCmd.AddCommand(productsDeleteCmd)
	listCmd.AddCommand(productsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// productsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// productsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
