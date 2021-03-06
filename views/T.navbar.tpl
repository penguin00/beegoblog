{{define "navbar"}}
<a class="navbar-brand" href="/">我的博客</a>
<div>
	<ul class="nav navbar-nav">
		<li {{if .isHome}}class="active"{{end}}><a href="/">首页</a></li>
		<li {{if .isCategory}}class="active"{{end}}><a href="/category">分类</a></li>
		<li {{if .isTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
	</ul>
</div>

<div class="pull-right">
	<ul class="nav navbar-nav">
		{{if .isLogin}}
		<li><a href="/login?exit=true">退出登录</a></li>
		{{else}}
		<li><a href="/login">管理员登录</a></li>
		{{end}}
		<li>{{.language}}</li>
	</ul>
</div>
{{end}}