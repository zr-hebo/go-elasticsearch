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

// KeywordTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3b09f9d8e90178243f8a340a7bc324aab152c602/specification/_types/analysis/tokenizers.ts#L62-L65
type KeywordTokenizer struct {
	BufferSize int     `json:"buffer_size"`
	Type       string  `json:"type,omitempty"`
	Version    *string `json:"version,omitempty"`
}

func (s *KeywordTokenizer) UnmarshalJSON(data []byte) error {

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

		case "buffer_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.BufferSize = value
			case float64:
				f := int(v)
				s.BufferSize = f
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s KeywordTokenizer) MarshalJSON() ([]byte, error) {
	type innerKeywordTokenizer KeywordTokenizer
	tmp := innerKeywordTokenizer{
		BufferSize: s.BufferSize,
		Type:       s.Type,
		Version:    s.Version,
	}

	tmp.Type = "keyword"

	return json.Marshal(tmp)
}

// NewKeywordTokenizer returns a KeywordTokenizer.
func NewKeywordTokenizer() *KeywordTokenizer {
	r := &KeywordTokenizer{}

	return r
}
