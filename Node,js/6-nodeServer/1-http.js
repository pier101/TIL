const http = require('http');
// const http2 = require('http'); // https

const server = http.createServer((req,res)=>{
    console.log('imcoming...');
    // console.log(req.headers);
    // console.log(req.httpVersion);
    // console.log(req.method);
    // console.log(req.url);
    const url = req.url;

    if (url  ==='/') {
        res.setHeader('Content-Type','text/html');
        res.write()
        res.write('welcome!');
    } else if(url ==='/courses'){
        res.write('Courses')
    } else {
        res.write('Not found')
    }

    res.end();
});

server.listen(8080);