<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>新增投资者</title>
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
  
  <div class="layui-side layui-bg-gray">
    <div class="layui-side-scroll">
      <!-- 左侧导航区域（可配合layui已有的垂直导航） -->
      <ul id="tree">
        <!--<li class="layui-nav-item"><a href="/getearlywarn">项目预警</a></li>
        <li class="layui-nav-item"><a href="/getnotifcationmessage">消息通知</a></li>-->
      </ul>
    </div>
  </div>
  <div class="layui-body">
    <!-- 内容主体区域 -->
    <div style="padding: 15px;">
		<form class="layui-form layui-form-pane1" action="" onsubmit="javascript:return false;">
		  <fieldset class="layui-elem-field layui-field-title site-title">
	      <legend><a name="color-design">新增投资者</a></legend>
	      </fieldset>
		  <div class="layui-form-item">
		    <div class="layui-inline">
		      <label class="layui-form-label">投资者姓名</label>
		      <div class="layui-input-inline" >
		        <input type="text" name="Sell_price" id="name" placeholder="请输入姓名" autocomplete="off" class="layui-input">
		      </div>
		    </div>
		  </div>
		  <div class="layui-form-item">
		    <div class="layui-inline">
		      <label class="layui-form-label">手机</label>
		      <div class="layui-input-inline" >
		        <input type="text" name="Sell_price" id="phone" placeholder="请输入手机" autocomplete="off" class="layui-input">
		      </div>
		    </div>
		  </div>
		  <div class="layui-form-item">
		    <div class="layui-inline">
		      <label class="layui-form-label">合约地址</label>
		      <div class="layui-input-inline">
			  <input type="text" name="Name" id="contract" placeholder="请输入合约地址" autocomplete="off" class="layui-input" style="width:300px;">
		      </div>
		    </div>
		  </div>
		  <div class="layui-form-item">
		    <div class="layui-inline">
		      <label class="layui-form-label">帐号地址</label>
		      <div class="layui-input-inline" >
		        <input type="text" name="Sell_price" id="address" placeholder="请输入帐号地址" autocomplete="off" class="layui-input" style="width:300px;">
		      </div>
		    </div>
		 <!-- <fieldset class="layui-elem-field layui-field-title site-title">
	      	<div class="layui-form-item">
		    <div class="layui-inline">
		      <label class="layui-form-label">钱包地址</label>
		      <div class="layui-input-inline" >
		        <input type="text" name="Sell_price" id="phone" placeholder="请输入手机" autocomplete="off" class="layui-input">
		      </div>
		    </div>
		  </div>
		  <div class="layui-form-item">
		    <div class="layui-inline">
		      <label class="layui-form-label">钱包标签</label>
		      <div class="layui-input-inline">
			  <input type="text" name="Name" id="contract" placeholder="请输入合约地址" autocomplete="off" class="layui-input" style="width:300px;">
		      </div>
		    </div>
		  </div>
		  </fieldset>
		  <fieldset class="layui-elem-field layui-field-title site-title">
	      <div class="layui-form-item">
		    <div class="layui-inline">
		      <label class="layui-form-label">项目投资额</label>
		      <div class="layui-input-inline" >
		        <input type="text" name="Sell_price" id="phone" placeholder="请输入手机" autocomplete="off" class="layui-input">
		      </div>
		    </div>
		  </div>
		  <div class="layui-form-item">
		    <div class="layui-inline">
		      <label class="layui-form-label">锁仓比例</label>
		      <div class="layui-input-inline">
			  <input type="text" name="Name" id="contract" placeholder="请输入合约地址" autocomplete="off" class="layui-input" style="width:300px;">
		      </div>
		    </div>
		  </div>
		  <div class="layui-form-item">
		    <div class="layui-inline">
		      <label class="layui-form-label">锁仓时长</label>
		      <div class="layui-input-inline" >
		        <input type="text" name="Sell_price" id="address" placeholder="请输入帐号地址" autocomplete="off" class="layui-input" style="width:300px;">
		      </div>
		    </div>
		  </fieldset>-->
		 <!--<div class="layui-form-item">
		    <div class="layui-inline">
		      <label class="layui-form-label">金额阈值</label>
		      <div class="layui-input-inline" >
		        <input type="text" name="Sell_price" id="value" placeholder="请输入金额" autocomplete="off" class="layui-input">
		      </div>
		    </div>
		  </div>-->
		  <div class="layui-form-item">
		    <div class="layui-input-block">
		      <button class="layui-btn" id="add">确认提交</button>
		<!--	  <input type="hidden" id="pic_path">-->
		      <button type="reset" class="layui-btn layui-btn-primary">取消</button>
		    </div>
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
	layui.use(['element','layer','jquery','table','tree'], function(){
	  var element = layui.element
		,form=layui.form
		,layer=layui.layer
		,$=layui.jquery
		,table=layui.table;
	  //layer.msg("你好");
	//tree 列表
	layui.tree({
	  elem: '#tree' //传入元素选择器'
	  ,nodes: [{ //节点
	    name: '项目管理'
		,href:'/pm'	
		,spread:true
	    ,children: [{
	      name: '实时数据'
		,href:'/realtimedata'
	    },{
	       name: '钱包数据'
		 ,href:'/wallet'
		 ,spread:true	
		  ,children: [{
	        name: '钱包监控'
			,href:'/walletmonitor'	
	      },{
	        name: '饼图'
			,href:'/walletpie'
	      },{
	        name: '钱包增长'
			,href:'/walletincrease'			
	      }]
	    },{
	      name: '基石投资者管理'
		 ,spread:true		  
		  ,href:'/getstockholder'
		  ,children: [{
	        name: '新增投资者'
		    ,href:'/addmonitor'
			
	      },{
	        name: '批量导入'
	      }]
	    },{
	      name: '项目预警'
		  ,href: '/getearlywarn'
	    },{
	      name: '数据分析'
		  ,children: [{
	        name: '流转图谱'
	      },{
	        name: '关联图谱'
	      },{
	        name: '持仓分布'
	      },{
	        name: '交易分析'
	      },{
	        name: '资产分析'
			,children: [{
	        name: '详情'
	      	}]
	      }]
	    }]		
	  },{
	    name: 'USTD汇率'
	    ,children: [{
	      name: '短信监控通知'	      
	    }]
	  },{
	    name: '系统设置'
	    ,children: [{
	      name: '账号续约'	      
	    }]
	  },{
	    name: '消息通知'
		,href:'/getnotifcationmessage'
	  }]
	});
//添加
	$('#add').on('click',function(){
	    var data={
			'name':$("#name").val(),
			'contract':$("#contract").val(),
			'address':$("#address").val(),
			'phone':$("#phone").val(),
			};
			$.ajax({
				type:"POST",
				contentType:"application/json;charset=utf-8",
				url:"/addmonitor_action",
				data:JSON.stringify(data),
				async:false,
				error:function(request){
					alert("post error")						
				},
				success:function(res){
					if(res.code==200){
						alert("新增成功")
						window.location.reload();						
					}else{
						alert("新增失败")
					}						
				}
			});							
		return false;
	}); 
	
			
  });

	
	
</script>

</body>
</html>