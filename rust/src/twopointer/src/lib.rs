#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}

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
}

fn is_alnum(s: char) -> bool {
    return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9');
}
