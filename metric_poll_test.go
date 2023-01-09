package fluency

import (
	"fmt"

	"github.com/SecurityDo/fluency-go/model"
)

func testMetricPollList(client *FluencyClient) {
	polls, err := client.ListMetricPoll()
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(polls)
}

func testMetricPollAdd(client *FluencyClient) {
	poll := &model.MetricPoll{
		Name:          "all",
		Description:   "default poll",
		PollInterval:  60,
		Namespace:     "AWS/EC2",
		Metric:        "CPUUtilization",
		Interval:      60,
		Dimensions:    []string{"InstanceId"},
		Stat:          "Sum",
		MinuteEmulate: true,
		Unit:          "", // optional
	}
	err := client.AddMetricPoll(poll)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
}

func testMetricPollDelete(client *FluencyClient) {
	err := client.DeleteMetricPoll("AWS.EC2.CPUUtilization.poll.all")
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
}
func testMetricPollAPI(client *FluencyClient) {
	//testMetricPollAdd(client)
	//testMetricPollList(client)

	testMetricPollDelete(client)
	testMetricPollList(client)
}
