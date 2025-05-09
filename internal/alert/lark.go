package alert

import (
	"fmt"
	"github.com/dromara/carbon/v2"
	"github.com/go-lark/lark"
	"log"
	"os"
)

var client *lark.Bot

// InitLark 初始化 Lark 客户端
func InitLark(webhookURL string) {
	client = lark.NewNotificationBot(webhookURL)
}

// SendAlert 发送普通告警消息
func SendAlert(title, content string) {
	if client == nil {
		log.Println("Lark client is not initialized")
		return
	}
	SendCardMsg(title, content)
}

func SendCardMsg(title, content string) (*lark.PostNotificationV2Resp, error) {
	if client == nil {
		return nil, fmt.Errorf("lark client is not initialized")
	}

	// 获取当前时间，并格式化
	timeStr := fmt.Sprintf("time: %s", carbon.Now(carbon.Shanghai).ToDateTimeString())

	// 获取本机 hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	b := lark.NewCardBuilder()
	card := b.Card(
		//b.Div(
		//	b.Field(b.Text("整排内容")),
		//	b.Field(b.Text("整排**Markdown**内容").LarkMd()),
		//),
		b.Div().
			Text(b.Text(content).LarkMd()),
		b.Hr(),
		b.Note().
			AddText(b.Text(fmt.Sprintf("%s (%s)", timeStr, hostname)).LarkMd()),
	).
		Wathet().
		Title(title)
	msgV4 := lark.NewMsgBuffer(lark.MsgInteractive)
	omV4 := msgV4.Card(card.String()).Build()
	res, err := client.PostNotificationV2(omV4)
	return res, err

}
