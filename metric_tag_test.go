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

func testMetricTagSearch(client *FluencyClient, metricBucket string, dimension string, tag string, prefix string) {
	tags, err := client.MetricTagSearch(metricBucket, dimension, tag, prefix)
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

	testMetricTagSearch(client, "AWS.EC2.CPUUtilization", "", "Name", "audit")
}
