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
// https://github.com/elastic/elasticsearch-specification/tree/9a0362eb2579c6604966a8fb307caee92de04270

// Package termsaggregationexecutionhint
package termsaggregationexecutionhint

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/9a0362eb2579c6604966a8fb307caee92de04270/specification/_types/aggregations/bucket.ts#L993-L998
type TermsAggregationExecutionHint struct {
	Name string
}

var (
	Map = TermsAggregationExecutionHint{"map"}

	Globalordinals = TermsAggregationExecutionHint{"global_ordinals"}

	Globalordinalshash = TermsAggregationExecutionHint{"global_ordinals_hash"}

	Globalordinalslowcardinality = TermsAggregationExecutionHint{"global_ordinals_low_cardinality"}
)

func (t TermsAggregationExecutionHint) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TermsAggregationExecutionHint) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "map":
		*t = Map
	case "global_ordinals":
		*t = Globalordinals
	case "global_ordinals_hash":
		*t = Globalordinalshash
	case "global_ordinals_low_cardinality":
		*t = Globalordinalslowcardinality
	default:
		*t = TermsAggregationExecutionHint{string(text)}
	}

	return nil
}

func (t TermsAggregationExecutionHint) String() string {
	return t.Name
}
