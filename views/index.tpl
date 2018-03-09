<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>以太网交易详细</title>
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
	  	<blockquote class="layui-elem-quote" style="margin-top:10px;">交易记录</blockquote>			
		<form class="layui-form layui-form-pane1" action="" onsubmit="javascript:return false;">
		 	<div class="layui-form-item">
				<div class="layui-inline">
				    <label class="layui-form-label">帐号地址</label>
				    <div class="layui-input-inline">
				    	<input type="text" class="layui-input" id="account"  style="width:300px;">
				    </div>
					<div class="layui-inline">
				        <button class="layui-btn" id="search" style="margin-left:200px;">查询</button>
				    </div>  
				</div> 
			</div>
		</form>
		<blockquote class="layui-elem-quote" style="margin-top:10px;">Get ERC20-Token Account Balance for TokenContractAddress</blockquote>			
		<table id="balanceList" lay-filter="room"></table>
		<blockquote class="layui-elem-quote" style="margin-top:10px;">Get a list of 'Normal' Transactions By Address</blockquote>			
		<table id="transactionList" lay-filter="transaction"></table>
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
	      {field:'blockNumber', title:'blockNumber', width:120}
		  ,{field:'timeStamp',  title:'timeStamp', width:120}
	      ,{field:'hash',  title:'hash', width:150}
		  ,{field:'blockHash',  title:'blockHash', width:150}
		  ,{field:'transactionIndex',  title:'transactionIndex', width:150}
		  ,{field:'from',  title:'from', width:150}
		  ,{field:'to',  title:'to', width:150}
		  ,{field:'value',  title:'value', width:150}
		  ,{field:'gas',  title:'gas', width:150}
		  ,{field:'gasPrice',  title:'gasPrice', width:150}
		  ,{field:'cumulativeGasUsed',  title:'cumulativeGasUsed', width:150}
		  ,{field:'gasUsed',  title:'gasUsed', width:150}
		  ,{field:'confirmations',  title:'confirmations', width:150}
	    ]]
	  });		
		

	//点击查询	
	$('#search').on('click',function(){
		table.reload('listReload', {
                    where: {
                        address:$("#account").val() ,
                    }
                });
				
		table.reload('listReload1', {
                    where: {
                        address:$("#account").val() ,
                    }
                });		
	});
	
	
});
</script>

</body>
</html>