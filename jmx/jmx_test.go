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

package jmx_test

import (
	"io/ioutil"
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/jmx/jmx"
	"github.com/sclevine/spec"
)

func testJMX(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx libcnb.BuildContext
		j   jmx.JMX
	)

	it.Before(func() {
		var err error

		ctx.Buildpack.Info.Version = "test-version"

		ctx.Layers.Path, err = ioutil.TempDir("", "debug")
		Expect(err).NotTo(HaveOccurred())

		j = jmx.NewJMX(ctx.Buildpack.Info)
	})

	it("contributes debug configuration", func() {
		layer, err := ctx.Layers.Layer("test-layer")
		Expect(err).NotTo(HaveOccurred())

		layer, err = j.Contribute(layer)
		Expect(err).NotTo(HaveOccurred())

		Expect(layer.Launch).To(BeTrue())
		Expect(layer.Profile["jmx"]).To(Equal(`PORT=${BPL_JMX_PORT:=5000}

printf "JMX enabled on port %s\n" "${PORT}"

export JAVA_OPTS="${JAVA_OPTS}
  -Djava.rmi.server.hostname=127.0.0.1
  -Dcom.sun.management.jmxremote.authenticate=false
  -Dcom.sun.management.jmxremote.ssl=false
  -Dcom.sun.management.jmxremote.port=${PORT}
  -Dcom.sun.management.jmxremote.rmi.port=${PORT}"
`))
	})
}
