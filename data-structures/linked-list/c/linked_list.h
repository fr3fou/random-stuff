typedef struct Node {
    int Value;
    struct Node *Next;
} Node;

void add(Node* n, int v) {
    Node *next = n->Next;

    Node temp;
    temp.Value = v;

    n->Next = &temp;
} 
