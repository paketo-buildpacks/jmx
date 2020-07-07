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
	"github.com/paketo-buildpacks/libpak/sherpa"

	_ "github.com/paketo-buildpacks/jmx/jmx/statik"
)

type JMX struct {
	LayerContributor libpak.LayerContributor
	Logger           bard.Logger
}

func NewJMX(info libcnb.BuildpackInfo) JMX {
	return JMX{LayerContributor: libpak.NewLayerContributor("JMX", info)}
}

//go:generate statik -src . -include *.sh

func (j JMX) Contribute(layer libcnb.Layer) (libcnb.Layer, error) {
	j.LayerContributor.Logger = j.Logger

	return j.LayerContributor.Contribute(layer, func() (libcnb.Layer, error) {
		s, err := sherpa.StaticFile("/jmx.sh")
		if err != nil {
			return libcnb.Layer{}, fmt.Errorf("unable to load jmx.sh\n%w", err)
		}

		layer.Profile.Add("jmx.sh", s)

		layer.Launch = true
		return layer, nil
	})
}

func (JMX) Name() string {
	return "jmx"
}
