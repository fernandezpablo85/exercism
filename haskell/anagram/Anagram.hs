{-# OPTIONS_GHC -Wall #-}
module Anagram(anagramsFor) where

import Data.List(sort)
import Data.Char(toLower)

anagramsFor :: String -> [String] -> [String]
anagramsFor orig = filter (isAnagram orig)

isAnagram :: String -> String -> Bool
isAnagram a b = notSame && anagram
  where notSame = lower a /= lower b
        anagram = let normA = norm a in seq normA (normA == norm b)
        norm = sort . lower
        lower = map toLower
