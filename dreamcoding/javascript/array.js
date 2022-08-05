// Array 

// 1. Declaratopm
const arr1 = new Array();
const arr2 = [1, 2];

// 2. Indext position
const fruits = ['🍎', '🍌'];
console.log(fruits);
console.log(fruits.length);
console.log(fruits[0]);
console.log(fruits[1]);

console.log(fruits[fruits.length - 1 ]);


// 3. Looping over an array
// print all fruits
// a. for
console.clear()
for (let i = 0; i < fruits.length; i++) {
    console.log(fruits[i]);
}

// b. for of
for(let fruit of fruits) {
    console.log(fruit)
}

// c. for each
console.clear()
// fruits.forEach( function (fruit, index, array) {
//     console.log(fruit, index);
// });

fruits.forEach((fruit) => console.log(fruit))

// 4. Addition, delete, copy 
// push: add an item to the end
fruits.push('🍓', '🍑');
console.log(fruits)
// pop: remove an item from the end
fruits.pop();
fruits.pop();
console.log(fruits);

// unshift : add an item to the beginning;
fruits.unshift('🍓', '🍋')
console.log(fruits)
// shift: remove an item from the beginning; 
fruits.shift()
fruits.shift()
console.log(fruits)

// note! shift, unshift are slower than pop, push
// splice : remove an item by index position
console.clear()
fruits.push('🍓', '🍑', '🍋')
console.log(fruits)
fruits.splice(1, 1);
console.log(fruits)
fruits.splice(1, 1, '🍏', '🍉')
console.log(fruits)

// two conbine two arrays
const fruits2 = ['🍐', '🥥'];
const newFruits = fruits.concat(fruits2);
console.log(newFruits)

// 5, Searching 
// find the index
console.clear()
console.log(fruits)
console.log(fruits.indexOf('🍎'))
console.log(fruits.indexOf('🍉'))
console.log(fruits.includes('🥥'))
console.log(fruits.indexOf('🥥'))

// lastIndexOf
console.clear()
fruits.push('🍎')
console.log(fruits);
console.log(fruits.indexOf('🍎'))
console.log(fruits.lastIndexOf('🍎'))
