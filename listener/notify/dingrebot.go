package notify

import (
	"encoding/json"
	"fmt"
	"github.com/ouqiang/supervisor-event-listener/event"
	"github.com/royeo/dingrobot"
)

type DingRobot struct{}

func (hook *DingRobot) Send(message event.Message) error {
	encodeMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}
	robot := dingrobot.NewRobot(Conf.WebHook.Url)
	title := "Nidus-Guard"
	text := fmt.Sprintf("```json\n %s ```", encodeMessage)
	atMobiles := []string{}
	isAtAll := false

	err = robot.SendMarkdown(title, text, atMobiles, isAtAll)

	return err
}
