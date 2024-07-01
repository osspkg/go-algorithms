module go.osspkg.com/x/app

go 1.20

replace (
	go.osspkg.com/x/algorithms => ../algorithms
	go.osspkg.com/x/config => ../config
	go.osspkg.com/x/console => ../console
	go.osspkg.com/x/context => ../context
	go.osspkg.com/x/env => ../env
	go.osspkg.com/x/errors => ../errors
	go.osspkg.com/x/log => ../log
	go.osspkg.com/x/sync => ../sync
	go.osspkg.com/x/syscall => ../syscall
	go.osspkg.com/x/test => ../test
)

require (
	go.osspkg.com/x/algorithms v1.4.0
	go.osspkg.com/x/config v0.5.0
	go.osspkg.com/x/console v0.5.0
	go.osspkg.com/x/context v0.5.0
	go.osspkg.com/x/env v0.5.0
	go.osspkg.com/x/errors v0.5.0
	go.osspkg.com/x/log v0.5.0
	go.osspkg.com/x/sync v0.5.1
	go.osspkg.com/x/syscall v0.5.0
	go.osspkg.com/x/test v0.5.0
)

require (
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)