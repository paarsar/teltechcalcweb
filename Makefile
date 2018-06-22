export GOPATH=${HOME}/workspace/go

grab-dependencies:
	go get -u github.com/beego/bee
	go get github.com/astaxie/beego
run:
	export PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin
	bee run