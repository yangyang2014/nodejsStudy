var cache = require('./cache').cache;
var s = cache.getCache('name1');
if (s) {
    console.log(s);
} else {
    cache.addCache('name1', 'content1', 20000);
}

//如果程序运行结束，进程会死亡。否则一直在执行，这里包括定时器和异步任务。

// var interval = setInterval(() => {
//     console.log(cache.getCache('name1'));
//     var i = 0;
//     if (i > 3) {
//         clearInterval(interval);
//     } else {
//         i++;
//     }
// }, 10000);

setTimeout(() => {
    console.log(cache.getCache('name1'));
}, 10000);