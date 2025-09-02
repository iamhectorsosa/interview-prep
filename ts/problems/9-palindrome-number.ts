export function isPalindrome(x: number): boolean {
  if (x < 0) return false;
  // return String(x) === String(x).split("").reverse().join("");
  let c = x;
  let r = 0;
  while (c > 0) {
    // Increase by magnitud of 10 and add the remainder or last digit
    r = r * 10 + (c % 10);
    // Divide by magnitud of 10 which drops the last digit
    c = Math.floor(c / 10);
  }

  return x === r;
}

export function isPalindromeStr(x: number): boolean {
  if (x < 0) return false;
  let r = x.toString().split("");
  let i = 0;
  let j = r.length - 1;
  while (i < j) {
    let t = r[i];
    r[i] = r[j];
    r[j] = t;
    i++;
    j--;
  }
  return x === parseInt(r.join(""));
}
