<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>消息通知</title>
  <meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <meta name="apple-mobile-web-app-status-bar-style" content="black"> 
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="format-detection" content="telephone=no">

  <link rel="stylesheet" href="/static/css/layui.css">

<style>
body{padding: 10px;}
</style>
</head>
<body>
	<div style="padding: 15px;">
	  	<blockquote class="layui-elem-quote" style="margin-top:10px;">消息通知</blockquote>					
		<div>
			<span class="layui-breadcrumb" lay-separator="|">										  					
				<i class="layui-icon">&#xe640;</i>
				<a id="dish_del">删除</a>
			</span>
		</div>
		<table id="MessageList" lay-filter="message"></table>		
	</div>
<br><br><br>


<script src="/static/layui.js"></script>
<!-- <script src="../build/lay/dest/layui.all.js"></script> -->

<script>
layui.use(['form','laydate','upload','jquery','layedit','element','table'], function(){
  var form = layui.form
  ,laydate=layui.laydate
  ,upload = layui.upload
  , $ = layui.jquery
  ,layedit=layui.layedit
  ,element=layui.element
  ,table=layui.table;
  	
	//table 渲染
	  table.render({
	    elem: '#balanceList'
	    ,height: 80
	    ,url: '/search/balance'//数据接口
	    //,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[ //表头
	      {field:'Status', title:'Status', width:120}
		  ,{field:'Message',  title:'Message', width:120}
	      ,{field:'Result',  title:'Result', width:300}
	    ]]
	  });
	//table 渲染
	  table.render({
	    elem: '#transactionList'
	    ,height: 300
	    ,url: '/search/transaction'//数据接口
	    ,page: true //开启分页
		,id: 'listReload1'
	    ,cols: [[ //表头
		   {type:'checkbox', fixed: 'left'}
	      ,{field:'blockNumber', title:'预警时间', width:120}
		  ,{field:'timeStamp',  title:'预警对象', width:120}
	      ,{field:'hash',  title:'预警类型', width:150}
		  ,{field:'blockHash',  title:'变动数量', width:150}
		  ,{field:'transactionIndex',  title:'占比', width:150}
		  ,{field:'from',  title:'哈希值', width:150}
	    ]]
	  });		
		

	//点击查询	
	$('#search').on('click',function(){
		table.reload('listReload', {
                    where: {
                        address:$("#account").val() ,
                    }
                });
	});
	
	
});
</script>

</body>
</html>