//似乎存在共享变量的问题，还有边界情况未考虑
console.log(getMaxByOneSwap('0'));
//打印函数与java中有区别
console.log(getMaxByOneSwap('5221131413'));
console.log(getMaxByOneSwap('5544331'));
console.log(getMaxByOneSwap('550113'));
console.log(getMaxByOneSwap('55311'));
console.log(getMaxByOneSwap('554311')); 
console.log(getMaxByOneSwap('313488'));

function getMaxByOneSwap(num) {
    if(num >= 0 && num <= 10){
        return num;
    }
    //将字符串拆分为一个个字符的数组
    digitArr = num.split("");
    //如何倒序
    digitArr.sort(function (a, b) { return b - a; });
   
    for (i = 0; i < digitArr.length; i++) {
        if (num[i] != digitArr[i]) {
            //var的作用域提升
            var swapValue = digitArr[i];
            var swapPos = i;
            break;
        }
    }
    for (i = num.length - 1; i >= 0; i--) {
        if (num[i] == swapValue) {
            swapPos2 = i;
            break;
        }
    }
    // console.log(swapPos, swapPos2);
    var resArr = num.split("");
    var temp = resArr[swapPos];
    resArr[swapPos] = resArr[swapPos2];
    resArr[swapPos2] = temp;
    //单字符数组转为字符串 与split的反操作
//     return resArr.toString();
    return resArr.join("");
 }