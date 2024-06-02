run: build
		@./bin/app

build:
		@go build -o bin/app cmd/app/main.go

css:
		tailwind -i ./public/tailwind.css -o ./public/styles.css --minify