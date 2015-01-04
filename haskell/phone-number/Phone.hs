{-# OPTIONS_GHC -Wall #-}

module Phone(areaCode, number, prettyPrint) where

import Data.Char(isDigit)

clean :: String -> String
clean = stripExtra1 . filter isDigit

stripExtra1 :: String -> String
stripExtra1 n = if extra1 then drop 1 n else n
  where extra1 = length n == 11 && head n == '1'

-- ideally these woulde be ifValid (take 3) and ifValid id
-- can't figure out how to do it (yet)
areaCode :: String -> String
areaCode n = ifValid (take 3) (clean n)

number :: String -> String
number n = ifValid id (clean n)

ifValid :: (String -> String) -> String -> String
ifValid f n
  | valid n = f n
  | otherwise = "0000000000"

valid :: String -> Bool
valid = (==10) . length

prettyPrint :: String -> String
prettyPrint n = area ++ first ++ "-" ++ second
  where
    area = "(" ++ areaCode n ++ ") "
    first = drop 3 . take 6 $ clean n
    second = drop 6 $ clean n
