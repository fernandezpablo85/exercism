class Bob {

  val hey: String => String = {
    case s if isEmpty(s) => "Fine. Be that way!"
    case s if isShout(s) => "Whoa, chill out!"
    case s if isQuestion(s) => "Sure."
    case _ => "Whatever."
  }

  private def isEmpty(message: String) = message.trim.isEmpty

  private def isQuestion(message: String) = message.endsWith("?")

  private def isShout(message: String) = {
    def allUppercase = message == message.toUpperCase()
    def containsLetters = message.exists(_.isLetter)

    allUppercase && containsLetters
  }
}