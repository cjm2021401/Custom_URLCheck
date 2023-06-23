package agent

import (
	"log"
	"mayf_url_monitor/alert"
	"mayf_url_monitor/influx"
	"mayf_url_monitor/model"
	"net/http"
	"strings"
	"time"
)

func Check_URlStatus(urlInfo model.UrlInfo, useDB bool, useSlack bool) {
	url := MakeUrl(urlInfo.Url)
	timeout := MakeTimeStamp(urlInfo.Timeout)

	response := model.Response{}
	response.Name = urlInfo.Name
	response.Url = url

	start := time.Now()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	resp, err := client.Do(req)
	elapsed := time.Since(start).Seconds()
	response.ResponseTime = elapsed
	if err != nil {
		log.Println(urlInfo.Url, "알수없음", err)
		if useDB {
			response.StatusCode = 999
			go influx.URLStatToInflux(response)
		}
		if useSlack {
			go alert.URLAlert(urlInfo, response, err)
		}
		return
	}

	response.StatusCode = int32(resp.StatusCode)
	log.Println(urlInfo.Url, resp.Status)
	if useDB {
		go influx.URLStatToInflux(response)
	}
	if useSlack {
		go alert.URLAlert(urlInfo, response, nil)
	}
	return
}

func MakeUrl(url string) string {
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		return url
	} else {
		return "http://" + url
	}
}

func MakeTimeStamp(timeout int32) int32 {
	if timeout == 0 || timeout > 5 {
		return 5
	} else {
		return timeout
	}
}
