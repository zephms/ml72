package main

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/go-cmd/cmd"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func RunCMD(cmdstr string, prefix string, wg *sync.WaitGroup)  {
	//
	defer wg.Done()
	prefix = "[" + prefix + "]\t"
	cmd := exec.Command("cmd")
	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()
	cmd.Start()
	in.Write([]byte(cmdstr))
	in.Close()
	//readOutPut(out)
	dec := mahonia.NewDecoder("gbk")

	lenOnce := 100
	var buff []byte
	for {
		carrier := make([]byte, 0, lenOnce)
		nn, err := out.Read(carrier[0:lenOnce])
		carrier = carrier[:nn]
		if err != nil{
			if err==io.EOF{
				//_,temp,_ := dec.Translate(carrier, true)
				//if string(temp)[len(string(temp))-1]=='\n' {
				//	fmt.Print(buff+string(temp))
				//} else {
				//	fmt.Println(buff+string(temp))
				//}
			} else {
				fmt.Println("其他报错 有空再来解决")
			}
			break
		}
		if string(carrier)[len(string(carrier))-1] == '\n'{
			s := Decode2str(append(buff, carrier...), dec)
			s = strings.Trim(s,"\r\n")
			fmt.Println(prefix,strings.ReplaceAll(s, "\r\n", "\r\n"+prefix))
			buff = nil
		} else {
			tempss := bytes.Split(carrier, []byte("\r\n"))
			if len(tempss)==1{
				buff = append(buff, tempss[0]...)
			} else if len(tempss)==2{
				fmt.Println(prefix,Decode2str(append(buff, tempss[0]...), dec))
				buff = tempss[1]
			} else {
				fmt.Println(prefix ,Decode2str(append(buff, tempss[0]...), dec))
				for i := 1;i<len(tempss)-1;i++{
					fmt.Println(prefix,Decode2str(tempss[i],dec))
				}
				buff = tempss[len(tempss)-1]
			}
		}
	}

}

//func main() {
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go runCMD("ping www.baidu.com\n", "ping1", &wg)
//	wg.Add(1)
//	go runCMD("ping www.google.com\n", "ping2", &wg)
//	wg.Wait()
//}

func TestTempRun(){ // 这个是一个测试的东西，可以给它删掉
	fmt.Println("testT")
	// Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Create Cmd with options
	envCmd := cmd.NewCmdOptions(cmdOptions, "./print-some-lines")

	// Print STDOUT and STDERR lines streaming from Cmd
	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		// Done when both channels have been closed
		// https://dave.cheney.net/2013/04/30/curious-channels
		for envCmd.Stdout != nil || envCmd.Stderr != nil {
			select {
			case line, open := <-envCmd.Stdout:
				if !open {
					envCmd.Stdout = nil
					continue
				}
				fmt.Println(line)
			case line, open := <-envCmd.Stderr:
				if !open {
					envCmd.Stderr = nil
					continue
				}
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	// Run and wait for Cmd to return, discard Status
	<-envCmd.Start()

	// Wait for goroutine to print everything
	<-doneChan
}