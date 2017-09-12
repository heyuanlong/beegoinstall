<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta http-equiv="Pragma" content="no-cache">
<meta http-equiv="Cache-Control" content="no-cache">
<meta http-equiv="Expires" content="0">
<title>resume</title>
<link href="/static/css/resume.css" type="text/css" rel="stylesheet" charset="utf-8" />
<link rel="stylesheet" href="/static/editor.md/css/editormd.preview.css" />

<script type="text/javascript">

</script>

</head>
<body>
  
  {{template "home/inc_header.tpl" .}}
  <div class="container">
    <div class="itemcontent">
      <div class="content_list">
          <div class="item">
            <div class="title"><a class="article_title" href="/article/{{.id}}">{{.title}}</a></div>
            <div class="content_attr">
              time:<span class="time">{{.time}}</span>
              by:<span class="author">{{.author}}</span>
              classify:<span class="classify"><a href="/category/{{.classifyid}}">{{.classifyname}}</a> </span>
            </div>
            <div id="content" class="content_n">
                  <textarea style="display:none;">{{.content}}</textarea>          
            </div>
          </div>
      </div>
      <div class="comment">
        <div class="comment_list">
            {{range $k,$v := .commentlist}}
              <div  class="comment_head">post by:{{$v.name}}|{{$v.time}}</div>
              <div>{{$v.content}}</div>
              <hr/>
            {{end}}
        </div>
        <div class="comment_fabiao">
          <form>
            <input type="hidden" name="articleid" value="{{.id}}">
            <input type="text" name="name" id="name" placeholder="名字"><br>
            <textarea name="content"  id="comment_content" ></textarea><br>
            <img src="/code" onclick="this.src='/code?'+Math.random()"> 
            <input type="text" name="code" id="code" placeholder="验证码">
            <input type="submit" id="fabiao" value="发表">
          </form>
        </div>
      </div>
    </div>
  </div>

  {{template "home/inc_footer.tpl" .}}

</body>

<script src="/static/js/jquery.min.js"></script>
<script src="/static/editor.md/lib/marked.min.js"></script>
<script src="/static/editor.md/lib/prettify.min.js"></script>
<script src="/static/editor.md/lib/raphael.min.js"></script>
<script src="/static/editor.md/lib/underscore.min.js"></script>
<script src="/static/editor.md/lib/sequence-diagram.min.js"></script>
<script src="/static/editor.md/lib/flowchart.min.js"></script>
<script src="/static/editor.md/lib/jquery.flowchart.min.js"></script>

<script src="/static/editor.md/editormd.js"></script>
<script type="text/javascript">
    $(function() {
          editormd.markdownToHTML("content", {
              htmlDecode      : "style,script,iframe",  // you can filter tags decode
              emoji           : true,
              taskList        : true,
              tex             : true,  // 默认不解析
              flowChart       : true,  // 默认不解析
              sequenceDiagram : true,  // 默认不解析
          });
    });

  $(function() {
    $("form").submit( function () {
      var articleid = $("input[name='articleid']").val();
      var name = $("input[name='name']").val();
      var content = $("textarea[name='content']").val();
      var code = $("input[name='code']").val();
      $.ajax({
        type: "post",
        url: "/comment",
        data: { articleid: articleid, name: name, content: content, code: code },
        dataType: "json",
        success: function(msg){
          console.log(msg);
          if (msg.result) {
            alert("发表成功");
            window.location = msg.refer;
          } else{
            alert("发表失败-"+msg.msg);
          };
        }
      });
      return false;
    });
  })

</script>

</html>

