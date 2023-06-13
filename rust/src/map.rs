#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        use super::Map;
        let res = Map::unequal_triplets(vec![1, 2, 3]);
        println!("{}", res);
    }
}

pub struct Map;

impl Map {
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        return false;
    }

    // 2475. 数组中不等三元组的数目
    pub fn unequal_triplets(nums: Vec<i32>) -> i32 {
        use std::collections::HashMap;
        let mut count = HashMap::new();
        for i in 0..nums.len() {
            count.entry(&nums[i]).and_modify(|e| { *e += 1 }).or_insert(1);
        }

        let mut res = 0;
        let mut n = nums.len();
        let mut t = 0;

        for (k, v) in &count {
            res = res + (t * v * (n - t - v));
            t = t + v;
        }

        return res as i32;
    }
}