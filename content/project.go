package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-k8s/ponzu/management/editor"
	"github.com/ponzu-k8s/ponzu/system/item"
)

type Project struct {
	item.Item

	Name           string   `json:"name"`
	ResearchTopics []string `json:"research-topics"`
	Researchers    []string `json:"researchers"`
	Labs           []string `json:"labs"`
	SplashImage    string   `json:"splash-image"`
	Summary        string   `json:"summary"`
	Description    string   `json:"description"`
}

// MarshalEditor writes a buffer of html to edit a Project within the CMS
// and implements editor.Editable
func (p *Project) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Project field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", p, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Tags("ResearchTopics", p, map[string]string{
				"label":       "ResearchTopics",
				"placeholder": "+ResearchTopics",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Researchers", p, map[string]string{
				"label": "Researchers",
			},
				"Person",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Labs", p, map[string]string{
				"label": "Labs",
			},
				"Lab",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.File("SplashImage", p, map[string]string{
				"label":       "SplashImage",
				"placeholder": "Upload the SplashImage here",
			}),
		},
		editor.Field{
			View: editor.Input("Summary", p, map[string]string{
				"label":       "Summary",
				"type":        "text",
				"placeholder": "Enter the Summary here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", p, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Project editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Project"] = func() interface{} { return new(Project) }
}

// String defines how a Project is printed. Update it using more descriptive
// fields from the Project struct type
func (p *Project) String() string {
	return fmt.Sprintf("Project: %s", p.UUID)
}
