var res = getMinMinute(['00:01', '00:23', '10:34','01:30']);
var res2 = getMinMinute(["23:59","00:00"]);
var res3 = getMinMinute(["01:02","12:00","09:03"]);

console.log(res);
console.log(res2);
console.log(res3);
// getMinMinute('00:01');
function getMinMinute(timeArr) {
    timeArr.sort();//排序的规则是什么？
    var minMinutesTemp = getMinusBetween2Time(timeArr[0],timeArr[timeArr.length-1]) + 24*60;
    // console.log(minMinutesTemp);
    for (i = 1; i < timeArr.length; i++){
        minusMinutes = getMinusBetween2Time(timeArr[i],timeArr[i-1]);
        if(minusMinutes < minMinutesTemp){
            minMinutesTemp = minusMinutes;
        }
    }
    return minMinutesTemp;
}

function getMinusBetween2Time(time1,time2) {
    var h1 =  parseInt(time1.split(":")[0]);//需要转为数值型再进行计算，否则出现字符串拼接的结果
    var m1 = parseInt(time1.split(":")[1]);
    var h2 =  parseInt(time2.split(":")[0]);
    var m2 = parseInt(time2.split(":")[1]);
    return (h1-h2)*60 + m1 - m2;
}
