class LLNode {
    constructor() {}
}

class LRUCache {
    constructor(capacity) {
        // Map of key to LLNode
        this.cache = {}
        this.head = null
        this.cap = capacity
        this.len = 0
    }

    get(key) {
        let node = this._get(key)
        if (node) {
            return node.value.value
        }
        return -1
    }

    _get(key) {
        if (!this.cache[key]) {
            return null
        }

        let node = this.cache[key]

        // No need to do anything if this is already the most recently used
        if (node === this.head) return node

        // Remove the node in the middle
        node.prev.next = node.next
        node.next.prev = node.prev

        // Make this the new head
        let previousHead = this.head
        this.head = node

        // Point the tail to the new head
        let tail = previousHead.prev
        tail.next = this.head

        // Move previous head as 2nd element
        this.head.next = previousHead
        this.head.prev = tail
        previousHead.prev = this.head

        return node
    }

    set(key, value) {
        let node = this._get(key)
        if (node) {
            node.value.value = value
            return
        }

        if (this.len >= this.cap) {
            // Delete the LRU element
            let lru = this.head.prev
            delete this.cache[lru.value.key]

            // Delete the head if it's the only one or make the new tail the node before 
            if (!this.head.next) {
                this.head = null
            } else {
                this.head.prev = lru.prev
            }
            this.len--
        }

        // Make the newly added node the new head
        let previousHead = this.head
        this.head = new LLNode()
        this.head.value = {
            key,
            value
        }

        // Wire up the tail and head to each other (or to itself if there are no other nodes)
        if (previousHead) {
            let tail = previousHead.prev
            tail.next = this.head

            // Move previous head as 2nd element
            this.head.next = previousHead
            this.head.prev = tail
            previousHead.prev = this.head
        } else {
            this.head.prev = this.head
        }

        this.cache[key] = this.head
        this.len++
    }
}
