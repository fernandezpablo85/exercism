module HelloWorld exposing (helloWorld)


helloWorld : Maybe String -> String
helloWorld maybe =
    case maybe of
        Just name ->
            "Hello, " ++ name ++ "!"

        Nothing ->
            "Hello, World!"
