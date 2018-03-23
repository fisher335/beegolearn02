{{template "header.tpl"}}
<table class="table table-responsive" align="center">
    <caption>信息查询</caption>

    <thead>
    <tr>
        <th>键</th>
        <th>值</th>
    </tr>
    </thead>
    <tbody>

    {{range $key, $val :=  .head}}
    <tr>
        <td>{{$key}}</td>
        <td>{{$val}}</td>
    </tr>
    {{end}}

    </tbody>
</table>
