build:
	rm -rf public && hugo --config hugo.json
	docker build -t reg.lab.st/echo .

push: build
	docker push reg.lab.st/echo

.PHONY: build push
