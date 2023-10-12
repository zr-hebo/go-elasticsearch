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

package explore

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package explore
//
// https://github.com/elastic/elasticsearch-specification/blob/3b09f9d8e90178243f8a340a7bc324aab152c602/specification/graph/explore/GraphExploreRequest.ts#L28-L72
type Request struct {

	// Connections Specifies or more fields from which you want to extract terms that are
	// associated with the specified vertices.
	Connections *types.Hop `json:"connections,omitempty"`
	// Controls Direct the Graph API how to build the graph.
	Controls *types.ExploreControls `json:"controls,omitempty"`
	// Query A seed query that identifies the documents of interest. Can be any valid
	// Elasticsearch query.
	Query *types.Query `json:"query,omitempty"`
	// Vertices Specifies one or more fields that contain the terms you want to include in
	// the graph as vertices.
	Vertices []types.VertexDefinition `json:"vertices,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Explore request: %w", err)
	}

	return &req, nil
}
