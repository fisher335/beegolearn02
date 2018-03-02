{{template "base.tpl"}}
<table class="table table-bordered" align="center" sortable="true">
    <caption>文件列表</caption>

    <thead>
    <tr>
        <th>键</th>
        <th>值</th>
        <th>值</th>
        <th>值</th>
        <th>值</th>
    </tr>
    </thead>
    <tbody>

    {{range .rows}}
    <tr>
        <td hidden>{{.Record_id}}</td>
        <td>{{.Name}}</td>
        <td>{{.Size}}</td>
    {{if eq .Recordtype   "1" }}
        <td>文件夹
    <td>
    {{else}}
        <td>文件</td>
    {{end}}
        <td>{{.Size}}</td>
        <td>{{.Fileextname}}</td>
    </tr>
    {{end}}
    </tbody>
</table>
