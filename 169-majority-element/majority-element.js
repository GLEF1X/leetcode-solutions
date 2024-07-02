/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
    nums.sort()
    const threshold = Math.ceil(nums.length / 2)
    let counter = 0;
    for (const num of nums) {
        counter++
        if (counter == threshold) {
            return num
        }
    }

    return 0
};