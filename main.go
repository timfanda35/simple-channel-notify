package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramToken)
	payload := fmt.Sprintf(`{"chat_id": "%s","text":"%s"}`, telegramChatID, message)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			panic(err)
	}
	defer resp.Body.Close()

	ioutil.ReadAll(resp.Body)
	log.Print("Telegram notified")
}

func main() {
	// Subcommands
	telegramCmd := flag.NewFlagSet("telegram", flag.ExitOnError)
	telegramMessage := telegramCmd.String("message", "This is a test message", "message")

	if len(os.Args) < 2 {
		log.Println("Expected 'telegram'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "telegram":
		telegramCmd.Parse(os.Args[2:])
		notifyTelegram(*telegramMessage)
	default:
		log.Println("Expected 'telegram'")
		os.Exit(1)
	}
}
