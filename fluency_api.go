package fluency

import (
	"encoding/json"
	"fmt"

	"github.com/SecurityDo/fluency-go/model"
	// "github.com/olivere/elastic"
)

func (r *FluencyClient) ListFPLReport() (entries []*model.FPLReport, err error) {

	res, err := r.serviceClient.Call("api/ds", "list_fpl_report", nil)
	if err != nil {
		fmt.Println("fail to parse list_fpl_report response!")
		return nil, err
	}
	var result struct {
		Entries []*model.FPLReport `json:"entries,omitempty"`
	}
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse fpl_report_dao response!")
		return nil, err
	}
	return result.Entries, nil

}
func (r *FluencyClient) GetFPLReport(name string) (entry *model.FPLReport, err error) {

	var getFPLReportRequest struct {
		Name string `json:"name,omitempty"`
	}

	getFPLReportRequest.Name = name

	res, err := r.serviceClient.Call("api/ds", "get_fpl_report", getFPLReportRequest)
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

/*
{
   "namespace": "AWS/AutoScaling",
   "interval": 60,
   "category": "default",
   "minuteEmulate": false,
   "createdOn": "0001-01-01T00:00:00Z"
}
*/

func (r *FluencyClient) MetricAWSListGroups() (groups []*model.MetricImportGroup, err error) {
	res, err := r.serviceClient.Call("api/ds", "metric_aws_list_groups", nil)
	if err != nil {
		fmt.Println("fail to call metric_aws_list_groups:", err.Error())
		return nil, err
	}
	var result model.MetricAWSListGroupsResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_aws_list_groups response!")
		return nil, err
	}
	return result.Groups, nil
}

/*
	{
	   "name": "BurstBalance",
	   "description": "Used with General Purpose SSD (gp2), Throughput Optimized HDD (st1), and Cold HDD (sc1) volumes only. Provides information about the percentage of I/O credits (for gp2) ...",
	   "unit": "Percent",
	   "dimensions": [ "VolumeId" ],
	   "bucket": "AWS.EBS.BurstBalance",
	   "namespace": "AWS/EBS",
	   "category": "default"
	}
*/
func (r *FluencyClient) MetricAWSListMetrics() (metrics []*model.MetricImportEntry, err error) {
	res, err := r.serviceClient.Call("api/ds", "metric_aws_list_metrics", nil)
	if err != nil {
		fmt.Println("fail to call metric_aws_list_metrics:", err.Error())
		return nil, err
	}
	var result model.MetricAWSListMetricsResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_aws_list_metrics response!")
		return nil, err
	}
	return result.Metrics, nil
}

func (r *FluencyClient) MetricAWSListAll() (groups []*model.MetricImportGroup, metrics []*model.MetricImportEntry, err error) {
	res, err := r.serviceClient.Call("api/ds", "metric_aws_list_all", nil)
	if err != nil {
		fmt.Println("fail to call metric_aws_list_all:", err.Error())
		return nil, nil, err
	}
	var result model.MetricAWSListAllResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_aws_list_all response!")
		return nil, nil, err
	}
	return result.Groups, result.Metrics, nil
}

func (r *FluencyClient) MetricAWSAddGroup(group *model.MetricImportGroup) (err error) {

	functionName := "metric_aws_add_group"

	input := &model.MetricAWSAddGroupRequest{
		Group: group,
	}

	_, err = r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}
	return nil
}

func (r *FluencyClient) MetricAWSAddMetric(metric *model.MetricImportEntry) (err error) {

	functionName := "metric_aws_add_metric"

	input := &model.MetricAWSAddMetricRequest{
		Metric: metric,
	}

	_, err = r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}
	return nil
}

func (r *FluencyClient) MetricAWSDeleteMetric(bucketName string) (err error) {

	functionName := "metric_aws_delete_metric"

	input := &model.MetricAWSDeleteMetricRequest{
		Name: bucketName,
	}

	_, err = r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}
	return nil
}

func (r *FluencyClient) MetricAWSDeleteMetricGroup(namespace string, category string) (err error) {

	functionName := "metric_aws_delete_group"

	input := &model.MetricAWSDeleteGroupRequest{
		Namespace: namespace,
		Category:  category,
	}

	_, err = r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}
	return nil
}

func (r *FluencyClient) MetricNotificationListActor() (entries []*model.ActorConfig, err error) {

	functionName := "metric_notification_actor_list"

	res, err := r.serviceClient.Call("api/ds", functionName, nil)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	var result model.MetricNotificationActorListResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_aws_list_all response!")
		return nil, err
	}
	return result.Entries, nil
}

