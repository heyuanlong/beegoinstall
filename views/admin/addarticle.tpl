<!DOCTYPE html>

<html>
  	<head>
    	<title>添加文章</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    	<link rel="stylesheet" href="/static/editor.md/css/editormd.css" />
    	<link rel="stylesheet" href="/static/css/admin.css" />
		
		<style>
		.center{
			display: block;
		    margin: 0 auto;
		    position: relative;
		    width: 100%;
		    text-align: center;
		}
		</style>
	</head>
  	<body>
  		<div class="container">
	  		{{template "admin/admin_menu.tpl" .}}
	  		<div class="main">
				<form action="/admin/addarticle" method="post" class="center">
					<input style="width: 90%;height:30px;" type="text" name="title" id="title" class="input" placeholder="title" value="{{.title}}"><br>
					<input type="hidden" name="status" id="status">
					<input type="hidden" name="id" value="{{.id}}">
					<input style="width: 90%;height:30px;" type="text" name="keywords" id="keywords" class="input" placeholder="keywords" value="{{.keywords}}"><br>
					<select  style="width: 90%;height:30px;"  name="classifyid" id="classifyid" >
						{{range $k,$v := .classify}}
							<option value="{{$v.id}}">{{$v.name}}</option>
						{{end}}
					</select>
					<div id="test-editormd">
						   <textarea name="content" style="display:none;">{{.content}}</textarea>
					</div>
					<input type="button" id="fabu" value="发布">
					<input type="button" id="caogao" value="保存为草稿">
				</form>
	 		</div>
 		</div>
	</body>

<script src="/static/js/jquery.min.js"></script>
<script src="/static/editor.md/editormd.min.js"></script>
<script type="text/javascript">

var testEditor;
$(function() {
    testEditor = editormd("test-editormd", {
        width   : "90%",
        height  : 640,
        syncScrolling : "single",
        path    : "/static/editor.md/lib/",
        imageUpload : true,
        imageFormats : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
        imageUploadURL : "/admin/uploadimg",
    });

    $("form").submit( function (event) {
		if (testEditor.getMarkdown() == "" || $("#title").val() == "") {
			event.preventDefault();
		}
		return ;
	})

	$("#fabu").click(function() {  
    	$("#status").val(1);
    	$("form").submit();
	}); 
	$("#caogao").click(function() {  
    	$("#status").val(0);
    	$("form").submit();
	}); 


	$("#classifyid").val({{.classifyid}});

});


</script>
</html>

