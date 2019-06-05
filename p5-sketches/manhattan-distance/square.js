class Square {
  constructor(x, y, thickness, size) {
    this.pos = createVector(x, y);
    this.thickness = thickness;
    this.size = size;
  }

  draw() {
    stroke(185);
    strokeWeight(squareThickness);
    fill(255);
    square(this.pos.x + this.thickness, this.pos.y + this.thickness, this.size);
  }
}
