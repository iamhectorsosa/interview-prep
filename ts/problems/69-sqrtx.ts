export function mySqrt(x: number): number {
  let leftBoundary = 0;
  let rightBoundary = x;

  while (leftBoundary < rightBoundary) {
    let middle = Math.ceil((leftBoundary + rightBoundary) / 2);
    if (middle * middle <= x) {
      leftBoundary = middle;
      continue;
    }
    rightBoundary = middle - 1;
  }
  return leftBoundary;
}
