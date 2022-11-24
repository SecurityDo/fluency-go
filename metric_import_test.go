package fluency

import (
	"fmt"

	"github.com/SecurityDo/fluency-go/model"
)

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
		Namespace:     "Fluency/Test",
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
