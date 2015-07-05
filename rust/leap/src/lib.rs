/// tests if a year is a leap year.
///
/// # Examples
///
/// ```
/// use leap;
///
/// assert_eq!(leap::is_leap_year(1990), false);
/// assert_eq!(leap::is_leap_year(1984), true);
/// ```
pub fn is_leap_year (y: i32) -> bool {
  let divisible_by = |n: i32| -> bool { y % n == 0 };
  (divisible_by(400)) || (divisible_by(4) && !divisible_by(10))
}

