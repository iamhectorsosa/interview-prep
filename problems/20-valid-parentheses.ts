export function isValid(s: string): boolean {
  if (!s.length) return false;

  const opPar = "({[";
  const clPar: Record<string, string> = {
    ")": "(",
    "}": "{",
    "]": "[",
  };

  const stack: Array<string> = [];
  for (let i = 0; i < s.length; i++) {
    switch (true) {
      case opPar.includes(s[i]):
        stack.push(s[i]);
        continue;
      case clPar[s[i]] === stack.at(-1):
        stack.pop();
        continue;
      default:
        return false;
    }
  }

  return !stack.length;
}
