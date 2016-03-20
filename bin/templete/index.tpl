<!DOCTYPE html>
<html>
<head>
	<title>LogPreview</title>
</head>
<body>

{{range .Content}}
<table width:100% style="border:1px solid black;">
	<tr>
		<td>Time</td>
		<td>{{.Time}}</td>
	</tr>
	<tr>
		<td>Type</td>
		<td>{{.Type}}</td>
	</tr>
	<tr>
		<td>Condition</td>
		<td>{{.Condition}}</td>
	</tr>
	<tr>
		<td>StackTrace</td>
		<td>{{.StackTrace}}</td>
	</tr>
	<tr>
		<td>OperatingSystem</td>
		<td>{{.OperatingSystem}}</td>
	</tr>
	<tr>
		<td>Device</td>
		<td>{{.Device}}</td>
	</tr>
	<tr>
		<td>DeviceIP</td>
		<td>{{.DeviceIP}}</td>
	</tr>
	<tr>
		<td>DeviceModel</td>
		<td>{{.DeviceModel}}</td>
	</tr>
</table>
<br>
{{end}}

</body>
</html>