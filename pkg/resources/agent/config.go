package agent

const BaseConfigName = "vector.toml"

var agentConfigTemplate = `
[sources.in]
type = "stdin"

[sinks.out]
inputs = ["in"]
type = "console"
encoding.codec = "text"
`
