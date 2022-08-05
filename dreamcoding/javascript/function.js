// Function 
// - fundamental building block in the program
// - subprogram can be used multiple times
// - performs a task or calculates a value 

// 1. Funtion declaration
// function name(param1, param2) { body... return; }
// one function == one thing
// naming: doSomething, command, verb
// e.g. createCardAndPoint -> createCard, createPoint
// function is object in JS 

function printHello() {
    console.log('Hello');
}
printHello();

function log(massage) {
    console.log(massage)
};

log("Hello@");
log(1234);

// 2. Parameters 
// premitive parameters: passed by value
// object parameters : passed by reference
function changeName(obj) {
    obj.name = 'coder';
}
const kylie = { name: 'ellie'};
changeName(kylie);
console.log(kylie);

// 3. Dafault parameters 
function showMessage(message, from = 'unkwnow') {
    console.log(`${message} by ${from}`)
}
showMessage("Hi")

// 4. Rest parameters 
function printAll(...args) {
    for (let i = 0; i < args.length; i++) {
        console.log(args.length)
        console.log(args[i])
    }
}
printAll('dream', 'coding', 'kylie');

function printAllOther(...args) {
    for (const arg of args) {
        console.log(arg);
    }
}
printAllOther('print', 'all', 'other');


function printAllAraay(...args) {
    args.forEach((arg) => console.log(arg))
}
printAllAraay('print', 'all', 'array');


// 5. Local scope
let globalMessage = 'global';   // global variable
// 밖에서는 안이 보이지 않고, 안에서만 밖을 볼 수 있다
function printMessage() {
    let message = 'hello';
    console.log(message);
    console.log(globalMessage);
    function printAnother() {
        console.log(message);
        let childMessage = 'child';
        console.log(childMessage);
    }
    printAnother()
    // return undefined
}
printMessage()


// 6. Return a value
function sum(a, b) {
    return a + b;
}
const result = sum(1, 2);
console.log(`sum: ${sum(1, 2)}`);
console.log(`result: ${result}`);

// 7. Early return, early exit
// bad
function upgradeUser(user) {
    if (user.point > 10) {
        //loug upgrade logic...
    }
}

// good
function upgradeUser(uer) {
    if (user.point <= 10) {
        return;
    }
    // long upgrade logig
}

// First-class function
// functions are treated like any other variable
// can be assigned as a value to variable
// can be passed as an argument to other functions
// can be returned by other function

// 1. Function expression
// a function declaration can be called earlier than it is defined (hoisted)
// a function expression is created when the the execution reaches it 
const print = function() { // anonymous function // function 뒤에 name이 있는 경우는 named function 
    console.log('this is anonymous function print')
};
print();
const printAgain = print;
printAgain();
const sumAgain = sum;
console.log(sumAgain(1, 3));

// 2. Callback function using function expression
// 상황에 따라 전달된 함수를 부르는것을 call back이라고 한다 
function randomQuiz(answer, printYes, printNo) {
    if (answer === 'love you') {
        printYes();
    }   else {
        printNo();
    }
}
// anonymous function
const printYes = function () {
    console.log('yes !');
};

// named function
// better debugging in debugger's stack traces
// recursions
const printNo = function print() {
    console.log("no !");
};

randomQuiz('wrong', printYes, printNo);
randomQuiz('love you', printYes, printNo);

// Arrow function
// always anonymous
const simplePrint = function () {
    console.log('simplePrint!')
};
const simplePrintArroy = () => console.log('simpplePrint!');
const add = (a, b) => a+b;
const addFcuntion = function(a, b) {
    return a+b
};

// IIFE : Immediately Invoked Function Expression
(function hello() {
    console.log("IIFE")
})();
// hello();
