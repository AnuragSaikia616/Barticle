/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "This command takes a input file and makes a summary of it",
	Long: `Use this commnad to input a txt file
  the command will make a summary of the content of the file
  and output a file with the summary.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check for url flag and handle urls
		if url, _ := cmd.Flags().GetString("url"); url != "" {

			// get scraped content fromt the url
			content, err := getScrapedURLContent(url)
			if err != nil {
				log.Fatal(err)
			}

			summary, _ := getSummary(content)
			fmt.Println(SummaryStyle.Render("url: ", url))
			fmt.Println(summary)
			fmt.Println()

		} else { // url flag absent

			// return if no arguments area provided
			if len(args) == 0 {
				fmt.Println("ERROR: You need to provide at least one argument!!!")
				return
			}
			// handle files
			for _, filepath := range args {
				if _, err := os.Stat(filepath); err != nil {
					log.Fatal("File does not exits: %s", filepath)
				}
				content, err := os.ReadFile(filepath)
				if err != nil {
					log.Fatal("Error reading file: %s", err)
				}

				summary, _ := getSummary(string(content))
				fmt.Println(SummaryStyle.Render("input file: ", filepath))
				fmt.Println(summary)
				fmt.Println()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	summaryCmd.Flags().String("url", "", "give a url to scrape and summarize")
}
