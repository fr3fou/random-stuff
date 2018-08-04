def bubblesort(arr)
    isNotSorted = true;
    while isNotSorted
        isNotSorted = false
        for i in 1..arr.length-1
            if arr[i-1] > arr[i]
                temp = arr[i]
                arr[i] = arr[i-1]
                arr[i-1] = temp
                isNotSorted = true
            end
        end
    end
end
#      4 5 6 3 8 11 12 1 2 15
arr = [4, 5, 12, 6, 3, 8, 11, 15, 1, 2]

bubblesort(arr)

puts arr