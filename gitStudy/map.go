var config = require("../config");
var path = require("path");
//var AgentKeepalive = require('agentkeepalive');
var bill = require("../rpc/bill")
var pfs = require("../util/pfs");
var log = require("../util/log");
var ossHelper = require("../fileHelper/oss")
var taskList = require("./upime_task_list")
var UploadTaskList = taskList.UploadTaskList;
var SaveRecordTaskList = taskList.SaveRecordTaskList
var uploadTaskList = new UploadTaskList();
var saveRecordTaskList = new SaveRecordTaskList();
var oss = config.oss;
var ROOM_STATUS = require("../const/RoomStatus");
var __base = oss.root + "liveclass/";
function sleep(mil) {
    return new Promise(function (a, r) {
        setTimeout(a, mil);
    });
}

var eventArray = [];
// console.log("eventArray");
// eventArray.push('default');
async function sync() {
    while (eventArray[0]) {
        await uplodData(eventArray[0]);
        runAsync(eventArray[0]);
        eventArray.shift();
    }
};

function enterData(data) {
    let length = eventArray.length;
    eventArray.push(data);
    if(length == 0){
        sync();
    }
}

async function uplodData(data) {
    function getOption(fileName) {
        if (fileName == "info.plist") {
            return { "mime": "application/json", headers: { "Content-Encoding": "gzip" } };
        }
        var noPublicFiles = [/\.pptx?$/];
        var index = noPublicFiles.findIndex((a) => {
            return a.test(fileName);
        })
        if (index < 0)
            return { headers: { 'x-oss-object-acl': 'public-read' } };
        return null;
    }
    async function doUpload(lessonName) {

        var remoteDir = __base + data.appId + (data.subId ? "/" + data.subId : "") + "/" + lessonName;
        var begin = Date.now();
        log.debug(`begin upload deal ${lessonName} to ${remoteDir} ${JSON.stringify(data)}`)
        var genPicutre = require("./picture/genPictures")
        await genPicutre(data.pages, data.dirname, data.meetingId);
        log.debug(`banshu ${data.meetingId} of ${data.appId}:${data.subId} generated`);
        await sleep(3000);
        let failedList = [];
        data.fileSize = await ossHelper.uploadDir(data.dirname, remoteDir, getOption, failedList);
        data.fileLocation = `/${oss.bucket}/${remoteDir}`;
        if (data.cover) {
            if (!(data.cover.startsWith("http://") || data.cover.startsWith("https://"))) {
                var stat = await pfs.stat(data.dirname + "/" + data.cover)
                data.coverSize = stat.size;
                data.cover = `${data.fileLocation}/${data.cover}`
            } else {
                // ts那边需要cover oss相对路径
                data.cover = ossHelper.toRelativePath(data.cover);
                data.coverSize = 0;
            }
        }
        data.duration = Math.round(data.duration / 1000);
        delete data.pages;


        if (failedList.length == 0) {
            await pfs.deleteDir(data.dirname);
            process.send({ meetingId: lessonName, status: ROOM_STATUS.UPLOADED });
        } else {
            failedList.forEach(function (value) {
                uploadTaskList.addToTask(value);
            });
        }
        await ossHelper.deleteDir(config.remoteTempDir + lessonName);//this is the temp ppt dir;
        try {
            var res = await bill.saveRecord(data);

            //todo:如果首次失败了就触发不到此方法了
            process.send({ meetingId: lessonName, status: ROOM_STATUS.END });
            log.info(`end upload deal ${lessonName} using ${Date.now() - begin} of ${data.appId}:${data.subId} `);
        } catch (error) {
            saveRecordTaskList.addToTask(data);
            log.error(`saveRecord failed and push to queue ${JSON.stringify(data)},${error.stack}`);
        }
    }
    console.log(data, '串行任务 begin...')
    await sleep(1000);
    console.log(data, '串行任务 end...')
    // var lessonName=path.basename(data.dirname)
    // try{
    //     log.info("receive upload msg, begin update room status,lesson_name:" + lessonName);
    //     //for test
    //     // process.send({meetingId:lessonName ,status: ROOM_STATUS.UPLOADED});
    //     await doUpload(lessonName);
    // }catch(err){
    //     log.error(`uploadData failed ${data.meetingId} of ${data.appId}:${data.subId} ${lessonName}`);
    //     log.error(err.stack);
    // }
}
async function runAsync(data){
    // return new Promise(function (resolve, reject) {
        console.log(data, '并行任务 begin...')
        await  sleep(5000);
        console.log(data, '并行任务 end...')
}



module.exports = enterData;
