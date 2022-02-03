const http = require('http');
const ejs = require('ejs');

const name = 'dongwook'
const courses = [{name:'HTML'},{name:'CSS'},{name:'Javascript'}]

const server = http.createServer((req,res)=>{   
    const url = req.url; //what?
    const method = req.method; //how? action?
    if (url === '/courses') {
        if (method === 'GET') {
            res.writeHead(200,{'Content-Type': 'application/json'});
            res.end(JSON.stringify(courses))
        } else if (method ==='POST'){
            const body = [];
            req.on('data',chunk=>{
                console.log(chunk);
                body.push(chunk);
            });

            req.on('end',()=>{
                const bodyStr = Buffer.concat(body).toString()
                console.log("bodyStr",bodyStr)
                const course = JSON.parse(bodyStr)
                courses.push(course);
                console.log(course);
                res.writeHead(201);
                res.end();
            })
        }
    }
//     res.setHeader('Content-Type','text/html')
//     if (url  ==='/') {
//         ejs.renderFile('./template/index.ejs',{name})
//         .then(data=>res.end(data))
//     } else if(url ==='/courses'){
//         ejs.renderFile('./template/courses.ejs',{courses})
//         .then(data=>res.end(data))
//     } else {
//         ejs.renderFile('./template/not-found.ejs',{name})
//         .then(data=>res.end(data))
//     }
});

server.listen(8080);