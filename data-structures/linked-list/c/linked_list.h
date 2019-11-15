#include<stdlib.h>
typedef struct Node {
    int Value;
    struct Node *Next;
} Node;

void add(Node* n, int v) {
    Node *next = n->Next;

    Node *temp = malloc(sizeof(Node));
    temp->Value = v;
    temp->Next = next;

    n->Next = temp;
} 

void reverse(Node* n) {
    Node *prev, *current, *next;

    current = n;
    next = n->Next;

    while (!next) {
        current->Next = prev;

        Node *temp = next->Next;
        next->Next = current;
        
        prev = current;
        current = next;
        next = temp;
    }

    n = current;
}
