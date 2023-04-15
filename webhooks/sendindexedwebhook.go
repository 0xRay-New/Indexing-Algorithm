package webhooks

import (
	"calculator/common"
	"strconv"
	"time"

	"github.com/akulsharma1/godiscord"
)

var indexWebhook = "webhoook removed"
func SendIndexedCollectionWebhook(data common.CollectionStruct) {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	ts := time.Now().In(loc).Format("2006-01-02 15:04:05")

	e := godiscord.NewEmbed("Indexed Collection "+data.CollectionName+" - "+data.Address, "Indexed "+strconv.Itoa(data.IndexCount)+"/"+strconv.Itoa(data.Count), "https://opensea.io/collection/"+data.CollectionName)
	if data.ImageURL != "" {
		e.SetThumbnail(data.ImageURL)
	}
	e.AvatarURL = "https://media.discordapp.net/attachments/722560149511995394/996875875511455856/bigicon.png?width=550&height=550"
	e.Username = "0xRay Helper"
	e.SetColor("008DBA")
	e.SetFooter("0xRay Software • "+"discord.gg/6DdBZHZsZa"+" • "+ts, "https://media.discordapp.net/attachments/722560149511995394/996875875511455856/bigicon.png?width=550&height=550")
	e.SendToWebhook(indexWebhook)

}