package fluency

import (
	"fmt"
)

func testMetricTagList(client *FluencyClient, metricBucket string, dimension string) {
	tags, err := client.MetricTagList(metricBucket, dimension)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(tags)

}

func testMetricTagSearch(client *FluencyClient, metricBucket string, dimension string, tag string, pattern string) {
	tags, err := client.MetricTagSearch(metricBucket, dimension, tag, pattern)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(tags)
}

func testMetricIDSearch(client *FluencyClient, metricBucket string, pattern string) {
	tags, err := client.MetricIDSearch(metricBucket, pattern)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(tags)
}

func testMetricTagAPI(client *FluencyClient) {
	//testMetricTagList(client, "", "InstanceId")
	//testMetricTagList(client, "AWS.EC2.CPUUtilization", "")

	//testMetricTagSearch(client, "AWS.EC2.CPUUtilization", "", "Name", "audit")
	//testMetricTagSearch(client, "AWS.ApplicationELB.UnHealthyHostCount", "", "fluencyAccountID", "")
	testMetricIDSearch(client, "AWS.ApplicationELB.UnHealthyHostCount", "torn")
}
