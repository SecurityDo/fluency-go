package fluency

import (
	"fmt"

	"github.com/SecurityDo/fluency-go/model"
)

func testMetricNotificationActorList(client *FluencyClient) {
	entries, err := client.MetricNotificationListActor()
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	for _, entry := range entries {
		PrettyPrintJSON(entry)
	}
}

func testMetricNotificationGetEndpoint(client *FluencyClient) {
	entry, err := client.MetricNotificationGetEndpoint("")
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(entry)
}

func testMetricNotificationSetDefaultSlackEndpoint(client *FluencyClient) {
	entry, err := client.MetricNotificationGetDefaultEndpoint("Slack", "Message")
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(entry)
	entry.ActorName = "Slack_FluencyDemo"
	entry.SetFieldDefault("channel", "#fluency_grid")
	err = client.MetricNotificationAddEndpoint(entry)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}

}

func testMetricNotificationSetDefaultPagerDutyEndpoint(client *FluencyClient) {
	entry, err := client.MetricNotificationGetDefaultEndpoint("PagerDuty", "Event")
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(entry)
	entry.ActorName = "PagerDuty_Fluency-internal"
	// entry.SetFieldDefault("channel", "#fluency_grid")
	err = client.MetricNotificationAddEndpoint(entry)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}

}
func testMetricNotificationGetDefaultEndpoint(client *FluencyClient) {
	entry, err := client.MetricNotificationGetDefaultEndpoint("Email", "Email")
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(entry)
	// first get default actor endpoint config
	// entry.Name =""
	entry.ActorName = "Email"
	entry.SetFieldDefault("to", "kun@security.do")
	err = client.MetricNotificationAddEndpoint(entry)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}

}

func testMetricNotificationDeleteEndpoint(client *FluencyClient, name string) {
	//err := client.MetricNotificationDeleteEndpoint("Email-Default")
	err := client.MetricNotificationDeleteEndpoint(name)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}

}

func testMetricNotificationAddAction(client *FluencyClient) {
	entry := &model.MetricAlertAction{
		ID: "12334",
		Patterns: []string{
			"ec2.cpu.fluency.highUsage",
		},
		Actions: []string{
			"trigger", "resolve",
		},
		Endpoints: []string{
			"Email-Default",
		},
	}
	err := client.MetricNotificationAddAction(entry)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}

}

func testMetricNotificationAddSlackAction(client *FluencyClient) {
	entry := &model.MetricAlertAction{
		ID: "12335",
		Patterns: []string{
			"ec2.cpu.fluency.highUsage",
		},
		Actions: []string{
			"trigger", "resolve",
		},
		Endpoints: []string{
			"Slack-Message-Default",
		},
	}
	err := client.MetricNotificationAddAction(entry)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}

}

func testMetricNotificationAddPagerDutyAction(client *FluencyClient) {
	entry := &model.MetricAlertAction{
		ID: "12336",
		Patterns: []string{
			"ec2.cpu.fluency.highUsage",
		},
		Actions: []string{
			"trigger", "resolve",
		},
		Endpoints: []string{
			"PagerDuty-Event-Default",
		},
	}
	err := client.MetricNotificationAddAction(entry)
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}

}
func testMetricNotification(client *FluencyClient) {
	// testMetricNotificationActorList(client)
	// testMetricNotificationGetEndpoint(client)
	//testMetricNotificationGetDefaultEndpoint(client)
	//testMetricNotificationAddAction(client)
	testMetricNotificationDeleteEndpoint(client, "PagerDuty-Event-Default")
	//testMetricNotificationDeleteEndpoint(client, "Slack-Message")
	//testMetricNotificationSetDefaultSlackEndpoint(client)
	//testMetricNotificationAddSlackAction(client)
	testMetricNotificationSetDefaultPagerDutyEndpoint(client)
	//testMetricNotificationAddPagerDutyAction(client)
}
