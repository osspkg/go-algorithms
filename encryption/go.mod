module go.osspkg.com/x/encryption

go 1.20

replace (
	go.osspkg.com/x/errors => ../errors
	go.osspkg.com/x/random => ../random
	go.osspkg.com/x/test => ../test
)

require (
	go.osspkg.com/x/errors v0.5.0
	go.osspkg.com/x/random v0.5.0
	go.osspkg.com/x/test v0.5.0
	golang.org/x/crypto v0.24.0
)
