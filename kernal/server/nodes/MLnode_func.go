package nodes

import (
	"dev/server/base"
	"encoding/json"
)

type MLnode_fun struct {
	server.MLnode
	args     args
	clientId int
}

type args struct {
	Names    [] string `json:"names"`
	ArgTypes [] string `json:"types"`
}

func NewMLnode_fun() *args {
	return &args{}
}

func (receiver MLnode_fun) loadArgs (jsonArg string) error {
	err := json.Unmarshal([]byte(jsonArg), &receiver.args)
	if err != nil {
		return err
	}
	return nil
}

func (receiver MLnode_fun) check() bool {
	return false
}