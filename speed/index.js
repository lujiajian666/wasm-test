function fib(max) {
  if (max < 2) {
    return 1
  }
  return fib(max - 1) + fib(max - 2)
}
function fib30Times(max) {
  for (let i = 0; i < 30; i++) {
    fib(max)
  }
  return 'end'
}