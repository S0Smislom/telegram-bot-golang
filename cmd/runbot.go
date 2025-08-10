/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"telegram-bot/internal/adapters/telegram"

	"github.com/spf13/cobra"
)

// runbotCmd represents the runbot command
var runbotCmd = &cobra.Command{
	Use:   "runbot",
	Short: "Run bot",
	Long:  `Starts telegram bot`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runbot called")
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Fatal(err)
		}
		bot, err := telegram.NewTelegramBotV2(configPath)
		if err != nil {
			log.Fatal(err)
		}
		if err := bot.Start(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runbotCmd)
	runbotCmd.Flags().StringP("config", "c", ".env", "Path to the config file")
}
