package fluency

import (
	"fmt"
)

/*
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
}*/

func testMetricResourceTagList(client *FluencyClient, resourceTypes []string) {
	tags, err := client.MetricResourceTagKeys(resourceTypes)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(tags)
}

func testMetricResourceTagValues(client *FluencyClient, resourceTypes []string, key string, pattern string) {
	tags, err := client.MetricResourceTagValues(resourceTypes, key, pattern)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(tags)
}

func testMetricTestTagFilter(client *FluencyClient, resourceTypes []string, key string, value string) {
	result, err := client.MetricTestTagFilter(resourceTypes, key, value)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(result)
}

func testMetricPollTest(client *FluencyClient, namespace string, metric string, dimensions []string, resourceTypes []string, key string, value string) {
	result, err := client.MetricPollTest(namespace, metric, dimensions, resourceTypes, key, value)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(result)
}
func testMetricResourceTagAPI(client *FluencyClient) {
	//testMetricTagList(client, "", "InstanceId")
	//testMetricResourceTagList(client, []string{"ec2:instance"})
	//testMetricResourceTagValues(client, []string{"ec2:instance"}, "lvdb-app", "analytic")

	//testMetricTestTagFilter(client, []string{"ec2:instance"}, "lvdb-app", "analytic")
	testMetricPollTest(client, "AWS/EC2", "CPUUtilization", []string{"InstanceId"}, []string{"ec2:instance"}, "lvdb-app", "analytic")

	// testMetricTagSearch(client, "AWS.EC2.CPUUtilization", "", "metric_stream", "")
	//testMetricTagSearch(client, "AWS.ApplicationELB.UnHealthyHostCount", "", "fluencyAccountID", "")
	//testMetricIDSearch(client, "AWS.ApplicationELB.UnHealthyHostCount", "torn")
}
