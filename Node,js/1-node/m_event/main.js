const {Logger} = require('./logger');
const emitter = new Logger();

emitter.on('log',(event)=>{
    console.log(event);
})

emitter.log(()=>{
    console.log('...doing something');
})