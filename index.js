(function() {
    // var promise = new Promise(function (resolve,reject){
    //     setTimeout(()=>{
    //        resolve();},200
    //     )
    // })

    var promise = new Promise(function(resolve,reject){
        setTimeout(()=>{reject(new Error('error occurs')),200})
    })

    console.log(promise);

    setTimeout(()=>{
        console.log(promise)},1000
    )

})();