SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64


go build -o ./build/erp ./main.go
go build -o ./build/LoginApp ./LoginApp/main.go
go build -o ./build/ProductApp ./ProductApp/main.go
go build -o ./build/UserApp ./UserApp/main.go
go build -o ./build/SysApp ./SysApp/main.go
go build -o ./build/OrderApp ./OrderApp/main.go
go build -o ./build/SrmApp ./SrmApp/main.go
go build -o ./build/PdApp ./PdApp/main.go
go build -o ./build/WmApp ./WmApp/main.go
go build -o ./build/SupplierApp ./SupplierApp/main.go
go build -o ./build/FmsApp ./FmsApp/main.go
go build -o ./build/NbomApp ./NbomApp/main.go
go build -o ./build/OaApp ./OaApp/main.go
go build -o ./build/artisan ./artisan.go
