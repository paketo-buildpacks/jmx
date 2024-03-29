# `gcr.io/paketo-buildpacks/jmx`
**This buildpack will soon be archived, JMX functionality is now available at build time by default and can be enabled with the runtime variable `$BPL_JMX_ENABLED`. See [here](https://paketo.io/docs/howto/java/#enable-jmx) for more information.**


The Paketo JMX Buildpack is a Cloud Native Buildpack that configures JMX for JVM applications.

## Behavior
This buildpack will participate if all the following conditions are met

* `$BP_JMX_ENABLED` is set

The buildpack will do the following:

* Contribute JMX configuration to `$JAVA_TOOL_OPTIONS`

## Configuration
| Environment Variable | Description
| -------------------- | -----------
| `$BP_JMX_ENABLED` | Whether to contribute JMX support
| `$BPL_JMX_ENABLED` | Whether to enable JMX support
| `$BPL_JMX_PORT` | What port the JMX connector will listen on. Defaults to `5000`.

## Publishing the Port
When starting an application with JMX enabled, a port must be published.  To publish the port in Docker, use the following command:

```bash
$ docker run --publish <LOCAL_PORT>:<REMOTE_PORT> ...
```

The `REMOTE_PORT` should match the `port` configuration for the application (`5000` by default).  The `LOCAL_PORT` can be any open port on your computer, but typically matches the `REMOTE_PORT` where possible.

Once the port has been published, your JConsole should connect to `localhost:<LOCAL_PORT>` for JMX access.

![JConsole Configuration](jconsole.png)

## License
This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0

