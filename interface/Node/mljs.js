//v0.1.0 (err)
var http = require('http'); 

   
var qs = require('querystring'); 
   

var options = { 
    hostname: '127.0.0.1', 
    port: 8083, 
    path: '/get/py-node/', 
    method: 'GET' 
}; 
// var req = http.request(options, function (res) { 
//     console.log('STATUS: ' + res.statusCode); 
//     console.log('HEADERS: ' + JSON.stringify(res.headers)); 
//     res.setEncoding('utf8'); 
//     res.on('data', function (chunk) { 
//         console.log('BODY: ' + chunk); 
//     }); 
// }); 
// req.on('error', function (e) { 
//     console.log('problem with request: ' + e.message); 
// }); 
// req.end();


class PipeManager{
    constructor(HP){
        this.host = HP[0],
        this.port = HP[1],
        console.log(this.port)

    }
    check(){
        var options = {
            hostname:this.host,
            port:this.port,
            path:'/check',
            method:'GET'
        }
        console.log(options)
        var req = http.request(options, function(res){
            console.log('STATUS: ' + res.statusCode); 
            console.log('HEADERS: ' + JSON.stringify(res.headers)); 
            res.setEncoding('utf8'); 
            res.on('data', function (chunk) { 
                console.log('BODY: ' + chunk); 
            }); 
        })
        req.on('error', function (e) { 
            console.log('problem with request: ' + e.message); 
        }); 
        req.end();
    }
}

var pipeManager = new PipeManager(["http://127.0.0.1", 8083])
pipeManager.check()

// var startTime = Date.now()
// var myfunction =function(startTime){
//    var timeNow = Date.now()
//    console.log('当前时间'+timeNow+'|| 时差：'+(timeNow-startTime));
// }
// var myInterval=setInterval(myfunction,5000,startTime);