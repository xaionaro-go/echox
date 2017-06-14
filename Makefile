dependency:
	go get github.com/kardianos/govendor
	govendor get github.com/spf13/hugo
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -update

build:
	cd website && rm -rf public && hugo
	for d in $(shell go list ./... | grep -v vendor); do \
		go build $$d; \
	done

serve:
	cd website && hugo serve

.PHONY: dependency build serve 
