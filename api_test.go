package fluency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"testing"

	"github.com/SecurityDo/fluency-go/model"
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

func metricAWSListAll(client *FluencyClient) {
	groups, metrics, err := client.MetricAWSListAll()
	if err != nil {
		panic(err.Error())
	}
	for _, group := range groups {
		PrettyPrintJSON(group)
	}
	for _, metric := range metrics {
		PrettyPrintJSON(metric)
	}

}

func testAddMetricGroup(client *FluencyClient) {
	group := &model.MetricImportGroup{
		Namespace:     "Fluebcy/Test",
		Category:      "default",
		Interval:      60,
		MinuteEmulate: true,
	}
	err := client.MetricAWSAddGroup(group)
	if err != nil {
		// fmt.Println(err.Error())
		panic(err.Error())
	}
}

func testAddMetric(client *FluencyClient) {
	metric := &model.MetricImportEntry{
		Namespace:   "Fluency/Test",
		Category:    "default",
		Name:        "CPURate",
		Description: "CPURate percent",
		Unit:        "Percent",
		Dimensions: []string{
			"InstanceId",
		},
	}
	err := client.MetricAWSAddMetric(metric)
	if err != nil {
		// fmt.Println(err.Error())
		panic(err.Error())
	}
}

func testDeleteMetric(client *FluencyClient) {
	err := client.MetricAWSDeleteMetric("Fluency.Test.CPURate")
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
}
func testDeleteMetricGroup(client *FluencyClient) {
	err := client.MetricAWSDeleteMetricGroup("Fluency/Test", "default")
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
}
func testGroupAPI(client *FluencyClient) {
	testAddMetricGroup(client)
	testAddMetric(client)
	testDeleteMetric(client)
	testDeleteMetricGroup(client)
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
	testGroupAPI(client)
}
