{-# OPTIONS_GHC -Wall #-}
module Anagram(anagramsFor) where

import Data.List(sort)
import Data.Char(toLower)

anagramsFor :: String -> [String] -> [String]
anagramsFor orig = filter (isAnagram orig)

isAnagram :: String -> String -> Bool
isAnagram a b = notSame && anagram
  where norm = sort . lower
        lower = map toLower
        normA = norm a
        notSame = lower a /= lower b
        anagram = normA == norm b

