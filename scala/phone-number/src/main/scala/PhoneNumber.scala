class PhoneNumber(input: String) {

  val digits = input.filter(_.isDigit)

  val number = digits.size match {
    case 11 if digits.head == '1' => digits.tail
    case 10 => digits
    case _ => "0" * 10
  }

  val (areaCode, first, second) =
    (number.substring(0,3), number.substring(3,6), number.substring(6,10))

  override val toString = s"($areaCode) $first-$second"
}