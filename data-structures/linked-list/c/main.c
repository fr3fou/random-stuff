#include<stdio.h>
#include "./linked_list.h"

int main(int argc, char *argv[]) {
    Node node; 
    node.value = 1;

    Node node1;
    node1.value = 2;

    node.Next = &node1;

    printf("%i\n", node.value);
    printf("%i\n", node.Next->value);

    return 0;
}
