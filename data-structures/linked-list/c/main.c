#include<stdio.h>
#include "./linked_list.h"

int main(void) {
    Node node; 
    node.Value = 1;

    Node node1;
    node1.Value = 2;

    node.Next = &node1;

    printf("%i\n", node.Value);
    printf("%i\n", node.Next->Value);
    printf("before adding\n");

    add(&node, 5);

    printf("after adding\n");
    printf("%i\n", node.Value);
    printf("%i\n", node.Next->Value);
    printf("%i\n", node.Next->Next->Value);

    return 0;
}
