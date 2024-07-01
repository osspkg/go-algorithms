
.PHONY: install
install:
	go install github.com/osspkg/devtool@latest


.PHONY: lint
lint:
	devtool lint

.PHONY: license
license:
	devtool license

.PHONY: tests
tests:
	devtool test

.PHONY: ci
ci: install license lint tests

.PHONY: go_work
go_work:
	go work use -r .
	go work sync

create_release:
	devtool tag