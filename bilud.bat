
// 设置windows编译环境
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
// 设置Linux编译环境
#SET CGO_ENABLED=0
#SET GOOS=linux
#SET GOARCH=amd64

//开始编译
go build -o ./build/tmaicerp.exe ./main.go
go build -o ./build/LoginApp.exe ./LoginApp/main.go
go build -o ./build/ProductApp.exe ./ProductApp/main.go
go build -o ./build/UserApp.exe ./UserApp/main.go
go build -o ./build/SysApp.exe ./SysApp/main.go
go build -o ./build/OrderApp.exe ./OrderApp/main.go
go build -o ./build/artisan.exe ./artisan.go