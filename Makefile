serve:
	cd website && hugo serve

build:
	cd website && rm -rf public && hugo

.PHONY: serve build
