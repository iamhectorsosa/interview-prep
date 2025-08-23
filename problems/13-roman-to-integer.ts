export function romanToInt(s: string): number {
  if (s.length < 1 || s.length > 15) return 0;

  const map: Record<string, number> = {
    I: 1,
    V: 5,
    X: 10,
    L: 50,
    C: 100,
    D: 500,
    M: 1000,
  };

  let res = 0;

  for (let i = 0; i < s.length; i++) {
    const currentVal = map[s[i]];
    if (!currentVal) return 0;

    const nextVal = map[s?.[i + 1]];
    if (nextVal && nextVal > currentVal) {
      res += nextVal - currentVal;
      i++;
      continue;
    }

    res += currentVal;
  }

  return res;
}
