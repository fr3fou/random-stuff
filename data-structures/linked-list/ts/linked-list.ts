import Node from "./node";

class LinkedList<T> {
  private head: Node<T> = null;

  constructor(val?: T) {
    if (val) {
      this.head = new Node(val);
    }
  }

  add(val: T) {
    if (this.head !== null) {
      this.head.add(val);
    } else {
      this.head = new Node(val);
    }
  }

  contains(val: T) {
    if (this.head !== null) {
      return this.head.contains(val);
    } else {
      return false;
    }
  }
}

export default LinkedList;
