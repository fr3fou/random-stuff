using System;
using System.Collections.Generic;

namespace Linked_List
{
    class Program
    {
        static void Main(string[] args)
        {
            LinkedList<int> linkedList = new LinkedList<int>();
            linkedList.AddSorted(29);
            linkedList.AddSorted(19);
            linkedList.AddSorted(9);
            linkedList.AddSorted(1);
            linkedList.Visualize();
        }
    }
}
