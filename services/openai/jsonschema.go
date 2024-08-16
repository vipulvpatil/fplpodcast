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
							"summary": jsonschema.Definition{
								Type: jsonschema.Array,
								Items: &jsonschema.Definition{
									Type: jsonschema.Object,
									Properties: map[string]jsonschema.Definition{
										"text": jsonschema.Definition{
											Type: jsonschema.String,
										},
										"speaker": jsonschema.Definition{
											Type: jsonschema.String,
										},
									},
									Required:             []string{"text", "speaker"},
									AdditionalProperties: false,
								},
							},
						},
						Required:             []string{"name", "web_name", "tags", "summary"},
						AdditionalProperties: false,
					},
				},
				"teams": jsonschema.Definition{
					Type: jsonschema.Array,
					Items: &jsonschema.Definition{
						Type: jsonschema.Object,
						Properties: map[string]jsonschema.Definition{
							"owner": jsonschema.Definition{
								Type: jsonschema.String,
							},
							"goalkeepers": jsonschema.Definition{
								Type: jsonschema.Array,
								Items: &jsonschema.Definition{
									Type: jsonschema.String,
								},
							},
							"defense": jsonschema.Definition{
								Type: jsonschema.Array,
								Items: &jsonschema.Definition{
									Type: jsonschema.String,
								},
							},
							"midfield": jsonschema.Definition{
								Type: jsonschema.Array,
								Items: &jsonschema.Definition{
									Type: jsonschema.String,
								},
							},
							"attack": jsonschema.Definition{
								Type: jsonschema.Array,
								Items: &jsonschema.Definition{
									Type: jsonschema.String,
								},
							},
						},
						Required:             []string{"owner", "goalkeepers", "defense", "midfield", "attack"},
						AdditionalProperties: false,
					},
				},
			},
			Required:             []string{"players", "teams"},
			AdditionalProperties: false,
		},
	},
}
