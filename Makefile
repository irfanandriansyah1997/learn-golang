init-apps:
	@echo 'Initial setup'
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.54.2
	git config core.hooksPath conf/hooks