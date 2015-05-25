{{define "bottom_scripts"}}
<!-- jQuery -->
<script src="/static/libs/jquery/dist/jquery.min.js"></script>

<!-- Bootstrap Core JavaScript -->
<script src="/static/libs/bootstrap/dist/js/bootstrap.min.js"></script>

<!-- Metis Menu Plugin JavaScript -->
<script src="/static/libs/metisMenu/dist/metisMenu.min.js"></script>

<!-- Morris Charts JavaScript -->
<script src="/static/libs/raphael/raphael-min.js"></script>
<script src="/static/libs/morrisjs/morris.min.js"></script>

<!-- Custom Theme JavaScript -->
<script src="/static/libs/startbootstrap-sb-admin-2/dist/js/sb-admin-2.js"></script>

{{with .Common.Scripts}}{{range .}}
<script src="{{.}}"></script>{{end}}{{end}}
{{end}}