// 监听事件的并行转串行
var uploadTaskArr = []

function enterUploadTaskArr(data) {
    var length = uploadTaskArr.length;
    uploadTaskArr.push(data);
    if(length == 0){
        dealUploadTask();
    }
}  

async function dealUploadTask() {
    while (uploadTaskArr[0]) {
        //串行执行
        await syncFunction();
        //并行执行
        asyncFunction();
        uploadTaskArr.shift();
    }
}

async function syncFunction(){

}

async function asyncFunction(){
    
}
module.exports=enterUploadTaskArr;
