{-# OPTIONS_GHC -Wall #-}
module Anagram(anagramsFor) where

import Data.List(sort)
import Data.Char(toLower)

anagramsFor :: String -> [String] -> [String]
anagramsFor orig = filter (isAnagram orig)

isAnagram :: String -> String -> Bool
isAnagram a b = notSame && anagram
  where notSame = lower a /= lower b
        anagram = norm a == norm b
        norm = sort . lower
        lower = map toLower
