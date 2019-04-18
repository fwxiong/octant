package component

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Text_Marshal(t *testing.T) {
	tests := []struct {
		name     string
		input    Component
		expected string
		isErr    bool
	}{
		{
			name: "general",
			input: &Text{
				Config: TextConfig{
					Text: "lorem ipsum",
				},
			},
			expected: `
            {
                "metadata": {
                  "type": "text"
                },
                "config": {
                  "value": "lorem ipsum"
                }
            }
`,
		},
		{
			name: "with title",
			input: &Text{
				base: newBase(typeText, TitleFromString("image")),
				Config: TextConfig{
					Text: "nginx:latest",
				},
			},
			expected: `
            {
                "metadata": {
									"type": "text",
									"title": [
										{
											"config": { "value": "image" },
											"metadata": { "type": "text" }
										}
									]
                },
                "config": {
                  "value": "nginx:latest"
                }
            }
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := json.Marshal(tc.input)
			isErr := (err != nil)
			if isErr != tc.isErr {
				t.Fatalf("Unexpected error: %v", err)
			}

			assert.JSONEq(t, tc.expected, string(actual))
		})
	}
}

func Test_Text_SupportsTitle(t *testing.T) {
	var c Component = NewText("text")

	_, ok := c.(TitleComponent)
	assert.True(t, ok)
}

func Test_Text_String(t *testing.T) {
	c := NewText("string")
	assert.Equal(t, "string", c.String())
}

func Test_Text_LessThan(t *testing.T) {
	cases := []struct{
		name string
		ts Text
		other Component
		expected bool
	}{
		{
			name: "is less",
			ts: *NewText("b"),
			other: NewText("c"),
			expected: true,
		},
		{
			name: "is not less",
			ts: *NewText("b"),
			other: NewText("a"),
			expected: false,
		},
		{
			name: "other is not a timestamp",
			ts: *NewText("b"),
			other: nil,
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T){
			got := tc.ts.LessThan(tc.other)
			assert.Equal(t, tc.expected, got)
		})
	}
}