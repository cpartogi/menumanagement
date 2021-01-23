dep: 
	go get
run:
	go build main.go && ./main

local:
	swag init -o api/docs go build main.go && ./main
