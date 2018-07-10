using System;
using System.Collections.Generic;
namespace BinaryTree
{
    class Tree<T> where T : IComparable // make sure it's a valid number (all numeric types implement IComparable)
    {
        public Node<T> Root { get; set; } // the top (bottom) of the tree 

        public Tree(Node<T> node)
        {
            this.Root = node; // assign passed node as root node
        }

        public Tree(T value)
        {
            this.Root = new Node<T>(value); // make a new node with the passed value, and set it as the root 
        }

        public void Add(Node<T> node)
        {
            try
            {
                this.Root.Add(node); // add passed node to the root
            }
            catch (System.NullReferenceException)
            {
                throw new Exception("Root must have an initial value.");
            }
        }

        public void Add(T value)
        {
            try
            {
                this.Root.Add(value); // add passed value to the root
            }
            catch (System.NullReferenceException)
            {
                throw new Exception("Root must have an initial value.");
            }
        }

        public bool Contains(Node<T> node)
        {
            return this.Root.Contains(node);
        }

        public bool Contains(T value)
        {
            return this.Root.Contains(value);
        }

        public override string ToString()
        {
            return this.Root.ToString();
        }
    }

    public class Node<T> where T : IComparable // make sure it's a valid number (all numeric types implement IComparable)
    {
        public T Value { get; set; } // current node value
        public Node<T> LeftNode { get; set; } = null;// the node to the left (smaller)
        public Node<T> RightNode { get; set; } = null; // the node to the right (bigger)
        private Comparer<T> comparer = Comparer<T>.Default; // used to compare values of generic types

        public Node(T value)
        {
            this.Value = value;
        }

        public void Add(Node<T> node)
        {
            if (comparer.Compare(node.Value, this.Value) == 0) // if the values are equal
            {
                throw new Exception("A node with the same value has already been added."); // throw an exception as binary trees don't have duplicate elements
            }

            if (comparer.Compare(node.Value, this.Value) < 0) // if the passed value is less than the current node's value
            {
                if (this.LeftNode == null) // check if there isn't a node to the left
                {
                    this.LeftNode = node; // and make a new node with the passed value
                }
                else // else there will be a node on the left
                {
                    this.LeftNode.Add(node); // recursively check if there is a node on the left 
                }
            }

            else // else the passed value is greater than the current node's value
            {
                if (this.RightNode == null) // check if there isn't a node to the right
                {
                    this.RightNode = node; // and make a new node with the passed value
                }
                else // else there will be a node on the right
                {
                    this.RightNode.Add(node); // recursively check if there is a node on the right 
                }
            }
        }

        public void Add(T value)
        {
            if (comparer.Compare(value, this.Value) == 0) // if the values are equal
            {
                throw new Exception("A node with the same value has already been added."); // throw an exception as binary trees don't have duplicate elements
            }

            if (comparer.Compare(value, this.Value) < 0) // if the passed value is less than the current node's value
            {
                if (this.LeftNode == null) // check if there isn't a node to the left
                {
                    this.LeftNode = new Node<T>(value); // and make a new node with the passed value
                }
                else // else there will be a node on the left
                {
                    this.LeftNode.Add(value); // recursively check if there is a node on the left 
                }
            }

            else // else the passed value is greater than the current node's value
            {
                if (this.RightNode == null) // check if there isn't a node to the right
                {
                    this.RightNode = new Node<T>(value); // and make a new node with the passed value
                }
                else // else there will be a node on the left
                {
                    this.RightNode.Add(value); // recursively check if there is a node on the right 
                }
            }
        }

        public bool Contains(Node<T> node)
        {
            if (comparer.Compare(node.Value, this.Value) == 0) // if the values are equal
            {
                return true; // we have found it
            }
            else if (comparer.Compare(node.Value, this.Value) < 0 && this.LeftNode != null) // if the passed value is less than the current node's value and if there is a node to the left
            {
                return this.LeftNode.Contains(node); // recursively search 
            }
            else if (comparer.Compare(node.Value, this.Value) > 0 && this.RightNode != null) // if the passed value is greater than the current node's value and if there is a node to the right
            {
                return this.RightNode.Contains(node); // recursively search 
            }
            return false;
        }

        public bool Contains(T value)
        {
            if (comparer.Compare(value, this.Value) == 0) // if the values are equal
            {
                return true; // we have found it
            }
            else if (comparer.Compare(value, this.Value) < 0 && this.LeftNode != null) // if the passed value is less than the current node's value and if there is a node to the left
            {
                return this.LeftNode.Contains(value); // recursively search 
            }
            else if (comparer.Compare(value, this.Value) > 0 && this.RightNode != null) // if the passed value is greater than the current node's value and if there is a node to the right
            {
                return this.RightNode.Contains(value); // recursively search 
            }
            return false;
        }

        public override string ToString()
        {
            string result = "";
            return result;
        }
    }
}
