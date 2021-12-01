default: gen_easyjson

gen_easyjson:
	go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest
	rm -rf *_easyjson.go
	easyjson -all -byte *.go
	(cd native/request/ && easyjson -all -byte *.go)
	(cd native/response/ && easyjson -all -byte *.go)

