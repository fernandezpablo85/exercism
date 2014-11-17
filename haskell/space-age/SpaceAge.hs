{-# OPTIONS_GHC -Wall #-}

module SpaceAge(Planet(..), ageOn) where

data Planet = Earth | Venus | Mercury | Mars | Jupiter | Saturn | Uranus | Neptune

ageOn :: Planet -> Float -> Float
ageOn planet days = days / period planet

period :: Planet -> Float
period planet = secondsPerYearInEarth * case planet of
  Earth -> 1
  Mercury -> 0.2408467
  Venus -> 0.61519726
  Mars -> 1.8808158
  Jupiter -> 11.862615
  Saturn -> 29.447498
  Uranus -> 84.016846
  Neptune -> 164.79132
  where secondsPerYearInEarth = 31557600
