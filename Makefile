export GOPATH=${HOME}/workspace/go

grab-dependencies:
	go get -u github.com/beego/bee
	go get github.com/astaxie/beego
start:
	export PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin
	${GOPATH}/bin/bee run -runmode aws-cloud9	