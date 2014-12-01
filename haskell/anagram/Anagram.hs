{-# OPTIONS_GHC -Wall #-}
module Anagram(anagramsFor) where

import Data.List(sort)
import Data.Char(toLower)

anagramsFor :: String -> [String] -> [String]
anagramsFor orig = filter (isAnagram orig (norm orig))

isAnagram :: String -> String -> String -> Bool
isAnagram orig a b = notSame && anagram
  where notSame = lower orig /= lower b
        anagram = a == norm b

norm :: String -> String
norm = sort . lower

lower :: String -> String
lower = map toLower

