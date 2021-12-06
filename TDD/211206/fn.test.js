const fn = require('./fn')


//---------------------------------------------------------
//matcher--toBe
test('1은 1이야',()=>{
    expect(1).toBe(1)
});

test('2더하기 3은 5야.',()=>{
    expect(fn.add(2,3)).toBe(5);
});
//실패 케이스(not추가)
test('2더하기 3은 5가 아니야!.',()=>{
    expect(fn.add(3,3)).not.toBe(5);
});


//---------------------------------------------------------
//matcher--toEqual
test('2더하기 3은 5야.',()=>{
    expect(fn.add(2,3)).toEqual(5);
});



//객체나 배열은 재귀적으로 돌면서 확인해줘야 되기 때문에
//toBe() 대신 toEqual() 사용
test('이름과 나이를 전달받아서 객체를 반환해줘',()=>{
    expect(fn.makeUser("wook",29)).toEqual({
        name: "wook",
        age: 29
    })
})
//보다 엄격하게 사용하려면 toEqual() 대신 toStrictEqual()
// test('이름과 나이를 전달받아서 객체를 반환해줘',()=>{
//     expect(fn.makeUser("wook",29)).toStrictEqual({
//         name: "wook",
//         age: 29
//     })
// })

//---------------------------------------------------------
//marchers--toBeNull
//        --toBeUndefined              
//        --toBeDefined              
test("null은 null입니다.",()=>{
    expect(null).toBeNull();
})

//---------------------------------------------------------
//marchers--toBeTruthy
//        --toBeFalsy              
test("0은 false 입니다.",()=>{
    expect(fn.add(1,-1)).toBeFalsy();
})

//matchers--toBeGreaterThan 크다
//        --toBeGreaterThanOrEqual 크거나 같다.              
//        --toBeLessThan 작다.              
//        --toBeLessThanOrEqual 작거나 같다.              