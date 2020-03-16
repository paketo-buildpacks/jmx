/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package jmx

import (
	"fmt"

	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
)

type Build struct {
	Logger bard.Logger
}

func (b Build) Build(context libcnb.BuildContext) (libcnb.BuildResult, error) {
	r := libpak.PlanEntryResolver{Plan: context.Plan}

	if _, ok, err := r.Resolve("jmx"); err != nil {
		return libcnb.BuildResult{}, fmt.Errorf("unable to resolve buildpack plan entry jmx\n%w", err)
	} else if !ok {
		return libcnb.BuildResult{}, nil
	}

	b.Logger.Title(context.Buildpack)

	j := NewJMX(context.Buildpack.Info)
	j.Logger = b.Logger

	return libcnb.BuildResult{Layers: []libcnb.LayerContributor{j}}, nil
}
