run: get swag 
	air

serve: compile
	./main


compile: get_packages

	@go build  -o main

get:
	@go get

swag:
	swag init


