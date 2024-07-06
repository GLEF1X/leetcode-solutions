/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let buyPrice = prices[0], profit = 0;

    for (const price of prices) {
        if (price < buyPrice) {
            buyPrice = price
        } else if (price > buyPrice) {
            profit += price - buyPrice
            buyPrice = price
        }
    }

    return profit
};