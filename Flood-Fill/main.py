from PIL import Image
import numpy as np
import sys

# get args
file_path = sys.argv[1]
x = int(sys.argv[2])
y = int(sys.argv[3])

# get img
print('Loading image...')
img = Image.open(file_path).convert('L')

# get dimensions
width = img.size[0]
height = img.size[1]

# we need 2 arrays - original and new
arr = np.array(img)

# convert 2d array values from (0 - 255) to (1s or 0s) 
for i in range(len(arr)):
    for j in range(len(arr[i])):
        if arr[i][j] == 255 and arr[i][j] == 255 and arr[i][j] == 255:
            arr[i][j] = 0
            arr[i][j] = 0
            arr[i][j] = 0
        else:
            arr[i][j] = 1
            arr[i][j] = 1
            arr[i][j] = 1

# validate
if x > width or x < 0 or y > height or y < 0:
    sys.exit('Invalid coordinates!')

print(f'{file_path} successfully loaded!')


# def flood_fill(arr, x, y):


# convert 2d array values from (1s or 0s) to (0 - 255)
for i in range(len(arr)):
    for j in range(len(arr[i])):
        if arr[i][j] == 1 and arr[i][j] == 1 and arr[i][j] == 1:
            arr[i][j] = 0
            arr[i][j] = 0
            arr[i][j] = 0
        else:
            arr[i][j] = 255
            arr[i][j] = 255
            arr[i][j] = 255

result = Image.fromarray(arr)
result.save('result.png')
