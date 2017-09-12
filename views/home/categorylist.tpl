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
        {{range $k,$v := .list}}
          <div class="item">
            <div class="title"><a class="article_title" href="/article/{{$v.id}}">{{$v.title}}</a></div>
            <div class="content_attr">
              time:<span class="time">{{$v.time}}</span>
              by:<span class="author">{{$v.author}}</span>
              classify:<span class="classify">{{$v.classifyname}}</span>
            </div>
          </div>
        {{end}}
      </div>

      <div class="page_forword">
        {{if .prev_page_flag}}
        <a class="prev_page" href="{{.prev_page}}" class="page-nav">上一页</a>
        {{end}}
        {{if .next_page_flag}}
        <a class="next_page" href="{{.next_page}}" class="page-nav">下一页</a>
        {{end}}
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
       {{range $k,$v := .list}}
          editormd.markdownToHTML("content_{{$v.id}}", {
              htmlDecode      : "style,script,iframe",  // you can filter tags decode
              emoji           : true,
              taskList        : true,
              tex             : true,  // 默认不解析
              flowChart       : true,  // 默认不解析
              sequenceDiagram : true,  // 默认不解析
          });
       {{end}}
    });
</script>

</html>
