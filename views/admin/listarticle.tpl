<!DOCTYPE html>

<html>
  	<head>
    	<title>文章列表</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<script src="/static/js/jquery.min.js"></script>
		<link rel="stylesheet" href="/static/css/admin.css" />
		<style>
		
			li{
				display: block;
   				min-width: 130px;
    			width: auto;
			}
		
		</style>
	</head>
  	<body>
  		<div class="container">
	  	  	{{template "admin/admin_menu.tpl" .}}
	  		<div class="main">
		  		<table>
					{{range $k,$v := .list}}
					<tr>
					<td>{{$v.time}}:</td>
					<td><a href="/admin/article/{{$v.id}}">{{$v.title}}</a></td>
					<td><a href="/admin/update_article/{{$v.id}}">修改</a></td>
					<td><a href="/admin/delete_article/{{$v.id}}">删除</a></td>
					{{if eq $v.status "0"}}
						<td><a href="/admin/issue_article/{{$v.id}}">发布</a></td>
					{{end}}
					{{if eq $v.status "1"}}
						<td><a href="/admin/draft_article/{{$v.id}}">转为草稿</a></td>
					{{end}}
					</tr>
					{{end}}
				</table>
			</div>
		</div>
	</body>
</html>

