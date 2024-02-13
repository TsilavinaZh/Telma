const fs = require('fs');

function generateNumbers(debut, fin) {
    let numbers = [];
    for (let i = debut; i <= fin; i++) {
        numbers.push("0" + String(i).padStart(10, '0'));
    }
    return numbers;
}

const telma34 = generateNumbers(3400000000, 3499999999);
const telma38 = generateNumbers(3800000000, 3899999999);

const data = {
    "telma34": telma34,
    "telma38": telma38
};

const jsonData = JSON.stringify(data, null, 2);
fs.writeFileSync('Telma.json', jsonData);

console.log("success.");