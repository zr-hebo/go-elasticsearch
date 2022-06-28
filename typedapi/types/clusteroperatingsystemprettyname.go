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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// ClusterOperatingSystemPrettyName type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/cluster/stats/types.ts#L241-L244
type ClusterOperatingSystemPrettyName struct {
	Count      int  `json:"count"`
	PrettyName Name `json:"pretty_name"`
}

// ClusterOperatingSystemPrettyNameBuilder holds ClusterOperatingSystemPrettyName struct and provides a builder API.
type ClusterOperatingSystemPrettyNameBuilder struct {
	v *ClusterOperatingSystemPrettyName
}

// NewClusterOperatingSystemPrettyName provides a builder for the ClusterOperatingSystemPrettyName struct.
func NewClusterOperatingSystemPrettyNameBuilder() *ClusterOperatingSystemPrettyNameBuilder {
	r := ClusterOperatingSystemPrettyNameBuilder{
		&ClusterOperatingSystemPrettyName{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterOperatingSystemPrettyName struct
func (rb *ClusterOperatingSystemPrettyNameBuilder) Build() ClusterOperatingSystemPrettyName {
	return *rb.v
}

func (rb *ClusterOperatingSystemPrettyNameBuilder) Count(count int) *ClusterOperatingSystemPrettyNameBuilder {
	rb.v.Count = count
	return rb
}

func (rb *ClusterOperatingSystemPrettyNameBuilder) PrettyName(prettyname Name) *ClusterOperatingSystemPrettyNameBuilder {
	rb.v.PrettyName = prettyname
	return rb
}