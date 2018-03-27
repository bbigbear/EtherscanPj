<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>实时数据</title>
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
		<div class="layui-tab">
		  <ul class="layui-tab-title">
		    <li>全部</li>
		    <li class="layui-this">1小时</li>
		    <li>24小时</li>
		    <li>一周</li>
		    <li>一月</li>
		  </ul>
		  <div class="layui-tab-content">
		    <div class="layui-tab-item">
		        <table id="AllList" lay-filter="all"></table>
<!--				<label class="layui-form-label">共{{.all_num}}笔</label>-->
		    </div>
		    <div class="layui-tab-item layui-show">
				<table id="HourList" lay-filter="hour"></table>
<!--				<label class="layui-form-label">共{{.hour_num}}笔</label>-->
			</div>
		    <div class="layui-tab-item">
				<table id="DayList" lay-filter="day"></table>
<!--				<label class="layui-form-label">共{{.day_num}}笔</label>-->
			</div>
		    <div class="layui-tab-item">
				<table id="WeekList" lay-filter="week"></table>
<!--				<label class="layui-form-label">共{{.week_num}}笔</label>-->
			</div>
		    <div class="layui-tab-item">
				<table id="MonthList" lay-filter="mouth"></table>
<!--				<label class="layui-form-label">共{{.month_num}}笔</label>-->
			</div>
		  </div>
		</div>				
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
	//hour
	//table 渲染
	  table.render({
	    elem: '#AllList'
	    ,height: 315
	    ,url: '/getrealtimedata?type=all' //数据接口
	    //,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[ //表头
	      {field:'transaction_hash', title:'TxHash', width:200}
		  ,{field:'timestamp',  title:'交易日期', width:180}
	      ,{field:'from_address',  title:'From', width:200}
		  ,{field:'to_address',  title:'To', width:200}
		  ,{field:'value',  title:'数量', width:200}
		  ,{field:'percent',  title:'总量占比', width:100}
		  ,{field:'status',  title:'预警状态', width:100}
		  ,{field:'transaction_status',  title:'交易状态', width:100}
	    ]]
	  });
	
	  table.render({
	    elem: '#HourList'
	    ,height: 315
	    ,url: '/getrealtimedata?type=hour' //数据接口
	    ,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[ //表头
	      {field:'transaction_hash', title:'TxHash', width:200}
		  ,{field:'timestamp',  title:'交易日期', width:180}
	      ,{field:'from_address',  title:'From', width:200}
		  ,{field:'to_address',  title:'To', width:200}
		  ,{field:'value',  title:'数量', width:200}
		  ,{field:'percent',  title:'总量占比', width:100}
		  ,{field:'status',  title:'预警状态', width:100}
		  ,{field:'transaction_status',  title:'交易状态', width:100}
	    ]]
	  });
	  table.render({
	    elem: '#DayList'
	    ,height: 315
	    ,url: '/getrealtimedata?type=day' //数据接口
	    ,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[ //表头
	      {field:'transaction_hash', title:'TxHash', width:200}
		  ,{field:'timestamp',  title:'交易日期', width:180}
	      ,{field:'from_address',  title:'From', width:200}
		  ,{field:'to_address',  title:'To', width:200}
		  ,{field:'value',  title:'数量', width:200}
		  ,{field:'percent',  title:'总量占比', width:100}
		  ,{field:'status',  title:'预警状态', width:100}
		  ,{field:'transaction_status',  title:'交易状态', width:100}
	    ]]
	  });
	  table.render({
	    elem: '#WeekList'
	    ,height: 315
	    ,url: '/getrealtimedata?type=week' //数据接口
	    ,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[ //表头
	      {field:'transaction_hash', title:'TxHash', width:200}
		  ,{field:'timestamp',  title:'交易日期', width:180}
	      ,{field:'from_address',  title:'From', width:200}
		  ,{field:'to_address',  title:'To', width:200}
		  ,{field:'value',  title:'数量', width:200}
		  ,{field:'percent',  title:'总量占比', width:100}
		  ,{field:'status',  title:'预警状态', width:100}
		  ,{field:'transaction_status',  title:'交易状态', width:100}
	    ]]
	  });
	  table.render({
	    elem: '#MonthList'
	    ,height: 315
	    ,url: '/getrealtimedata?type=month' //数据接口
	    ,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[ //表头
	      {field:'transaction_hash', title:'TxHash', width:200}
		  ,{field:'timestamp',  title:'交易日期', width:180}
	      ,{field:'from_address',  title:'From', width:200}
		  ,{field:'to_address',  title:'To', width:200}
		  ,{field:'value',  title:'数量', width:200}
		  ,{field:'percent',  title:'总量占比', width:100}
		  ,{field:'status',  title:'预警状态', width:100}
		  ,{field:'transaction_status',  title:'交易状态', width:100}
	    ]]
	  });	
		
					
			
  });

	
	
</script>

</body>
</html>