// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/899364a63e7415b60033ddd49d50a30369da26d7

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// CoordinatorStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/enrich/stats/types.ts#L29-L35
type CoordinatorStats struct {
	ExecutedSearchesTotal int64  `json:"executed_searches_total"`
	NodeId                string `json:"node_id"`
	QueueSize             int    `json:"queue_size"`
	RemoteRequestsCurrent int    `json:"remote_requests_current"`
	RemoteRequestsTotal   int64  `json:"remote_requests_total"`
}

func (s *CoordinatorStats) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "executed_searches_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ExecutedSearchesTotal = value
			case float64:
				f := int64(v)
				s.ExecutedSearchesTotal = f
			}

		case "node_id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return err
			}

		case "queue_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.QueueSize = value
			case float64:
				f := int(v)
				s.QueueSize = f
			}

		case "remote_requests_current":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.RemoteRequestsCurrent = value
			case float64:
				f := int(v)
				s.RemoteRequestsCurrent = f
			}

		case "remote_requests_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RemoteRequestsTotal = value
			case float64:
				f := int64(v)
				s.RemoteRequestsTotal = f
			}

		}
	}
	return nil
}

// NewCoordinatorStats returns a CoordinatorStats.
func NewCoordinatorStats() *CoordinatorStats {
	r := &CoordinatorStats{}

	return r
}