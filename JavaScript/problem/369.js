function solution(num) {
    let res = 0;
    for (let i = 0; i <= num; i++) {
        if (
            String(i).includes(3) ||
            String(i).includes(6) ||
            String(i).includes(9)
        ) {
            res++;
        }
    }
    return res;
}

console.log(solution(10))
console.log(solution(33))
console.log(solution(15))
console.log(solution(36))
