package model

import (
	"encoding/json"
	"time"

	elastic "gopkg.in/olivere/elastic.v3"
)

const ( // iota is reset to 0
)

// chart, table,
type ReportPanel struct {
	Key         string `json:"key,omitempty"`
	DataTable   string `json:"dataTable,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	PanelType   string `json:"panelType,omitempty"`

	/*
		Table     json.RawMessage `json:"table,omitempty"`
		Histogram json.RawMessage `json:"histogram,omitempty"`
		Topn      json.RawMessage `json:"topn,omitempty"`
		Chart     json.RawMessage `json:"chart,omitempty"`
		Pie       json.RawMessage `json:"pie,omitempty"`
		GeoMap    json.RawMessage `json:"geoMap,omitempty"`
		Counter   json.RawMessage `json:"counter,omitempty"`
	*/
	PanelConfig json.RawMessage `json:"panelConfig,omitempty"`
	//TableConfig *TablePanelConfig `json:"tableConfig,omitempty"`
	//ChartConfig *ChartPanelConfig `json:"chartConfig,omitempty"`
}

// attach to every fplTask
// either copy from the saved FPL Report
// or created by user from scratch
// this config could be saved together as one FPL Report
type TaskReportConfig struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	Panels []*ReportPanel  `json:"panels,omitempty"`
	Layout json.RawMessage `json:"layout,omitempty"`
}

type FPLReport struct {
	ID         int64  `json:"id"`
	Repository string `json:"repository,omitempty"`
	Group      string `json:"group"`

	Name         string               `json:"name,omitempty"`
	Description  string               `json:"description,omitempty"`
	FPL          string               `json:"fpl"`
	ReportConfig *TaskReportConfig    `json:"reportConfig,omitempty"`
	Arguments    []*FPLReportArgument `json:"arguments,omitempty"`
	UpdatedOn    time.Time            `json:"updatedOn"`

	ScheduleFlag bool          `json:"scheduleFlag"`
	Schedule     *ScheduleTime `json:"schedule,omitempty"`

	LastRun time.Time `json:"-"`
}

type FPLReportArgument struct {
	Name         string  `json:"name,omitempty"`
	Description  string  `json:"description,omitempty"`
	Value        *string `json:"value,omitempty"`
	DefaultValue string  `json:"defaultValue,omitempty"`
	Optional     bool    `json:"optional,omitempty"`
	Type         string  `json:"type,omitempty"` // integer, float, string, boolean
	IsList       bool    `json:"isList,omitempty"`
}

type ScheduleTime struct {
	Hour     int    `json:"hour"`
	Minute   int    `json:"minute"`
	WeekDay  int    `json:"weekday"`
	Day      int    `json:"day"`
	Timezone string `json:"timezone"`
	Interval string `json:"interval"`
}

// interval in seconds
type MetricImportGroup struct {
	Namespace     string `json:"namespace"`
	Interval      int64  `json:"interval"`
	Category      string `json:"category"`
	MinuteEmulate bool   `json:"minuteEmulate"`

	// will be set by server api
	CreatedOn time.Time `json:"createdOn"`
}

type MetricImportEntry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Unit        string `json:"unit"`

	// must match with the incoimg data
	Dimensions []string `json:"dimensions"`
	// if the target dimensions differ from original dimensions, optional
	KeyDimensions []string `json:"keyDimensions,omitempty"`

	// default valueField is "sum"
	ValueField string `json:"valueField,omitempty"`

	// if differ from the template  aws.ec2.cpuutinization
	// FluencyMetric string `json:"fluencyMetric,omitempty"`
	BucketName string `json:"bucket,omitempty"`

	Namespace string `json:"namespace"`
	Category  string `json:"category"`

	// will be set by server api
	CreatedOn time.Time `json:"createdOn"`
}

type MetricAWSListGroupsResponse struct {
	Groups []*MetricImportGroup `json:"groups"`
}
type MetricAWSListMetricsResponse struct {
	Metrics []*MetricImportEntry `json:"metrics"`
}
type MetricAWSListAllResponse struct {
	Groups  []*MetricImportGroup `json:"groups"`
	Metrics []*MetricImportEntry `json:"metrics"`
}

type MetricAWSAddGroupRequest struct {
	Group *MetricImportGroup `json:"group"`
}

type MetricAWSAddMetricRequest struct {
	Metric *MetricImportEntry `json:"metric"`
}

type MetricAWSDeleteGroupRequest struct {
	Namespace string `json:"namespace"`
	Category  string `json:"category"`
}

type MetricAWSDeleteMetricRequest struct {
	Name string `json:"name"`
}

type pagerdutyConfig struct {
	Token          string `json:"token"`
	IntegrationKey string `json:"integrationKey"`
}

type emailConfig struct {
}

type slackConfig struct {
	Token    string `json:"token"`
	Team     string `json:"team"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type ActorConfig struct {
	Name     string `json:"name"`
	Customer string `json:"customer"`
	Type     string `json:"type"`

	BuiltIn bool `json:"builtIn,omitempty"`

	Description string `json:"description"`
	Disabled    bool   `json:"disabled"`

	Pagerduty *pagerdutyConfig `json:"PagerDuty"`
	Email     *emailConfig     `json:"Email"`
	Slack     *slackConfig     `json:"Slack"`
	//Notification *notificationConfig `json:"Notification"`
	//SentinelOne  *sentinelOneConfig  `json:"SentinelOne"`
	//Peplink      *peplinkConfig      `json:"Peplink"`
	//UEBA *UEBAConfig `json:"UEBA"`

	// Config json.RawMessage `json:"config,omitempty"`
	// actorHandle ActorHandle `json:"-"`

	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

type MetricNotificationActorListResponse struct {
	Entries []*ActorConfig `json:"entries"`
}

type FieldAttribute struct {
	Description  string   `json:"description"`
	Field        string   `json:"field"`
	Enums        []string `json:"enums"`
	Default      string   `json:"default"`
	Formula      string   `json:"formula"`
	Optional     bool     `json:"optional"`
	DataType     string   `json:"dataType"` // string, integer, bool
	HtmlTemplate string   `json:"htmlTemplate"`
}

type EventActionEntry struct {
	Actor      string            `json:"actor"`
	ActorName  string            `json:"actorName"`
	Action     string            `json:"action"`
	Name       string            `json:"name"`
	Attributes []*FieldAttribute `json:"attributes"`
	// Attributes map[string]*FieldAttribute `json:"attributes"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

func (r *EventActionEntry) SetFieldDefault(field string, defaultValue string) {
	for _, attribute := range r.Attributes {
		if attribute.Field == field {
			attribute.Default = defaultValue
		}
	}
}

type MetricNotificationEndpointDaoRequestArgs struct {
	Id    string            `json:"id,omitempty"`
	Entry *EventActionEntry `json:"entry,omitempty"`
}

type MetricNotificationEndpointDaoRequest struct {
	Action string
	Args   *MetricNotificationEndpointDaoRequestArgs `json:"args"`
}

type MetricNotificationEndpointDaoResponse struct {
	Entry *EventActionEntry `json:"entry"`
}

type MetricAlertAction struct {
	ID       string   `json:"id"`
	Patterns []string `json:"patterns"`

	// trigger/resolve
	Actions   []string `json:"actions"`
	Endpoints []string `json:"endpoints"`
}

type MetricNotificationActionDaoRequestArgs struct {
	Id    string             `json:"id,omitempty"`
	Entry *MetricAlertAction `json:"entry,omitempty"`
}

type MetricNotificationActionDaoRequest struct {
	Action string
	Args   *MetricNotificationActionDaoRequestArgs `json:"args"`
}

type MetricNotificationActionDaoResponse struct {
	Entry *EventActionEntry `json:"entry"`
}

type MetricTag struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// search tags by  metric OR dimension name
type MetricTagsRequest struct {
	Metric    string `json:"metric,omitempty"`
	Dimension string `json:"dimension,omitempty"`
}

type MetricTagsResponse struct {
	Tags []*MetricTag `json:"tags"`
}

// search tags by  metric OR dimension name
type MetricTagSearchRequest struct {
	Metric    string `json:"metric"`
	Dimension string `json:"dimension"`
	Tag       string `json:"tag"`
	Pattern   string `json:"pattern"`
}

type MetricTagSearchResponse struct {
	Entries []string `json:"entries"`
}

type SimpleSearchOption struct {
	SearchStr  string `json:"searchStr,omitempty"`
	RangeFrom  int64  `json:"range_from"`
	RangeTo    int64  `json:"range_to"`
	RangeField string `json:"range_field"`

	FetchOffset int           `json:"fetchOffset"`
	FetchLimit  int           `json:"fetchLimit"`
	SortField   string        `json:"sortField,omitempty"`
	SortOrder   string        `json:"sortOrder,omitempty"`
	Field       string        `json:"field,omitempty"`
	FieldTerms  []interface{} `json:"fieldTerms,omitempty"`
}

type FilterEntry struct {
	Field      string        `json:"field"`
	Terms      []interface{} `json:"terms"`
	FilterType string        `json:"filterType"`
}

// Name: Interval: Key:
type DateFacetEntry struct {
	Name string `json:"name,omitempty"`
	// "day","minute", "hour", "week", "month" or "1.5h"
	Interval string `json:"interval,omitempty"`
	// timestamp field
	Key string `json:"key,omitempty"`
	// optional
	Value string `json:"value,omitempty"`
	// optional
	FilterField string `json:"filterField,omitempty"`
	// optional
	FilterTerm interface{} `json:"filterTerm,omitempty"`
}

type FacetEntry struct {
	Field string `json:"field"`
	Order string `json:"order"`
	Size  int    `json:"size"`
}

type FacetsOption struct {
	DateFacets     []*DateFacetEntry `json:"dateFacets"`
	Facets         []*FacetEntry     `json:"facets"`
	MustFilters    []*FilterEntry    `json:"mustFilters"`
	MustNotFilters []*FilterEntry    `json:"mustNotFilters"`
}

type SimpleFacetSearchOption struct {
	SimpleSearchOption
	Facets *FacetsOption `json:"facets,omitempty"`
}

type MetricIncidentSearchRequest struct {
	Options *SimpleFacetSearchOption `json:"options"`
}

type MetricIncidentSearchResponse elastic.SearchResult
