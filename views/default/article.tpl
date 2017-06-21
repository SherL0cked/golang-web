<!DOCTYPE html>

<head>

    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Clean Blog - Start Bootstrap Theme</title>

    <!-- Bootstrap Core CSS -->
    <link href="static/css/bootstrap.min.css" rel="stylesheet">

    <!-- Additional fonts for this theme -->
    <link href="static/css/font-awesome.min.css" rel="stylesheet" type="text/css">
    <link href='https://fonts.googleapis.com/css?family=Lora:400,700,400italic,700italic' rel='stylesheet' type='text/css'>
    <link href='https://fonts.googleapis.com/css?family=Open+Sans:300italic,400italic,600italic,700italic,800italic,400,300,600,700,800' rel='stylesheet' type='text/css'>

    <!-- Custom styles for this theme -->
    <link href="static/css/clean-blog.css" rel="stylesheet">
	<!-- Animate.css-->
	<link href="static/css/animate.css" rel="stylesheet"> 

    <!-- Temporary navbar container fix until Bootstrap 4 is patched -->
    <style>
    .navbar-toggler {
        z-index: 1;
    }
    
    @media (max-width: 576px) {
        nav > .container {
            width: 100%;
        }
    }
    </style>

</head>

<body>

    <!-- Navigation -->
    <nav class="navbar fixed-top navbar-toggleable-md navbar-light" id="mainNav">
        <div class="container">
            <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
                Menu <i class="fa fa-bars"></i>
            </button>
            <a class="navbar-brand page-scroll" href="../index.html">找法务</a>
            <div class="collapse navbar-collapse" id="navbarResponsive">
                <ul class="navbar-nav ml-auto">
                    <li class="nav-item">
                        <a class="nav-link page-scroll" href="../">律师简介</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link page-scroll" href="../articles&page1">经典案例</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link page-scroll" href="../journals.html">资讯速递</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link page-scroll" href="../contact.html">服务项目</a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>

	
	 <header class="intro-header" style="background-image: url('static/img/home-bg.jpg')">
        <div class="container">
            <div class="row">
                <div class="col-lg-8 offset-lg-2 col-md-10 offset-md-1">
                    <div class="site-heading">
                        <h1>{{.keywords}}</h1>
                        <span class="subheading" hidden>A Blog Theme by Start Bootstrap</span>
                    </div>
                </div>
            </div>
        </div>
    </header>
		
	<div class="container">
        <div class="row">
            <div class="col-lg-8 offset-lg-2 col-md-10 offset-md-1">
                {{range $k, $v := .Mapped}}
				<div class="post-preview">
                    <a href="articles/articles{{range $.ID}}{{index $v .}}{{end}}">
                        <h2 class="post-title">
						    {{range $.Title}}
								{{index $v .}}
							{{end}}
                        </h2>
						<h3 class="post-subtitle">
                            {{range $.Subtitle}}
								{{index $v .}}
							{{end}}
                        </h3>
                    </a>
                 <p class="post-meta">Posted by {{range $.Name}}{{index $v .}}{{end}}</a> on {{range $.Date}}{{index $v .}}{{end}}</p>
                </div>
                <hr>
				{{end}}
                
                <!-- Pager -->
                <div class="clearfix">
                    <a class="btn btn-secondary float-right" href="../articles&page{{.Next}}">下一页 &rarr;</a>
					<a class="btn btn-secondary float-left" href="../articles&page{{.Previous}}">&larr; 上一页</a>
                </div>
            </div>
        </div>
    </div>

    <hr>
		
</body>

</html>
