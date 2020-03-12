# `jmx`
The Paketo JMX Buildpack is a Cloud Native Buildpack that configures JMX for JVM applications.

## Behavior
This buildpack will participate if all of the following conditions are met

* `$BP_JMX_ENABLED` is set

The buildpack will do the following:

* Contribute JMX configuration to `$JAVA_OPTS` 

## Configuration 
| Environment Variable | Description
| -------------------- | -----------
| `$BP_JMX_ENABLED` | Whether to contribute JMX support
| `$BPL_JMX_PORT` | What port the JMX connector will listen on. Defaults to `5000`. 

## Creating SSH Tunnel
After starting an application with JMX enabled, an SSH tunnel must be created to the container.  To create that SSH container, execute the following command:

```bash
$ cf ssh -N -T -L <LOCAL_PORT>:localhost:<REMOTE_PORT> <APPLICATION_NAME>
```

The `REMOTE_PORT` should match the `port` configuration for the application (`5000` by default).  The `LOCAL_PORT` must match the `REMOTE_PORT`.

Once the SSH tunnel has been created, your JConsole should connect to `localhost:<LOCAL_PORT>` for JMX access.

![JConsole Configuration](jconsole.png)

## License
This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
