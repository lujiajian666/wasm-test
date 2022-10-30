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

function genGar(time, len) {
  let list
  for(let i = 0; i < time; i++) {
    list = Array(len).fill(null)
  }
  return list
}
// 泄漏
// function genGar() {
//   let list = Array(1000000).fill(null)
//   function closesure() {
//     list = Array(1).fill(null)
//   }
//   function getList(len) {
//     return Array(len).fill(null)
//   }
//   return getList
// }