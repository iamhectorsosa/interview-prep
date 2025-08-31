export function strStr(haystack: string, needle: string): number {
  if (!haystack.length && !needle.length) return 0;
  // return haystack.indexOf(needle);
  for (let i = 0; i <= haystack.length - needle.length; i++) {
    if (haystack.slice(i, i + needle.length) === needle) {
      return i;
    }
  }

  return -1;
}
