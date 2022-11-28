package fluency

import (
	"encoding/json"
	"fmt"

	"github.com/SecurityDo/fluency-go/model"
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

func (r *FluencyClient) MetricTagSearch(bucket string, dimension string, tag string, prefix string) (entries []string, err error) {

	input := &model.MetricTagSearchRequest{
		Metric:    bucket,
		Dimension: dimension,
		Tag:       tag,
		Pattern:   prefix,
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
		fmt.Println("fail to parse metric_tag_list response!")
		return nil, err
	}
	return result.Entries, nil
}
