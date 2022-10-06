module.exports = {
  '/main.wasm': {
    target: 'http://127.0.0.1:8000',
    on: {
      proxyReq: (proxyReq, req, res) => {
        console.debug(req)
      },
      proxyRes: (proxyRes, req, res) => {
        console.debug(req)
        proxyRes.headers['Content-Type'] = 'application/wasm';
      },
    }
  }
}