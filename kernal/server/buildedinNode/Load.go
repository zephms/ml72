package buildedinNode

import server "dev/server/base"

func LoadBulidInNode(ser *MLserver)  {

	var mlExecNode = BILNode_ml_exec{
		*server.NewMLnode("/ml/exec", "bil", ),
	}
	ser.Db["/ml/exec"] =

}

