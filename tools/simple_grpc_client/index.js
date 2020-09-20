const path = require('path')
const Mali = require('mali')

const PROTO_PATH = path.resolve(__dirname, './greeter.proto')

async function sayHello (ctx) {
    ctx.res = { message: 'Hello '.concat(ctx.req.name) }
}

function main () {
    const app = new Mali(PROTO_PATH)
    app.use({ sayHello })
    app.start('127.0.0.1:50001')
}

main()
