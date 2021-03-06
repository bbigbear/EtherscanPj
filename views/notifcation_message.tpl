<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>消息通知</title>
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
		<blockquote class="layui-elem-quote" style="margin-top:10px;">消息通知</blockquote>					
		<div>
			<span class="layui-breadcrumb" lay-separator="|">										  					
				<i class="layui-icon">&#xe640;</i>
				<a id="del">删除</a>
			</span>
		</div>
		<table id="MessageList" lay-filter="message"></table>				
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
	//自动加载
	$(function(){
		if({{.campus}}!=""){
			$("#campus").val({{.campus}});			
			form.render('select');	
		}				
	});
	//tree 列表
	layui.tree({
	  elem: '#tree' //传入元素选择器'
	  ,nodes: [{ //节点
	    name: '项目管理'
	    ,children: [{
	      name: '实时数据'
	    },{
	      name: '钱包数据'
		  ,children: [{
	        name: '钱包监控'
	      },{
	        name: '饼图'
	      },{
	        name: '钱包增长'
	      }]
	    },{
	      name: '基石投资者管理'
		  ,href:'/getstockholder'
		  ,children: [{
	        name: '新增投资者'
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
	    ,page: true //开启分页
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
	
	//批量删除
	$('#del').on('click',function(){				
		var str="";
		var checkStatus=table.checkStatus('listReload')
		,data=checkStatus.data;
		if(data.length==0){
			alert("请选择要删除的数据")
		}else{
			for(var i=0;i<data.length;i++){
				str+=data[i].Id+",";
			}
			layer.confirm('是否删除这'+data.length+'条数据?',{icon:3,title:'提示'},function(index){
				//window.location.href="/v1/delmultidata?id="+str+"";
				$.ajax({
					type:"POST",
					url:"/delnotifcationdata",
					data:{
						id:str	
					},
					async:false,
					error:function(request){
						alert("post error")						
					},
					success:function(res){
						if(res.code==200){
							alert("删除成功")	
							//重载表格
							table.reload('listReload', {							  
							});												
						}else{
							alert("删除失败")
						}						
					}					
				});				
				layer.close(index);
			});
		}
		return false;
	});	
			
  });

	
	
</script>

</body>
</html>