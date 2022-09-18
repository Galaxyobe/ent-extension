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

package annotation

import (
	"strconv"
	"strings"
)

// CreateTableStatus implements entc.Annotation. default is true.
type CreateTableStatus bool

// Name of the annotation. Used by the codegen templates.
func (CreateTableStatus) Name() string {
	return "CreateTableStatus"
}

// Options of entsql.Annotation.
func (s CreateTableStatus) Options() string {
	return s.Name() + " = " + strconv.FormatBool(bool(s))
}

func (s CreateTableStatus) Disabled() bool {
	return bool(s) == false
}

func NewCreateTableStatus(any interface{}) CreateTableStatus {
	var ok = true
	switch v := any.(type) {
	case bool:
		return CreateTableStatus(v)
	case string:
		if v == "" {
			break
		}
		ok, _ = strconv.ParseBool(v)
	}
	return CreateTableStatus(ok)
}

func ParseCreateTableStatusFromOptions(options string) CreateTableStatus {
	options = strings.ReplaceAll(options, " = ", "=")
	items := strings.Split(options, " ")
	var s CreateTableStatus = true
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		if kv[0] == s.Name() {
			v, err := strconv.ParseBool(kv[1])
			if err != nil {
				continue
			}
			s = CreateTableStatus(v)
		}
	}
	return s
}
