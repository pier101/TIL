function solution(num) {
    while (
        num.toString().length % 2 !== 0 ||
        num.toString().split("").slice(0, num.toString().length / 2).reduce((prev, cur) => prev * parseInt(cur), 1) !==
        num.toString().split("").slice(num.toString().length / 2).reduce((prev, cur) => prev * parseInt(cur), 1)
    ) {
        num++;
    }
    return num;
}

console.log(solution(21));
console.log(solution(103));
console.log(solution(235386));