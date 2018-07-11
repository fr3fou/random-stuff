using System;
using System.Collections.Generic;

namespace Linked_List {
    class Node<T> where T : IComparable {
        public T Data;
        public Node<T> Next;
        Comparer<T> comparer = Comparer<T>.Default;
        public Node (T data) {
            this.Data = data;
            Next = null;
        }

        public void Add (T data) {
            if (Next == null) {
                Next = new Node<T> (data);
            } else {
                Next.Add (data);
            }
        }

         public void AddSorted (T data) {

            if (Next == null) {
                Next = new Node<T> (data);
            } else if(comparer.Compare(data, Next.Data) < 0 ) {
                Node<T> tempNode = new Node<T>(data);
                tempNode.Next = Next;
                Next = tempNode;
            }
            else{
                Next.AddSorted(data);
            }
        }

        public void Visualize () {
            System.Console.WriteLine ("[{0}]", this.Data);
            if (Next != null) {
                Next.Visualize ();
            }
        }
    }

    class LinkedList<T> where T : IComparable {
        public Node<T> headNode;
        Comparer<T> comparer = Comparer<T>.Default;
        public LinkedList() {
            headNode = null;
        }

        public LinkedList(T data) {
            headNode = new Node<T> (data);
        }

        public void AddToEnd(T data)
        {
            if(headNode==null)
            {
                headNode = new Node<T>(data);
            }
            else
            {
                headNode.Add(data);
            }
        }

         public void AddToBeginning(T data)
        {
            if(headNode==null)
            {
                headNode = new Node<T>(data);
            }
            else
            {
                Node<T> tempNode = new Node<T>(data);   // In order to add a node to the beginning of the list
                tempNode.Next = headNode;               // we will have to point the next node to be our head node
                headNode = tempNode;                    // then we say that our head node (headNode) will be pointing to our new "temp" node (tempNode) since every head node of a Linked List is only a reference to the first node.
            }
        }

        public void AddSorted(T data)
        {
            if(headNode==null) headNode= new Node<T>(data);
            else if(comparer.Compare(data, headNode.Data)<0)
            {
                AddToBeginning(data);
            }
            else
            {
                headNode.AddSorted(data);
            }
        }
         public void Visualize()
        {
            if(headNode!=null)
            {
                headNode.Visualize();
            }
        }

    }
}