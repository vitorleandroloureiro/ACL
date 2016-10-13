package models

import (
	"github.com/golangit/dic/container"
)

var Dic container.Container;

func NewDic() {
	Dic = container.New()
	Dic.Register("db", NewCouchbase)

}