#include<stdio.h>
#include<stdbool.h>
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

    printf("after adding\n");
    print(node);

    node = reverse(node);

    printf("after reversing\n");
    print(node);

    bool contains2 = contains(node, 2);

    if(contains2) {
        printf("the list contains 2\n");
    } else {
        printf("the list doesn't contain 2\n");
    }

    bool contains9 = contains(node, 9);

    if(contains9) {
        printf("the list contains 9\n");
    } else {
        printf("the list doesn't contain 9\n");
    }

    add_to_end(node, 10);

    printf("after adding to end\n");
    print(node);

    return 0;
}
