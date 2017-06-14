serve:
	cd website && hugo serve

build:
	cd website && rm -rf public && hugo

dependency:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -update

test:
	echo "" > coverage.txt
	for d in $(shell go list ./... | grep -v vendor); do \
		go test -race -coverprofile=profile.out -covermode=atomic $$d; \
		[ -f profile.out ] && cat profile.out >> coverage.txt && rm profile.out; \
	done

.PHONY: serve build dependency test
