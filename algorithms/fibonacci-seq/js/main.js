function* fib(n) {
  let prev = 1
  let current = 1
  yield prev
  yield current
  for (let i = 0; i < n; i++) {
    let nth = prev + current
    prev = current
    current = nth
    yield nth
  }
}

const f = fib(100)
for (var i = 0; i < 100; i++) {
  console.log(f.next().value)
}
