<!DOCTYPE html>

<html>
<head>
  <title>Install</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link href="/static/css/search.css" type="text/css" rel="stylesheet" charset="utf-8" />


  <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
  <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
  <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
  <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">
</head>

<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>



<body>

<div class="contain">

  <img src="/static/img/header.png">

  <form action="search" method="post">
    <input  class="form-input" type="text" name="data" placeholder="input" value="{{.value}}">
    <button class="form-button" type="submit">search 一下</button>
  </form>

  <div class="content">
    {{range $k,$v := .list}}
    <div class="item">
      <div class="title"><a class="article_title" href="/article/{{$v.id}}">{{$v.title}}......</a></div>
    </div>
    {{end}}
  </div>


</div>

</body>
</html>
