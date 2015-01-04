class PhoneNumber(input: String) {

  val digits = input.filter(_.isDigit)

  val number = digits.size match {
    case 11 if digits.head == '1' => digits.tail
    case 10 => digits
    case _ => PhoneNumber.Invalid
  }

  val areaCode = number.substring(0,3)

  override val toString = {
    val first = number.substring(3,6)
    val second = number.substring(6,10)
    s"($areaCode) $first-$second"
  }
}

object PhoneNumber {
  val Invalid = "0000000000"
}