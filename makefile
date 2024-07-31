run: get swag 
	air

serve: compile
	./main


compile: get

	@go build  -o main

get:
	@go get

swag:
	swag init