func (r *FluencyClient) MetricNotificationGetEndpoint(name string) (entry *model.EventActionEntry, err error) {

	input := &model.MetricNotificationEndpointDaoRequest{
		Action: "get",
		Args: &model.MetricNotificationEndpointDaoRequestArgs{
			Id: name,
		},
	}

	functionName := "metric_notification_endpoint_dao"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	var result model.MetricNotificationEndpointDaoResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_notification_endpoint_dao response!")
		return nil, err
	}
	return result.Entry, nil
}

// "PagerDuty","Event"
// "Email":"Email"
// "Slack":"Message"
func (r *FluencyClient) MetricNotificationGetDefaultEndpoint(actor string, action string) (entry *model.EventActionEntry, err error) {

	input := &model.MetricNotificationEndpointDaoRequest{
		Action: "get",
		Args: &model.MetricNotificationEndpointDaoRequestArgs{
			Id: "",
			Entry: &model.EventActionEntry{
				Actor:  actor,
				Action: action,
			},
		},
	}

	functionName := "metric_notification_endpoint_dao"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	var result model.MetricNotificationEndpointDaoResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_notification_endpoint_dao response!")
		return nil, err
	}
	return result.Entry, nil
}

func (r *FluencyClient) MetricNotificationAddEndpoint(entry *model.EventActionEntry) (err error) {

	input := &model.MetricNotificationEndpointDaoRequest{
		Action: "add",
		Args: &model.MetricNotificationEndpointDaoRequestArgs{
			Entry: entry,
		},
	}

	functionName := "metric_notification_endpoint_dao"

	_, err = r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

func (r *FluencyClient) MetricNotificationDeleteEndpoint(name string) (err error) {

	input := &model.MetricNotificationEndpointDaoRequest{
		Action: "delete",
		Args: &model.MetricNotificationEndpointDaoRequestArgs{
			Id: name,
		},
	}

	functionName := "metric_notification_endpoint_dao"

	_, err = r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

func (r *FluencyClient) MetricNotificationAddAction(entry *model.MetricAlertAction) (err error) {

	input := &model.MetricNotificationActionDaoRequest{
		Action: "add",
		Args: &model.MetricNotificationActionDaoRequestArgs{
			Entry: entry,
		},
	}

	functionName := "metric_notification_action_dao"

	_, err = r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

func (r *FluencyClient) MetricNotificationDeleteAction(id string) (err error) {

	input := &model.MetricNotificationActionDaoRequest{
		Action: "delete",
		Args: &model.MetricNotificationActionDaoRequestArgs{
			Id: id,
		},
	}

	functionName := "metric_notification_action_dao"

	_, err = r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

// get tag by dimension or one specific metric bucket
func (r *FluencyClient) MetricTagList(bucket string, dimension string) (tags []*model.MetricTag, err error) {

	input := &model.MetricTagsRequest{
		Metric:    bucket,
		Dimension: dimension,
	}

	functionName := "metric_tag_list"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	var result model.MetricTagsResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_tag_list response!")
		return nil, err
	}
	return result.Tags, nil
}

func (r *FluencyClient) MetricTagSearch(bucket string, dimension string, tag string, pattern string) (entries []string, err error) {

	input := &model.MetricTagSearchRequest{
		Metric:    bucket,
		Dimension: dimension,
		Tag:       tag,
		Pattern:   pattern,
	}

	functionName := "metric_tag_search"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	var result model.MetricTagSearchResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_tag_search response!")
		return nil, err
	}
	return result.Entries, nil
}

func (r *FluencyClient) MetricIDSearch(bucket string, pattern string) (entries []string, err error) {

	input := &model.MetricIDSearchRequest{
		Metric:  bucket,
		Pattern: pattern,
	}

	functionName := "metric_id_search"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	var result model.MetricIDSearchResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_id_search response!")
		return nil, err
	}
	return result.Entries, nil
}

func (r *FluencyClient) MetricIncidentSearch(options *model.SimpleFacetSearchOption) (result *model.MetricIncidentSearchResponse, err error) {

	input := &model.MetricIncidentSearchRequest{
		Options: options,
	}

	functionName := "metric_incident_search"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	//var result model.MetricIncidentSearchResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_incident_search response!")
		return nil, err
	}
	return result, nil
}

func (r *FluencyClient) MetricIncidentUpdate(input *model.MetricIncidentUpdateRequest) (err error) {

	functionName := "metric_incident_update"

	_, err = r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

func (r *FluencyClient) MetricAlertList() (result *model.MetricAlertListResponse, err error) {

	// limit default is 30
	// default range is last 8 hours
	input := &model.MetricAlertListRequest{}

	functionName := "metric_alert_list"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	// var result model.MetricAlertListResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_alert_list response!")
		return nil, err
	}
	return result, nil
}

