var array = ['11']
console.log('begin...');
//如果异步中死循环一直使用array数组，那么就无法在
(async function test(){
    while(true){
        var a = eventArray.shift();
        if(a){
            console.log(a);
        }
    }
})();
array.push('aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa');
console.log(array);
setTimeout(()=>{array.push('aaaaaaaaaaaaaaaaa')},2000);
setTimeout(()=>{console.log(array); array.shift(); console.log(array);array.push('a1'); console.log(array);},5000);


// 获取promise状态，使其休眠。

// while true方法如果不调用promise方法，那么会导致cpu占用比例特别高,
while (true) {
    console.log("test");
}