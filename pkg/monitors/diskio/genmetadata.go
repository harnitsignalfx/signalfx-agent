// Code generated by monitor-code-gen. DO NOT EDIT.

package diskio

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "disk-io"

var groupSet = map[string]bool{}

const (
	diskMergedRead     = "disk_merged.read"
	diskMergedWrite    = "disk_merged.write"
	diskOctetsAvgRead  = "disk_octets.avg_read"
	diskOctetsAvgWrite = "disk_octets.avg_write"
	diskOctetsRead     = "disk_octets.read"
	diskOctetsWrite    = "disk_octets.write"
	diskOpsAvgRead     = "disk_ops.avg_read"
	diskOpsAvgWrite    = "disk_ops.avg_write"
	diskOpsPending     = "disk_ops.pending"
	diskOpsRead        = "disk_ops.read"
	diskOpsTotal       = "disk_ops.total"
	diskOpsWrite       = "disk_ops.write"
	diskTimeAvgRead    = "disk_time.avg_read"
	diskTimeAvgWrite   = "disk_time.avg_write"
	diskTimeRead       = "disk_time.read"
	diskTimeWrite      = "disk_time.write"
)

var metricSet = map[string]monitors.MetricInfo{
	diskMergedRead:     {Type: datapoint.Counter},
	diskMergedWrite:    {Type: datapoint.Counter},
	diskOctetsAvgRead:  {Type: datapoint.Gauge},
	diskOctetsAvgWrite: {Type: datapoint.Gauge},
	diskOctetsRead:     {Type: datapoint.Counter},
	diskOctetsWrite:    {Type: datapoint.Counter},
	diskOpsAvgRead:     {Type: datapoint.Gauge},
	diskOpsAvgWrite:    {Type: datapoint.Gauge},
	diskOpsPending:     {Type: datapoint.Gauge},
	diskOpsRead:        {Type: datapoint.Counter},
	diskOpsTotal:       {Type: datapoint.Gauge},
	diskOpsWrite:       {Type: datapoint.Counter},
	diskTimeAvgRead:    {Type: datapoint.Gauge},
	diskTimeAvgWrite:   {Type: datapoint.Gauge},
	diskTimeRead:       {Type: datapoint.Counter},
	diskTimeWrite:      {Type: datapoint.Counter},
}

var defaultMetrics = map[string]bool{
	diskOpsAvgRead:  true,
	diskOpsAvgWrite: true,
	diskOpsRead:     true,
	diskOpsTotal:    true,
	diskOpsWrite:    true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "disk-io",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
