class Node<T extends number> {
  next: Node<T> = null;
  value: T;

  constructor(val: T) {
    this.value = val;
  }

  add(val: T) {
    if (this.next !== null) {
      this.next.add(val);
    } else {
      this.next = new Node(val);
    }
  }

  contains(val: T) {
    if (this.value === val) {
      return true;
    } else if (this.next !== null) {
      return this.next.contains(val);
    } else {
      return false;
    }
  }

  traverse() {
    const result: string[] = [];

    if (this.next !== null) {
      result.push(`${this.value} -> `);
      result.push(this.next.traverse());
    } else {
      result.push(`${this.value}`);
    }

    return result.join("");
  }
}

export default Node;
