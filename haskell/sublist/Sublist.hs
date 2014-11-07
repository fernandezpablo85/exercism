{-# OPTIONS_GHC -Wall #-}

module Sublist (sublist, Sublist(..)) where

data Sublist = Sublist | Equal | Superlist | Unequal deriving (Show, Eq)

sublist :: Eq a => [a] -> [a] -> Sublist
sublist a b
  | a == b = Equal
  | contains a b = Sublist
  | contains b a = Superlist
  | otherwise = Unequal

contains :: Eq a => [a] -> [a] -> Bool
contains _ [] = False
contains [] _ = True
contains a b = a `isPrefixOf` b || contains a (tail b)

isPrefixOf :: Eq a => [a] -> [a] -> Bool
isPrefixOf [] _ = True
isPrefixOf _ [] = False
isPrefixOf (a:as) (b:bs) = a == b && as `isPrefixOf` bs