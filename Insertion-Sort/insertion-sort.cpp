// g++ insertion-sort.cpp -o a ; ./a.exe
#include <iostream>
#include <array>
using namespace std;

template <size_t arrSize>
void insertion_sort(array<int, arrSize> &arr)
{
    int size = arr.size();

    for (int i = 1; i < size; i++)
    {
        int num = arr[i];
        for (int j = i - 1; j >= 0; j--)
        {
            int prev_num = arr[j];
            if (num > prev_num)
            {
                break;
            }
            if (num < prev_num)
            {
                arr[j + 1] = prev_num;
                arr[j] = num;
            }
        }
    }
}

int main()
{
    // 1 2 3 4 5 6 8 11 12 15 <-- goal
    // 4 5 12 6 3 8 11 15 1 2 <-- original
    array<int, 10> arr = {10, 9, 8, 7, 6, 5, 4, 3, 2, 1};
    insertion_sort(arr);
    int size = arr.size();
    for (size_t i = 0; i < size; i++)
    {
        cout << arr[i] << endl;
    }
    return 0;
}
