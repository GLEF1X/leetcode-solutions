/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
  const table = {};

  for (const num of nums) {
    table[num] = (table[num] ?? 0) + 1;
  }

  let pair = [0, 0];
  for (const [key, val] of Object.entries(table)) {
    if (val > pair[1]) {
      pair = [key, val];
    }
  }

  return pair[0];
};