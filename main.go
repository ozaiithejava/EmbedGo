// main.go
package main

import (
	"fmt"
	"yourmodulepath/discordembed" // Değiştirmeniz gereken kısım
)

func main() {
	// Replace 'YOUR_DISCORD_WEBHOOK_URL' with your actual Discord Webhook URL
	webhookURL := "YOUR_DISCORD_WEBHOOK_URL"
	embed := discordembed.NewDiscordEmbed(webhookURL)

	// Example usage:
	embedOptions := discordembed.EmbedOptions{
		Title:       "Test Embed",
		Description: "This is a test embed message.",
		Color:       0x00ff00, // Green color
		Fields: []discordembed.EmbedField{
			{Name: "Field 1", Value: "Value 1"},
			{Name: "Field 2", Value: "Value 2"},
		},
	}

	err := embed.SendEmbed(embedOptions)
	if err != nil {
		fmt.Println("Error sending embed:", err)
		return
	}

	fmt.Println("Embed message sent successfully.")
}
