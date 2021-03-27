#!/bin/bash


gen_proto() {
  rm -rf /tmp/protos && mkdir -p /tmp/protos

  for f in `find src |grep ".proto"` ; do
    # echo "gen proto File -> $f";
    protoc -Isrc -I/usr/local/include -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:/tmp/protos \
    --govalidators_out=/tmp/protos \
    --grpc-gateway_out=logtostderr=true:/tmp/protos \
    $f
  done ;

  cp -rf /tmp/protos/github.com/flamefatex/protos/goout .
}

gen_swagger() {
  mkdir -p swagger
  for f in `find src |grep ".proto"` ; do
    # echo "gen swagger File -> $f";
    protoc -Isrc -I/usr/local/include -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --swagger_out=logtostderr=true:swagger \
    $f
  done ;
}

projects=(
"flamefatex/example-api"
)


gen_swagger_mixin() {
  for project in "${projects[@]}" ; do
    swagger_src_dir=swagger/${project}
    swagger_src_title_file=${swagger_src_dir}/title.json

    swagger_target_dir=swagger_mixin/${project}
    mkdir -p ${swagger_target_dir}

    files=`find ${swagger_src_dir}  -type f -name "*.json"`
    cp swagger/title.json ${swagger_src_title_file}
    if [ "$(uname)" == "Darwin" ]; then
      sed -i ''  's~{project_name}~'"${project}"'~g' ${swagger_src_title_file}
    elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
      # GNU/Linux platform
      sed -i  's~{project_name}~'"${project}"'~g' ${swagger_src_title_file}
    fi
    docker run --rm -it -v `pwd`:/tmp/protos -w /tmp/protos 192.168.0.103:8080/tcgroup/goswagger:latest -q mixin ${swagger_src_dir}/title.json ${files} -o ${swagger_target_dir}/swagger.json || true
  done ;
}

gen_doc() {
  mkdir -p doc

  protoc -Isrc -I/usr/local/include -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --doc_out=./doc --doc_opt=markdown,doc.md src/*/*/*.proto src/*/*/*/*.proto src/*/*/*/*/*.proto

  protoc -Isrc -I/usr/local/include -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --doc_out=./doc --doc_opt=html,doc.html src/*/*/*.proto src/*/*/*/*.proto src/*/*/*/*/*.proto

  protoc -Isrc -I/usr/local/include -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --doc_out=./doc --doc_opt=json,doc.json src/*/*/*.proto src/*/*/*/*.proto src/*/*/*/*/*.proto

  protoc -Isrc -I/usr/local/include -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --doc_out=./doc --doc_opt=docbook,doc.docbook src/*/*/*.proto src/*/*/*/*.proto src/*/*/*/*/*.proto
}

gen_mock() {
  mkdir -p gooutmock
  for f in `find goout/* |grep ".pb.go"` ; do
    target_pb_file=${f:6}
    services=$(cat $f |grep interface|grep -o "[A-Z].*Client")
    if  [ $services ]; then
      echo "mockgen -source=${f} -destination=gooutmock/${target_pb_file} $services"
      mockgen -source=${f} -destination=gooutmock/${target_pb_file} $services
    fi;

  done ;
}

if [ "all" == "$1" ] || [ "" == "$1" ]; then
  gen_proto
  gen_swagger
  gen_doc
  gen_mock
else
  gen_$1
fi;
