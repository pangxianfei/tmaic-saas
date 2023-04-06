protoc --go_out=./LoginApp/events/protocol_model ./LoginApp/events/protocol_model/*.proto
protoc --go_out=./LoginApp/jobs/proto3 ./LoginApp/jobs/proto3/*.proto
protoc --go_out=./FmsApp/jobs/proto3 ./FmsApp/jobs/proto3/*.proto