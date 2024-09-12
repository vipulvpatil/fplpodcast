package openai

import (
	ai "github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

var responseFormat = ai.ChatCompletionResponseFormat{
	Type: "json_schema",
	JSONSchema: &ai.ChatCompletionResponseFormatJSONSchema{
		Name:   "podcast_analysis",
		Strict: true,
		Schema: jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"players": {
					Type: jsonschema.Array,
					Items: &jsonschema.Definition{
						Type: jsonschema.Object,
						Properties: map[string]jsonschema.Definition{
							"web_name": {
								Type: jsonschema.String,
							},
							"tags": {
								Type: jsonschema.Array,
								Items: &jsonschema.Definition{
									Type: jsonschema.String,
								},
							},
							"summaries": {
								Type: jsonschema.Array,
								Items: &jsonschema.Definition{
									Type: jsonschema.String,
								},
							},
						},
						Required:             []string{"web_name", "tags", "summaries"},
						AdditionalProperties: false,
					},
				},
				"speakers": {
					Type: jsonschema.Array,
					Items: &jsonschema.Definition{
						Type: jsonschema.String,
					},
				},
			},
			Required:             []string{"players", "speakers"},
			AdditionalProperties: false,
		},
	},
}
