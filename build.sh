# Windows 64 bit
env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/windows/hoauth.exe

# Mac OS 64 bit
env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/darwin/hoauth

# Linux 64 bit
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/linux/hoauth

chmod +x build/linux/hoauth
chmod +x build/darwin/hoauth
chmod +x build/windows/hoauth.exe

