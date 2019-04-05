import LinkedList from "./linked-list";

let val = 5;
const list = new LinkedList(val);

list.add(4);
list.add(1);
list.add(9);

console.log(list.contains(1));
console.log(list.contains(0));

console.log(list.traverse());

// console.log(JSON.stringify(list, null, 2));
