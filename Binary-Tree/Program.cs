using System;
using BinaryTree;
namespace Binary_Tree_Exercise
{
    class Program
    {
        static void Main(string[] args)
        {
            Tree<int> tree = new Tree<int>(50);
            Random random = new Random();
            for (int i = 0; i < 30; i++)
            {
                tree.Add(random.Next(0,100));
            }
            Console.WriteLine(tree.Contains(58));
        }
    }
}
