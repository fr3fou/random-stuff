#include<stdio.h>
#include "./linked_list.h"

int main(void) {
    Node *node = malloc(sizeof(Node));
    node->Value = 1;
    node->Next = NULL;

    Node *node1 = malloc(sizeof(Node));
    node1->Value = 3;
    node1->Next = NULL;

    node->Next = node1;

    printf("before adding\n");
    print(node);

    add(node, 2);
    add(node, 3);
    add(node, 4);
    add(node, 5);

    printf("after adding\n");
    print(node);

    node = reverse(node);

    printf("after reversing\n");
    print(node);

    return 0;
}
