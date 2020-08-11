PORT="${BPL_JMX_PORT:=5000}"

printf "JMX enabled on port %s\n" "${PORT}"

export JAVA_TOOL_OPTIONS="${JAVA_TOOL_OPTIONS}
  -Djava.rmi.server.hostname=127.0.0.1
  -Dcom.sun.management.jmxremote.authenticate=false
  -Dcom.sun.management.jmxremote.ssl=false
  -Dcom.sun.management.jmxremote.port=${PORT}
  -Dcom.sun.management.jmxremote.rmi.port=${PORT}"
