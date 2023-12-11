const fs = require('fs');

const content = fs.readFileSync('input_day1', 'utf-8')
const lines = content.split('\n')
let digits = []
Array.from({length: lines.length}, (v, i) => {
    const line = lines[i]
    let digit = ""
    Array.from({length: line.length}, (v, j) => {
        if (!isNaN(parseInt(line[j]))) {
            digit += line[j]
        }
    })
    digits.push(parseInt(digit[0] + digit[digit.length-1]))
})
console.log(digits.reduce((acc, cur) => acc + cur, 0))