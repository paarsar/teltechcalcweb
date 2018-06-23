export GOPATH=${HOME}/workspace/go

grab-dependencies:
	go get -u github.com/beego/bee
	go get github.com/astaxie/beego
	go get github.com/stretchr/testify
run:
	export PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin
	bee run