package server

type MLnode struct {
	Path string
	MLtype string
	Exec *func(*MLnode)(cmd []byte)
	Data interface{}
}



func NewMLnode(path string, mltype string, data interface{}) *MLnode {
	return &MLnode{
		Path: path,
		MLtype: mltype,
		Data: data,
	}
}

//func (receiver MLnode) Exec(cmd []byte) []byte {
//	if cmd==nil{
//		return nil
//	}
//	var bytes bytes2.Buffer
//	bytes.WriteString("接收到指令")
//	bytes.Write(cmd)
//	return bytes.Bytes()
//}

