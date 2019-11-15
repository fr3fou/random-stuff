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
