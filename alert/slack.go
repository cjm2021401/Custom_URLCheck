package alert

import (
	"log"
	"mayf_url_monitor/config"
	"mayf_url_monitor/model"
	"strconv"

	"github.com/slack-go/slack"
)

func URLAlert(urlinfo model.UrlInfo, response model.Response, response_err error) {
	for _, a := range urlinfo.Alert {
		switch {
		case a == "2xx" && (response.StatusCode >= 200 && response.StatusCode < 300):
			SlackAlert(response)
		case a == "4xx" && (response.StatusCode >= 400 && response.StatusCode < 500):
			SlackAlert(response)
		case a == "5xx" && (response.StatusCode >= 500 && response.StatusCode < 600):
			SlackAlert(response)

		case a == "unknown" && (response_err != nil):
			UnknownAlert(response, response_err)
		}
	}
}

func SlackAlert(response model.Response) {

	attachment := slack.Attachment{
		Pretext: response.Name + "'s Status Code is " + strconv.FormatInt(int64(response.StatusCode), 10),
		Text:    "URL :" + response.Url + "\nResponse Time: " + strconv.FormatFloat(response.ResponseTime, 'f', -1, 32) + "\n StatusCode : " + strconv.FormatInt(int64(response.StatusCode), 10),
	}
	api := slack.New(config.Env.Slack.Token)
	_, _, err := api.PostMessage(
		config.Env.Slack.Channel,
		slack.MsgOptionText("", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(false),
	)
	if err != nil {
		log.Println("Cant Send Slack Message wit ", err, response)
	}

}

func UnknownAlert(response model.Response, responseErr error) {

	attachment := slack.Attachment{
		Pretext: response.Name + "' is UnknownState",
		Text:    "URL :" + response.Url + "\nResponse Time: " + strconv.FormatFloat(response.ResponseTime, 'f', -1, 32) + "\nStatus: " + responseErr.Error(),
	}
	api := slack.New(config.Env.Slack.Token)
	_, _, err := api.PostMessage(
		config.Env.Slack.Channel,
		slack.MsgOptionText("", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(false),
	)
	if err != nil {
		log.Println("Cant Send Slack Message wit ", err, response)
	}
}
