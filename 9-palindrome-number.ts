export function isPalindrome(x: number): boolean {
  if (x < 0) return false;
  // return String(x) === String(x).split("").reverse().join("");
  let c = x;
  let y = 0;
  while (c > 0) {
    y = y * 10 + (c % 10);
    c = Math.floor(c / 10);
  }

  return x === y;
}
