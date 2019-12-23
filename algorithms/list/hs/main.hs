data List a = Empty | Fork a (List a) 

cons :: a -> List a -> List a
cons = Fork

mkEmpty :: List a
mkEmpty = Empty 

foldl' :: (a -> b -> b) -> b -> List a -> b
foldl' _ d Empty = d
foldl' f d (Fork x Empty) = f x d
foldl' f d (Fork x xs) = foldl' f (f x d) xs
-- f (f (f (f (a1 d) a2) a3) a4)

foldr' :: (a -> b -> b) -> b -> List a -> b
foldr' _ d Empty = d
foldr' f d (Fork x Empty) = f x d 
foldr' f d (Fork x xs) = f x (foldr' f d xs)
-- f (a1 f (a2 f (a3 f (a4 d))))
