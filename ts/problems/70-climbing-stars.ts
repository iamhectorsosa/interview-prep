export function climbStairs(n: number): number {
  let a = 1;
  let b = 0;
  for (let i = 0; i < n; i++) {
    let temp = a;
    a = a + b;
    b = temp;
  }
  return a;
}
