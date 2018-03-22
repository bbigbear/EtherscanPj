<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>项目管理</title>
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
       <!-- <li class="layui-nav-item"><a href="/getearlywarn">项目预警</a></li>
        <li class="layui-nav-item"><a href="/getnotifcationmessage">消息通知</a></li>-->
      </ul>
    </div>
  </div>
  <div class="layui-body">
    <!-- 内容主体区域 -->
    <div style="padding: 15px;">					
		 <div class="layui-row layui-col-space5" style="margin-bottom:50px;">
		    <div class="layui-col-md4" style="width:150px;">
		      <div class="layui-bg-gray"  style="height:100px;;">钱包</div>
		    </div>
		    <div class="layui-col-md4" style="width:150px;margin-left:50px;">
		      <div class="layui-bg-gray"  style="height:100px;;">基石投资者</div>
		    </div>
		    <div class="layui-col-md4" style="width:150px;margin-left:50px;">
		      <div class="layui-bg-gray"  style="height:100px;;">一小时流通量</div>
		    </div>
		 </div>
		<fieldset class="layui-elem-field layui-field-title site-title">
	      <legend><a name="color-design">项目预警</a></legend>
	    </fieldset>
		<table id="MessageList" lay-filter="message"></table>
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

			
	  //table 渲染
	  table.render({
	    elem: '#MessageList'
	    ,height: 315
	    ,url: '/getnotifcationdata' //数据接口
	    //,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[ //表头
		  {type:'checkbox', fixed: 'left'}
	      ,{field:'Time', title:'预警时间', width:160}
		  ,{field:'Target',  title:'预警对象', width:250}
	      ,{field:'Style',  title:'预警类型', width:120}
		  ,{field:'Num',  title:'变动数量', width:150}
		  ,{field:'Percent',  title:'占比', width:150}
		  ,{field:'Hash',  title:'哈希值', width:250}
	    ]]
	  });
							
  });
	
</script>

</body>
</html>