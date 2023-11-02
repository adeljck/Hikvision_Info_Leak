export LDFLAGS='-s -w '

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -trimpath -o Hikvision_Info_Leak_linux_amd64 main.go
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="$LDFLAGS" -trimpath -o Hikvision_Info_Leak_windows_386.exe  main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="$LDFLAGS" -trimpath -o Hikvision_Info_Leak_windows_amd64.exe  main.go
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags="$LDFLAGS" -trimpath -o Hikvision_Info_Leak_windows_arm64.exe  main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="$LDFLAGS" -trimpath -o Hikvision_Info_Leak_darwin_amd64 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="$LDFLAGS" -trimpath -o Hikvision_Info_Leak_darwin_arm64 main.go

upx -9 Hikvision_Info_Leak_linux_amd64
upx -9 Hikvision_Info_Leak_windows_386.exe
upx -9 Hikvision_Info_Leak_windows_amd64.exe
upx -9 Hikvision_Info_Leak_windows_arm64.exe
upx -9 Hikvision_Info_Leak_darwin_amd64
upx -9 Hikvision_Info_Leak_darwin_arm64

zip Hikvision_Info_Leak_linux_amd64.zip Hikvision_Info_Leak_linux_amd64
zip Hikvision_Info_Leak_windows_386.zip Hikvision_Info_Leak_windows_386.exe
zip Hikvision_Info_Leak_windows_amd64.zip Hikvision_Info_Leak_windows_amd64.exe
zip Hikvision_Info_Leak_windows_arm64.zip Hikvision_Info_Leak_windows_arm64.exe
zip Hikvision_Info_Leak_darwin_amd64.zip Hikvision_Info_Leak_darwin_amd64
zip Hikvision_Info_Leak_darwin_arm64.zip Hikvision_Info_Leak_darwin_arm64

rm -f Hikvision_Info_Leak_linux_amd64
rm -f Hikvision_Info_Leak_windows_386.exe
rm -f Hikvision_Info_Leak_windows_amd64.exe
rm -f Hikvision_Info_Leak_windows_arm64.exe
rm -f Hikvision_Info_Leak_darwin_amd64
rm -f Hikvision_Info_Leak_darwin_arm64