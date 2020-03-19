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
	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
)

type JMX struct {
	LayerContributor libpak.LayerContributor
	Logger           bard.Logger
}

func NewJMX(info libcnb.BuildpackInfo) JMX {
	return JMX{LayerContributor: libpak.NewLayerContributor("JMX", info)}
}

func (j JMX) Contribute(layer libcnb.Layer) (libcnb.Layer, error) {
	j.Logger.Body(bard.FormatUserConfig("BPL_JMX_PORT", "the port the JVM will listen on", "5000"))

	j.LayerContributor.Logger = j.Logger

	return j.LayerContributor.Contribute(layer, func() (libcnb.Layer, error) {
		layer.Profile.Add("jmx", `PORT=${BPL_JMX_PORT:=5000}

printf "JMX enabled on port %%s\n" "${PORT}"

export JAVA_OPTS="${JAVA_OPTS}
  -Djava.rmi.server.hostname=127.0.0.1
  -Dcom.sun.management.jmxremote.authenticate=false
  -Dcom.sun.management.jmxremote.ssl=false
  -Dcom.sun.management.jmxremote.port=${PORT}
  -Dcom.sun.management.jmxremote.rmi.port=${PORT}"
`)

		layer.Launch = true
		return layer, nil
	})
}

func (JMX) Name() string {
	return "jmx"
}
