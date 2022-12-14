// Promise is a JavaScript object for asynchronous operation 
// state: pending -> fullfilled or rejected 
// Producer vs Consumer 

// 1. Producer
// when new Promise is created, the executor runs automatically 
const promise = new Promise((resolve, reject) => {
    // doing some heavy work (network, readfiles)
    console.log('doing something...')
    setTimeout(()=> {
        resolve('ellie');
        reject(new Error('no network'));
    }, 2000)
})


// 2. Consumers: then, chatch, finally 
promise
    .then((value)=> {
        console.log(value);
    })
    .catch(error => {
        console.log(error);
    })
    .finally(()=> {
        console.log('finally')
    })

// 3. Promise chaining
const fetchNumber = new Promise((resolve, reject) => {
    setTimeout(()=> resolve(1), 1000);

})


fetchNumber
.then(num => num * 2)
.then(num => num * 3)
.then(num => {
    return new Promise((resolve, reject) => {
        setTimeout(()=> resolve(num - 1), 1000);
    })
})
.then(num => console.log(num))



// 4, Error Handiling
const getHen = () => 
    new Promise((resolve, reject) => {
        setTimeout(()=> resolve('π'),  1000)
    });

const getEgg = hen =>
    new Promise((resolve, reject) => {
        // setTimeout(()=> resolve(`${hen} => π₯`), 1000)
        setTimeout(() => reject(new Error(`error! ${hen} => π₯`), 1000))
    })

const cook = egg => 
    new Promise((resolve, reject) => {
        setTimeout(()=> resolve(`${egg} => π³`), 1000)
    })


// getHen()
// .then(hen => getEgg(hen))
// .then(egg => cook(egg))
// .then(meal => console.log(meal))

// κΉλνκ² μΆλ ₯νκΈ°

getHen() 
    .then(getEgg)
    .catch(error => {
        return 'π₯'
    })
    .then(cook)
    .then(console.log)
    .catch(console.log)



