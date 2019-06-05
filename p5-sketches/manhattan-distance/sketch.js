const canvasWidth = 800;
const squareThickness = 25;
const canvasHeight = canvasWidth;
const squareAmount = 6;
const squareSize = Math.ceil(canvasHeight / squareAmount);
const squares = [];

let start;
let end;
let startOrEnd = false;
let euclideanDistance;
let manhattanDistance;
let distP;

function setup() {
  console.log(squareSize);
  const c = createCanvas(
    canvasWidth + 3 * squareThickness,
    canvasHeight + 3 * squareThickness,
  );

  for (let i = 0; i < canvasWidth; i += squareSize) {
    for (let j = 0; j < canvasHeight; j += squareSize) {
      squares.push(new Square(i, j, squareThickness, squareSize));
    }
  }

  distP = createP();

  c.mousePressed(handleClick);
}

function draw() {
  background(255);

  for (const square of squares) {
    square.draw();
  }

  if (start && end) {
    // draw euclidean distance
    stroke(0, 255, 0);
    strokeWeight(10);
    line(start.x, start.y, end.x, end.y);

    // draw manhattan distance

    // start is to the left of end
    let maxLines = manhattanDistance / squareSize; 
    if (start.x <= end.x) {
      // start is under end
      if (start.y <= end.y) {
        let prevEnd = createVector(start.x, start.y);
        stroke(0,0,255);
        for (let i = 0; i <= maxLines; i++) {
          if (i % 2 == 0) {
            line(prevEnd.x, prevEnd.y, prevEnd.x + squareSize, prevEnd.y)
            prevEnd.x = prevEnd.x + squareSize;
          } else {
            line(prevEnd.x, prevEnd.y, prevEnd.x, prevEnd.y + squareSize)
            prevEnd.y = prevEnd.y + squareSize;
          }
        }
      } else {
      }
    } else {
      // start is under end
      if (start.y <= end.y) {
      } else {
      }
    }
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

  if (start && end) {
    euclideanDistance = floor(dist(start.x, start.y, end.x, end.y));
    manhattanDistance = floor(abs(start.x - end.x) + abs(start.y - end.y));

    distP.html(`
     Euclidean Distance: ${euclideanDistance}
     </br>
     Manhattan Distance: ${manhattanDistance}
    `);
  }
}
