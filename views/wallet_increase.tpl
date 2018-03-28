<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>钱包增长</title>
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
      <li class="layui-nav-item"><a href="/logout">退出</a></li>
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
		<div id='main' style='width:600px;height:400px;'></div>				
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
<script src="http://echarts.baidu.com/build/dist/echarts-all.js"></script>
	<script>
          //基于准备好的DOM，初始化echarts实例
        var myChart = echarts.init(document.getElementById('main'));
          //指定图表的配置项和数据
        var option = {
            title:{
                text:'钱包增长'
            },
            //提示框组件
            tooltip:{
                //坐标轴触发，主要用于柱状图，折线图等
                trigger:'axis'
            },
            //图例
            legend:{
                data:['个数']
            },
            //横轴
            xAxis:{
				name:'七天内',
                data:[
					{{range .maps}}
					{{.timestamp}},
					{{end}}
				]
            },
            //纵轴
            yAxis:{},
            //系列列表。每个系列通过type决定自己的图表类型
            series:[{
                name:'个数',
                //折线图
                type:'line',
                data:[
					{{range .maps}}
					{{.address_num}},
					{{end}}
				]
            }]
        };
        //使用刚指定的配置项和数据显示图表
        myChart.setOption(option);
    </script>
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
		
			
  });
	
</script>

</body>
</html>