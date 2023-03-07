## Fetch project dependencies and tidy Go modules
.PHONY: dep
dep:
	GO111MODULE=on go mod vendor
	GO111MODULE=on go mod tidy
