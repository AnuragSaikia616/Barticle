/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Barticle",
	Short: "A text summarization tool",
	Long: `Barticle is a command-line interface (CLI) tool for efficient text summarization, 
  equipped with a question-and-answer system for quick, interactive insights. With a simple command-line prompt, 
  users can input long documents or text files, and Barticle instantly generates a streamlined summary, 
  highlighting the core information. 
  Additionally, its Q&A feature allows users to ask questions about the text directly from the terminal, 
  making it easy to pinpoint specific details without sifting through the entire document. `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			RunBarticleTUI()
		} else {
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Barticle.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
