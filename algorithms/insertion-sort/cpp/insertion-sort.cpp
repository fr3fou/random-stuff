// g++ insertion-sort.cpp -o a ; ./a.exe
#include <iostream>
#include <array>
using namespace std;

template<typename T, std::size_t arr_size>
void print_arr(array<T, arr_size> &arr) {
    size_t size = arr.size();
    
    for(size_t i = 0; i < size; i++) {
        cout << arr[i] << endl;
    }
}

template<typename T, std::size_t arr_size>
void insertion_sort(array<T, arr_size> &arr) {
    size_t size = arr.size();

    for (size_t i = 1; i < size; i++) {
        T num = arr[i];
        size_t j = i - 1;
         
        while (j >= 0 && num < arr[j]) {
            arr[j + 1] = arr[j];
            arr[j] = num;
            j--;
        }
    }
}

int main() {
    // 1 2 3 4 5 6 8 11 12 15 <-- goal
    // 4 5 12 6 3 8 11 15 1 2 <-- original
    array<int, 10> arr = {4, 5, 12, 6, 3, 8, 11, 15, 1, 2};
    insertion_sort(arr);
    print_arr(arr);
    return 0;
}
