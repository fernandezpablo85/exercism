object Hamming {
  type Distance = Int

  def compute (first: String, second: String): Distance = {
    def differences(chars: (Char, Char)) = chars._1 != chars._2

    (first zip second).count(differences)
  }
}