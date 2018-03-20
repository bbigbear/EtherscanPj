<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
	</head>
	<body>
		<div class="container" style="width:300px;padding-top:150px">
			<form class="form-signin" method="POST" action="/login">
		        <h2 class="form-signin-heading">监控管理系统</h2>
		        <label for="inputAccount" class="sr-only">帐号</label>
		        <input id="inputAccount" class="form-control" placeholder="帐号" required autofocus name="inputAccount">
		        <label for="inputPassword" class="sr-only">密码</label>
		        <input type="password" id="inputPassword" class="form-control" placeholder="密码" required name="inputPassword">
		        <div class="checkbox">
		          <label>
		            <input type="checkbox" value="remember-me">记住密码
		          </label>
		        </div>
		        <button class="btn btn-lg btn-primary btn-block" type="submit">登陆</button>
	      	</form>
		</div>
	</body>
</html>
