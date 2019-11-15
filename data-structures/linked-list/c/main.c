#include<stdio.h>
#include "./linked_list.h"

int main(void) {
    Node node; 
    node.Value = 1;

    Node node1;
    node1.Value = 3;

    node.Next = &node1;

    printf("before adding\n");
    printf("%i\n", node.Value);
    printf("%i\n", node.Next->Value);

    add(&node, 2);

    printf("after adding\n");
    printf("%i\n", node.Value);
    printf("%i\n", node.Next->Value);
    printf("%i\n", node.Next->Next->Value);

    /* reverse(&node); */
    print(&node);

    printf("after reversing\n");
    printf("%i\n", node.Value);
    printf("%i\n", node.Next->Value);
    printf("%i\n", node.Next->Next->Value);


    return 0;
}
