module go.osspkg.com/x/config

go 1.20

replace go.osspkg.com/x/test => ./../test

require (
	go.osspkg.com/x/test v0.3.0
	gopkg.in/yaml.v3 v3.0.1
)
