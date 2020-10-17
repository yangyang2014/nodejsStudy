const Koa = require('koa');
const app = new Koa();
var fs = require("fs")



// app.use((ctx,next) => {
//     const start = Date.now();
//     return next().then(()=>{
//         const ms = Date.now() - start;
//         console.log(`${ctx.method}  ${ctx.url} - ${ms}ms`);
//     });
// });

app.use(async (ctx, next) => {
    const start = Date.now();
    await next();
    const ms = Date.now() - start;
    console.log(`${ctx.method} ${ctx.url} - ${ms}ms`)
});
 // response
app.use(ctx => {
    ctx.body = 'Hello Koa';
  });
  
// app.use(async (ctx, next) => {await next();})

// app.use(async (ctx, next) => {
//     ctx.assert(ctx.request.accepts('xml'),406);
//     await next();
// })

// app.use(async (ctx, next) => {
//     await next();
//     ctx.response.type = 'xml';
//     ctx.response.body = fs.createReadStream('really_large.xml')
// });

app.listen(3000);