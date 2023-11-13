build:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/go-scaffold
	GOOS=windows GOARCH=amd64 go build -o bin/windows/go-scaffold.exe


zip: 
	cd bin/linux && zip go-scaffold-linux.zip go-scaffold
	cd bin/windows && zip go-scaffold-windows.zip go-scaffold.exe


clean:
	cd bin/linux && rm -f go-scaffold
	cd bin/windows && rm -f go-scaffold.exe