# Kratos Project Template

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## layout_template
- 项目布局模板：kratos+grpc+gin+wire+zap(log)
- 项目初始化
```
    - 拷贝模板并重命名(项目、mod)
    - make Makefile init
```
- 项目开发
```
    - 创建proto文件: kratos proto add [proto file path]
    - 生成proto-go代码: kratos proto client [proto file path]
    - 生成proto-service代码: kratos proto service [proto file path] -t [dest server file]
    - 完善 data/biz/service/server 的代码逻辑，并配置各模块的ProviderSet
    - wire注入: make Makefile generate
```
- 项目编译
```
    - 编译: make Makefile build
    - kratos run: kratos run  或  make Makefile run_kratos
```

## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

