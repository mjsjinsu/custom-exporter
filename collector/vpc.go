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
	vpcsubSystem = "vpc"
)

var (
	vpcDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, vpcsubSystem, "created"),
		"Fin Cloud VPC labels converted to Prometheus labels.",
		[]string{"vpcNo", "vpcName", "ipv4CidrBlock", "regionCode", "createDate"}, nil,
	)
)

type vpcCollector struct {
	desc   *prometheus.Desc
	logger log.Logger
}

func init() {
	registerCollector("vpc", defaultEnabled, NewVpcCollector)
}

// NewTimeCollector returns a new Collector exposing the current system time in
// seconds since epoch.
func NewVpcCollector(logger log.Logger) (Collector, error) {
	return &vpcCollector{
		desc:   vpcDesc,
		logger: logger,
	}, nil
}

func (c *vpcCollector) Update(ch chan<- prometheus.Metric) error {

	var status float64 = 0
	result, err := getVpcList()
	level.Debug(c.logger).Log("msg", "Return VPC List", "Run", result)

	if err != nil {
		return err
	}
	for _, v := range *result {
		if v.VpcStatus.Code == "RUN" {
			status = 1
		}
		ch <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue, status, v.VpcNo, v.VpcName, v.Ipv4CidrBlock, v.RegionCode, v.CreateDate)
	}
	return nil
}
