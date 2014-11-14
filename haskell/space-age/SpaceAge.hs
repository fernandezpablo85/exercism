{-# OPTIONS_GHC -Wall #-}

module SpaceAge(Planet(..), ageOn) where

data Planet = Earth | Venus | Mercury | Mars | Jupiter | Saturn | Uranus | Neptune

ageOn :: Planet -> Float -> Float
ageOn planet days = days / (period planet)

period :: Planet -> Float
period Earth = 31557600
period Mercury = period Earth * 0.2408467
period Venus = period Earth * 0.61519726
period Mars = period Earth * 1.8808158
period Jupiter = period Earth * 11.862615
period Saturn = period Earth * 29.447498
period Uranus = period Earth * 84.016846
period Neptune = period Earth * 164.79132