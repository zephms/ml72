package top.acbs.nocode0;

/*
* 生成输出结果文件 例如这段代码:
* with open("openID71.json", mode="w", encoding="utf8") as fp:
    json.dump(res, fp)
* */

public class Output {
    String mode; // listJson, csv, xlsx
    Reptile father;

    public String getMode() {
        return mode;
    }

    public void setMode(String mode) {
        this.mode = mode;
    }

    public String getEncoding() {
        return encoding;
    }

    public void setEncoding(String encoding) {
        this.encoding = encoding;
    }

    public String getOutFileName() {
        return outFileName;
    }

    public void setOutFileName(String outFileName) {
        this.outFileName = outFileName;
    }

    String encoding;
    String outFileName;

    public Output(String mode, Reptile father){
        this.mode = mode;
        this.encoding = "utf8";
        this.outFileName = "data/outputResult"; // 只记录不包括后缀名部分
        this.father = father;
    }

    public String compile() {
        this.father.imports.addImports("json");
        StringBuilder sb = new StringBuilder();
        switch (this.mode) {
            case "listJson":
                sb.append(String.format("with open(\"%s.json\", mode=\"w\", encoding=\"%s\") as fp:\n",
                        this.outFileName,this.encoding));
                sb.append("\tjson.dump(res, fp)\n");
                break;
            case "csv":
                System.out.println("doing1");
            case "xlsx":
                System.out.println("doing2");
            default:
                System.out.println("doing3");
        }
        return sb.toString();
    }
}
