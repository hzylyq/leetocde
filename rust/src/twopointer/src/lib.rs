mod array;
mod stack;
mod map;
mod grady;

pub struct Solution;

impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let mut sgood = s
            .chars()
            .filter(|c| c.is_alphanumeric())
            .map(|c| c.to_ascii_lowercase());

        while let (Some(a), Some(b)) = (sgood.next(), sgood.next_back()) {
            if a != b {
                return false;
            }
        }

        return true;
    }

    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        if m == 0 && n == 0 {
            return;
        }

        let mut idx = (m + n - 1) as usize;
        let mut i = m - 1;
        let mut j = n - 1;

        while i >= 0 && j >= 0 {
            if nums1[i as usize] < nums2[j as usize] {
                nums1[idx] = nums2[j as usize];
                j -= 1;
            } else {
                nums1[idx] = nums1[i as usize];
                i -= 1;
            }
            idx -= 1;
        }

        while j >= 0 {
            nums1[idx] = nums2[j as usize];
            idx -= 1;
            j -= 1;
        }
    }
}

fn is_alnum(s: char) -> bool {
    return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9');
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let mut vec = &[1, 2, 3, 0, 0, 0].repeat(6);
    }
}