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
		<form class="layui-form layui-form-pane1" action="" onsubmit="javascript:return false;">	  
		  <blockquote class="layui-elem-quote">变动通知</blockquote>
		  <div class="layui-form-item">
			<input type="checkbox" name="switch" lay-skin="switch" lay-text="是|否" checked>
		  </div>
		  <blockquote class="layui-elem-quote">单笔交易</blockquote>
		  <div class="layui-form-item">
				<div class="layui-form-mid">单笔交易超过</div>
			    <div class="layui-input-inline" style="width: 30px;">
					 <input type="text" name="Sell_price" id="sell_price" value="{{.Price}}" autocomplete="off" class="layui-input">
				</div>
				<div class="layui-form-mid">or</div>
				<div class="layui-input-inline" style="width: 30px;">
					 <input type="text" name="Sell_price" id="sell_price" value="{{.Price}}" placeholder="%" autocomplete="off" class="layui-input">
				</div>
		  </div>
		  <blockquote class="layui-elem-quote">累计交易</blockquote>
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
		  </div>
		  <div class="layui-form-item">
		      <button class="layui-btn" id="add">保存</button>
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
	//自动加载
	$(function(){
		if({{.campus}}!=""){
			$("#campus").val({{.campus}});			
			form.render('select');	
		}				
	});
	
	//获取下拉列表
	form.on('select(campus_select)',function(data){
		//layer.msg(data)
		console.log(data.value);
		window.location.href="/v1/dining_room?campus="+data.value;
		
	});
	
	//点击新增按钮
	$('#addroom').on('click',function(){
		//layer.msg("点击添加按钮");
		var cp=$("#campus").val();
		//iframe窗
		layer.open({
		  type: 2,
		  title: '新增餐厅',
		  //closeBtn: 0, //不显示关闭按钮
		  shadeClose: true,
		  shade: false,
		  area: ['893px', '600px'],
		 // offset: 'rb', //右下角弹出
		  //time: 2000, //2秒后自动关闭
		  maxmin: true,
		  anim: 2,
		  content: ['/v1/dining_room/add?campus='+cp], //iframe的url，no代表不显示滚动条
		  cancel: function(index, layero){ 
		  if(confirm('确定要关闭么')){ //只有当点击confirm框的确定时，该层才会关闭
		    layer.close(index)
			//window.parent.location.reload();
			//重载表格
			table.reload('listReload', {});
		  }
		  return false; 
		  },
		});
	});
	
	  //table 渲染
	  table.render({
	    elem: '#roomList'
	    ,height: 315
	    ,url: '/v1/dining_room/getdata' //数据接口
	    ,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[ //表头
		  {field: 'RoomPicPath', title: '窗口图片', width:'11%',height:'20%'
			,templet:function(d){
				var list=d.RoomPicPath.split(',')
				//alert(list.length)
				if(list.length!=1){
					for(var i=0;i<list.length-1;i++){
						return '<img src="'+'/'+list[i]+'">'				
					}
				}else{
					return ""	
				}						
			}}
	      ,{field:'Name', title:'窗口名称', width:120}
	      ,{field:'Status',  title:'状态', width:120}
	      ,{field:'Time', title:'供应时段', width:120}
		  ,{fixed: 'right', title:'操作',width:200, align:'center', toolbar: '#barDemo'}
	    ]]
	  });		
		//监听工具条
		table.on('tool(room)', function(obj){ //注：tool是工具条事件名，test是table原始容器的属性 lay-filter="对应的值"
		    var data = obj.data //获得当前行数据
		    ,layEvent = obj.event; //获得 lay-event 对应的值
		    if(layEvent === 'edit'){
		      //layer.msg('查看操作');		
			  layer.open({
			  type: 2,
			  title: '查看菜品',
			  //closeBtn: 0, //不显示关闭按钮
			  shadeClose: true,
			  shade: false,
			  area: ['893px', '600px'],
			 // offset: 'rb', //右下角弹出
			  //time: 2000, //2秒后自动关闭
			  maxmin: true,
			  anim: 2,
			  content: ['/v1/dining_room/edit?id='+data.Id], //iframe的url，no代表不显示滚动条
			  cancel: function(index, layero){ 
			  if(confirm('确定要关闭么')){ //只有当点击confirm框的确定时，该层才会关闭
			    layer.close(index)
			  }
			  return false; 
			  },
		});
	    } else if(layEvent === 'del'){
	      layer.confirm('真的删除行么', function(index){
	        var jsData={'id':data.Id}
			$.post('/v1/dining_room/del', jsData, function (out) {
                if (out.code == 200) {
                    layer.alert('删除成功了', {icon: 1},function(index){
                        layer.close(index);
                        table.reload({});
                    });
                } else {
                    layer.msg(out.message)
                }
            }, "json");
			obj.del(); //删除对应行（tr）的DOM结构
	        layer.close(index);
	        //向服务端发送删除指令
	      });
	    } else if(layEvent === 'stop'){
	      layer.msg('编辑操作');
		  layer.open({
			  type: 2,
			  title: '编辑菜品',
			  //closeBtn: 0, //不显示关闭按钮
			  shadeClose: true,
			  shade: false,
			  area: ['893px', '600px'],
			 // offset: 'rb', //右下角弹出
			  //time: 2000, //2秒后自动关闭
			  maxmin: true,
			  anim: 2,
			  content: ['/v1/dish/edit_show?id='+data.Id], //iframe的url，no代表不显示滚动条
			  cancel: function(index, layero){ 
			  if(confirm('确定要关闭么')){ //只有当点击confirm框的确定时，该层才会关闭
			    layer.close(index)
			  }
			  return false; 
			  },
		});					  
	    }
	  });	
			
  });

	
	
</script>

</body>
</html>