package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	// "path/filepath"
	// "unsafe"
	"io/ioutil"
	"time"
	//修改环境变量和新增src目录结构
	"os"
	// "picture"
	// "oss"
	// "room"
	// "task"
)

//改为常量
var base = getConfigValue("oss", "root") + "liveclass/"

func getConfigValue(v1 string, v2 string) string {
	return ""
}

type Statics struct {
	meetingId         uint32
	meetingType       string
	appId             string
	callback          string
	subId             string
	mediaType         string
	duration          int64
	userCount         string
	listenerCount     string
	speakerCount      string
	userCS            string
	teacherCS         string
	speakerShareTime  string
	listenerShareTime string
	beginTime         string
	endTime           string
	recordBegin       string
	dirname           string
	recordAvator      string
	videoStream       string
	richMedia         string
	cover             string
	statistics        string
	refFiles          string
	pages             []pageInfo
	attachment        string
	test              string
	recordId          string
	saved             string
	topic             string
	hasMedia          string
	fileSize 		  int
	fileLocation      string
}

func main() {
	fmt.Println("main func...")
	//处理入参
	param := Statics{meetingId: 11, dirname: "/目录1/目录2/文件"}
	deal(param)
}

func deal(data Statics) {
	//获取文件名  returns the last portion of a path
	// lessonName := filepath.Base(data.dirname)

	// begin := Date.now()
	//  golang中没有三目运算符，只能用ifelse
	// var remoteDir string
	// if len(data.subId) == 0 {
		// remoteDir := base + data.appId + "/" + lessonName
	// } else {
		// remoteDir := base + data.appId + "/" + data.subId + "/" + lessonName
	// }
	// fmt.Println(lessonName)
	//
	fd, err := os.Open("C:/Users/yangyangwang/Desktop/testfile/dev_info2.plist")
	if err != nil {
		fmt.Println("os open plist error",err)
		return 
	}

	rawContent, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println("ioutil read file error",err)
		return 
	}

	info := &CmdInfos{}
	var pages = json.Unmarshal(rawContent, info)
	genPicture(data.pages, data.dirname, data.meetingId)
	// ossDeal(data,remoteDir,lessonName)
	// saveRecordTask(data,lessonName,begin)
}



type Info struct {
	infotype   string
	cmds        []interface{}
	devicetype string
	version    float32
}

type pageInfo struct {
	imgs interface{}
}
func saveRecordTask() {
	// try{
		// var res = await bill.saveRecord(data);
		// todo:如果首次失败了就触发不到此方法了
		// process.send({meetingId:lessonName ,status: ROOM_STATUS.END,appId:data.appId,subId:data.subId});
		// log.info(`end upload deal ${lessonName} using ${Date.now()-begin} of ${data.appId}:${data.subId} `);
	// }catch(error){
		// saveRecordTaskList.addToTask(data);
		// log.error(`saveRecord failed and push to queue ${JSON.stringify(data)},${error.stack}`);
	// }
}
//TODO 待确定
var imageBuffer = map[string]string{}

// [

// [[0,0,0],34,1,1,1,"",[0.0,0.0],[968.0,726.0]],
// [[1,0,0],36,1],
// [[1,1,0],6,0.0],
// [[1,1,0],22,1,"paper_bg_blackboard.png",1,"border_bg_blackboard.png","#2e3038"],
// [[1,1,1059],8,1,1,[205.73,200.01],2,[82,146,191]],
// [[1,1,1088],8,1,1,[205.73,206.49]],
// [[1,1,1091],8,1,1,[205.73,215.84]],

//生成图片，并保存到指定文件夹中
func genPicture(pages []pageInfo, dirname string, meetingId uint32) []string {
	imagedir := dirname + "/banshu"
	//TODO node js mkdir使用了bind 还没理解，待思考
	os.MkdirAll(imagedir, os.ModePerm)
	// await pfs.mkdir(imagedir/*,{ recursive: true } /*node8 is not supported*/);
	names := []string{}

	// info := Info{
	// 	infotype:   "comment",
	// 	devicetype: "upimeserver",
	// 	version:    1.0,
	// }

	// for i := 0; i < len(pages); i++ {
		p := picture.NewPicture(pages[0], dirname, i, imageBuffer, meetingId)
		// try {
		//TODO 考虑判断异常的事情 保持持串行
		// TODO 返回值代表什么？
		// imageFile := p.GetPicture(imagedir)
		// //TODO 表示空
		// pages[i].imgs = 0
		// //结构体判断空
		// if unsafe.Sizeof(imageFile) != 0 {
		// 	imageFileName := imageFile.Name
		// 	//FIXME 无法追加到切片
		// 	names = append(names, imageFileName)
		// 	var cmdTemp = []interface{};
		// 	cmdTemp[0] = {0, 215, i, 0, "attach" + i, imageFile.size, imageFile.name}
		// 	info.cmds = append(info.cmds, cmdTemp)
		// }

	// }
	// const util = require('util');
	// gzipAndSave(imagedir,info)
	return names
}

func gzipAndSave(imageDir string, info Info) {
	// const gzip =util.promisify(require("zlib").gzip);
	// var ziped_info=await gzip(JSON.stringify(info));
	// await saveFile(imagedir+"/info.plist",ziped_info);

	fw, err := os.Create(imageDir + "/info.plist") // 创建gzip包文件，返回*io.Writer
	if err != nil {
		log.Fatalln(err)
		//TODO 直接中止
	}
	defer fw.Close()
	//返回byte[]
	infos, errs := json.Marshal(info)
	if errs != nil {
		log.Fatalln(errs)
		//TODO 直接中止
	}
	// 实例化心得gzip.Writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 写入数据到zip包
	_, err = gw.Write(infos)
	if err != nil {
		log.Fatalln(err)
		//TODO 直接中止
	}
}

func ossDeal(data Statics,remoteDir string,lessonName string) {
	time.Sleep(time.Duration(3)*time.Second)
	failedList := []interface{};
	data.fileSize = oss.UploadDir(data.dirname,remoteDir,failedList);
	data.fileLocation=`/${oss.bucket}/${remoteDir}`;
	if(len(data.cover) != 0){
		if !(strings.HasPrefix("data.cover", "http://") || strings.HasPrefix("data.cover", "https://")) {
			//FIXME stat待实现
			// var stat = await pfs.stat(data.dirname + "/" + data.cover)
			data.coverSize = stat.size;
			data.cover = data.fileLocation+"/"+data.cover
		} else {
			// ts那边需要cover oss相对路径
			data.cover = oss.ToRelativePath(data.cover);
			data.coverSize = 0;
		}
	}
	data.duration = Math.round( );
	//实现四舍五入
	data.duration = int(math.Floor(data.duration / 1000 + 0.5))
	//TODO FIXME将变量的属性删除
	data.pages = []pageInfo;

	if len(failedList.length)==0 {
		// 删除目录
		os.Remove(data.dirname);
		tempData := map[string]interface{}{
			meetingId:lessonName,
			status: ROOM_STATUS.UPLOADED,
			appId:data.appId,
			subId:data.subId,
		}
		room.Deal(tempData)
	}else{
		for _,value := range failedList {
			task.AddToTask(value);
		}
	}
	oss.DeleteDir(config.remoteTempDir+lessonName);//this is the temp ppt dir;s
}

const INIT = 0 //初始
const CREATED = 10  //已创建
const CONTLLER_ENTER = 20 //控制者，一般是老师
const RECORDING = 30 //正在录制
const RECORDING_END = 40//录制结束
const UPLOADED = 50//已上传
const END = 60     //已生成账单
const END_NO_SAVED = 61//未生成账单结束