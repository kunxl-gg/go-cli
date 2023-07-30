/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var comment string = "1"

		if(len(args) > 0 && args[0] != ""){
			comment = args[0]
		}

		URL := "https://jsonplaceholder.typicode.com/comments/?postId=" + comment

		fmt.Println("Trying to get the commnad for ")

		// Get the data
		res, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}

		// Close the connnection
		defer res.Body.Close()

		if res.StatusCode == 200{
			body, err := ioutil.ReadAll(res.Body)
			if err != nil{
				fmt.Println(err)
			}

			fmt.Println(string(body))
		}



	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
