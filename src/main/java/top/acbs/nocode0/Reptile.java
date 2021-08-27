package top.acbs.nocode0;

public class Reptile {
    public Imports imports;
    DoFunction doFunction;
    MultiTool multiTool;
    Output output;

    public Reptile () {
        imports = new Imports();
        imports.addImports("tool");
        imports.addImports("requests");
        imports.addImports("random");
//        imports.addImports("json");

        doFunction = new DoFunction();
        multiTool = new MultiTool(20, 5);
        output = new Output("listJson", this);

    }

    public String  compile(){
        StringBuilder sb = new StringBuilder();
        sb.append(doFunction.compile());
        sb.append(doFunction.compileReader());
        sb.append(multiTool.compile());
        sb.append(output.compile());
        sb.insert(0, imports.compile());
        return sb.toString();
    }

    public void showCompile(){
        System.out.println(this.compile());
    }
}
