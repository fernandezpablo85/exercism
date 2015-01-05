{-# OPTIONS_GHC -Wall #-}
module Phone(areaCode, number, prettyPrint) where
import Data.Char(isDigit)

number :: String -> String
number n
  | size == 10 = digits
  | size == 11 && head digits == '1' = tail digits
  | otherwise = "0000000000"
  where
    digits = filter isDigit n
    size = length digits

areaCode :: String -> String
areaCode ns = area
  where (area, _, _) = pieces ns

pieces :: String -> (String, String, String)
pieces ns = (area, first, second)
  where (area, rest) = splitAt 3 $ number ns
        (first, second) = splitAt 3 rest

prettyPrint :: String -> String
prettyPrint n = "(" ++ area ++ ") " ++ first ++ "-" ++ second
  where (area, first, second) = pieces n