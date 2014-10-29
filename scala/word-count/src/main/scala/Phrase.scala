class Phrase(content: String) {
  import Phrase._
  lazy val words = content.toLowerCase.split(OnlyWordsRegex)
  lazy val wordCount = words.foldLeft(EmptyMap)(CountWords)
}

object Phrase {
  type WordCount = Map[String, Int]

  private val OnlyWordsRegex = "[^a-z0-9']+"
  private val EmptyMap = Map.empty[String, Int].withDefaultValue(0)
  private val CountWords: (WordCount, String) => WordCount = {
    case (map, word) =>
      val count = map(word)
      map + (word -> (count + 1))
  }
}