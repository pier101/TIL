function hello() {
    console.log(this)
    console.log(this === global);
}
hello();

class A {
    constructor(num){
        this.num = num;
    }
    member(){
        console.log("---class---")
        console.log(this)
    }
}
const a = new A(1);
a.member();

console.log("---global scope---")
console.log(this)

/*
    이처럼 this는 어디서 쓰이느냐에 따라 달리 쓰임
    this === module.exports
*/