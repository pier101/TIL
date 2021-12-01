const Calculator = require('../calculator')

describe('Calculator', ()=>{
    let cal;
    beforeEach(()=>{
        cal = new Calculator();
    })
    it('inits with 0',()=>{
        //const cal = new Calculator();
        expect(cal.value).toBe(0);
    });
    
    it('sets',()=>{
        //const cal = new Calculator(); //코드 중복
        cal.set(9);
        expect(cal.value).toBe(9)
    })
    it('clear',()=>{
        //const cal = new Calculator(); //코드 중복
        cal.set(9);
        cal.clear();
        expect(cal.value).toBe(0)
    })
    it('adds',()=>{
        //const cal = new Calculator(); //코드 중복
        cal.set(1);
        cal.add(2);
        expect(cal.value).toBe(3)
    })
    it('multiply',()=>{
        //const cal = new Calculator(); //코드 중복
        cal.set(5);
        cal.multiply(4);
        expect(cal.value).toBe(20)
    })
    // ('divides',()=>{
    //     //const cal = new Calculator(); //코드 중복
    //     it('0/0 === NaN',()=>{
    //             cal.divide(0);
    //             expect(cal.value).toBe(NaN);
    //         });
    //     it('1/0 === Infinity',()=>{
    //             cal.set(1);
    //             cal.divide(0);
    //             expect(cal.value).toBe(Infinity);
    //         });w
    //     })
    //     it("4/2 === 2", () => {
    //         cal.set(4);
    //         cal.divide(2);
    //         expect(cal.value).toBe(2);
    //     });
    describe("divides", () => {
        it("0/0 === NaN", () => {
           cal.divide(0);
           expect(cal.value).toBe(NaN);
        });
        it("1/0 === Infinity", () => {
           cal.set(1);
           cal.divide(0);
           expect(cal.value).toBe(Infinity);
        });
        it("4/2 === 2", () => {
           cal.set(4);
           cal.divide(2);
           expect(cal.value).toBe(2);
        });
     });
  
})