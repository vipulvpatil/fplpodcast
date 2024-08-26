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
				"players": jsonschema.Definition{
					Type: jsonschema.Array,
					Items: &jsonschema.Definition{
						Type: jsonschema.Object,
						Properties: map[string]jsonschema.Definition{
							"name": jsonschema.Definition{
								Type: jsonschema.String,
							},
							"web_name": jsonschema.Definition{
								Type: jsonschema.String,
							},
							"tags": jsonschema.Definition{
								Type: jsonschema.Array,
								Items: &jsonschema.Definition{
									Type: jsonschema.String,
								},
							},
							"summaries": jsonschema.Definition{
								Type: jsonschema.Array,
								Items: &jsonschema.Definition{
									Type: jsonschema.String,
								},
							},
						},
						Required:             []string{"name", "web_name", "tags", "summaries"},
						AdditionalProperties: false,
					},
				},
			},
			Required:             []string{"players"},
			AdditionalProperties: false,
		},
	},
}
