export type TreeNode = {
  val: number;
  left: TreeNode;
  right: TreeNode;
} | null;

export function inorderTraversal(root: TreeNode | null): number[] {
  if (!root) return [];

  const val: number[] = [];
  val.push(...inorderTraversal(root.left));
  val.push(root.val);
  val.push(...inorderTraversal(root.right));

  return val;
}
