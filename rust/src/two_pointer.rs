pub struct TwoPointer {}

impl TwoPointer {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        let mut slow = 0;
        for i in 0..nums.len() {
            if nums[i] != nums[slow as usize] {
                slow += 1;
                nums[slow as usize] = nums[i];
            }
        }

        return slow + 1;
    }

    pub fn remove_duplicates2(nums: &mut Vec<i32>) -> i32 {
        let mut slow = 0;
        for i in 0..nums.len() {
            if slow < 2 || nums[i] != nums[slow - 2 as usize] {
                nums[slow] = nums[i];
                slow += 1;
            }
        }
        return slow as i32;
    }

    pub fn remove_duplicates_n(nums: &mut Vec<i32>, k :i32) -> i32 {
        let mut u = 0;
        for i in 0..nums.len() {
            if u < k || nums[i] != nums[u - k as usize] {
                nums[u] = nums[i];
                u += 1;
            }
        }
        return u as i32;
    }
}


#[cfg(test)]
mod tests {
    use crate::two_pointer::TwoPointer;

    #[test]
    fn it_works() {
        use super::TwoPointer;
        let mut vec = vec![1, 1, 2];

        let res = TwoPointer::remove_duplicates(&mut vec);
        println!("{}", res);
    }
}