use core::num;

pub struct Array;

impl Array {
    pub fn apply_operations(nums: Vec<i32>) -> Vec<i32> {
        let mut j = 0;
        let mut nums_clone = nums.clone();

        for i in 0..nums_clone.len() {
            if i + 1 < nums_clone.len() && nums_clone[i] == nums_clone[i + 1] {
                nums_clone[i] *= 2;
                nums_clone[i + 1] = 0;
            }

            if nums_clone[i] != 0 {
                let temp = nums_clone[i];
                nums_clone[i] = nums_clone[j];
                nums_clone[j] = temp;
                j += 1;
            }
        }

        return nums_clone;
    }

    // 27. 移除元素
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        let mut j = 0;
        for i in 0..nums.len() {
            if nums[i] != val {
                nums[j] = nums[i];
                j += 1;            
            }
        }
        return j as i32;
    }
}
