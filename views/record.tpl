{{template "header.tpl"}}
<table class="table table-bordered" align="center" sortable="true">
    <caption>文件列表</caption>

    <thead>
    <tr>
        <th hidden>rocordid</th>
        <th>NO.</th>
        <th>名称</th>
        <th>大小</th>
        <th>类型</th>
        <th>后缀</th>
        <th>最后修改时间</th>
    </tr>
    </thead>
    <tbody>

    {{range $index,$row := .rows}}
    <tr>

        <td hidden>{{.Record_id}}</td>
         <td>{{$index|addindex}}</td>
        <td>{{.Name}}</td>
        <td>{{.Size}}</td>
    {{if eq .Recordtype   "1" }}
        <td>文件夹</td>

    {{else}}
        <td>文件</td>
    {{end}}

    <td>{{.Fileextname}}</td>
        <td>{{.Createtime}}</td>

    </tr>
    {{end}}
    </tbody>
</table>
{{template "footer.tpl"}}