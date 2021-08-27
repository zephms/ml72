package top.acbs.nocode0;

/*
* 生成多线程处理部分, 例如:
*   print("start")
    ct = tool.ControlThread(data, YvYue, 10)
    ct.startThreads()
    ct.join()
    print("finish")
    res = ct.getResults()
* */

public class MultiTool {
    int threadNum;
    int groupNum;

    public MultiTool(int threadNum, int groupNum){
        this.threadNum = threadNum;
        this.groupNum = groupNum;
    }

    public String compile() {
        StringBuilder sb = new StringBuilder();
        sb.append("print(\"start\")\n");
        sb.append(String.format("ct = tool.ControlThread(dataInput, runFunction, threadnum=%d, groupnum=%d)\n",
                this.threadNum, this.groupNum));
        sb.append("ct.startThreads()\n" +
                "ct.join()\n" +
                "print(\"finish\")\n" +
                "res = ct.getResults()\n");
        return sb.toString();
    }
}
