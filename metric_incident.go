package fluency

import (
	"fmt"
	"time"

	"github.com/SecurityDo/fluency-go/model"
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

}*/

func testMetricIncidentSearch(client *FluencyClient) {
	now := time.Now().Unix() * 1000
	options := &model.SimpleFacetSearchOption{
		SimpleSearchOption: model.SimpleSearchOption{
			RangeFrom:   now - 24*3600000,
			RangeTo:     now,
			RangeField:  "timestamp",
			SortField:   "timestamp",
			SortOrder:   "desc",
			FetchOffset: 0,
			FetchLimit:  2,
		},
		Facets: &model.FacetsOption{},
	}
	facets := []*model.FacetEntry{
		{
			Field: "status",
			Order: "count",
			Size:  20,
		},
		{
			Field: "state",
			Order: "count",
			Size:  20,
		},
		{
			Field: "name",
			Order: "count",
			Size:  20,
		},
		{
			Field: "displayName",
			Order: "count",
			Size:  20,
		},
	}
	dateFacets := []*model.DateFacetEntry{
		{
			Name: "dateHistogram",
			Key:  "timestamp",
		},
	}
	options.Facets.Facets = facets
	options.Facets.DateFacets = dateFacets
	result, err := client.MetricIncidentSearch(options)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(result)
}

func testMetricIncidentRequest(client *FluencyClient) {

	input := &model.MetricIncidentUpdateRequest{
		ID:        "20221130_SQSQueueBacklog_s3_notification_fluency-metricstream-uswest2",
		UpdatedOn: 1669914796000,
		Status:    "acknowledged", // or empty for comment
		Username:  "kun@fluencysecurity.com",
		Comment:   "",
	}

	err := client.MetricIncidentUpdate(input)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}

}
func testMetricIncidentAPI(client *FluencyClient) {

	// testMetricIncidentSearch(client)
	testMetricIncidentRequest(client)

}
