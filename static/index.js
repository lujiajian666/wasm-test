function fib(max) {
  const slice = [1,1]
  if (max < 3) {
    return slice[max - 1]
  }
  for (let i = 2; i < max; i++) {
    [slice[0], slice[1]] = [slice[1], slice[0]+ slice[1]]
  }
  return slice[1];
}
function fib2(max) {
  if (max < 2) {
    return 1
  }
  return fib2(max - 1) + fib2(max - 2)
}
function fib3(max) {
  for (let i = 0; i < 30; i++) {
    fib2(max)
  }
  return 'end'
}
function genGar(max) {
	return Array(max).fill(null)
}