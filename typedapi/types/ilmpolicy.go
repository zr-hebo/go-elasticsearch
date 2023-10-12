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
)

// IlmPolicy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3b09f9d8e90178243f8a340a7bc324aab152c602/specification/ilm/_types/Policy.ts#L23-L26
type IlmPolicy struct {
	Meta_  Metadata `json:"_meta,omitempty"`
	Phases Phases   `json:"phases"`
}

func (s *IlmPolicy) UnmarshalJSON(data []byte) error {

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

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return err
			}

		case "phases":
			if err := dec.Decode(&s.Phases); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIlmPolicy returns a IlmPolicy.
func NewIlmPolicy() *IlmPolicy {
	r := &IlmPolicy{}

	return r
}
