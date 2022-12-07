package fluency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

type clientConfig struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

func newClientConfig(configFile string) (*clientConfig, error) {
	r := new(clientConfig)
	fileBytes, e := ioutil.ReadFile(configFile)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return nil, e
	}
	jsonError := json.Unmarshal(fileBytes, r)
	if jsonError != nil {
		fmt.Println("client parsing error:", jsonError.Error())
		return nil, jsonError
	}
	return r, nil
}

func loadLocalFile(filename string) string {
	_, filename, _, _ = runtime.Caller(0)
	folder := path.Dir(filename)
	fmt.Println("Current folder: ", folder)
	localFile := path.Join(folder, filename)
	return localFile
}

func getFplReport(client *FluencyClient) {
	entry, err := client.GetFPLReport("O365_Administrator_Listing")
	if err != nil {
		panic(err.Error())
	}
	PrettyPrintJSON(entry)

}

func listFplReport(client *FluencyClient) {
	entries, err := client.ListFPLReport()
	if err != nil {
		panic(err.Error())
	}
	for _, entry := range entries {
		PrettyPrintJSON(entry)
	}

}

func Test(t *testing.T) {

	// handle := NewGeoHandle()
	clientConfig, err := newClientConfig("/etc/api_test_config.json")
	if err != nil {
		panic(err.Error())
	}
	client := NewFluencyClient(clientConfig.URL, clientConfig.Token)

	//client := NewFluencyClient("https://terpvue.cloud.fluencysecurity.com", "")
	// getFplReport(client)
	// listFplReport(client)
	//metricAWSListAll(client)
	// testGroupAPI(client)
	//testMetricNotification(client)
	//testMetricTagAPI(client)
	// testMetricIncidentAPI(client)
	// testMetricRuleTemplateAPI(client)
	testMetricTagAPI(client)
}
