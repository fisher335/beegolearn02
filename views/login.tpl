<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>登录</title>
    <meta http-equiv="X-UA-Compatible" content="IE=Edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0,user-scalable=no">
    <link href="../static/css/bootstrap.min.css" rel="stylesheet">
    <script src="../static/js/bootstrap.min.js"></script>
</head>
<body>
<style> .bgi {
    z-index: -1;
    position: fixed;
    top: 0px;
    left: 0px;
    height: 100%;
    width: 100%;
    background-size: cover;
    background-color: #525252;
    background-image: url(/static/img/{{.pic}}.jpg);
}

.login {
    position: absolute;
    width: 300px;
    margin-top: -100px;
    margin-left: -150px;
    left: 50%;
    top: 40%;
    background-color: rgba(255, 255, 255, 0.9);
} </style>
<div class="container">
    <div class="login panel panel-primary">
        <div class="panel-heading text-center">用户登录</div>
        <div class="panel-body">
            <form method="post" accept-charset="utf-8" action="">
                <div class="input-group"><span class="input-group-addon">用户：</span><input type="text" name="user"
                                                                                          class="form-control"></div>
                <br>
                <div class="input-group"><span class="input-group-addon">密码：</span><input type="password" name="passwd"
                                                                                          class="form-control"></div>
                <br><a class="btn btn-success" href='/login/'>忘记密码</a>
                <button style="float: right;" type="submit" class="btn btn-primary">登录</button>
            </form>
        </div>
    </div>
    <div class="bgi"></div>
</div>
</body>
</html>