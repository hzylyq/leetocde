#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);

        use super::Solution;
        let res = Solution::product_except_self(vec![1, 2]);
        assert_eq!(res[0], 2)
    }
}

pub struct Solution;

impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut r: Vec<i32> = vec![1; nums.len()];
        let mut l: Vec<i32> = vec![1; nums.len()];
        let mut res: Vec<i32> = vec![1; nums.len()];
        for i in 1..nums.len() {
            l[i] = l[i - 1] * nums[i - 1];
        }
        for i in (0..nums.len()-1).rev() {
            r[i] = r[i + 1] * nums[i + 1];
        }
        for i in 0..nums.len() {
            res[i] = l[i] * r[i];
        }

        return res;
    }
}