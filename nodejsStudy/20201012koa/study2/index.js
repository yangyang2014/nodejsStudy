const fs = require('fs');
const game = require ('./game')
const koa = require('koa')
const mount = require ('koa-mount')

var playerWinCount = 0 
var lastPlayerAction = null
var sameCount = 0;

const app = new koa();

app.use(

    mount("/favicon.ico",function(ctx) {
        ctx.status =200
    })

)

const gameKoa = new koa()
app.use(
    mount('/game',gameKoa)
)
gameKoa.use(
    async function(ctx,next) {
        if(playerWinCount >)
    }

)