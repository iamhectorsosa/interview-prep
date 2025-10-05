export type TreeNode = {
  val: number;
  left: TreeNode;
  right: TreeNode;
} | null;

export function isSameTree(p?: TreeNode, q?: TreeNode): boolean {
  if (!p && !q) return true;
  if (!p || !q) return false;
  if (p?.val !== q?.val) return false;
  return isSameTree(p?.left, q?.left) && isSameTree(p?.right, q?.right);
}
