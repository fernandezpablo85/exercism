package foodchain

import "strings"

const (
  TestVersion = 1
  start = `I know an old lady who swallowed a `
  end = `I don't know why she swallowed the fly. Perhaps she'll die.`
  wiggle = `wriggled and jiggled and tickled inside her.`
)

var animals = [...]string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var comments = [...]string{
  `It ` + wiggle,
  `How absurd to swallow a bird!`,
  `Imagine that, to swallow a cat!`,
  `What a hog, to swallow a dog!`,
  `Just opened her throat and swallowed a goat!`,
  `I don't know how she swallowed a cow!`,
  `She's dead, of course!`}

func Verse(n int) string {
  n = n - 1
  verse := make([]string, 0, 10)
  verse = append(verse, start + animals[n] + ".")

  // first verse has no comment.
  if n > 0 {
    verse = append(verse, comments[n-1])
  }

  // last verse ends prematurely.
  if n == 7 {
    return strings.Join(verse, "\n")
  }
  for i := n; i > 0; i-- {
    trail := "."
    if i == 2 && n > 1 {
      trail = ` that ` + wiggle
    }
    verse = append(verse, `She swallowed the ` + animals[i] + ` to catch the `+ animals[i-1] + trail)
  }
  verse = append(verse, end)
  return strings.Join(verse, "\n")
}

func Verses(ns ...int) string {
  verses := make([]string, 0, 0)
  for i := 0; i < len(ns); i++ {
    verses = append(verses, Verse(ns[i]))
  }
  return strings.Join(verses, "\n\n")
}

func Song() string {
  return Verses(1,2,3,4,5,6,7,8)
}