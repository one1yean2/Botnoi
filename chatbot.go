package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/httphandler"
)

type Characters []struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Origin  string `json:"origin"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Gender  string `json:"gender"`
	Species string `json:"species"`
}

func callGameAPI() (Characters, error) {

	var characters Characters
	url := "https://api.sampleapis.com/rickandmorty/characters"

	resp, err := http.Get(url)
	if err != nil {
		return characters, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&characters); err != nil {
		return characters, err
	}

	return characters, nil

}
func main() {

	handler, err := httphandler.New(
		"2861992c4b99c99ac2e2685650b4d1f3",
		"/rfGvM8g9fRbwl7IZkDeaWsM2JU+fm1owagFT3nvskCmNTEJ0W2D0zFNB8ElopfIglEqoznBWG1nhrLKhA8QUHQUwxG15BcwoNpzjzqdRXFBYPYS3OasQCxMMgUHsyRx2Rn8boWU7Y2CIo+0F8AqpAdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			log.Print(err)
			return
		}
		var quickReply []*linebot.QuickReplyButton
		quickReply = append(quickReply, linebot.NewQuickReplyButton("https://webstockreview.net/images/smartphone-icon-png-6.png", linebot.NewMessageAction("ติดต่อ", "ติดต่อ")))
		quickReply = append(quickReply, linebot.NewQuickReplyButton("https://th.bing.com/th/id/OIP.oMyTWRvmNfKDUi0qGownGwHaHa?w=893&h=894&rs=1&pid=ImgDetMain", linebot.NewMessageAction("Rick & Morty", "ตัวละคร R&M")))
		var characters, _ = callGameAPI()

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if message.Text == "เริ่มนต้น" {

						var textMessages []linebot.SendingMessage
						textMessages = append(textMessages,
							linebot.NewTextMessage("สวัสดีครับผม"))
						textMessages = append(textMessages,
							linebot.NewTextMessage("นี่คือร้านการ์ตูน NongYean(น้องยีนส์) ครับ\nกดปุ่มข้างล่างเพื่อดำเนินการต่อได้เลยนะครับ").WithQuickReplies(&linebot.QuickReplyItems{quickReply}))

						if _, err := bot.ReplyMessage(event.ReplyToken, textMessages...).Do(); err != nil {
							log.Print(err)
						}
					}
					if message.Text == "ติดต่อ" {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ติดต่อน้องยีนส์ได้ที่ เบอร์ : 082-6351247").WithQuickReplies(&linebot.QuickReplyItems{quickReply})).Do(); err != nil {
							log.Print(err)
						}
					}
					if message.Text == "ตัวละคร R&M" {

						var columns []*linebot.CarouselColumn

						for i := 0; i < 10; i++ {
							columns = append(columns, linebot.NewCarouselColumn(characters[i].Image, characters[i].Name,
								characters[i].Type+" "+characters[i].Origin,
								linebot.NewMessageAction("รายละเอียด", "รายละเอียด "+characters[i].Name)))
						}
						carousel := linebot.NewCarouselTemplate(columns...)
						template := linebot.NewTemplateMessage("Rick & Morty", carousel)
						if _, err := bot.ReplyMessage(event.ReplyToken, template).Do(); err != nil {
							log.Print(err)
						}
					}
					if strings.HasPrefix(message.Text, "รายละเอียด") {

						characterName := strings.TrimPrefix(message.Text, "รายละเอียด ")
						for i := 0; i < 10; i++ {
							if characters[i].Name == characterName {
								var textMessages []linebot.SendingMessage
								textMessages = append(textMessages,
									linebot.NewTextMessage("ชื่อ : "+characters[i].Name))
								textMessages = append(textMessages,
									linebot.NewTextMessage("เผ่า : "+characters[i].Type))
								textMessages = append(textMessages,
									linebot.NewTextMessage("ที่อยู่ : "+characters[i].Origin))
								textMessages = append(textMessages,
									linebot.NewTextMessage("เพศ : "+characters[i].Gender))
								textMessages = append(textMessages,
									linebot.NewTextMessage("สถานะ : "+characters[i].Status).WithQuickReplies(&linebot.QuickReplyItems{quickReply}))
								if _, err := bot.ReplyMessage(event.ReplyToken, textMessages...).Do(); err != nil {
									log.Print(err)
								}
								break
							}
						}
					}
				}
			}
		}
	})

	http.Handle("/callback", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
