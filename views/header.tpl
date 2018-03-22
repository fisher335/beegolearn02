<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head lang="en">
    <meta http-equiv="Content-Type" content="text/html" charset="UTF-8">
    <meta name='viewport' content="width=device-width,initial-scale=1">
    <link href="../static/css/bootstrap.min.css" rel='stylesheet' type="text/css">

    <title>beego测试</title>
</head>

<nav class="navbar navbar-inverse" role="navigation">
    <div class="container-fluid">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse"
                    data-target="#example-navbar-collapse">
                <span class="sr-only">切换导航</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="/home/" >网盘测试</a>
        </div>
        <div class="collapse navbar-collapse" id="example-navbar-collapse">
            <ul class="nav navbar-nav nav">
                <li><a href="/test/"><span class="glyphicon glyphicon-baby-formula">测试</span></a></li>
                <li><a href="/file/"><span class="glyphicon glyphicon-file">文件</span></a></li>
                <li><a href="/upload/"><span class="glyphicon glyphicon-upload">上传</span></a></li>
                <li><a href="/list/"><span class="glyphicon glyphicon-alert">请求头</span></a></li>
                <!--<li><a href="/mongo/"><span class="glyphicon glyphicon-floppy-disk">mongo</span></a> </li>-->
            </ul>
            <form class="navbar-form navbar-right" role="search">
                <div class="form-group">
                    <input type="text" class="form-control" placeholder="Search">
                </div>
                <button type="submit" class="btn btn-default">提交</button>
            </form>

        </div>

    </div>
</nav>


<div class="panel-body">