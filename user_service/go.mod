module github.com/inonsdn/gacha-system/user_service

go 1.23.5

require (
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/inonsdn/gacha-system/proto v0.0.0
	google.golang.org/grpc v1.72.2
)

replace github.com/inonsdn/gacha-system/proto => ../proto

require (
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
