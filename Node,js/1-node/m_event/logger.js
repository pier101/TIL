const EventEmitter = require('events');

//emitter 객채 별로 적용되서 모든 곳 적용시 class화 해서 사용
// const emitter = new EventEmitter(); (x)
class Logger extends EventEmitter {
    log(callback){
        this.emit('log','started...');
        callback()
        this.emit('log','ended!')
    }
}

module.exports = {Logger}