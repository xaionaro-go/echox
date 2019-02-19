IMAGE = labstack/echo-web

build:
	rm -rf public && hugo
	docker build -t $(IMAGE) .

push: build
	docker push $(IMAGE)
