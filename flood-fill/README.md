# Flood-Fill

Simple python program to fill a black & white image using the flood fill algorithm.

## Requirements

1. Python

    Download from [here!](https://www.python.org/)

2. Pillow

    Install using your py package manager.

```sh
pip install Pillow
```

## Usage

1. Place a black & white image in the same directory as main.py

2. Run main.py using a terminal of your choice. Provide the image name + coordinates of where to start the flood fill.

```sh
py main.py simple50.png 15 30
```

This will start a flood fill on `simple50.png` at X: `15`, Y: `30`. The output file will be in the same dir under the name `simple50_filled.png`