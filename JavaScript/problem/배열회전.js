const solution = (matrix, k) => {
    // if (matrix.length === 0) return [];

    let result = [];
    for (let i = 0; i < matrix[0].length; i++) {
        result.push([]);
    }
    for (let j = 0; j < matrix[0].length; j++) {
        for (let i = 0; i < matrix.length; i++) {
            result[j].push(matrix[i][j]);
        }
        result[j].reverse();
    }

    if (k % 4 === 0) {
        return matrix;
    } else if (k === undefined || k % 4 === 1) {
        return result;
    } else if (k % 4 === 2) {
        for (let j = 0; j < matrix.length; j++) {
            matrix[j].reverse();
        }
        matrix.reverse();
        return matrix;
    } else {
        for (let j = 0; j < result.length; j++) {
            result[j].reverse();
        }
        result.reverse();
        return result;
    }
};

console.log(
    solution(
        [
            [1, 2],
            [3, 4],
        ],
        3
    )
);

// [[1,2], [3,4]]
//  j = 0
//     matrix[0][0] // 1
//     matrix[1][0] // 3
// j = 1
//     matrix[0][1] // 2
//     matrix[1][1] // 4
/*
[[1,2],[3,4]] ==> 0
[[3,1],[4,2]] ==> 1
[[4,3],[2,1]] ==> 2
[[2,4],[1,3]] ==> 3
*/
