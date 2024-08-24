build:
	env GOOS=darwin GOARCH=amd64  go build -o /usr/local/bin/kim-util .

	env GOOS=darwin GOARCH=amd64  go build -o ./bin/kim-util-macos-amd64 .
	env GOOS=darwin GOARCH=arm64  go build -o ./bin/kim-util-macos-arm64 .
	env GOOS=windows GOARCH=386 go build -o ./bin/kim-util-windows-386.exe .
	env GOOS=windows GOARCH=amd64 go build -o ./bin/kim-util-windows-amd64.exe .
	env GOOS=linux GOARCH=amd64   go build -o ./bin/kim-util-linux-amd64 .
	env GOOS=linux GOARCH=arm64   go build -o ./bin/kim-util-linux-arm64 .

	zip -v ./bin/kim-util.zip ./bin/kim-util-macos-amd64 ./bin/kim-util-macos-arm64 ./bin/kim-util-windows-386.exe ./bin/kim-util-windows-amd64.exe ./bin/kim-util-linux-amd64 ./bin/kim-util-linux-arm64

