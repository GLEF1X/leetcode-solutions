/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
    let count = 0, candidate = 0;

    for (const num of nums) {
        if (count === 0) {
            candidate = num
        }

        if (num === candidate) {
            count++
        } else {
            count--
        }
    }

    return candidate
 };