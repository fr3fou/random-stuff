import LinkedList from "./linked-list";

const list = new LinkedList(5);

list.add(4);
list.add(1);
list.add(9);

console.log(list.contains(1));
console.log(list.contains(0));

console.log(JSON.stringify(list, null, 2));
