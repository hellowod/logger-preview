<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>LD</title>
    <style>
        table {
            width: 100%;
            border-collapse: collapse;
        }
        table, th, td {
            border: 1px solid black;
        }
        th, td {
            padding: 8px;
            text-align: left;
        }
    </style>
</head>
<body>
    <h1>LD-LOG</h1>
    <table>
		<tr>
			<th>Time</th>
            <th>Type</th>
            <th>Condition</th>
            <th>StackTrace</th>
            <th>OperatingSystem</th>
            <th>Device</th>
            <th>DeviceIP</th>
            <th>DeviceModel</th>
        </tr>
        {{range .Content}}
        <tr>
			<td>{{.Time}}</td>
            <td>{{.Type}}</td>
            <td>{{.Condition}}</td>
            <td>{{.StackTrace}}</td>
            <td>{{.OperatingSystem}}</td>
            <td>{{.Device}}</td>
            <td>{{.DeviceIP}}</td>
            <td>{{.DeviceModel}}</td>
		</tr>
        {{end}}
    </table>
</body>
</html>