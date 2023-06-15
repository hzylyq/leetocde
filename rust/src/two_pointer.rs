pub struct TwoPointer {}

impl TwoPointer {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        let mut slow = 0;
        for i in 0..nums.len() {
            if nums[i] != nums[slow as usize] {
                nums[slow as usize] = nums[i];
                slow += 1;
            }
        }

        return slow + 1;
    }
}


#[cfg(test)]
mod tests {
    use crate::two_pointer::TwoPointer;

    #[test]
    fn it_works() {
        use super::TwoPointer;
        let mut vec=vec![1,1,2];

        let res = TwoPointer::remove_duplicates(&mut vec);
        println!("{}", res);
    }
}