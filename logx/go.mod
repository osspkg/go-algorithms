module go.osspkg.com/x/logx

go 1.20

replace (
	go.osspkg.com/x/syncing => ../syncing
	go.osspkg.com/x/test => ../test
)

require (
	github.com/mailru/easyjson v0.7.7
	go.osspkg.com/x/syncing v0.5.1
	go.osspkg.com/x/test v0.5.0
)

require github.com/josharian/intern v1.0.0 // indirect
