package appmanage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

func ReadConfig2Map(project string) map[string]interface{} {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := path.Join(pwd, "app", project, "/configs/config.json")
	config := make(map[string]interface{})
	file, err := os.Open(path)
	if err != nil {
		panic(nil)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(nil)
	}
	if err := json.Unmarshal(content, &config); err != nil {
		panic(err)
	}
	return config
}
