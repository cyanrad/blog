// canvas variables
let canvas;

// box variables
const boxCount = 20;
let boxesPosition = [];
let availableRows = [];

const spawnAtCount = 20;
let counter = 0;

// colors
let emerald; // const
let darkZink; // const

function setup() {
  // canvas
  canvas = createCanvas(windowWidth, windowHeight);
  canvas.position(0, 0);
  canvas.style('z-index', '-1');

  // color
  emerald = color(52, 211, 153);
  darkZink = color(24, 24, 27);

  // shapes
  noStroke();
  for (let i = 0; i < 20; i++) {
    boxesPosition.push(-1);
  }
}

function windowResized() {
  resizeCanvas(windowWidth, windowHeight);
}

function draw() {
  counter++;
  boxLen = canvas.height / boxCount;

  // drawing
  background(darkZink);

  if (counter == spawnAtCount) {
    row = random(0, 20);
    boxesPosition[row] = 0;
    drawBlockChain(0, 0, boxLen, 255, 7);
  }
}

// this is really not effecient
// and kinda bad, it does many things
function drawBlockChain(x, y, boxLen, maxRectOpacity, trailCount) {
  let opacityStep = maxRectOpacity / trailCount;

  for (let i = 0; i < trailCount; i++) {
    emeraldTransparent = color(
      red(emerald),
      green(emerald),
      blue(emerald),
      maxRectOpacity - i * opacityStep,
    );
    fill(emeraldTransparent);

    rect(x + i * boxLen, y, boxLen);
  }
}
