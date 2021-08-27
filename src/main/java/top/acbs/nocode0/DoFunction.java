package top.acbs.nocode0;

/*
* 完成访问函数部分, 下设数据处理部分和访问部分
* */

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class DoFunction {
    ArrayList<String> variables= new ArrayList<String>();

    String handleCode = "";

    Map<String, String> headers = new HashMap<String, String>();
    Map<String, String> params = new HashMap<String, String>();
    String url;

    public DoFunction () {
    }

    public void readFromFile(){
        // 读取 excel 或者是 csv 文件

        // 记录参与的变量:
        variables.add("日期");
        variables.add("名字");

    }

    public void addHeaders(String k, String v){
        headers.put(k, v);
    }
    public void addParams(String k, String v){
        params.put(k, v);
    }

    public String compile(){
        StringBuilder sb = new StringBuilder();
        sb.append("def runFunction(d):\n");

        // 变量处理 也就是 handleCode 的解析
        for (int i = 0; i < this.variables.size(); i++) {
            this.handleCode = this.handleCode.replace("$" + this.variables.get(i), "d[" + i + "]");
        }
        String[] codeLines = handleCode.split("\n");
        for (String s : codeLines) {
            sb.append("\t");
            sb.append(s);
            sb.append("\n");
        }

        // header的处理 (如果value带引号,则在字典中体现)
        sb.append("\theaders={\t");
        for(String key: headers.keySet()){
            sb.append(String.format("\t\t'%s':%s,\n", key, this.headers.get(key)));
        }
        sb.append("\t}\n");

        // params的处理 (如果value带引号,则在字典中体现)
        sb.append("\tparams={\t");
        for(String key: params.keySet()){
            sb.append(String.format("\t\t'%s':%s,\n", key, this.params.get(key)));
        }
        sb.append("\t}\n");

        // url的处理
        sb.append("\turl='").append(this.url).append("'\n");

        // 实际请求 todo 返回值怎么介入我还得思考一下
        sb.append("\ttry:\n" +
                "\t\tresponse=requests.post(url=url, headers=headers, data=params)\n" +
                "\t\tprint(response.text)\n" +
                "\texcept Exception:\n" +
                "\t\tprint(\"出现故障\")\n" +
                "\treturn 0\n");
        return sb.toString();
    }

    public String compileReader() {
        return "doing\n";
    }
}


