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
// https://github.com/elastic/elasticsearch-specification/tree/3b09f9d8e90178243f8a340a7bc324aab152c602

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// BucketMetricValueAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3b09f9d8e90178243f8a340a7bc324aab152c602/specification/_types/aggregations/Aggregate.ts#L233-L236
type BucketMetricValueAggregate struct {
	Keys []string `json:"keys"`
	Meta Metadata `json:"meta,omitempty"`
	// Value The metric value. A missing value generally means that there was no data to
	// aggregate,
	// unless specified otherwise.
	Value         Float64 `json:"value,omitempty"`
	ValueAsString *string `json:"value_as_string,omitempty"`
}

func (s *BucketMetricValueAggregate) UnmarshalJSON(data []byte) error {

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

		case "keys":
			if err := dec.Decode(&s.Keys); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "value":
			if err := dec.Decode(&s.Value); err != nil {
				return err
			}

		case "value_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ValueAsString = &o

		}
	}
	return nil
}

// NewBucketMetricValueAggregate returns a BucketMetricValueAggregate.
func NewBucketMetricValueAggregate() *BucketMetricValueAggregate {
	r := &BucketMetricValueAggregate{}

	return r
}
