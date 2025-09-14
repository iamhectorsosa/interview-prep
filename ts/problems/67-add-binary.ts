export function addBinary(a: string, b: string): string {
  let carry = 0;
  let result = "";
  let i = a.length - 1;
  let j = b.length - 1;

  while (i >= 0 || j >= 0) {
    let sum = carry;

    if (i >= 0) {
      sum += parseInt(a[i]);
      i--;
    }

    if (j >= 0) {
      sum += parseInt(b[j]);
      j--;
    }

    carry = Math.floor(sum / 2);
    let char = sum % 2;
    result = String(char) + result;
  }

  if (carry > 0) {
    return "1" + result;
  }
  return result;
}
