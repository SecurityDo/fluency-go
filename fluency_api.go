package fluency

import (
	"encoding/json"
	"fmt"

	"github.com/SecurityDo/fluency-go/fsb"
	"github.com/SecurityDo/fluency-go/model"
)

func (r *FluencyClient) GetFPLReport(name string) (entry *model.FPLReport, err error) {

	var getFPLReportRequest struct {
		Name string `json:"name,omitempty"`
	}

	getFPLReportRequest.Name = name

	jnode, _ := fsb.NewJNodeInterface(getFPLReportRequest)
	res, err := r.serviceClient.Call("api/ds", "get_fpl_report", jnode)
	if err != nil {
		fmt.Println("fail to parse get_fpl_report response!")
		return nil, err
	}
	var ruleRes struct {
		Entry *model.FPLReport `json:"entry,omitempty"`
	}
	err = json.Unmarshal(res.GetBytes(), &ruleRes)
	if err != nil {
		fmt.Println("fail to parse fpl_report_dao response!")
		return nil, err
	}
	return ruleRes.Entry, nil

}