// fullkey  $dayIndex_$alert_$key
func (r *FluencyClient) MetricAlertGet(fullkey string) (result *model.MetricAlertGetResponse, err error) {
	// default range is last 8 hours
	input := &model.MetricAlertGetRequest{
		ID: fullkey,
	}

	functionName := "metric_alert_get"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	// var result model.MetricAlertListResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_alert_list response!")
		return nil, err
	}
	return result, nil
}

func (r *FluencyClient) MetricListRuleTemplate() (entries []*model.RuleTemplate, err error) {

	functionName := "metric_list_rule_template"

	res, err := r.serviceClient.Call("api/ds", functionName, nil)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	var result model.MetricListRuleTemplateResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_aws_list_all response!")
		return nil, err
	}
	return result.Entries, nil
}

func (r *FluencyClient) MetricGetRuleTemplate(name string) (entry *model.RuleTemplate, err error) {

	input := &model.MetricGetRuleTemplateRequest{
		Name: name,
	}

	functionName := "metric_get_rule_template"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	var result model.MetricGetRuleTemplateResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_aws_list_all response!")
		return nil, err
	}
	return result.Entry, nil
}

func (r *FluencyClient) MetricAddRuleFromTemplate(name string, desc string, severity string, template *model.RuleTemplate) (err error) {

	input := &model.MetricAddRuleFromTemplateRequest{
		Name:        name,
		Description: desc,
		Severity:    severity,
		Template:    template,
	}

	functionName := "metric_add_rule_from_template"

	_, err = r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

func (r *FluencyClient) ListMetricPoll() (entries []*model.MetricPoll, err error) {

	input := &model.MetricPollDaoRequest{
		Action: "list",
		Args:   &model.MetricPollDaoRequestArgs{},
	}
	res, err := r.serviceClient.Call("api/ds", "metric_poll_dao", input)
	if err != nil {
		fmt.Println("fail to parse metric_poll_dao response!")
		return nil, err
	}
	var result struct {
		Entries []*model.MetricPoll `json:"entries,omitempty"`
	}
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_poll_dao response!")
		return nil, err
	}
	return result.Entries, nil

}

func (r *FluencyClient) GetMetricPoll(name string) (entry *model.MetricPoll, err error) {

	input := &model.MetricPollDaoRequest{
		Action: "get",
		Args: &model.MetricPollDaoRequestArgs{
			Id: name,
		},
	}

	res, err := r.serviceClient.Call("api/ds", "metric_poll_dao", input)
	if err != nil {
		fmt.Println("fail to parse metric_poll_dao response!")
		return nil, err
	}
	var ruleRes struct {
		Entry *model.MetricPoll `json:"entry,omitempty"`
	}
	err = json.Unmarshal(res.GetBytes(), &ruleRes)
	if err != nil {
		fmt.Println("fail to parse metric_poll_dao response!")
		return nil, err
	}
	return ruleRes.Entry, nil

}

func (r *FluencyClient) DeleteMetricPoll(name string) (err error) {

	input := &model.MetricPollDaoRequest{
		Action: "delete",
		Args: &model.MetricPollDaoRequestArgs{
			Id: name,
		},
	}

	_, err = r.serviceClient.Call("api/ds", "metric_poll_dao", input)
	if err != nil {
		fmt.Println("fail to parse metric_poll_dao response!")
		return err
	}
	return nil

}

func (r *FluencyClient) AddMetricPoll(entry *model.MetricPoll) (err error) {

	input := &model.MetricPollDaoRequest{
		Action: "add",
		Args: &model.MetricPollDaoRequestArgs{
			Entry: entry,
		},
	}

	_, err = r.serviceClient.Call("api/ds", "metric_poll_dao", input)
	if err != nil {
		fmt.Println("fail to parse metric_poll_dao response!")
		return err
	}
	return nil

}

func (r *FluencyClient) MetricResourceTagKeys(resourceTypes []string) (keys []string, err error) {

	input := &model.ListResourceTagKeysRequest{
		ResourceTypes: resourceTypes,
	}

	functionName := "metric_get_resource_tag_keys"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	var result model.ListResourceTagKeysResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_get_resource_tag_keys response!")
		return nil, err
	}
	return result.Keys, nil
}

func (r *FluencyClient) MetricResourceTagValues(resourceTypes []string, key string, pattern string) (values []string, err error) {

	input := &model.ListResourceTagValuesRequest{
		ResourceTypes: resourceTypes,
		Key:           key,
		Pattern:       pattern,
	}

	functionName := "metric_get_resource_tag_values"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	var result model.ListResourceTagValuesResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_get_resource_tag_values response!")
		return nil, err
	}
	return result.Values, nil
}

