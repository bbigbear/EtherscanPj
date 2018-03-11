<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>项目监控</title>
<link rel="stylesheet" href="/static/css/layui.css">
</head>
<body class="layui-layout-body">
<div class="layui-layout layui-layout-admin">
  <div class="layui-header">
    <div class="layui-logo">项目监控v1.0</div>
    <ul class="layui-nav layui-layout-right">
      <li class="layui-nav-item">
        <a href="javascript:;">
          <img src="../static/img/admin.jpg" class="layui-nav-img">
          用户
        </a>
        <dl class="layui-nav-child">
          <dd><a href="">基本资料</a></dd>
          <dd><a href="">安全设置</a></dd>
        </dl>
      </li>
      <li class="layui-nav-item"><a href="">退出</a></li>
    </ul>
  </div>
  
  <div class="layui-side layui-bg-black">
    <div class="layui-side-scroll">
      <!-- 左侧导航区域（可配合layui已有的垂直导航） -->
      <ul class="layui-nav layui-nav-tree"  lay-filter="test">
        <li class="layui-nav-item"><a href="/getearlywarn">项目预警</a></li>
        <li class="layui-nav-item"><a href="/getnotifcationmessage">消息通知</a></li>
      </ul>
    </div>
  </div>
  <div class="layui-body">
    <!-- 内容主体区域 -->
    <div style="padding: 15px;">
		<form class="layui-form layui-form-pane1" action="" onsubmit="javascript:return false;">	  
		  <blockquote class="layui-elem-quote">变动通知</blockquote>
		  <div class="layui-form-item">
			<input id="notify" type="checkbox" name="switch" lay-skin="switch" lay-text="是|否" lay-filter="notify" checked>
		  </div>
		  <blockquote class="layui-elem-quote">单笔交易</blockquote>
		  <div class="layui-form-item">
				<div class="layui-form-mid">单笔交易超过</div>
			    <div class="layui-input-inline" style="width: 50px;">
					 <input type="text" name="single_num" id="single_num" value="{{.Price}}" autocomplete="off" class="layui-input">
				</div>
				<div class="layui-form-mid">or</div>
				<div class="layui-input-inline" style="width: 50px;">
					 <input type="text" name="single_percent" id="single_percent" value="{{.Price}}" placeholder="%" autocomplete="off" class="layui-input">
				</div>
		  </div>
		  <!--<blockquote class="layui-elem-quote">累计交易</blockquote>
		  <div class="layui-form-item">
				<div class="layui-form-mid">累计24小时交易超过</div>
			    <div class="layui-input-inline" style="width: 30px;">
					 <input type="text" name="Sell_price" id="sell_price" value="{{.Price}}" autocomplete="off" class="layui-input">
				</div>
				<div class="layui-form-mid">or</div>
				<div class="layui-input-inline" style="width: 30px;">
					 <input type="text" name="Sell_price" id="sell_price" value="{{.Price}}" placeholder="%" autocomplete="off" class="layui-input">
				</div>
		  </div>
		  <div class="layui-form-item">
				<div class="layui-form-mid">累计</div>
				<div class="layui-input-inline" style="width: 30px;">
					 <input type="text" name="Sell_price" id="sell_price" value="{{.Price}}" autocomplete="off" class="layui-input">
				</div>
				<div class="layui-form-mid">小时交易超过</div>
			    <div class="layui-input-inline" style="width: 30px;">
					 <input type="text" name="Sell_price" id="sell_price" value="{{.Price}}" autocomplete="off" class="layui-input">
				</div>
				<div class="layui-form-mid">or</div>
				<div class="layui-input-inline" style="width: 30px;">
					 <input type="text" name="Sell_price" id="sell_price" value="{{.Price}}" placeholder="%" autocomplete="off" class="layui-input">
				</div>
		  </div>
		  <blockquote class="layui-elem-quote">短信通知</blockquote>
		  <div class="layui-form-item">
				<div class="layui-form-mid">短信通知到</div>
			    <div class="layui-input-inline" style="width: 30px;">
					 <input type="text" name="Sell_price" id="sell_price" value="{{.Price}}" autocomplete="off" class="layui-input">
				</div>
		  </div>-->
		  <div class="layui-form-item">
		      <button class="layui-btn" id="save">保存</button>
		  </div>
		</form>					
	</div>
  </div>
  
  <div class="layui-footer">
    <!-- 底部固定区域 -->
    ©2018 项目监控. All Rights Reserved
  </div>
</div>
<style>
	.layui-tab-title li:first-child > i {
		display: none;
		disabled:true
	}
</style>

<script src="/static/layui.js"></script>
<!--<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>-->
<script>
	//JavaScript代码区域
	layui.use(['element','layer','jquery','table'], function(){
	  var element = layui.element
		,form=layui.form
		,layer=layui.layer
		,$=layui.jquery
		,table=layui.table;
	  //layer.msg("你好");

	//保存按钮
	$("#save").on('click',function(){
		//layer.msg("点击保存")
		
		if($("input[type='checkbox']").is(':checked')){
			//layer.msg("选中")
			var single_num=$("#single_num").val()
			var single_percent=$("#single_percent").val()
			var status="start"
			if(single_num!=""||single_percent!=""){
				alert("保存成功")	
				window.location.href="/getearlywarn?action="+status+"&sn="+single_num+"&sp="+single_percent;
			}else{
				layer.msg("请输入数字或者百分数")
			}
		
		}else{
			var status="stop"
			//layer.msg("未选中")
			alert("保存成功")
			window.location.href="/getearlywarn?action="+status;
		}
		
	});
	//监听checkbox 
	
	
	
			
			
			
  });

	
	
</script>

</body>
</html>