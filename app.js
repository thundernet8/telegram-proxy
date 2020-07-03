const Koa = require('koa');
const proxy = require('koa-proxy');

const app = new Koa();
app.use((ctx, next) => {
    if (ctx.path === '/') {
        return ctx.body = 'Hello';
    }
    console.log(`Proxy request: ${ctx.path}`);
    const p = proxy({
        host: 'https://api.telegram.org',
        jar: true,
    });
    return p(ctx, next);
});

app.listen(9099, () => {
    console.log('App started at port 9099');
});