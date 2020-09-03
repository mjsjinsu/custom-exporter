// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !notime

package collector

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	subnetsubSystem = "subnet"
)

var (
	subnetDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subnetsubSystem, "created"),
		"Fin Cloud Subnet labels converted to Prometheus labels.",
		[]string{"vpcNo", "subnetNo", "subnetName", "subnet", "subnetType", "zoneCode", "createDate"}, nil,
	)
)

type subnetCollector struct {
	desc   *prometheus.Desc
	logger log.Logger
}

func init() {
	registerCollector("subnet", defaultEnabled, NewSubnetCollector)
}

// NewTimeCollector returns a new Collector exposing the current system time in
// seconds since epoch.
func NewSubnetCollector(logger log.Logger) (Collector, error) {
	return &subnetCollector{
		desc:   subnetDesc,
		logger: logger,
	}, nil
}

func (c *subnetCollector) Update(ch chan<- prometheus.Metric) error {

	var status float64 = 0
	result, err := getSubnetList()
	level.Debug(c.logger).Log("msg", "Return Subnet List", "Run", result)

	if err != nil {
		return err
	}
	for _, v := range *result {
		if v.SubnetStatus.Code == "RUN" {
			status = 1
		}
		ch <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue, status, v.VpcNo, v.SubnetNo, v.SubnetName, v.Subnet, v.SubnetType.Code, v.ZoneCode, v.CreateDate)
	}
	return nil
}
