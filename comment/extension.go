/*
 * Copyright 2022 Galaxyobe.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package comment

import (
	_ "embed"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

var (
	//go:embed comment.tmpl
	commentTmpl string
)

// Extension Generate Getter ent Extension.
type Extension struct {
	entc.DefaultExtension
}

func NewCommentExtension() entc.Extension {
	return Extension{}
}

// Templates of the extensions.
func (Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("comment").Parse(commentTmpl)),
	}
}
