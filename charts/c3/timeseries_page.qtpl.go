// This file is automatically generated by qtc from "timeseries_page.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line timeseries_page.qtpl:1
package c3

//line timeseries_page.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line timeseries_page.qtpl:1
import "github.com/grokify/gotilla/time/timeutil"

//line timeseries_page.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line timeseries_page.qtpl:2
func StreamTimeseriesPage(qw422016 *qt422016.Writer, pageData TimeseriesPageData) {
	//line timeseries_page.qtpl:2
	qw422016.N().S(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/d3/3.5.6/d3.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/c3/0.4.15/c3.min.js"></script>
    <link type="text/css" rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/c3/0.4.15/c3.min.css">

	<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
    <link type="text/css" rel="stylesheet" href="https://cdn.datatables.net/1.10.16/css/jquery.dataTables.min.css">
    <script src="https://cdn.datatables.net/1.10.16/js/jquery.dataTables.min.js"></script>
    <style>
    p,h1,h2,th,td{
      font-family: Arial, Helvetica, sans-serif;
    }
    </style>
</head>
<body>

<h1>`)
	//line timeseries_page.qtpl:22
	qw422016.E().S(pageData.Title)
	//line timeseries_page.qtpl:22
	qw422016.N().S(`</h1>

<p>URL: <a href="`)
	//line timeseries_page.qtpl:24
	qw422016.N().S(pageData.URL)
	//line timeseries_page.qtpl:24
	qw422016.N().S(`">`)
	//line timeseries_page.qtpl:24
	qw422016.E().S(pageData.URL)
	//line timeseries_page.qtpl:24
	qw422016.N().S(`</a></p>

`)
	//line timeseries_page.qtpl:26
	qw422016.N().S(TimeseriesHTML(pageData.Charts[0]))
	//line timeseries_page.qtpl:26
	qw422016.N().S(`

`)
	//line timeseries_page.qtpl:28
	qw422016.N().S(TimeseriesHTML(pageData.Charts[1]))
	//line timeseries_page.qtpl:28
	qw422016.N().S(`

<h2>Growth</h2>

<table id="growthTable">
	<thead>
		<th>Quarter</th>
		<th>Total</th>
		<th>Growth</th>
		<th>YoY</th>
		<th>QoQ</th>
	</thead>
	<tbody>
	`)
	//line timeseries_page.qtpl:41
	for _, point := range pageData.XoxPoints {
		//line timeseries_page.qtpl:41
		qw422016.N().S(`
        <tr>
        	<td>`)
		//line timeseries_page.qtpl:43
		qw422016.E().S(point.Time.Format(timeutil.ISO8601YM))
		//line timeseries_page.qtpl:43
		qw422016.N().S(`</td>
        	<td>`)
		//line timeseries_page.qtpl:44
		qw422016.N().D(int(point.AggregateValue))
		//line timeseries_page.qtpl:44
		qw422016.N().S(`</td>
        	<td>`)
		//line timeseries_page.qtpl:45
		qw422016.N().D(int(point.Value))
		//line timeseries_page.qtpl:45
		qw422016.N().S(`</td>
        	<td>`)
		//line timeseries_page.qtpl:46
		qw422016.N().FPrec(point.YoYAggregate*100.0, 0)
		//line timeseries_page.qtpl:46
		qw422016.N().S(`%</td>
        	<td>`)
		//line timeseries_page.qtpl:47
		qw422016.N().FPrec(point.QoQAggregate*100.0, 0)
		//line timeseries_page.qtpl:47
		qw422016.N().S(`%</td>
        </tr>
    `)
		//line timeseries_page.qtpl:49
	}
	//line timeseries_page.qtpl:49
	qw422016.N().S(`

	</tbody>
</table>

<script>
$(document).ready(function(){
    $('#growthTable').DataTable({
    "lengthMenu": [ [-1], ["All"] ]
    });
});
</script>

</body></html>

`)
//line timeseries_page.qtpl:64
}

//line timeseries_page.qtpl:64
func WriteTimeseriesPage(qq422016 qtio422016.Writer, pageData TimeseriesPageData) {
	//line timeseries_page.qtpl:64
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line timeseries_page.qtpl:64
	StreamTimeseriesPage(qw422016, pageData)
	//line timeseries_page.qtpl:64
	qt422016.ReleaseWriter(qw422016)
//line timeseries_page.qtpl:64
}

//line timeseries_page.qtpl:64
func TimeseriesPage(pageData TimeseriesPageData) string {
	//line timeseries_page.qtpl:64
	qb422016 := qt422016.AcquireByteBuffer()
	//line timeseries_page.qtpl:64
	WriteTimeseriesPage(qb422016, pageData)
	//line timeseries_page.qtpl:64
	qs422016 := string(qb422016.B)
	//line timeseries_page.qtpl:64
	qt422016.ReleaseByteBuffer(qb422016)
	//line timeseries_page.qtpl:64
	return qs422016
//line timeseries_page.qtpl:64
}
