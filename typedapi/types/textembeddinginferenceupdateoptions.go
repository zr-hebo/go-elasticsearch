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

// TextEmbeddingInferenceUpdateOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/ml/_types/inference.ts#L351-L355
type TextEmbeddingInferenceUpdateOptions struct {
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string                       `json:"results_field,omitempty"`
	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

// TextEmbeddingInferenceUpdateOptionsBuilder holds TextEmbeddingInferenceUpdateOptions struct and provides a builder API.
type TextEmbeddingInferenceUpdateOptionsBuilder struct {
	v *TextEmbeddingInferenceUpdateOptions
}

// NewTextEmbeddingInferenceUpdateOptions provides a builder for the TextEmbeddingInferenceUpdateOptions struct.
func NewTextEmbeddingInferenceUpdateOptionsBuilder() *TextEmbeddingInferenceUpdateOptionsBuilder {
	r := TextEmbeddingInferenceUpdateOptionsBuilder{
		&TextEmbeddingInferenceUpdateOptions{},
	}

	return &r
}

// Build finalize the chain and returns the TextEmbeddingInferenceUpdateOptions struct
func (rb *TextEmbeddingInferenceUpdateOptionsBuilder) Build() TextEmbeddingInferenceUpdateOptions {
	return *rb.v
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *TextEmbeddingInferenceUpdateOptionsBuilder) ResultsField(resultsfield string) *TextEmbeddingInferenceUpdateOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

func (rb *TextEmbeddingInferenceUpdateOptionsBuilder) Tokenization(tokenization *NlpTokenizationUpdateOptionsBuilder) *TextEmbeddingInferenceUpdateOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}