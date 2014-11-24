class Anagram(text: String) {

  val cannonicalText = cannonical(text)

  def cannonical(string: String): String =
    string.toLowerCase.sorted

  def matches(candidates: Seq[String]): Seq[String] = {
    def isAnagram(candidate: String) =
      cannonical(candidate) == cannonicalText

    def sameWord(candidate: String) =
      candidate.toLowerCase == text.toLowerCase

    def mustBeIncluded(candidate: String, rest: Seq[String]) = {
      val alreadyIncluded =
        rest.map(_.toLowerCase).contains(candidate.toLowerCase)

      isAnagram(candidate) && !sameWord(candidate) && !alreadyIncluded
    }

    def pickAnagrams(rest: Seq[String], candidate: String): Seq[String] = {
      if (mustBeIncluded(candidate, rest))
        rest ++ Seq(candidate)
      else
        rest
    }

    candidates.foldLeft(Seq.empty[String])(pickAnagrams)
  }
}