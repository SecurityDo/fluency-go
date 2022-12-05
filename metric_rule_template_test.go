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

func testMetricRuleTemplateAPI(client *FluencyClient) {

	// testMetricRuleTemplateList(client)
	testMetricRuleTemplateGet(client)
}
