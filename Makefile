build:
	GOOS=darwin GOARCH=amd64 go build -o app
	GOOS=windows GOARCH=amd64 go build -o app.exe

release:
	GOOS=linux GOARCH=amd64 go build -o multilingo
	zip handler.zip ./multilingo