#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}

pub struct Grady;

impl Grady {
    pub fn jump(nums: Vec<i32>) -> i32 {
        use std::cmp::max;

        let mut step = 0;
        let mut end = 0;
        let mut max_position = 0;

        for i in 0..(nums.len() - 1) {
            max_position = max(max_position, i + nums[i] as usize);
            if i == end {
                end = max_position;
                step += 1;
            }
        }

        return step;
    }
}