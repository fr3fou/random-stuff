a = [1..5]
b = [6..10]
main :: IO ()
main = print [(x, y) | x <- a, y <- b]
