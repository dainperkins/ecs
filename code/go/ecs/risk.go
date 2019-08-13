// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

// Risk fields represent the normalized comparative sensitivity of internal
// resources or the potential risk of threats, vulnerabilities, and IOCs.
// At the event level normalized risk is calculated based on any risk metrics
// available for lookup for the given event (e.g. a user, source, destination
// in a firewall log, a client and file in a endpoint file/process log). The
// final event score is calculated by summing the squares of each avaialble
// risk score and dividing by the total multiple of available scores. As an
// example:  ( user.risk + source.risk + destination.risk ) / ( user.risk +
// source.risk + destination.risk )
// Risk labels are for labeling the type of business risk represented by a
// given resource, typically in reference to the data housed in an asset or
// zone, the level of access of a given user or group, or the risk represented
// by a given threat or IOC.
type Risk struct {
	// locally relevant label to describe the risk type of asset or data
	// identified either specific to a local data classification system or
	// regulatory/compliance applicability
	Label string `ecs:"label"`

	// comparative risk score quantitavely scoring the level of risk of the
	// asset or data, or risk represented by a given threat (typically provided
	// by the observer or IOC provider)
	Score int64 `ecs:"score"`
}
