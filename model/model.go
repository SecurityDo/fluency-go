package model

import (
	"encoding/json"
	"time"
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
