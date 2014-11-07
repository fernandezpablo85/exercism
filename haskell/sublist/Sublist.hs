{-# OPTIONS_GHC -Wall #-}

module Sublist (sublist, Sublist(..)) where

data Sublist = Sublist | Equal | Superlist | Unequal deriving (Show, Eq)

sublist :: Eq a => [a] -> [a] -> Sublist
sublist a b
  | a == b = Equal
  | a == [] = Sublist
  | b == [] = Superlist
  | contains a b = Sublist
  | contains b a = Superlist
  | otherwise = Unequal

contains :: Eq a => [a] -> [a] -> Bool
contains a b
  | empty a = False
  | empty b = False
  | otherwise = if (a == take (length a) b) then True else contains a (tail b)
    where empty xs = length xs == 0