package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
		# Telegram

		Parameters:
		- message

	  Environment variables:
	  - NOTIFY_TELEGRAM_TOKEN
	  - NOTIFY_TELEGRAM_CHAT_ID
*/
func notifyTelegram(message string) {
	// Load environment variables
	telegramToken, ok := os.LookupEnv("NOTIFY_TELEGRAM_TOKEN")
	if !ok {
		log.Fatal("NOTIFY_TELEGRAM_TOKEN environment variable is unset")
	}
	telegramChatID, ok := os.LookupEnv("NOTIFY_TELEGRAM_CHAT_ID")
	if !ok {
		log.Fatal("NOTIFY_TELEGRAM_CHAT_ID environment variable is unset")
	}

	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramToken)
	payload := fmt.Sprintf(`{"chat_id": "%s","text":"%s"}`, telegramChatID, message)

	req, _ := http.NewRequest("POST", endpoint, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.ReadAll(resp.Body)
	log.Print("Telegram notified")
}

/*
		# Hangouts Chat

		Parameters:
		- message

	  Environment variables:
	  - NOTIFY_HANGOUTS_CHAT_WEBHOOK
*/
func notifyHangoutsChat(message string) {
	// Load environment variables
	hangoutsChatWebhook, ok := os.LookupEnv("NOTIFY_HANGOUTS_CHAT_WEBHOOK")
	if !ok {
		log.Fatal("NOTIFY_HANGOUTS_CHAT_WEBHOOK environment variable is unset")
	}

	payload := fmt.Sprintf(`{"text":"%s"}`, message)

	req, _ := http.NewRequest("POST", hangoutsChatWebhook, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.ReadAll(resp.Body)
	log.Print("Hangouts Chat notified")
}

/*
		# Slack

		Parameters:
		- message

	  Environment variables:
	  - NOTIFY_SLACK_WEBHOOK
*/
func notifySlack(message string) {
	// Load environment variables
	hangoutsChatWebhook, ok := os.LookupEnv("NOTIFY_SLACK_WEBHOOK")
	if !ok {
		log.Fatal("NOTIFY_SLACK_WEBHOOK environment variable is unset")
	}

	payload := fmt.Sprintf(`{"text":"%s"}`, message)

	req, _ := http.NewRequest("POST", hangoutsChatWebhook, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.ReadAll(resp.Body)
	log.Print("Slack notified")
}

/*
		# Discord

		Parameters:
		- message

	  Environment variables:
	  - NOTIFY_DISCORD_WEBHOOK
*/
func notifyDiscord(message string) {
	// Load environment variables
	hangoutsChatWebhook, ok := os.LookupEnv("NOTIFY_DISCORD_WEBHOOK")
	if !ok {
		log.Fatal("NOTIFY_DISCORD_WEBHOOK environment variable is unset")
	}

	payload := fmt.Sprintf(`{"content":"%s"}`, message)

	req, _ := http.NewRequest("POST", hangoutsChatWebhook, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.ReadAll(resp.Body)
	log.Print("Discord notified")
}

func main() {
	// Subcommands
	telegramCmd := flag.NewFlagSet("telegram", flag.ExitOnError)
	telegramMessage := telegramCmd.String("message", "This is a test message", "message")

	hangoutsChatCmd := flag.NewFlagSet("hangoutschat", flag.ExitOnError)
	hangoutsChatMessage := hangoutsChatCmd.String("message", "This is a test message", "message")

	slackCmd := flag.NewFlagSet("slack", flag.ExitOnError)
	slackMessage := slackCmd.String("message", "This is a test message", "message")

	discordCmd := flag.NewFlagSet("discord", flag.ExitOnError)
	discordMessage := discordCmd.String("message", "This is a test message", "message")

	if len(os.Args) < 2 {
		log.Println("Expected 'telegram', 'hangoutschat', 'slack', or 'discord'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "telegram":
		telegramCmd.Parse(os.Args[2:])
		notifyTelegram(*telegramMessage)
	case "hangoutschat":
		hangoutsChatCmd.Parse(os.Args[2:])
		notifyHangoutsChat(*hangoutsChatMessage)
	case "slack":
		slackCmd.Parse(os.Args[2:])
		notifySlack(*slackMessage)
	case "discord":
		discordCmd.Parse(os.Args[2:])
		notifyDiscord(*discordMessage)
	default:
		log.Println("Expected 'telegram', 'hangoutschat', 'slack', or 'discord'")
		os.Exit(1)
	}
}
