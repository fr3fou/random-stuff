def selection_sort(arr):
    for i in range(len(arr)):
        PIVOT = i
        for j in range(i+1, len(arr)):
            if arr[j] < arr[PIVOT]:
                PIVOT = j
        TEMP = arr[i]
        arr[i] = arr[PIVOT]
        arr[PIVOT] = TEMP
    return arr


ARR = [4, 5, 12, 6, 3, 8, 11, 15, 1, 2]
ARR = selection_sort(ARR)
print(ARR)
