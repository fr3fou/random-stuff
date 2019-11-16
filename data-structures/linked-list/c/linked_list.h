#include<stdbool.h>
#include<stdlib.h>

typedef struct Node {
    int Value;
    struct Node *Next;
} Node;

void add(Node *n, int v) {
    Node *next = n->Next;

    Node *temp = malloc(sizeof(Node));
    temp->Value = v;
    temp->Next = next;

    n->Next = temp;
} 

void add_to_end(Node *n, int v) {
    if(n->Value == 0) {
        n->Value = v;
        return;
    }

    if (n->Next == NULL) {
        Node *temp = malloc(sizeof(Node));
        temp->Value = v;
        n->Next = temp;
        return;
    }

    add_to_end(n->Next, v);
}

void print(Node *n) {
    for (Node *next = n; next != NULL; next = next->Next) {
        // end case
        if (next->Next == NULL) {
            printf("%d -> NULL\n", next->Value);
            return;
        }

        // normal case
        printf("%d -> %d\n", next->Value, next->Next->Value);
    }
}

Node *reverse(Node *n) {
    Node *prev, *current, *next;

    prev = NULL;
    current = n;
    next = n->Next;

    while (next != NULL) {
        current->Next = prev;
        Node *temp = next->Next;
        next->Next = current;
        
        prev = current;
        current = next;
        next = temp;
    }

    return current;
}

bool contains(Node *n, int v) {
    for (Node *next = n; next != NULL; next = next->Next) {
        if(next->Value == v) {
            return true;
        }
    }

    return false;
}
