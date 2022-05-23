// Object 프로퍼티 열거
var sports = {
    soccer : "축구",
    baseball:  "야구"
};

for (const item in sports) {
    console.log(item)
    // console.log(sports[item])
}