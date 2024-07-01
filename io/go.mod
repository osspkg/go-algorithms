module go.osspkg.com/x/io

go 1.20

replace (
	go.osspkg.com/x/errors => ../errors
	go.osspkg.com/x/sync => ../sync
	go.osspkg.com/x/test => ../test
)

require (
	go.osspkg.com/x/errors v0.3.1
	go.osspkg.com/x/sync v0.3.0
	go.osspkg.com/x/test v0.3.0
	gopkg.in/yaml.v3 v3.0.1
)
