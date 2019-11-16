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

