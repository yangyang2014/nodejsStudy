// //使用循环判断时间的方法，cpu使用占比30%
// function sleep(second) {
//     //获取日期的毫秒数
//     var start = +new Date();
//     console.log(start);
//     while(true){
//         var now = +new Date();
//         if(now - start >= second*1000) {
//             return;
//         }
//     }
// }
// console.log('enter sleep ');
// var times = 1;
// //等待所有程序跑完后，才会执行setInterval的方法
// setInterval(()=>{console.log('sleep...',times++,'s')},1000)
// sleep(10);
// console.log('over sleep ');


// //引用包
// var sleep = require('sleep');
// sleep.sleep(5);
// console.log('sleep 5s');
// //使用addon
//promise

//使用promise async await实现sleep 
function sleep(deley) {
    return new Promise(function(resolve,reject){ 
        setTimeout(() => {resolve()}, deley);
    })
}


// async function test(){
//     await sleep(5000)
//     console.log('sleep 5s')
// }

// test();
// console.log('after sleep');
//
//实现方式二 then 如果存在多个sleep，那么存在多层嵌套的问题
sleep(1000).then(()=>{
    console.log('test');
    sleep(500).then(console.log('aaaa'));}
)


