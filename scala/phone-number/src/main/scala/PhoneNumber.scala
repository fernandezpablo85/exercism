class PhoneNumber(input: String) {
  val onlyNumbers = input.filter(_.isDigit)
  val clean = if (onlyNumbers.startsWith("11")) onlyNumbers.substring(1) else onlyNumbers
  val isValid = clean.size >= 10 && clean.size < 11
  val number = if (isValid) clean else PhoneNumber.Invalid
  val areaCode = clean.substring(0,3)

  override def toString = {
    val first = clean.substring(3,6)
    val second = clean.substring(6,10)
    s"($areaCode) $first-$second"
  }
}

object PhoneNumber {
  val Invalid = "0000000000"
}