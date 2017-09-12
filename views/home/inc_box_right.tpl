    <div class="box_right">
      <div>
        <div class="box_title">最热博文</div>
        <ul class="box_ul">
          {{range $k,$v := .hottest}}
          <li>
            <a class="box_right_article_title" href="/article/{{$v.id}}">>{{$v.title}}</a>
            <span class="view_count">{{$v.count}}</span>
          </li>
          {{end}}
        </ul>
      </div>

      <div>
          <div class="box_title">分类</div>
          <ul class="box_ul">
            {{range $k,$v := .classify}}
            <li>
              <a class="box_right_article_title" href="/category/{{$v.classifyID}}">>{{$v.name}}</a>
              <span class="view_count">{{$v.count}}</span>
            </li>
            {{end}}
          </ul>
      </div>
    </div>
    