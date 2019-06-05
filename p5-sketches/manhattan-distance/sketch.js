const canvasWidth = 800;
const squareSize = 100;
const squareThickness = 25;
const canvasHeight = canvasWidth;
const squareAmount = canvasHeight / squareSize;

let squares = [];
let start;
let end;
let startOrEnd = false;

function setup() {
  const c = createCanvas(
    canvasWidth + 3 * squareThickness,
    canvasHeight + 3 * squareThickness
  );

  for (let i = 0; i < canvasWidth; i += squareSize) {
    for (let j = 0; j < canvasHeight; j += squareSize) {
      squares.push(new Square(i, j, squareThickness, squareSize));
    }
  }

  c.mousePressed(handleClick);
}

function draw() {
  background(255);
  for (const square of squares) {
    square.draw();
  }

  if (start) {
    fill(0);
    noStroke();
    circle(start.x, start.y, 50);
  }

  if (end) {
    fill(0);
    noStroke();
    circle(end.x, end.y, 50);
  }
}

function handleClick() {
  let x, y;
  let bestDistance = Infinity;

  for (const square of squares) {
    let currentDistance, currentX, currentY;

    // top left corner
    currentX = square.pos.x;
    currentY = square.pos.y;
    currentDistance = dist(mouseX, mouseY, currentX, currentY);
    if (currentDistance < bestDistance) {
      bestDistance = currentDistance;
      x = currentX + squareThickness;
      y = currentY + squareThickness;
    }

    // top right corner
    currentX = square.pos.x + squareSize;
    currentY = square.pos.y;
    currentDistance = dist(mouseX, mouseY, currentX, currentY);
    if (currentDistance < bestDistance) {
      bestDistance = currentDistance;
      x = currentX + squareThickness;
      y = currentY + squareThickness;
    }

    // bottom right corner
    currentX = square.pos.x + squareSize;
    currentY = square.pos.y + squareSize;
    currentDistance = dist(mouseX, mouseY, currentX, currentY);
    if (currentDistance < bestDistance) {
      bestDistance = currentDistance;
      x = currentX + squareThickness;
      y = currentY + squareThickness;
    }

    // bottom left corner
    // bottom right corner
    currentX = square.pos.x;
    currentY = square.pos.y + squareSize;
    currentDistance = dist(mouseX, mouseY, currentX, currentY);
    if (currentDistance < bestDistance) {
      bestDistance = currentDistance;
      x = currentX + squareThickness;
      y = currentY + squareThickness;
    }
  }

  if (startOrEnd) {
    start = createVector(x, y);
  } else {
    end = createVector(x, y);
  }

  startOrEnd = !startOrEnd;
}
