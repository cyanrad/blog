// By Roni Kaufman
// https://ronikaufman.github.io
// Changes: made it reset on resize and changed colors & opacity

let paths = [],
  points = [],
  pointIdx = 0,
  canvas = null;
const MAX_DEPTH = 3,
  N = 5, // block count (don't set too high)
  POINTS_PER_FRAME = 1;
COLOR_OPACITY = 50; // game speed

function setup() {
  // canvas
  canvas = createCanvas(windowHeight, windowHeight);
  canvas.position(windowWidth / 2 - windowHeight / 2, 0);
  canvas.style('z-index', '-1');

  strokeCap(PROJECT);
  findPaths([0, 0], [[0, 0]]);

  // color
  emerald = color(52, 211, 153, COLOR_OPACITY);
  indigo = color(99, 102, 241, COLOR_OPACITY);
  darkBlue = color(33, 53, 71, COLOR_OPACITY);
  darkZink = color(24, 24, 27); // bg color

  let margin = 50;
  let width = windowHeight;
  strokeWeight(((1 / 2) * (width - 2 * margin)) / pow(N, MAX_DEPTH));
  rsfc(margin, margin, width - 2 * margin, true, true, 0);

  background(darkZink);
}

function draw() {
  for (let k = 0; k < POINTS_PER_FRAME; k++) {
    let p1 = points[pointIdx],
      p2 = points[pointIdx + 1];
    stroke(rainbow(pointIdx / points.length));
    line(p1.x, p1.y, p2.x, p2.y);

    pointIdx++;
    if (pointIdx == points.length - 1) {
      noLoop();
      break;
    }
  }
}

function windowResized() {
  paths = [];
  points = [];
  pointIdx = 0;

  setup();
}

function possibleNeighbors([i, j]) {
  let possibilities = [];
  if (i % 2 == 0 && j < N - 1) possibilities.push([i, j + 1]);
  if (i % 2 == 1 && j > 0) possibilities.push([i, j - 1]);
  if (j % 2 == 0 && i < N - 1) possibilities.push([i + 1, j]);
  if (j % 2 == 1 && i > 0) possibilities.push([i - 1, j]);
  return possibilities;
}

function inArray([i, j], arr) {
  for (let e of arr) {
    if (e[0] == i && e[1] == j) return true;
  }
  return false;
}

// find all paths in a N*N grid, going from top-left to bottom-right and through all points
function findPaths(p, visited) {
  let neighbors = possibleNeighbors(p);
  if (neighbors.length == 0) {
    if (visited.length == sq(N)) paths.push(visited);
    return;
  }
  for (let neigh of neighbors) {
    if (!inArray(neigh, visited)) findPaths(neigh, [...visited, neigh]);
  }
}

// random space-filling curve
function rsfc(x0, y0, s, topToBottom, leftToRight, depth) {
  if (depth == MAX_DEPTH) {
    points.push({ x: x0 + s / 2, y: y0 + s / 2 });
    return;
  }

  let newS = s / N;
  let idx1 = topToBottom ? 0 : 1;
  let idx2 = leftToRight ? 0 : 1;
  let path = random(paths);

  for (let [i, j] of path) {
    let x = leftToRight ? i * newS : (N - i - 1) * newS;
    let y = topToBottom ? j * newS : (N - j - 1) * newS;
    rsfc(x0 + x, y0 + y, newS, i % 2 == idx1, j % 2 == idx2, depth + 1);
  }
}

function rainbow(t) {
  let palette = [emerald, indigo, darkBlue];
  let i = floor(palette.length * t);
  let amt = fract(palette.length * t);
  return lerpColor(
    color(palette[i % palette.length]),
    color(palette[(i + 1) % palette.length]),
    amt,
  );
}
