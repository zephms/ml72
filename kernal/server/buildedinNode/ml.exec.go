package buildedinNode

import server "dev/server/base"

type BILNode_ml_exec struct {
	server.MLnode
}

func BILNode_ml_exec()  {

}

func (receiver BILNode_ml_exec) Exec(cmd []byte) []byte {
	return []byte("这是ml exec node")
}
