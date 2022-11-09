function fib(max) {
  if (max < 2) {
    return 1
  }
  return fib(max - 1) + fib(max - 2)
}
function fib30Times(max) {
  let sum = 0
  for (let i = 0; i < 30; i++) {
    sum += fib(max)
  }
  return sum
}