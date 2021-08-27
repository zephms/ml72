package top.acbs.nocode0;

import java.util.ArrayList;

public class Imports {
    private ArrayList<String> imports;

    public Imports () {
        imports = new ArrayList();
    }

    public void addImports(String s){
        if (!imports.contains(s)){
            this.imports.add(s);
        }
    }

    public String compile(){
        StringBuilder sb = new StringBuilder();
        for (String s : imports) {
            sb.append("import ");
            sb.append(s);
            sb.append("\n");
        }
        return sb.toString();
    }

    public void peek(){
        System.out.println(this.compile());
    }
}
