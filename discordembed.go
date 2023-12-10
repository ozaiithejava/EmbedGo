// discordembed.go
package discordembed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// EmbedField struct represents a field in the embed message.
type EmbedField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// EmbedOptions struct represents options for creating an embed message.
type EmbedOptions struct {
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Color       int          `json:"color,omitempty"`
	Fields      []EmbedField `json:"fields,omitempty"`
}

// DiscordEmbed struct represents the Discord Embed module.
type DiscordEmbed struct {
	webhookURL string
}

// NewDiscordEmbed creates a new DiscordEmbed instance.
func NewDiscordEmbed(webhookURL string) *DiscordEmbed {
	return &DiscordEmbed{webhookURL: webhookURL}
}

// SendEmbed sends an embed message to Discord.
func (de *DiscordEmbed) SendEmbed(embedOptions EmbedOptions) error {
	payload := map[string]interface{}{
		"embeds": []EmbedOptions{embedOptions},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(de.webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
