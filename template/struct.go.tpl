package {{.Models}}
//Models 为包名（传参时候决定）

import (
"fmt"
{{$ilen := len .Imports}}{{if gt $ilen 0}}{{range .Imports}}"{{.}}"{{end}}{{end}}
)

{{range .Tables}}
    //Mapper 数据库的表
    var {{Mapper .Name}}Model = &{{Mapper .Name}}{}
    //struct这段是xorm官方给的模版
    type {{Mapper .Name}} struct {
    {{$table := .}}
    {{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
    {{end}}
    }

    {{range .ColumnsSeq}}{{$col := $table.GetColumn .}}
    func (m *{{Mapper $table.Name}}) Get{{Mapper $col.Name}}() (val {{Type $col}}) {
    if m == nil {
    return
    }
    return m.{{Mapper $col.Name}}
    }
    {{end}}

    func (m *{{Mapper .Name}}) String() string {
    return fmt.Sprintf("%#v", m)
    }

    func (m *{{Mapper .Name}}) TableName() string {
    return "{{.Name}}"
    }

{{end}}