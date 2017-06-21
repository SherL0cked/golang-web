<!DOCTYPE html>
<html lang= "zh-CN">

<head>
  <meta charset="UTF-8">
  <title>文本更新</title>
  <script src="/"></script>
  <link href="http://netdna.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.css" rel="stylesheet">
  <script src="http://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.4/jquery.js"></script> 
  <script src="http://netdna.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.js"></script> 
  <link href="http://cdnjs.cloudflare.com/ajax/libs/summernote/0.8.4/summernote.css" rel="stylesheet">
  <script src="http://cdnjs.cloudflare.com/ajax/libs/summernote/0.8.4/summernote.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/summernote/0.8.4/lang/summernote-zh-CN.js"></script>
  <!-- Latest compiled and minified CSS -->
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-select/1.10.0/css/bootstrap-select.min.css">

<!-- Latest compiled and minified JavaScript -->
<script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-select/1.10.0/js/bootstrap-select.min.js"></script>

</head>

<body>


<div class="input-group input-group-lg" id="Subjects">
	 <input type="text" class="form-control" id="subject" placeholder="标题">
	<select id="usertype" name="usertype" class="selectpicker show-tick form-control" multiple data-live-search="false">文章类型
				<option value="0">经典案例</option>
				<option value="1">咨询速递</option>
	</select>
</div>
 
 <div id="summernote">
	  <p>正文</p>
 </div>

<!--div>
	Field2: <input type="text" id="field2">
	<input type="submit" value="提交" onclick="postForm()"/>
</div-->

<form id="Articles">
    <input name="Subjects" id = "st" type="text" hidden/>
	<input name="Content" id = "ct" type="text" hidden/>
    <button type="button" class="btn btn-lg btn-default" onclick="postForm()">提交</button>
</form>


<div>
<h1>{{.test}}</h1>
</div>

<script>
    
	$(document).ready(function() {
        $('#summernote').summernote({
		lang : "zh-CN",
		height :500,
		minHeight: null,             // set minimum height of editor
		maxHeight: null,             // set maximum height of editor
		focus: true                  // set focus to editable area after initializing summernote
		});
    });
	
	var postForm = function() {
		var content = $('#summernote').summernote('code')
		document.getElementById("ct").value=content
		document.getElementById("st").value=document.getElementById("subject").value
	}

 </script>
  
</body>

</html>