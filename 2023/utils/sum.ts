export function sum(nums: number[]) {
  return nums.reduce((sum, curr) => sum + curr, 0);
}
