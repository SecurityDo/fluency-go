package fluency

import (
	"fmt"
)

func testMetricRuleTemplateList(client *FluencyClient) {
	entries, err := client.MetricListRuleTemplate()
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	for _, e := range entries {
		PrettyPrintJSON(e)
	}
}

func testMetricRuleTemplateGet(client *FluencyClient) {
	entry, err := client.MetricGetRuleTemplate("SuddenChange")
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	PrettyPrintJSON(entry)
}

func testMetricAddRuleFromTemplate(client *FluencyClient) {
	entry, err := client.MetricGetRuleTemplate("Ec2InstanceError")
	if err != nil {
		fmt.Println(err.Error())
		return
		// panic(err.Error())
	}
	//PrettyPrintJSON(entry)

	entry.Filter.Default = "tag(\"lvdb-app\") == \"analytic\""
	client.MetricAddRuleFromTemplate("FluencyInstanceError", "check fluency nodes", entry.Severity.Default, entry)

}

func testMetricRuleTemplateAPI(client *FluencyClient) {

	// testMetricRuleTemplateList(client)
	// testMetricRuleTemplateGet(client)

	testMetricAddRuleFromTemplate(client)
}
