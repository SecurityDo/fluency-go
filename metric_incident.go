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

func testMetricAlertListRequest(client *FluencyClient) {

	result, err := client.MetricAlertList()
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}

	PrettyPrintJSON(result)
}

/*
	{
	   "signals": {
	      "key": "fluency-feed-redist-techdata.fifo",
	      "values": [
	         1, 1, 1, 0, 0, 0
	      ],
	      "flags": [
	         1, 1, 0, 0, 0, 0, 0, 0
	       ]
	   },
	   "metrics": [
	      {
	         "bucket": "AWS.SQS.ApproximateAgeOfOldestMessage",
	         "unit": "Seconds",
	         "values": [
	            345370, 345431, 345486, 345571,
	            345457, 0, 0, 0, 0,0
	         ],
	         "flags": [
	            1, 1, 1, 1, 1, 0, 0, 0, 0, 0
	         ]
	      }
	   ],
	   "slotCount": 481,
	   "slots": [
	      1670069040,
	      1670069100,
	      1670097780,
	      1670097840
	   ],
	   "rangeFrom": 1670069040,
	   "rangeTo": 1670097840
	}
*/
func testMetricAlertGetRequest(client *FluencyClient) {

	// "20221202_SQSQueueBacklog_fluency-feed-redist-techdata.fifo"
	// "20221202_SQSConsumerStopped_s3_notification_fluency-metricstream-uswest2"
	result, err := client.MetricAlertGet("20221202_SQSConsumerStopped_s3_notification_fluency-metricstream-uswest2")
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(result)
}

func testMetricIncidentAPI(client *FluencyClient) {

	// testMetricIncidentSearch(client)
	// testMetricIncidentRequest(client)
	testMetricAlertGetRequest(client)
}
