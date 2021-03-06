// Copyright (C) 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io"
	"strings"
	"text/template"

	"github.com/google/gapid/core/text/copyright"
	"github.com/google/gapid/core/text/reflow"
)

// EntryParameter represents a parameter for a specific stringtable entry.
type EntryParameter struct {
	Identifier string // Name of the parameter.
	Type       string // Type of the parameter. Default is "string".
}

// Entry represents a single stringtable entry.
type Entry struct {
	Key        string           // Name of the stringtable key.
	Parameters []EntryParameter // List of parameters.
}

// EntryList is a sortable slice of entries.
type EntryList []Entry

// Len returns the number of elements in the list.
func (l EntryList) Len() int { return len(l) }

// Less returns true if the i'th element's key is less than the j'th element's
// key.
func (l EntryList) Less(i, j int) bool { return l[i].Key < l[j].Key }

// Swap switches the i'th and j'th element.
func (l EntryList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

// context is the template execution context.
type context struct {
	Entries   []Entry
	Package   string
	Copyright string
}

// Execute executes the template with the given list of entries.
func Execute(templateRoutine, templateStr, pkg string, entries []Entry, w io.Writer) error {
	funcs := map[string]interface{}{
		"ToCamel": underscoredToCamel,
	}

	tmpl, err := template.New(templateRoutine).Funcs(funcs).Parse(templateStr)
	if err != nil {
		return err
	}

	ctx := context{
		Entries:   entries,
		Package:   pkg,
		Copyright: copyright.Build("generated_by", copyright.Info{Year: "2016", Tool: "stringgen"}),
	}

	r := reflow.New(w)
	r.Indent = "\t"
	defer r.Flush()
	return tmpl.Execute(r, ctx)
}

func underscoredToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(strings.ToLower(parts[i]))
	}
	return strings.Join(parts, "")
}
