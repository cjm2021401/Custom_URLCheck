package agent

import (
	"encoding/json"
	"io/ioutil"
	"mayf_url_monitor/model"
	"os"
)

func GetUrlList() (model.UrlInfoList, error) {
	urlInfoList := model.UrlInfoList{}
	path, _ := os.Getwd()
	info, err := os.Open(path + "/url.json")
	if err != nil {
		return urlInfoList, err
	}
	defer info.Close()

	bytes, _ := ioutil.ReadAll(info)
	err = json.Unmarshal(bytes, &urlInfoList)
	if err != nil {
		return urlInfoList, err
	}
	return urlInfoList, nil

}
