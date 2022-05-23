/* 내 풀이 */
function solution(arr) {
    var answer = [];
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === arr[i - 1]) {
            continue;
        }
        answer.push(arr[i]);
    }
    return answer;
}

console.log(solution([4, 4, 4, 3, 3]));

// /* 찾아본거 */
// function _solution(arr) {
//     return arr.filter((val, index) => val != arr[index + 1]);
// }
function _solution(arr) {
    return Array.from(new Set(arr))
}

console.log(_solution(
    [5, 4, 4, 3, 5]));
