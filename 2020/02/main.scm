(load "lib.scm")

(display "part one: ")
(println (length (filter part-one-password-valid? (map string->entry input))))

(display "part two: ")
(println (length (filter part-two-password-valid? (map string->entry input))))
