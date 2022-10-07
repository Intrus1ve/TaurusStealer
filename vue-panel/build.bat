set GOOS=linux
set GOARCH=386
go build -o ./build/taurus_v1_4.panel -ldflags="-w -s" -trimpath ./backend/main.go
C:\Windows\system32\cmd.exe /c D:\7-Zip\7zG.exe a -tzip taurus_v1_4.zip ./build/* -mx9