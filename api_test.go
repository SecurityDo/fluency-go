package fluency

import "testing"

func Test(t *testing.T) {

	// handle := NewGeoHandle()
	client := NewFluencyClient("https://terpvue.cloud.fluencysecurity.com", "")
	entry, err := client.GetFPLReport("O365_Administrator_Listing")
	if err != nil {
		panic(err.Error())
	}
	PrettyPrintJSON(entry)
	// testInvite(handle)
}
