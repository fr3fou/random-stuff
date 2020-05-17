const rl = require("readline").createInterface({
  input: process.stdin,
  output: process.stdout,
});
let l = 0;
let shift, str;
rl.on("line", (line) => {
  if (l == 0) shift = line;
  if (l == 1) {
    str = line;
    console.log(solve(parseInt(shift), str.split(" ")).join(" "));
  }
  l++;
});

function solve(shift, arr) {
  const ret = new Array(arr.length);
  for (let i = 0; i < arr.length; i++) {
    ret[(arr.length + (i + (shift % arr.length))) % arr.length] = arr[i];
  }
  return ret;
}
