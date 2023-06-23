package main

import (
	"log"
	"mayf_url_monitor/agent"
	"mayf_url_monitor/config"
	"mayf_url_monitor/influx"
	"time"
)

func main() {
	err := config.GetEnvironmentVariable()
	if err != nil {
		log.Fatal(err)
	}
	checkDB := false
	if config.Env.InfluxDB.Url != "" && config.Env.InfluxDB.Token != "" {
		influx.InfluxInit()
		checkDB = true
	}
	checkSlack := true
	if config.Env.Slack.Token == "" || config.Env.Slack.Channel == "" {
		checkSlack = false
	}

	for {
		urlList, err := agent.GetUrlList()
		if err != nil {
			log.Panic(err)
		}

		for _, urlinfo := range urlList.UrlInfos {
			go agent.Check_URlStatus(urlinfo, checkDB, checkSlack)
		}
		time.Sleep(1 * time.Minute)
	}
}
