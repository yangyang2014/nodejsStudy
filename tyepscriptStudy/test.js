// var aliossHelper = require("ali-oss");
var testStr = "1aa1";
console.log(testStr);
if (testStr !== "1aa") {
    console.log("is undefined");
}
else {
    console.log("1aaa ok");
}
var config = {
    oss: {
        root: "",
        accessKeyId: "LTAI4hbpgkC1mW8F",
        accessKeySecret: "cIw4ZF3OZzgSSycksITzN5dd5ELHLE",
        roleArn: "acs:ram::1845455964094256:role/oss-readonly-role",
        bucket: "file-plaso",
        region: "oss-cn-hangzhou",
        internal: true,
        bkpath: "http://hz-public-files.oss-cn-hangzhou-internal.aliyuncs.com/resource/upime/bgpaper/"
    },
    domain: "confnode3.plaso.cn"
};
var oss;
if (config.oss && config.oss.region) {
    oss = config.oss;
}
;
var policy;
var ossConfig = {
    accessKeyId: oss.accessKeyId,
    accessKeySecret: oss.accessKeySecret
};
// var STS=require("ali-oss").STS;
// var sts=new STS(ossConfig);
policy = {
    "Version": "1",
    "Statement": [{
            "Effect": "Allow",
            //TODO Action 数组需要putObject和getObject的权限
            "Action": ["oss:PutObject", "oss:GetObject"],
            "Resource": "acs:oss:*:*:" + oss.bucket + "/" + config.oss.root + "*"
        }]
};
var message = "Hello World";
console.log(message);
// var ossController={
function getStsInfo(params, req) {
    console.log(params, ",req=", req);
    // }
    //         //todo query如何使用
    //         // var queryparam = params.query
    //         //op=xxxx&signature=xxxx 
    //         //校验签名是否正确
    //         // op的值是uploadTempFile,返回ossToken
    //         //将浮点数变成整数
    //         // var now=(~~(Date.now()/1000));
    //         // if(now<expire){
    //         //     return validStsInfo;
    //         // }
    //         var expire:number = 3600
    //         // var token = await sts.assumeRole(config.oss.roleArn, policy, expire, config.domain);
    //         // var creds=token.credentials;
    //         // stsInfo.id=creds.AccessKeyId;
    //         // stsInfo.secret=creds.AccessKeySecret;
    //         // stsInfo.token=creds.SecurityToken;
    //         // //TODO 怎样获取path
    //         // // me.remoteHistoryDir=`${root}liveclass/${me.__appId}${me.subId ? "/" + me.subId : ""}/${basename}
    //         // stsInfo.path = 
    //         // stsInfo.bucket = 
    //         // stsInfo.region = 
    var tempFilePath = "temp";
    var stsInfo = {
        // "ossid": creds.AccessKeyId,
        // "osskey": creds.AccessKeySecret,
        // "osstoken": creds.SecurityToken,
        "ossid": "",
        "osskey": "",
        "osstoken": "",
        "path": tempFilePath,
        "bucket": oss.bucket,
        "region": oss.region
    };
    var stsInfoResp = {
        "code": 0,
        "obj": stsInfo
    };
    return stsInfoResp;
}
// }
var res = getStsInfo("11", "22");
console.log(res);
exports["default"] = [
    ["/getUploadToken", getStsInfo]
];
