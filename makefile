run:
	go run sandbox/sandbox.go

build:
	go build -o sandbox sandbox/sandbox.go

# Build the server
buildserver:
	go build -o live tools/live/live.go

# clear the output folder
clear:
	rm -rf output/*

# Run  output folder as site
serve:
	./live output