// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package renderer

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/xigxog/crd-ref-docs/config"
	"github.com/xigxog/crd-ref-docs/templates"
	"github.com/xigxog/crd-ref-docs/types"
	"sigs.k8s.io/controller-tools/pkg/crd/markers"
)

type MarkdownRenderer struct {
	conf *config.Config
	*Functions
}

func NewMarkdownRenderer(conf *config.Config) (*MarkdownRenderer, error) {
	baseFuncs, err := NewFunctions(conf)
	if err != nil {
		return nil, err
	}
	return &MarkdownRenderer{conf: conf, Functions: baseFuncs}, nil
}

func (m *MarkdownRenderer) Render(gvd []types.GroupVersionDetails) error {
	funcMap := combinedFuncMap(funcMap{prefix: "markdown", funcs: m.ToFuncMap()}, funcMap{funcs: sprig.TxtFuncMap()})

	var tpls fs.FS
	if m.conf.TemplatesDir != "" {
		tpls = os.DirFS(m.conf.TemplatesDir)
	} else {
		sub, err := fs.Sub(templates.Root, "markdown")
		if err != nil {
			return err
		}
		tpls = sub
	}

	tmpl, err := loadTemplate(tpls, funcMap)
	if err != nil {
		return err
	}

	f, _ := createOutFile(m.conf.OutputPath, "out.md")
	defer f.Close()

	return tmpl.ExecuteTemplate(f, mainTemplate, gvd)
}

func (m *MarkdownRenderer) ToFuncMap() template.FuncMap {
	return template.FuncMap{
		"GroupVersionID":     m.GroupVersionID,
		"RenderExternalLink": m.RenderExternalLink,
		"RenderGVLink":       m.RenderGVLink,
		"RenderLocalLink":    m.RenderLocalLink,
		"RenderType":         m.RenderType,
		"RenderTypeLink":     m.RenderTypeLink,
		"RenderValidation":   m.RenderValidation,
		"SafeID":             m.SafeID,
		"ShouldRenderType":   m.ShouldRenderType,
		"TypeID":             m.TypeID,
		"RenderFieldDoc":     m.RenderFieldDoc,
	}
}

func (m *MarkdownRenderer) ShouldRenderType(t *types.Type) bool {
	return t != nil && (t.GVK != nil || len(t.References) > 0)
}

func (m *MarkdownRenderer) RenderType(t *types.Field) string {
	var sb strings.Builder
	enum := t.Markers.Get("kubebuilder:validation:Enum")
	switch {
	case enum != nil:
		sb.WriteString("enum[")
		for i, v := range enum.(markers.Enum) {
			sb.WriteString("`")
			sb.WriteString(v.(string))
			if i != len(enum.(markers.Enum))-1 {
				sb.WriteString("`, ")
			} else {
				sb.WriteString("`")
			}
		}
		sb.WriteString("]")

	case t.Type.Kind == types.MapKind:
		sb.WriteString("map{")
		sb.WriteString(m.RenderTypeLink(t.Type.KeyType))
		sb.WriteString(", ")
		sb.WriteString(m.RenderTypeLink(t.Type.ValueType))
		sb.WriteString("}")

	case t.Type.Kind == types.SliceKind:
		sb.WriteString(m.RenderTypeLink(t.Type.UnderlyingType))
		sb.WriteString(" array")

	default:
		sb.WriteString(m.RenderTypeLink(t.Type))
	}

	return sb.String()
}

func (m *MarkdownRenderer) RenderValidation(f *types.Field) string {
	fmt.Printf("%#v\n", f)

	var v []string
	if f.Markers.Get("kubebuilder:validation:Required") != nil {
		v = append(v, "required")
	}
	if m := f.Markers.Get("kubebuilder:validation:MinLength"); m != nil {
		v = append(v, fmt.Sprintf("minLength: %d", m))
	}
	if m := f.Markers.Get("kubebuilder:validation:Minimum"); m != nil {
		v = append(v, fmt.Sprintf("min: %.0f", m.(markers.Minimum).Value()))
	}
	if m := f.Markers.Get("kubebuilder:validation:Maximum"); m != nil {
		v = append(v, fmt.Sprintf("max: %.0f", m.(markers.Maximum).Value()))
	}
	if m := f.Markers.Get("kubebuilder:validation:Pattern"); m != nil {
		v = append(v, fmt.Sprintf("pattern: %s", m))
	}
	if m := f.Markers.Get("kubebuilder:validation:Format"); m != nil {
		v = append(v, fmt.Sprintf("format: %s", m))
	}

	return strings.Join(v, ", ")
}

func (m *MarkdownRenderer) RenderTypeLink(t *types.Type) string {
	text := m.SimplifiedTypeName(t)

	link, local := m.LinkForType(t)
	if link == "" {
		return text
	}

	if local {
		return m.RenderLocalLink(text)
	} else {
		return m.RenderExternalLink(link, text)
	}
}

func (m *MarkdownRenderer) RenderLocalLink(text string) string {
	anchor := strings.ToLower(
		strings.NewReplacer(
			" ", "-",
			".", "",
			"/", "",
			"(", "",
			")", "",
		).Replace(text),
	)
	return fmt.Sprintf("[%s](#%s)", text, anchor)
}

func (m *MarkdownRenderer) RenderExternalLink(link, text string) string {
	return fmt.Sprintf("[%s](%s)", text, link)
}

func (m *MarkdownRenderer) RenderGVLink(gv types.GroupVersionDetails) string {
	return m.RenderLocalLink(gv.GroupVersionString())
}

func (m *MarkdownRenderer) RenderFieldDoc(text string) string {
	// Escape the pipe character, which has special meaning for Markdown as a way to format tables
	// so that including | in a comment does not result in wonky tables.
	out := strings.ReplaceAll(text, "|", "\\|")

	// Replace newlines with 2 line breaks so that they don't break the Markdown table formatting.
	return strings.ReplaceAll(out, "\n", "<br /><br />")
}
