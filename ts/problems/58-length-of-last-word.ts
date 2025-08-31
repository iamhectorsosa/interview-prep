export function lengthOfLastWord(s: string): number {
  // return s.trimEnd().split(" ").pop()?.length || 0;
  if (!s.length) return 0;
  let length = 0;
  for (let i = s.length - 1; i >= 0; i--) {
    if (s[i] !== " ") length++;
    if (length && s[i] === " ") return length;
  }
  return length;
}
