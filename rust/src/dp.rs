pub struct Dp;

impl Dp {
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
