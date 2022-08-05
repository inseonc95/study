// 1. String concatenation 
console.log('my' + ' cat');
console.log('1' + 2);
console.log(`string literals: 1 + 2 = ${1 + 2}`);

// 2. Numeric operators
console.log(1 + 1); // add
console.log(1 - 1); // substract
console.log(1 / 1); // divide
console.log(1 * 1); // multiply
console.log(5 % 2); // remainder
console.log(2 ** 3); // exponentiation

// 3. Increment and decrement operators
let counter = 2;
const preIncrement = ++counter;
console.log(`preIncrement: ${preIncrement}, counter: ${counter}`);
// counter = counter + 1;
// preIncrement = counter;
const postIncrement = counter++;
console.log(`postIncrement: ${postIncrement}, counter: ${counter}`);
// postIncrement = counter;
// counter = counter + 1; 
const preDecrement = --counter;
console.log(`preDecrement: ${preDecrement}, counter: ${counter}`);
// counter = counter - 1
// preDecrement = counter 
const postDecrement = counter--;
console.log(`postDecrement: ${postDecrement}, counter: ${counter}`);
// postDecrement = counter 
// counter = counter - 1 

// 4. Assignment operator
let x = 3;
let y = 6;
x += y; // x = x+y
x -= y; // x = x-y
x *= y; // x = x*y
x /= y; // x = x/y

// 5. Comparison operators
console.log(10 < 6);    // less than
console.log(10 <= 6);    // less than or equal
console.log(10 > 6);    // greater than
console.log(10 >= 6);   // greater than or equal

// 6. Logical operators; || (or), && (and), ! (not)
const value1 = false;
const value2 = 4 > 2;
console.log(`value1: ${value1}, value2: ${value2}`);
function check() {
    for (let i=0; i<10; i++) {
        // wasting time
        console.log('😀')
    }
    return true
}
// || (or)
console.log(`or: ${value1 || value2 || check()}`);
// && (and)
console.log(`and: ${value1 && value2 && check()}`);
// often use to compress long if-statement 
// nullableObject && nullableObject.something
// null object가 null이 아닌 경우에만 null object의 something이라는 value를 가져온다
// if (nullableObject != null) {
//     nullableObject.something;
// }

// ! (not)
console.log(value1);
console.log(!value1);

// 7. Equality 
console.log(`Equality`)
const stringFive = '5';
const numberFive = 5;
// == loose equality, with type conversion
console.log(stringFive == numberFive);  // type을 변경해서 비교 
console.log(stringFive != numberFive);
// === strict equality, no type conversion
console.log(stringFive === numberFive);
console.log(stringFive !== numberFive); 

// object equality by reference 
console.log(`object equality`)
const ellie1 = { name: 'ellie'};
const ellie2 = { name: 'ellie'};
const ellie3 = ellie1;
console.log(ellie1 == ellie2);
console.log(ellie1 === ellie2);
console.log(ellie1 == ellie3);
console.log(ellie1 === ellie3); 

// eqaulity - puzzler
console.log(`test`)
console.log(0 == false); // t (false로 간주될 수 있다)
console.log(0 === false); // f (0은 boolean type이 아니기 때문에)
console.log('' == false); // t
console.log('' === false); // f (boolean type이 아니기 때문에)
console.log(null == undefined); // t (같은 것으로 간주)
console.log(null === undefined);    // f

// 8. Conditional operators: if
console.log(`Conditional operators`)
// if, else if, else
const Myname = 'kylie';
if (Myname === 'ellie') {
    console.log('Welcome, Ellie!!') 
} else if (Myname === 'coder') {
    console.log('You are amazing coder')
} else {
    console.log('nukwon')
}

// 9. Ternary operator: ? 
// (if를 더 간단하게 쓸 수 있는)
// condition ? value1 : value2
console.log(Myname == 'kylie' ? 'yes' : 'no')

// 10. Switch statement
// use for multiple if check
// use for enum-like value check
// use for multiple type checks in TS 
console.log(`Switch statement`)
browser = 'IE';
switch (browser) {
    case 'IE':
        console.log('go away!');
        break;
    case 'Chrome':
    case 'Firefox':
        console.log('love you');
        break;
    default:
        console.log('same all');
        break;
}

// 11. Loop
// while loop, while the condition is truthy
// body code is executed 
console.log(`loop`)
let i = 3;
while (i > 0) {
    console.log(`while: ${i}`);
    i--;
}

// do while loop, body code is executed first
// then check the condition
// block을 실행한 후에 조건이 맞는지 안맞는지 확인 


console.log(`do`)
// i = 3;
do {
    console.log(`do while: ${i}`);
    i--;
} while (i > 0);

// for loop, for(begin; condition; step)
for (i = 3; i > 0; i--) {
    console.log(`for: ${i}`);
}

for (let i = 3; i > 0; i = i-2) {
    // inline variable declaration
    console.log(`inline variable for: ${i}`);
}

// nested loop
for (let i = 0; i < 10; i++) {
    for (let j = 0; j < 10; j++) {
        console.log(`i: ${i}, j: ${j}`)
    }
}

// break, continue
console.log("quiz")
// Q1. iterate from 0 to 10 and print only even numbers
i = 0
n = 0
while (i < 5) {
    i++;
    if (i == 3) {
        continue
    }
    n += 1;
    console.log(i)
}


i = 0
while (i < 10) {
    
    i++;
    if (i % 2 ==1) {
        continue
    }
    else {
        console.log(i)
    }    
};


for (i = 0; i < 10; i++) {
    if (i % 2 ==1) {
        continue
    }
    else {
        console.log(i+2)
    }
}


// Q2. iterate from 0 to 10 and print nunbers until reaching 8 (use break)

for (i = 0; i < 10; i++) {
    if (i > 8) {
        break
    }
    else {
        console.log(i)
    }
}

i = 0
while (i < 10) {
    if (i > 8) {
        break
    }
    else {
        console.log(i)
    }
    i++;
}
