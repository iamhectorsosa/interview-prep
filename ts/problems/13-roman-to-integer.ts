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

  // const integers = s.split("").map((c) => map[c]);
  // return integers.reduce(
  //   (acc, x, i) => (x < integers[i + 1] ? acc - x : acc + x),
  //   0,
  // );

  let res = 0;

  for (let i = 0; i < s.length; i++) {
    const val = map[s[i]];
    if (!val) return 0;

    if (map[s?.[i + 1]] > val) {
      res -= val;
      continue;
    }

    res += val;
  }

  return res;
}
