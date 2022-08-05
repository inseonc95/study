// number, string, boolean, null, undefined
let number = 2;
let number2 = number;

number2 = 3

console.log(number)
console.log(number2)


// object 
let obj = {
    name : 'ellie',
    age : 28,
}

let obj2 = obj
obj2.name = 'kylie'

console.log(obj)
console.log(obj2)


class Counter {
    constructor(runEveryFiveTimes) {
        this.counter = 0 
        this.callback = runEveryFiveTimes
    }
    increase(runIf5times) {
        this.counter++;
        console.log(this.counter)
        if(this.counter % 5 === 0) {
            this.callback && this.callback(this.counter)
        }
    }
}


function printSomething(num) {
    console.log(`yo!! ${num}`)
}
function alertNum(num) {
    alert(`wow!! ${num}`)
}

const coolCounter = new Counter(alertNum);
coolCounter.increase()
coolCounter.increase()
coolCounter.increase()
coolCounter.increase()
coolCounter.increase()

const printCounter = new Counter(printSomething);
const alertCounter = new Counter(alertNum)
