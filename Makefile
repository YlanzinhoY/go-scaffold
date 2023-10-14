build:
	mkdir -p "downloads"
	GOOS=linux GOARCH=amd64 go build -o downloads/linux/go-scaffold
	GOOS=windows GOARCH=amd64 go build -o downloads/windows/go-scaffold.exe


zip: 
	cd downloads/linux && zip go-scaffold-linux.zip go-scaffold
	cd downloads/windows && zip go-scaffold-windows.zip go-scaffold.exe


clean:
	cd downloads/linux && rm -f go-scaffold
	cd downloads/windows && rm -f go-scaffold.exe