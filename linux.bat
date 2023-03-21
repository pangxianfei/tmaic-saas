SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64


go build -o ./build/tmaicerp ./main.go
go build -o ./build/LoginApp ./LoginApp/main.go
go build -o ./build/ProductApp ./ProductApp/main.go
go build -o ./build/UserApp ./UserApp/main.go
go build -o ./build/SysApp ./SysApp/main.go
go build -o ./build/OrderApp ./OrderApp/main.go
go build -o ./build/artisan ./artisan.go