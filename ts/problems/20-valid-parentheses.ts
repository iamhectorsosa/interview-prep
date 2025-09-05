export function isValid(s: string): boolean {
  if (s.length < 2) return false;

  const pairs: Record<string, string> = {
    ")": "(",
    "}": "{",
    "]": "[",
  };

  const stack: Array<string> = [];
  for (const char of s) {
    switch (char) {
      case "(":
      case "{":
      case "[":
        stack.push(char);
        continue;
      case ")":
      case "}":
      case "]":
        if (stack.length > 0 && pairs[char] === stack.at(-1)) {
          stack.pop();
          continue;
        }
      default:
        return false;
    }
  }

  return !stack.length;
}
