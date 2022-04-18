package mattermost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/megaease/easeprobe/global"
	"github.com/megaease/easeprobe/notify/base"
	"github.com/megaease/easeprobe/report"
	"log"
	"net/http"
)

// NotifyConfig is the slack notification configuration
type NotifyConfig struct {
	base.DefaultNotify `yaml:",inline"`
	WebhookURL         string `yaml:"webhook"`
}

// Kind return the type of Notify
func (c *NotifyConfig) Kind() string {
	return c.MyKind
}

// Config configures the slack notification
func (c *NotifyConfig) Config(gConf global.NotifySettings) error {
	c.MyKind = "mattermost"
	c.Format = report.Mattermost
	c.SendFunc = c.SendMattermostAlert
	c.DefaultNotify.Config(gConf)
	log.Printf("Notification [%s] - [%s] configuration: %+v", c.MyKind, c.Name, c)
	return nil
}

type MattermostMsg struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

func (c *NotifyConfig) SendMattermostAlert(title, msg string) error {
	log.Println(fmt.Sprintf("send error msg to mattermost: %v: %s", title, msg))
	p := MattermostMsg{
		Username: "QEX Monitor",
		Text:     fmt.Sprintf(":alert: %s", msg),
	}
	payload, _ := json.Marshal(p)
	if _, err := http.Post(c.WebhookURL, "application/json", bytes.NewBuffer(payload)); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}