# EmbedGo
embed messages in go for discord


## Kullanım:
```Go
package main

import (
	"fmt"
	"import in my github repo"
)

func main() {
	// Discord Webhook URL'sini belirtin
	webhookURL := "YOUR_DISCORD_WEBHOOK_URL"
	embed := discordembed.NewDiscordEmbed(webhookURL)

	// EmbedOptions kullanarak basit bir embed mesajı gönderme
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
```