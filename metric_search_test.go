package fluency

import (
	"testing"

	"github.com/SecurityDo/fluency-go/model"
)

func testSearch(client *FluencyClient) {

	//client.MetricSearch("metric(\"AWS.EC2.CPUUtilization\")", nil)

	options := &model.MetricSearchOptions{}

	aggOptions := &model.AggregationOptions{}

	options.Aggregations = aggOptions

	aggOptions.MustFilters = append(aggOptions.MustFilters, &model.MetricFilterEntry{
		Field: "region",
		Terms: []string{"us-east-1", "us-east-2"},
	})

	/*
		aggOptions.Facets = append(aggOptions.Facets, &model.MetricFacetEntry{
			Field: "lvdb-app",
		})
		aggOptions.Facets = append(aggOptions.Facets, &model.MetricFacetEntry{
			Field: "lvdb-account",
		}) */
	client.MetricSearch("metric(\"AWS.EC2.CPUUtilization\")", options)
}

func TestMetricSearch(t *testing.T) {

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
	// testMetricTagAPI(client)

	//testMetricPollAPI(client)
	testSearch(client)
}
