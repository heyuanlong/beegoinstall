<!DOCTYPE html>

<html>
  	<head>
    	<title>登录</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<script src="/static/js/jquery.min.js"></script>
		<style>
		.center{
			display: block;
		    margin: 0 auto;
		    position: relative;
		    width: 300px;
		    text-align: center;
		}
		.center .input{
			width: 100%;
		}
		</style>
	</head>
  	<body>
  		<div class="main">
			<form action="/admin/login" method="post" class="center">
				<input type="text" name="username" id="" class="input" placeholder="usernmae"><br>
				<input type="password" name="password" id="" class="input" placeholder="password"><br>
				<input type="submit" value="登录">
			</form>
 		</div>
	</body>
</html>
<script>
	$(document).ready(function (e) {
		$("form").submit( function () {
			var user_name = $("input[name='username']").val();
			var pass_word = $("input[name='password']").val();
			$.ajax({
				type: "post",
				url: "/admin/login",
				data: { username: user_name, password: pass_word },
				dataType: "json",
				success: function(msg){
					console.log(msg);
					if (msg.result) {
						alert("登录成功-"+msg.msg);
						window.location = msg.refer;
					} else{
						alert("登录失败-"+msg.msg);
					};
				}
			});
			return false;
		});
	})
</script>
