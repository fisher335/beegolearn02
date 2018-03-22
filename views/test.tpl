{{template "header.tpl"}}

<form class="form-horizontal col-md-10" role="form" action="" method="POST" id="info">
    <div class="form-group">
        <label for="name" class="col-md-3 align-left control-label">名字</label>
        <div class="col-md-4">
            <input type="text" class="form-control" id="name" name='name' placeholder="请输入姓名" autofocus="true">
        </div>
    </div>
    <div class="form-group">
        <label for="depart" class="col-md-3 align-left control-label ">部门</label>
        <div class="col-md-4">
            <input type="text" class="form-control" id="depart" name="depart" placeholder="请输入部门">
        </div>
    </div>
    <div class="form-group">
        <label for="card" class="col-md-3 align-left control-label">身份证</label>
        <div class="col-md-4">
            <input type="text" class="form-control" id="card" name="card" placeholder="身份证号"
                   aria-placeholder="aria-placeholder">
        </div>
    </div>
    <div class="form-group">
        <label for="other" class="col-md-3 align-left control-label">其他</label>
        <div class="col-md-4">
            <input type="text" class="form-control" id="other" name="other" placeholder="其他信息">
        </div>
        <span class="help-block">这里填些注意事项。</span>
    </div>
    <div class="col-md-offset-3">
        <label for="name">默认的复选框实例</label>
        <div class="checkbox">
            <label><input type="checkbox" name='check' value="option1">选项 1</label>
        </div>
        <div class="checkbox">
            <label><input type="checkbox" name='check' value="option2">选项 2</label>
        </div>
        <div class="checkbox">
            <label><input type="checkbox" name='check' value="option3">选项 2</label>
        </div>
        <div class="checkbox">
            <label><input type="checkbox" name='check' value="option4">选项 2</label>
        </div>


    </div>
    <!--这里是单选按钮  -->
    <div class="col-md-offset-3">
        <label for="name">默认单选按钮的实例</label>
        <div class="radio">
            <label>
                <input type="radio" name="optionsRadios" id="optionsRadios1" value="option1" checked> 选项 1
            </label>
        </div>
        <div class="radio">
            <label>
                <input type="radio" name="optionsRadios" id="optionsRadios2" value="c">选项 2 - 选择它将会取消选择选项 1
            </label>
        </div>
    </div>

    <div class="col-md-offset-3">
        <label>默认下拉框的实例</label>
        <div class="dropdown">
            <select name="select1" id="select_test">
                <option value="">--请选择--</option>
                <option value="python">python</option>
                <option value="java">java</option>
            </select>
        </div>
    </div>
    <br>
    <!--radio结束  -->
    <div class="form-group">
        <div class="col-md-offset-3 col-md-3">
            <button type="button" class="btn btn-primary" onclick="subinfo()">提交信息</button>
        </div>
    </div>
    <div class="form-group">
        <div class="col-md-offset-3 col-md-4">
            <button type="button" class="btn btn-primary" onclick="checkAll(true)">全选</button>
            <button type="button" class="btn btn-primary" onclick="checkAll(false)">全不选</button>
            <button type="button" class="btn btn-primary" onclick="doselect()">操作下拉框</button>
            <button type="button" class="btn btn-primary" onclick="doradio()">操作radio</button>
        </div>
    </div>
</form>
<script type="text/javascript">

    function subinfo() {
        if (check() == true) {
            if (window.confirm("你确认要提交吗？")) {
                $('#info').submit();
            }
            else {
                location.href = "/index/"
            }
        }
    }

    function check() {
        if ($("#depart").val().trim() == '' || $("#name").val().trim() == "") {
            alert("姓名和部门不能为空");
            return false;
        }
        else {
            return true;
        }
    }

    function checkAll(val) {

        $("input[name = 'check']").each(function () {
            $(this).prop('checked', val);
        })
    }

    function doselect() {
        $("#select_test").val('java');
    }

    function doradio() {
        $("#optionsRadios2").prop("checked", true);
    }

</script>
{{template "footer.tpl"}}