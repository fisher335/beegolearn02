{{template "header.tpl"}}

<table class="table table-bordered" align="center" sortable = "true" >
	<caption>视频文件</caption>

   <thead>
      <tr>
         <th>文件名称</th>
          <th>操作</th>
      </tr>
   </thead>
   <tbody>
  {{range $file:=  .file}}

          <tr>
                <td>{{$file}}</td>
             <td><a href="/player/?file=1111" target="_blank">播放</a></td>
          </tr>


{{end}}
   </tbody>
</table>