
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}

pub struct Array;

impl Array {
    pub fn apply_operations(nums: Vec<i32>) -> Vec<i32> {
        let mut j = 0;
        for i in 0..nums.len() {
            if i + 1 < nums.len() && nums[i] == nums[i + 1] {
                nums[i] *= 2;
                nums[i + 1] = 0;
            }

            if nums[i] != 0 {
                j += 1;

                // nums[i], nums[j] = nums[j], nums[i];
            }
        }

        return nums;
    }
}
