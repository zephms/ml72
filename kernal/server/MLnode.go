package server

type MLnode struct {
	Path string
	MLtype string
	Data interface{}
}

const MLNodeType_CHANNEL = "CHAN"
const MLNodeType_SIGNAL = "SIG"

const MLNodeChanDefaultCap = 3

func NewMLnode(path string, mltype string, data interface{}) *MLnode {
	return &MLnode{
		Path: path,
		MLtype: mltype,
		Data: data,
	}
}