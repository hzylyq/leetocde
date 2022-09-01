#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}

impl Solution {
    pub fn final_prices(prices: Vec<i32>) -> Vec<i32> {
        for i in 0..prices.len() {
            for j in i+1..prices.len() {
                if prices[j] < prices[i] {
                    prices[i] = prices[i] - prices[j];
                }
            }
        }

        return prices;
    }
}
