module go.osspkg.com/x/ioutils

go 1.20

replace (
	go.osspkg.com/x/errors => ../errors
	go.osspkg.com/x/syncing => ../syncing
	go.osspkg.com/x/test => ../test
)

require (
	go.osspkg.com/x/errors v0.5.0
	go.osspkg.com/x/syncing v0.5.1
	go.osspkg.com/x/test v0.5.0
	gopkg.in/yaml.v3 v3.0.1
)
