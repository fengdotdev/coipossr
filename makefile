run:
	go run sandbox/sandbox.go


# Build the server
buildserver:
	go build -o live live/live.go

# clear the output folder
clear:
	rm -rf output/*

# Run  output folder as site
serve:
	./live output