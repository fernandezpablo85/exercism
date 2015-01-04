{-# OPTIONS_GHC -Wall #-}

module Phone(areaCode, number, prettyPrint) where

import Data.Char(isDigit)

number :: String -> String
number n
  | length digits == 10 = digits
  | length digits == 11 && head digits == '1' = tail digits
  | otherwise = "0000000000"
 where digits = filter isDigit n

areaCode :: String -> String
areaCode = take 3 . number

prettyPrint :: String -> String
prettyPrint n = area ++ first ++ "-" ++ second
  where
    area = "(" ++ areaCode n ++ ") "
    first = drop 3 . take 6 $ number n
    second = drop 6 $ number n