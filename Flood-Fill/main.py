from PIL import Image
import numpy as np
import sys

file_path = sys.argv[1]
x = int(sys.argv[2])
y = int(sys.argv[3])

print('Loading image...')
img = Image.open(file_path)

width = img.size[0]
height = img.size[1]

if x > width or x < 0 or y > height or y < 0:
    sys.exit('Invalid coordinates!')

print(f'{file_path} successfully loaded!')








