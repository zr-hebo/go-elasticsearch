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

// BucketCorrelationFunctionCountCorrelation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/_types/aggregations/pipeline.ts#L129-L132
type BucketCorrelationFunctionCountCorrelation struct {
	// Indicator The indicator with which to correlate the configured `bucket_path` values.
	Indicator BucketCorrelationFunctionCountCorrelationIndicator `json:"indicator"`
}

// NewBucketCorrelationFunctionCountCorrelation returns a BucketCorrelationFunctionCountCorrelation.
func NewBucketCorrelationFunctionCountCorrelation() *BucketCorrelationFunctionCountCorrelation {
	r := &BucketCorrelationFunctionCountCorrelation{}

	return r
}