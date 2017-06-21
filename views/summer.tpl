<!DOCTYPE html>
<html lang= "zh-CN">

<head>
  <meta charset="UTF-8">
  <title>上传文件</title>
  <script src="/"></script>
  <link href="http://netdna.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.css" rel="stylesheet">
  <script src="http://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.4/jquery.js"></script> 
  <script src="http://netdna.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.js"></script> 
  <link href="http://cdnjs.cloudflare.com/ajax/libs/summernote/0.8.4/summernote.css" rel="stylesheet">
  <script src="http://cdnjs.cloudflare.com/ajax/libs/summernote/0.8.4/summernote.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/summernote/0.8.4/lang/summernote-zh-CN.js"></script>

</head>

<body>
<div style="padding: 10px;">
	<div class="input-group input-group-lg" id="Subjects">
	  <input type="text" class="form-control" id="field1" placeholder="标题">
	</div>
 </div>
 
 <div id="summernote">
	  <p></p>
 </div>

<!--div>
	Field2: <input type="text" id="field2">
	<input type="submit" value="提交" onclick="postForm()"/>
</div-->

<form id="Articles">
    <input name="Subjects" id = "subject" type="text" hidden/>
	<input name="Content" id = "ct" type="text" hidden/>
    <button onclick="postForm()">复制文本</button>
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
		alert(content)
	}

 </script>
  
</body>

</html>