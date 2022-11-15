package fluency

import "testing"

func getFplReport(client *FluencyClient) {
	entry, err := client.GetFPLReport("O365_Administrator_Listing")
	if err != nil {
		panic(err.Error())
	}
	PrettyPrintJSON(entry)

}

func listFplReport(client *FluencyClient) {
	entries, err := client.ListFPLReport()
	if err != nil {
		panic(err.Error())
	}
	for _, entry := range entries {
		PrettyPrintJSON(entry)
	}

}

func Test(t *testing.T) {

	// handle := NewGeoHandle()
	client := NewFluencyClient("https://terpvue.cloud.fluencysecurity.com", "")
	// testInvite(handle)
	// getFplReport(client)
	listFplReport(client)
}
