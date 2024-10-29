bin:
	CGO_CFLAGS="-Wno-nullability-completeness" GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o dnscheck.dll -buildmode=c-shared