func (r *FluencyClient) MetricTestTagFilter(resourceTypes []string, key string, value string) (result *model.MetricTestTagFilterResponse, err error) {

	input := &model.MetricTestTagFilterRequest{
		Selector: &model.TagSelectorT{
			ResourceTypes: resourceTypes,
			MustFilters: []*model.MetricTagFilter{
				{
					Field: key,
					Terms: []string{value},
				},
			},
		},
	}

	functionName := "metric_test_tag_filter"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	//var result model.ListResourceTagValuesResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_test_tag_filter response!")
		return nil, err
	}
	return result, nil
}

/*
{
   "total": 507,
   "count": 55,
   "results": [
      {
         "account": "fluencyDefault",
         "region": "eu-central-1",
         "total": 86,
         "count": 3,
         "hits": [
            "i-0dd8537646188caf1(venn_main)",
            "i-0c0cbd693cc13b1b9(jyskeuro_main)",
            "i-043a0a645718cdd93(jet-server-1)"
         ]
      },
      {
         "account": "fluencyDefault",
         "region": "eu-west-1",
         "total": 6,
         "count": 1,
         "hits": [
            "i-0d884fb8a44bf3979(deciphex_main)"
         ]
      },
      {
         "account": "fluencyDefault",
         "region": "ap-east-1",
         "total": 10,
         "count": 3,
         "hits": [
            "i-0c9556d0210c13750(malaysia_main)",
            "i-0230881fec74233fa(hk2_main)",
            "i-01b194343869b67d1(hk1_main)"
         ]
      },
      {
         "account": "fluencyDefault",
         "region": "eu-north-1",
         "total": 0,
         "count": 0,
         "hits": null
      },
      {
         "account": "fluencyDefault",
         "region": "af-south-1",
         "total": 94,
         "count": 3,
         "hits": [
            "i-0b28e0ee7f8ca3a63(ampath_main)",
            "i-09e8600fac18a9aab(newmantle_main)",
            "i-071260cbcacc697de(securicom_main)"
         ]
      },
      {
         "account": "fluencyDefault",
         "region": "us-west-2",
         "total": 23,
         "count": 5,
         "hits": [
            "i-0da61517dc69d9287(eventcore_main)",
            "i-0b1c183e25ba0642e(timely_main)",
            "i-0bc2262226b65dad4(gdrgroup_main)",
            "i-074a3c2a153ed9753(cyberclan_main)",
            "i-0a2c3255fa6d21341(gdrnetlist_main)"
         ]
      },
      {
         "account": "fluencyDefault",
         "region": "us-east-1",
         "total": 123,
         "count": 10,
         "hits": [
            "i-0ca639c82f96ffc15(dms-server-1)",
            "i-0b71f4b6b761a70a5(gupholstery_main)",
            "i-0bf5696c1370c6365(gaston_main)",
            "i-0e0115fb14acee2d8(intermed_main)",
            "i-02f41d22f0e3a1463(ametros_main)",
            "i-0cb90a81db56d37b2(redgategroup_main)",
            "i-07deb6ddb10c8a955(terplab-server-1)",
            "i-02088b2d8419ca71f(praxis_main)",
            "i-012e0376d05c55563(jcinc_main)"
         ]
      }
   ]
}
*/

func (r *FluencyClient) MetricPollTest(namespace string, metric string, dimensions []string, resourceTypes []string, key string, value string) (result *model.MetricPollTestResponse, err error) {

	input := &model.MetricPollTestRequest{
		Entry: &model.MetricPoll{
			Namespace:  namespace,
			Metric:     metric,
			Dimensions: dimensions,
			TagSelector: &model.TagSelectorT{
				ResourceTypes: resourceTypes,
				MustFilters: []*model.MetricTagFilter{
					{
						Field: key,
						Terms: []string{value},
					},
				},
			},
		},
	}

	functionName := "metric_poll_test"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	//var result model.ListResourceTagValuesResponse
	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_poll_test response!")
		return nil, err
	}
	return result, nil
}

func (r *FluencyClient) MetricSearch(fpl string, options *model.MetricSearchOptions) (result *model.MetricSearchResponse, err error) {

	input := &model.MetricSearchRequest{
		FPL:     fpl,
		Options: options,
	}

	functionName := "metric_search"

	res, err := r.serviceClient.Call("api/ds", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Println("fail to parse metric_poll_test response:", err.Error())
		return nil, err
	}

	PrettyPrintJSON(result)

	return result, nil

}
