class Phrase(content: String) {

  lazy val words = content.toLowerCase.split(Phrase.OnlyWordsRegex)

  lazy val wordCount = words.foldLeft(Phrase.EmptyMap) { case (map, word) =>
    val count = map.get(word).getOrElse(0)
    map + (word -> (count + 1))
  }
}

object Phrase {
  private val OnlyWordsRegex = "[^a-z0-9']+"
  private val EmptyMap = Map.empty[String, Int]
}