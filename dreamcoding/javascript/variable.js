'use strict';

// let globalName = 'global_name';
// let test = "kk";
// {
//     let name = 'debbie';
//     console.log(name); 
//     name = 'kylie';
//     console.log(name);
// }
// console.log(test);
// console.log(name);
// console.log(globalName);

// let myname = 'debbie';
// console.log(myname);

// let test = 'kylie';
// console.log(test);


// 3. Constant, r(read only)
// use const whenever possible
// only use lte if variavle needs to change
const daysInWeek = 7;
const maxNumber = 5;

// Note ! 
// Immutable data type: primitive types, frozen objects (i.e. object.freeze())
// Mutable data type : all objects by default are mutable in JS
// favor immutable data type always for a few reasons
// - security
// - thread safety
// - reduce human mistakes 



// 4. Variable type
// primitive, single item: number, string, boolean, null, undefined, symbol
// object, box container
// function, first class function 
const count = 17;
const size = 17.1;
console.log(`value: ${count}, type: ${typeof count}`);
console.log(`value: ${size}, type: ${typeof size}`);


// number - special numeric values: infinity, -infinity, NaN
const infinity = 1 / 0;
const negativeInfinity = -1 / 0;
const nAn = 'not a number' / 2
console.log(infinity);
console.log(negativeInfinity);
console.log(nAn);

// bigInt (fairly new, don't use it yet)
const bigInt = 123456789012345678901234567890n;
console.log(`value: ${bigInt}, type: ${typeof bigInt}`);
Number.MAX_SAFE_INTEGER;

// string
const char = 'c';
const brendan = 'brendan';
const greeting = 'hello ' + brendan;
console.log(`value: ${greeting}, type: ${typeof greeting}`);
const helloBob = `hi ${brendan}`; 
console.log(`value: ${helloBob}, type: ${typeof helloBob}`);

// boolean
// false : 0, null, undefinde, NaN, ''
// true : any other value
const canRead = true;
const test = 3 < 1 ;   // false 
console.log(`value: ${canRead}, type: ${typeof canRead}`);
console.log(`value: ${test}, type: ${typeof test}`)

// null
let nothing = null;
console.log(`value: ${nothing}, type: ${typeof nothing}`);

// undefined 
let x;
console.log(`value: ${x}, type: ${typeof x}`);

// symbol, create unique identifiers for objects
const symbol1 = Symbol('id');
const symbol2 = Symbol('id');
console.log(symbol1 === symbol2);

const gSymbol1 = Symbol.for('id');
const gSymbol2 = Symbol.for('id');
console.log(gSymbol1 === gSymbol2); 
console.log(`value: ${symbol1.description}, type: ${typeof symbol1.description}`)

// object, real-life object, data structure 
const ellie = { name: 'ellie', age: 20};
ellie.name = 21;

// dynamic typing: dynamoically typed laguage
let text = 'hello';
console.log(text.charAt(0));
console.log(`value: ${text}, type: ${typeof text}`);
text = 1;
console.log(`value: ${text}, type: ${typeof text}`);
text = '7'+5;
console.log(`value: ${text}, type: ${typeof text}`);
text = '8' / '2';
console.log(`value: ${text}, type: ${typeof text}`);
console.log(text.charAt(0));

