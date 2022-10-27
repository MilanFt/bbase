package discord

import (
	"errors"
	"strings"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/snowflake/v2"
)

var (
	webhookID     snowflake.ID
	webhookToken  string
	WebhookClient webhook.Client
)

func InitializeWebhook(url string) error {
	urlSplit := strings.Split(url, "/")
	var err error
	webhookID, err = snowflake.Parse(urlSplit[5])
	if err != nil {
		return err
	}
	webhookToken = urlSplit[6]
	WebhookClient = webhook.New(webhookID, webhookToken)
	return nil
}

func SuccessWebhook(
	productImg string,
	site string,
	product string,
	size string,
	price string,
	email string,
) error {
	if WebhookClient == nil {
		return errors.New("webhook client is nil")
	}

	_, err := WebhookClient.CreateEmbeds(
		[]discord.Embed{
			discord.NewEmbedBuilder().
				SetTitle("Successful Checkout").
				SetThumbnail(productImg).
				SetColor(632534).
				AddField("Site", site, true).
				AddField("Product", product, true).
				AddField("Size", size, true).
				AddField("Price", price, true).
				AddField("Email", "||"+email+"||", true).
				// Footer details, change it to your bot's name and icon
				SetFooter("BBase", "https://freepngimg.com/save/15778-penguin-png-images/999x1278").
				SetTimestamp(time.Now()).
				Build(),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func FailedWebhook(
	productImg string,
	site string,
	product string,
	size string,
	price string,
	email string,
) error {
	if WebhookClient == nil {
		return errors.New("webhook client is nil")
	}

	_, err := WebhookClient.CreateEmbeds(
		[]discord.Embed{
			discord.NewEmbedBuilder().
				SetTitle("Failed Checkout").
				SetThumbnail(productImg).
				SetColor(14027017).
				AddField("Site", site, true).
				AddField("Product", product, true).
				AddField("Size", size, true).
				AddField("Price", price, true).
				AddField("Email", "||"+email+"||", true).
				// Footer details, change it to your bot's name and icon
				SetFooter("BBase", "https://freepngimg.com/save/15778-penguin-png-images/999x1278").
				SetTimestamp(time.Now()).
				Build(),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
