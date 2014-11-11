{-# OPTIONS_GHC -Wall #-}

module Bob (responseFor) where

import Data.List (isSuffixOf)
import Data.Char (isUpper, isLetter, isSpace)

responseFor :: String -> String
responseFor msg
    | isBlank = "Fine. Be that way!"
    | isShout = "Whoa, chill out!"
    | isQuestion = "Sure."
    | otherwise = "Whatever."
    where
      isBlank = all isSpace msg
      isQuestion = "?" `isSuffixOf` msg
      isShout = (any isLetter msg) && (all isUpper letters)
      letters = filter isLetter msg